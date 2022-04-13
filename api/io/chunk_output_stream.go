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
    "github.com/apache/ozone-go/api/errors"
    "github.com/apache/ozone-go/api/proto/datanode"
    "github.com/apache/ozone-go/api/utils"

    log "github.com/sirupsen/logrus"
)

// ChunkOutputStream TODO
type ChunkOutputStream struct {
    blockId          *datanode.DatanodeBlockID
    checksumType     *datanode.ChecksumType
    chunkInfo        *datanode.ChunkInfo
    xceiverClient    dnClient.XceiverClientSpi
    bytesPerChecksum uint32
    chunkSize        uint64
    length           uint64
    buffer           *limitBuffer
    index            int
    closed           bool
    written          bool
    logIndex         uint64
}

func newChunkOutputStream(chunkName string, chunkSize uint64, blockId *datanode.DatanodeBlockID,
    checksumType *datanode.ChecksumType, bytesPerChecksum uint32, chunkIndex int,
    xceiverClient dnClient.XceiverClientSpi) *ChunkOutputStream {

    chunkInfo := &datanode.ChunkInfo{
        ChunkName: utils.PointString(chunkName),
        Offset:    utils.PointUint64(0),
        Len:       utils.PointUint64(0),
    }

    c := &ChunkOutputStream{
        blockId:          blockId,
        checksumType:     checksumType,
        bytesPerChecksum: bytesPerChecksum,
        chunkInfo:        chunkInfo,
        chunkSize:        chunkSize,
        length:           0,
        xceiverClient:    xceiverClient,
        index:            chunkIndex,
        logIndex:         0,
        closed:           false,
        written:          false,
    }
    c.buffer = newLimitBuffer(c.chunkSize)
    return c
}

// ChunkInfo TODO
func (c *ChunkOutputStream) ChunkInfo() *datanode.ChunkInfo {
    return c.chunkInfo
}

func (c *ChunkOutputStream) checkWritable() error {
    if c.written {
        return fmt.Errorf("chunk: %s has written to container", *c.chunkInfo.ChunkName)
    }

    if c.closed {
        return errors.ClosedPipeErr
    }
    return nil
}

// Write TODO
func (c *ChunkOutputStream) Write(p []byte) (n int, err error) {
    if err = c.checkWritable(); err != nil {
        return 0, err
    }
    n, err = c.buffer.Write(p)
    c.length += uint64(n)

    return n, err
}

// Close TODO
func (c *ChunkOutputStream) Close() error {
    if c.closed {
        return nil
    }
    if err := c.WriteToContainer(); err != nil {
        log.Error("WriteToContainer close err", err)
        return err
    }

    c.written = true
    c.closed = true
    // Clear buffer when writer is closed
    c.buffer = nil
    return nil
}

// Remaining TODO
func (c *ChunkOutputStream) Remaining() uint64 {
    return c.chunkSize - c.length
}

// HasRemaining TODO
func (c *ChunkOutputStream) HasRemaining() bool {
    return c.chunkSize > c.length
}

// Len TODO
func (c *ChunkOutputStream) Len() uint64 {
    return c.length
}

// IsClosed TODO
func (c *ChunkOutputStream) IsClosed() bool {
    return c.closed
}

// WriteToContainer TODO
func (c *ChunkOutputStream) WriteToContainer() error {
    if c.written {
        return nil
    }
    var data []byte
    if c.length == 0 {
        data = []byte{}
    } else {
        data = c.buffer.Bytes()
    }
    checksum := dnClient.NewChecksumOperatorComputer(c.checksumType, c.bytesPerChecksum)
    if err := checksum.ComputeChecksum(data); err != nil {
        return err
    }
    checksumData := checksum.Checksum

    c.chunkInfo.ChecksumData = checksumData

    c.chunkInfo.Offset = utils.PointUint64(uint64(c.index) * c.chunkSize)
    c.chunkInfo.Len = &c.length
    reply, err := dnClient.WriteChunk(c.xceiverClient, c.chunkInfo, c.blockId, data)
    if err != nil {
        return err
	}
	c.written = true
	c.logIndex = reply.GetLogIndex()
	return nil
}
