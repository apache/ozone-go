// Package io TODO
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package io

import (
    "fmt"
    dnClient "github.com/apache/ozone-go/api/datanodeclient"
    ozoneError "github.com/apache/ozone-go/api/errors"
    "github.com/apache/ozone-go/api/proto/datanode"
    "github.com/apache/ozone-go/api/proto/hdds"

    log "github.com/sirupsen/logrus"
    "google.golang.org/protobuf/proto"
)

// BlockOutputStream TODO
type BlockOutputStream struct {
    blockId             *datanode.DatanodeBlockID
    blockData           *datanode.BlockData
    pipelineOperator    *dnClient.PipelineOperator
    commitWatcher       *CommitWatcher
    xceiverManager      *dnClient.XceiverClientManager
    xceiverClient       dnClient.XceiverClientSpi
    ExcludeListOperator *dnClient.ExcludeListOperator
    length              uint64
    offset              uint64
    chunkOffset         uint64
    checksumType        *datanode.ChecksumType
    blockSize           uint64
    bytesPerChecksum    uint32
    chunkSize           uint64
    blockIndex          int
    // the effective length of data flushed so far
    totalDataFlushedLength uint64

    // effective data write attempted so far for the block
    writtenDataLength     uint64
    totalAckDataLength    uint64
    streamBufferFlushSize uint64 // 16MB
    streamBufferMaxSize   uint64 // 32MB
    chunks                []*ChunkOutputStream
    chunkIndex            int
    ackIndex              int
    closed                bool
    Success               bool
    putBlockCommitIndex   uint64
}

// NewBlockOutputStream TODO
func NewBlockOutputStream(blockId *datanode.DatanodeBlockID, xceiverManager *dnClient.XceiverClientManager,
    xceiverClient dnClient.XceiverClientSpi, checksumType *datanode.ChecksumType,
    blockSize uint64, bytesPerChecksum uint32, chunkSize uint64, streamBufferFlushSize uint64,
    streamBufferMaxSize uint64, pipeline *dnClient.PipelineOperator) *BlockOutputStream {
    var metadata []*datanode.KeyValue
    kvMetaData := &datanode.KeyValue{Key: proto.String("TYPE"), Value: proto.String("KEY")}
    metadata = append(metadata, kvMetaData)
    blockData := &datanode.BlockData{
        BlockID:  blockId,
        Metadata: metadata,
        Chunks:   make([]*datanode.ChunkInfo, 0),
    }

    commitWatcher := &CommitWatcher{
        xceiverClient:              xceiverClient,
        CommitIndex2flushedDataMap: make(map[uint64]int),
    }
    return &BlockOutputStream{
        xceiverManager:        xceiverManager,
        xceiverClient:         xceiverClient,
        blockId:               blockId,
        blockData:             blockData,
        blockSize:             blockSize,
        pipelineOperator:      pipeline,
        length:                0,
        checksumType:          checksumType,
        bytesPerChecksum:      bytesPerChecksum,
        chunkSize:             chunkSize,
        chunks:                make([]*ChunkOutputStream, 0),
        ExcludeListOperator:   dnClient.NewExcludeListOperator(),
        streamBufferFlushSize: streamBufferFlushSize,
        streamBufferMaxSize:   streamBufferMaxSize,
        putBlockCommitIndex:   0,
        ackIndex:              -1,
        blockIndex:            0,
        closed:                false,
        Success:               true,
        commitWatcher:         commitWatcher,
    }

}

func (b *BlockOutputStream) allocateChunk() (int, *ChunkOutputStream) {
    nextIndex := len(b.chunks)
    chunkName := fmt.Sprintf("%d_chunk_%d", b.blockId.GetLocalID(), nextIndex+1)
    writer := newChunkOutputStream(chunkName, b.chunkSize, b.blockId, b.checksumType, b.bytesPerChecksum, nextIndex,
        b.xceiverClient)
    b.chunks = append(b.chunks, writer)
    return nextIndex, writer
}

func (b *BlockOutputStream) allocateChunkIfNeed() *ChunkOutputStream {

    var current *ChunkOutputStream
    if len(b.chunks) == 0 {
        b.chunkIndex, current = b.allocateChunk()
        return current
    }

    current = b.chunks[b.chunkIndex]
    if !current.HasRemaining() {
        b.chunkIndex, current = b.allocateChunk()
        return current
    }
    return current
}

// BlockId TODO
func (b *BlockOutputStream) BlockId() *datanode.DatanodeBlockID {
    return b.blockId
}

// Pipeline TODO
func (b *BlockOutputStream) Pipeline() *hdds.Pipeline {
    return b.pipelineOperator.Pipeline
}

func (b *BlockOutputStream) handleFullBuffer() error {
    return b.watchForCommit(true)
}

func (b *BlockOutputStream) setClosedAndSuccess(closed, success bool) {
    b.closed = true
    b.Success = false
}

// Write TODO
func (b *BlockOutputStream) Write(p []byte) (int, error) {
    if b.closed {
        return 0, ozoneError.ClosedPipeErr
    }
    writtenLen := 0
    remain := len(p)
    for remain > 0 {
        currentChunk := b.allocateChunkIfNeed()
        var endPos int
        if currentChunk.Remaining() < uint64(remain) {
            endPos = writtenLen + int(currentChunk.Remaining())
        } else {
            endPos = writtenLen + remain
        }
        log.Debug(fmt.Sprintf("%d write to chunk %s remain:%d start:%d end:%d, chunk remain: %d", b.blockIndex,
            currentChunk.chunkInfo.GetChunkName(), remain, writtenLen, endPos, currentChunk.Remaining()))
        n, err := currentChunk.Write(p[writtenLen:endPos])
        writtenLen += n
        remain -= n
        if err != nil && err != ozoneError.BufferFullErr {
            log.Error(fmt.Sprintf("%d write data to chunk buffer error: %s", b.blockIndex, err.Error()))
            return writtenLen, err
        }

        if !currentChunk.HasRemaining() {
            writingIndex := b.chunkIndex
            log.Debug(fmt.Sprintf("%d start write chunk %s to container", b.blockIndex, currentChunk.chunkInfo.GetChunkName()))
            if we := b.WriteChunkToContainerAsync(writingIndex); we != nil {
                log.Error(fmt.Sprintf("write chunk %d to container error: %s", writingIndex, we.Error()))
                return writtenLen, we
            }

            log.Debug(fmt.Sprintf("%d end write chunk %s to container", b.blockIndex, currentChunk.chunkInfo.GetChunkName()))

            // 当达到streamBufferFlushSize的时候，flush chunk，put block
            if b.shouldFlushChunk() {
                if pe := b.executePutBlock(false); pe != nil {
                    log.Error(fmt.Sprintf("execute put block error: %s", pe.Error()))
                    return writtenLen, pe
                }
                b.totalDataFlushedLength = b.writtenDataLength
                b.commitWatcher.UpdateCommitInfoMap(b.putBlockCommitIndex, writingIndex)
            }

            // 当达到streamBufferMaxSize时，等待之前的所有异步操作结束，获取commitIndex，释放内存
            if b.totalDataFlushedLength%b.streamBufferMaxSize == 0 {
                log.Debug(fmt.Sprintf("%d start handle full buffer and chunk name is %s",
                    b.blockIndex, currentChunk.chunkInfo.GetChunkName()))
                if he := b.handleFullBuffer(); he != nil {
                    log.Error(fmt.Sprintf("handle full buffer error: %s", he.Error()))
                    return writtenLen, he
                }
                log.Debug(fmt.Sprintf("%d end handle full buffer and chunk name is %s",
                    b.blockIndex, currentChunk.chunkInfo.GetChunkName()))
            }
        }
    }
    return writtenLen, nil
}

func (b *BlockOutputStream) executePutBlock(eof bool) error {
    log.Debug("put block eof ", eof)
    return b.putBlockAsync(eof)
}

func (b *BlockOutputStream) shouldFlushChunk() bool {
    return b.writtenDataLength%b.streamBufferFlushSize == 0
}

func (b *BlockOutputStream) getCurrentChunk() *ChunkOutputStream {
    return b.chunks[b.chunkIndex]
}

func (b *BlockOutputStream) releaseBuffer(commitIndex uint64) {
    for logIndex, chunkIndex := range b.commitWatcher.CommitIndex2flushedDataMap {
        if logIndex <= commitIndex && b.ackIndex <= chunkIndex {
            b.ackIndex = chunkIndex
        }
    }

    for i := 0; i <= b.ackIndex; i++ {
        b.totalAckDataLength += b.chunks[i].length
        if err := b.chunks[i].Close(); err != nil {
            log.Warnf("Release chunk %d error:%s", i, err.Error())
        }
    }
}

func (b *BlockOutputStream) watchForCommit(bufferFull bool) error {
    var reply *dnClient.XceiverClientReply
    var err error
    if bufferFull {
        reply, err = b.commitWatcher.WatchOnLastIndex()
    } else {
        reply, err = b.commitWatcher.WatchOnFirstIndex()
    }
    if err != nil {
        return err
    } else {
        index := reply.GetLogIndex()
        b.releaseBuffer(index)
        if len(reply.Datanodes) > 0 {
            b.ExcludeListOperator.AddDataNodes(reply.Datanodes)
        }
        return nil
    }
}

// Close TODO
func (b *BlockOutputStream) Close() error {

    if b.closed && b.Success {
        return nil
    }

    if err := b.flush(true); err != nil {
        log.Error("flush close error", err)
        return err
    }
    b.Clean()
    return nil
}

// Clean TODO
func (b *BlockOutputStream) Clean() {
    if b.closed && b.Success {
        return
    }
    b.closed = true
    b.Success = true
    for _, writer := range b.chunks {
        if writer.written {
            if err := writer.Close(); err != nil {
                log.Warnf("Close block output stream %s error: %s", writer.blockId.String(), err.Error())
            }
        }
    }

}

// IsClosed TODO
func (b *BlockOutputStream) IsClosed() bool {
    return b.closed
}

// Len TODO
func (b *BlockOutputStream) Len() uint64 {
    return b.length
}

// HasRemaining TODO
func (b *BlockOutputStream) HasRemaining() bool {
    if b.closed {
        return false
    }
    return b.blockSize > b.length
}

// Remaining TODO
func (b *BlockOutputStream) Remaining() uint64 {
    if b.closed {
        return 0
    }
    return b.blockSize - b.length
}

func (b *BlockOutputStream) putBlockAsync(eof bool) error {
    reply, err := dnClient.PutBlock(b.xceiverClient, b.blockData, eof)
    if err != nil {
        return err
    }
    b.putBlockCommitIndex = reply.GetLogIndex()
    return nil
}

// WriteChunkToContainerAsync TODO
func (b *BlockOutputStream) WriteChunkToContainerAsync(writerIndex int) error {
    if b.chunks[writerIndex].written {
        return nil
    }
    log.Debug(fmt.Sprintf("start writer chunk_%d", writerIndex+1))
    if err := b.chunks[writerIndex].WriteToContainer(); err != nil {
        log.Error(fmt.Sprintf("write %d th chunk to container error", writerIndex+1), err)
        return err
    }
    log.Debug(fmt.Sprintf("end writer chunk_%d", writerIndex+1))
    b.blockData.Chunks = append(b.blockData.Chunks, b.chunks[writerIndex].ChunkInfo())
    b.writtenDataLength += b.chunks[writerIndex].length
    return nil
}

func (b *BlockOutputStream) getNotFlushedAckData() []byte {
    data := make([]byte, 0)
    for _, writer := range b.chunks {
        if !writer.IsClosed() {
            data = append(data, writer.buffer.Bytes()...)
        }
    }
    return data
}

func (b *BlockOutputStream) flush(close bool) error {
    if len(b.chunks) == 0 {
        return nil
    }
    if b.totalDataFlushedLength == 0 ||
        b.totalDataFlushedLength < b.writtenDataLength ||
        b.chunks[b.chunkIndex].HasRemaining() {
        if b.chunks[b.chunkIndex].HasRemaining() {
            if err := b.WriteChunkToContainerAsync(b.chunkIndex); err != nil {
                return err
            }
        }
        if err := b.executePutBlock(close); err != nil {
            return err
        }
        b.totalDataFlushedLength = b.writtenDataLength
        b.commitWatcher.UpdateCommitInfoMap(b.putBlockCommitIndex, b.chunkIndex)
        log.Debug(fmt.Sprintf("flush block close: %v, UpdateCommitInfoMap", close))
    } else if close {
        if err := b.executePutBlock(true); err != nil {
            return err
        }
    }
    log.Debug(fmt.Sprintf("flush block close: %v", close))
	return b.watchForCommit(false)
}
