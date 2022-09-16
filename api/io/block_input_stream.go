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
    "io"
    dnClient "github.com/apache/ozone-go/api/datanodeclient"
    "github.com/apache/ozone-go/api/proto/datanode"
    "github.com/apache/ozone-go/api/proto/hdds"
    "github.com/apache/ozone-go/api/proto/ozone"
    "sync"

    log "github.com/sirupsen/logrus"
)

// BlockInputStream TODO
type BlockInputStream struct {
    blockId          *hdds.BlockID
    location         *ozone.KeyLocation
    xceiverManager   *dnClient.XceiverClientManager
    xceiverClient    dnClient.XceiverClientSpi
    pipelineOperator *dnClient.PipelineOperator
    chunks           []*ChunkInputStream
    length           int64
    offset           int64
    blockIndex       int
    chunkIndex       int
    verifyChecksum   bool
    closed           bool
    initialized      bool
    mutex            sync.Mutex
}

// ConvertBlockIdToDataNodeBlockID TODO
func ConvertBlockIdToDataNodeBlockID(bid *hdds.BlockID) *datanode.DatanodeBlockID {
    id := datanode.DatanodeBlockID{
        ContainerID:           bid.ContainerBlockID.ContainerID,
        LocalID:               bid.ContainerBlockID.LocalID,
        BlockCommitSequenceId: bid.BlockCommitSequenceId,
    }
    return &id
}

// NewBlockInputStream TODO
func NewBlockInputStream(location *ozone.KeyLocation, xceiverManager *dnClient.XceiverClientManager,
    xceiverClient dnClient.XceiverClientSpi, pipelineOperator *dnClient.PipelineOperator,
    verifyChecksum bool) *BlockInputStream {
    return &BlockInputStream{
        blockId:          location.GetBlockID(),
        location:         location,
        xceiverManager:   xceiverManager,
        pipelineOperator: pipelineOperator,
        xceiverClient:    xceiverClient,
        chunks:           make([]*ChunkInputStream, 0),
        length:           int64(location.GetLength()),
        offset:           0,
        chunkIndex:       0,
        blockIndex:       0,
        verifyChecksum:   verifyChecksum,
        closed:           false,
        initialized:      false,
        mutex:            sync.Mutex{},
    }
}

func (b *BlockInputStream) initialize() error {
    b.mutex.Lock()
    defer b.mutex.Unlock()
    if b.closed {
        log.Debug(fmt.Sprintf("block %d closed", b.blockIndex))
        return nil
    }
    if b.initialized {
        return nil
    }
    log.Debug(fmt.Sprintf("block %d initialized", b.blockIndex))
    dnBlockId := ConvertBlockIdToDataNodeBlockID(b.blockId)
    blockData, err := dnClient.GetBlock(b.xceiverClient, dnBlockId)
    if err != nil {
        return err
    }
    cks := blockData.GetBlockData().GetChunks()
    b.chunks = make([]*ChunkInputStream, 0)
    for _, ck := range cks {
        chunk := newChunkInputStream(ck, dnBlockId, b.xceiverClient, b.verifyChecksum)
        b.chunks = append(b.chunks, chunk)
    }
    b.initialized = true
    b.closed = false
    return nil
}

// Read TODO
func (b *BlockInputStream) Read(buff []byte) (readLen int, err error) {
    if err = b.initialize(); err != nil {
        log.Error("block init error", err)
        return readLen, err
    }

    if b.closed {
        return readLen, io.EOF
    }

    remain := len(buff)
    log.Debug("block buff remain ", remain)
    for remain > 0 {
        chunk := b.chunks[b.chunkIndex]
        n, err := chunk.Read(buff[readLen:])
        readLen += n
        remain -= n
        b.offset += int64(n)
        if err != nil && err != io.EOF {
            return readLen, err
        }

        if !chunk.hasRemainData() {
            chunk.Close()
            b.chunkIndex++
            if b.chunkIndex >= len(b.chunks) {
                break
            }
        }
    }
    log.Debug("block read len: ", readLen, " total read len: ", b.offset, " remain to read: ", b.remainDataLen())
    return readLen, nil
}

// Seek TODO
func (b *BlockInputStream) Seek(offset int64, whence int) (int64, error) {

    if !b.initialized {
        if err := b.initialize(); err != nil {
            return 0, err
        }
    }

    var pos int64
    switch whence {
    case 0:
        pos = offset
    case 1:
        pos = b.offset + offset
    case 2:
        pos = b.length + offset
    default:
        return int64(b.Off()), fmt.Errorf("blockReader: %s invalid whence: %d", b.blockId.String(), whence)
    }

    if pos < 0 || pos > b.length {
        return int64(b.Off()), fmt.Errorf("blockReader: %v invalid resulting position: %d, block length: %v",
            b.blockId.String(), pos, b.Len())
    }

    b.offset = pos

    totalLen := uint64(0)
    for idx, reader := range b.chunks {
        totalLen += reader.Len()
        if uint64(b.offset) < totalLen {
            b.chunkIndex = idx
            if !b.chunks[b.chunkIndex].initialized {
                if err := b.chunks[b.chunkIndex].initialize(); err != nil {
                    return 0, err
                }
            }
            b.chunks[b.chunkIndex].offset = uint64(b.offset - int64(totalLen-reader.Len()))
            if _, err := b.chunks[b.chunkIndex].Seek(0, 1); err != nil {
                return 0, err
            }
            break
        }
    }

    return pos, nil
}

// Close TODO
func (b *BlockInputStream) Close() {
    if b.closed {
        return
    }
    b.Reset()
    b.closed = true
    log.Debug(fmt.Sprintf("block %d close", b.blockIndex))
}

// Reset TODO
func (b *BlockInputStream) Reset() {
    for _, reader := range b.chunks {
        reader.Reset()
    }
    b.chunkIndex = 0
}

func (b *BlockInputStream) hasRemainData() bool {
    return uint64(b.offset) < b.Len()
}

// Len TODO
func (b *BlockInputStream) Len() uint64 {
    return uint64(b.length)
}

// Off TODO
func (b *BlockInputStream) Off() uint64 {
    return uint64(b.offset)
}

func (b *BlockInputStream) remainDataLen() int64 {
	return b.length - b.offset
}
