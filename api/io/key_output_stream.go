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
    "github.com/apache/ozone-go/api/omclient"
    "github.com/apache/ozone-go/api/proto/datanode"
    "github.com/apache/ozone-go/api/proto/hdds"
    "github.com/apache/ozone-go/api/proto/ozone"
    "github.com/apache/ozone-go/api/utils"
    "path"
    "strings"
    "sync"
    "time"

    log "github.com/sirupsen/logrus"
)

// StreamAction TODO
type StreamAction string

var (
    // StreamActionFull TODO
    StreamActionFull StreamAction = "FULL"
    // StreamActionClose TODO
    StreamActionClose StreamAction = "CLOSE"
    // StreamActionFlush TODO
    StreamActionFlush StreamAction = "FLUSH"
)

// KeyOutputStream TODO
// put key
type KeyOutputStream struct {
    keyInfo               *ozone.KeyInfo
    omClient              *omclient.OmClient
    xceiverManager        *dnClient.XceiverClientManager
    clientId              *uint64
    keyLocation           *ozone.KeyLocation
    keyLocations          []*ozone.KeyLocation
    checksumType          *datanode.ChecksumType
    blocks                []*BlockOutputStream
    fileName              string
    volumeName            string
    bucketName            string
    keyName               string
    path                  string
    length                uint64
    offset                uint64
    blockSize             uint64
    chunkSize             uint64
    bytesPerChecksum      uint32
    blockIndex            int
    closed                bool
    streamBufferFlushSize uint64 // 16MB
    streamBufferMaxSize   uint64 // 32MB
    currentWritingBytes   uint64
    excludeList           *hdds.ExcludeListProto
    mux                   sync.Mutex
}

// NewKeyOutputStream TODO
func NewKeyOutputStream(keyInfo *ozone.KeyInfo, xceiverManager *dnClient.XceiverClientManager, volume string,
    bucket string, key string,
    omClient *omclient.OmClient, clientId *uint64, length int64,
    checksumType *datanode.ChecksumType, blockSize uint64, chunkSize uint64, bytesPerChecksum uint32,
    streamBufferFlushSize uint64, streamBufferMaxSize uint64) (*KeyOutputStream, error) {
    k := &KeyOutputStream{
        omClient:              omClient,
        xceiverManager:        xceiverManager,
        keyInfo:               keyInfo,
        fileName:              path.Base(key),
        volumeName:            volume,
        bucketName:            bucket,
        keyName:               key,
        clientId:              clientId,
        path:                  path.Clean(strings.Join([]string{"", volume, bucket, key}, "/")),
        length:                uint64(length),
        checksumType:          checksumType,
        blockSize:             blockSize,
        chunkSize:             chunkSize,
        bytesPerChecksum:      bytesPerChecksum,
        streamBufferFlushSize: streamBufferFlushSize,
        streamBufferMaxSize:   streamBufferMaxSize,
        keyLocations:          make([]*ozone.KeyLocation, 0),
        excludeList: &hdds.ExcludeListProto{
            Datanodes:    make([]string, 0),
            ContainerIds: make([]int64, 0),
            PipelineIds:  make([]*hdds.PipelineID, 0),
        },
    }
    keyLocationsList := keyInfo.GetKeyLocationList()
    locationsList := keyLocationsList[len(keyLocationsList)-1]
    locations := locationsList.GetKeyLocations()
    k.keyLocation = locations[len(locations)-1]

    if err := k.addBlockLocation(); err != nil {
        return nil, err
    }
    return k, nil
}

func (k *KeyOutputStream) allocateBlockIfNeed() (b *BlockOutputStream, err error) {
    if k.blockIndex < 0 {
        if b, err = k.allocateBlock(); err != nil {
            return nil, err
        }
    } else {
        b = k.blocks[k.blockIndex]
        // block没有空间，申请下一个block
        if !b.HasRemaining() {
            if b, err = k.allocateBlock(); err != nil {
                return nil, err
            }
        }
    }
    return b, nil
}

func (k *KeyOutputStream) allocateBlock() (*BlockOutputStream, error) {
    resp, err := k.omClient.AllocateBlock(k.volumeName, k.bucketName, k.keyName, k.clientId, k.GetExcludeList())
    if err != nil {
        return nil, err
    }
    k.keyLocation = resp.GetKeyLocation()
    err = k.addBlockLocation()
    length := len(k.blocks)
    block := k.blocks[length-1]
    if err != nil {
        k.blocks = k.blocks[:length-1]
        block.Clean()
        return nil, err
    }
    k.blockIndex = length - 1
    block.blockIndex = length - 1
    block.offset = k.offset
    log.Debug("block index", block.blockIndex, "offset", block.offset)
    return block, nil
}

func (k *KeyOutputStream) addBlockLocation() error {
    // blockLocation *hadoop_ozone.BlockLocation
    blockLocation := k.keyLocation
    k.keyLocations = append(k.keyLocations, blockLocation)
    // Param: topology just same to ozone java xceiverClient.
    pipeline, err := dnClient.NewPipelineOperator(blockLocation.GetPipeline())
    if err != nil {
        return err
    }
    xceiverClient, err := k.xceiverManager.GetClient(pipeline, false)
    if err != nil {
        return err
    }
    block := NewBlockOutputStream(
        ConvertBlockIdToDataNodeBlockID(blockLocation.BlockID),
        k.xceiverManager,
        xceiverClient,
        k.checksumType,
        *blockLocation.Length,
        k.bytesPerChecksum,
        k.chunkSize,
        k.streamBufferFlushSize,
        k.streamBufferMaxSize,
        pipeline,
    )

    k.blocks = append(k.blocks, block)
    return nil
}

// Write TODO
func (k *KeyOutputStream) Write(buff []byte) (int, error) {
    if k.closed {
        return 0, nil
    }
    remain := len(buff)
    writeLen := 0
    log.Debug("key buff remain ", remain)
    for remain > 0 {
        block, err := k.allocateBlockIfNeed()
        if err != nil {
            return writeLen, err
        }
        n, err := block.Write(buff)
        if err != nil {
            k.AddDataNodesToExcludeList(block)
            block.Clean()
            return writeLen, err
        }
        block.length += uint64(n)
        writeLen += n
        remain -= n
        k.offset += uint64(n)
        if !block.HasRemaining() {
            log.Debug("block has no remain")
            err := block.Close()
            if err != nil {
                k.AddDataNodesToExcludeList(block)
                block.Clean()
                return n, err
            }
        }
    }
    log.Debug("key write len: ", writeLen, " total write len: ", k.offset, " remain to write: ", k.remainDataLen())
    return writeLen, nil
}

// AddDataNodesToExcludeList TODO
func (k *KeyOutputStream) AddDataNodesToExcludeList(writer *BlockOutputStream) {
    if len(writer.ExcludeListOperator.ToExcludeList().Datanodes) > 0 {
        for _, s := range writer.ExcludeListOperator.ToExcludeList().Datanodes {
            k.AddDNToExcludeList(s)
        }
    }
}

// Close TODO
func (k *KeyOutputStream) Close() error {
    if k.closed {
        return nil
    }
    if err := k.handleFlushOrClose(); err != nil {
        log.Error("file writer handleFlushOrClose err", err)
        return err
    }
    if len(k.blocks) > 0 {
        log.Debug("complete key")
        k.complete()
    }
    k.blocks = nil
    k.closed = true
    return nil
}

func (k *KeyOutputStream) handleFlushOrClose() error {
    for _, writer := range k.blocks {
        if !writer.closed {
            if err := writer.Close(); err != nil {
                log.Error(k.keyName, "handle flush or close err", err)
                if len(writer.ExcludeListOperator.ToExcludeList().Datanodes) > 0 {
                    for _, s := range writer.ExcludeListOperator.ToExcludeList().Datanodes {
                        k.AddDNToExcludeList(s)
                    }
                }
                return err
            } else {
                // k.offset += writer.length
            }
        }
    }
    return nil
}

func (k *KeyOutputStream) complete() {
    for i, writer := range k.blocks {
        if writer.length > 0 {
            k := k.keyLocations[i]
            k.Offset = utils.PointUint64(writer.offset)
            k.Length = utils.PointUint64(writer.length)
        }
    }
    keyLocations := k.keyLocations
    retry := 1
    var err error
    for retry <= 3 {
        _, err = k.omClient.CommitKey(k.volumeName, k.bucketName, k.keyName, k.clientId, keyLocations, uint64(k.length))
        if err != nil {
            log.Warn(fmt.Sprint("complete file error after retry ", retry), err)
            retry += 1
            time.Sleep(100 * time.Millisecond)
        } else {
            return
        }
    }
    log.Error(fmt.Sprint("complete file error after retry ", retry), err)
	panic(err)
}

func (k *KeyOutputStream) remainDataLen() uint64 {
	return k.length - k.offset
}
