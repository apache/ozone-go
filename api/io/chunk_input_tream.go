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
    "sync"

    log "github.com/sirupsen/logrus"
)

// ChunkInputStream TODO
type ChunkInputStream struct {
    chunkInfo      *datanode.ChunkInfo
    blockId        *datanode.DatanodeBlockID
    checksum       *dnClient.ChecksumOperator
    xceiverClient  dnClient.XceiverClientSpi
    buffer         *limitBuffer
    length         uint64
    offset         uint64
    closed         bool
    initialized    bool
    verifyChecksum bool
    mutex          sync.Mutex
}

func newChunkInputStream(chunkInfo *datanode.ChunkInfo,
    blockId *datanode.DatanodeBlockID,
    xceiverClient dnClient.XceiverClientSpi,
    verifyChecksum bool) *ChunkInputStream {
    return &ChunkInputStream{
        chunkInfo:      chunkInfo,
        blockId:        blockId,
        xceiverClient:  xceiverClient,
        length:         chunkInfo.GetLen(),
        checksum:       nil,
        verifyChecksum: verifyChecksum,
        offset:         0,
        initialized:    false,
    }
}

func (c *ChunkInputStream) initialize() error {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    if c.initialized {
        return nil
    }
    log.Debug(fmt.Sprintf("chunk %s initialized, length:%d offset:%d",
        c.chunkInfo.GetChunkName(), c.chunkInfo.GetLen(), c.chunkInfo.GetOffset()))
    if err := c.readChunkDataFromDataNode(); err != nil {
        return nil
    }
    log.Debug("readChunkDataFromDataNode")

    c.initialized = true
    c.closed = false
    return nil
}

// Read TODO
func (c *ChunkInputStream) Read(buff []byte) (n int, err error) {
    readLen := 0
    if c.closed {
        return 0, err
    }
    if err := c.initialize(); err != nil {
        return readLen, err
    }
    remain := len(buff)
    log.Debug("chunk buff remain ", remain)
    for remain > 0 {
        n, err := c.buffer.Read(buff)
        readLen += n
        remain -= n
        c.offset += uint64(n)
        if err != nil {
            if err == io.EOF {
                break
            }
            return readLen, err
        }
    }
    log.Debug("chunk read len: ", readLen, " total read len: ", c.offset, " remain to read: ", c.remainDataLen())
    return readLen, nil
}

func (c *ChunkInputStream) readChunkDataFromDataNode() error {
    var err error
    var resp *datanode.ReadChunkResponseProto
    resp, err = dnClient.ReadChunk(c.xceiverClient, c.chunkInfo, c.blockId)
    if err != nil {
        log.Error("read chunk error:", err)
    }
    c.offset = resp.GetChunkData().GetOffset()
    c.buffer = newLimitBuffer(uint64(len(resp.GetData())))
    if _, err = c.buffer.Write(resp.GetData()); err != nil {
        return err
    }
    log.Debug("chunk size: ", c.buffer.Len())
    if c.verifyChecksum {
        checksum := dnClient.NewChecksumOperatorVerifier(c.chunkInfo.GetChecksumData())
        if err = checksum.VerifyChecksum(c.buffer.Bytes(), 0); err != nil {
            return err
        }
    }
    return nil
}

// Seek TODO
func (c *ChunkInputStream) Seek(offset int64, whence int) (int64, error) {
    if !c.initialized {
        if err := c.initialize(); err != nil {
            return 0, err
        }
    }

    var pos int64
    switch whence {
    case 0:
        pos = offset
    case 1:
        pos = int64(c.offset) + offset
    case 2:
        pos = int64(c.Len()) + offset
    default:
        return int64(c.offset), fmt.Errorf("chunkReader: %v invalid whence: %d", *c.chunkInfo.ChunkName, whence)
    }

    if !c.bufferHasPosition(uint64(pos)) {
        return int64(c.offset), fmt.Errorf("chunkReader: %v invalid resulting position: %d, chunk length: %v",
            *c.chunkInfo.ChunkName, pos, c.Len())
    }

    c.offset = uint64(pos)
    bufferOffset := c.Len() - c.remainDataLen()
    if c.offset < bufferOffset {
        return 0, fmt.Errorf("offset  %d lower buffer offset %d", c.offset, bufferOffset)
    }
    diff := c.offset - bufferOffset
    c.buffer.Next(int(diff))
    return pos, nil
}

// Reset TODO
func (c *ChunkInputStream) Reset() {
    c.buffer = nil
    c.offset = 0
    c.initialized = false
    c.closed = false
}

// Off TODO
func (c *ChunkInputStream) Off() uint64 {
    return c.offset
}

// GetChunkInfo TODO
func (c *ChunkInputStream) GetChunkInfo() *datanode.ChunkInfo {
    return c.chunkInfo
}

// Len TODO
func (c *ChunkInputStream) Len() uint64 {
    return c.chunkInfo.GetLen()
}

func (c *ChunkInputStream) bufferHasData() bool {
    return c.buffer != nil && c.buffer.Len() > 0
}

func (c *ChunkInputStream) bufferHasPosition(pos uint64) bool {
    if c.bufferHasData() {
        return pos >= c.offset && pos < c.Len()
    }
    return false
}

func (c *ChunkInputStream) hasRemainData() bool {
    return c.buffer.Len() > 0
}

func (c *ChunkInputStream) remainDataLen() uint64 {
    return c.buffer.Len()
}

// Close TODO
func (c *ChunkInputStream) Close() {
    c.Reset()
    c.closed = true
    log.Debug(fmt.Sprintf("%s close", c.chunkInfo.GetChunkName()))
}
