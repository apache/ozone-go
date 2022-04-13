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
    dnClient "github.com/apache/ozone-go/api/datanodeclient"
    "github.com/apache/ozone-go/api/omclient"
    "github.com/apache/ozone-go/api/proto/datanode"
    "github.com/apache/ozone-go/api/proto/hdds"
    "github.com/apache/ozone-go/api/proto/ozone"
    "reflect"
    "sync"
    "testing"
)

func TestKeyOutputStream_AddDataNodesToExcludeList(t *testing.T) {
    type fields struct {
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
        streamBufferFlushSize uint64
        streamBufferMaxSize   uint64
        currentWritingBytes   uint64
        excludeList           *hdds.ExcludeListProto
        mux                   sync.Mutex
    }
    type args struct {
        writer *BlockOutputStream
    }
    tests := []struct {
        name   string
        fields fields
        args   args
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            k := &KeyOutputStream{
                keyInfo:               tt.fields.keyInfo,
                omClient:              tt.fields.omClient,
                xceiverManager:        tt.fields.xceiverManager,
                clientId:              tt.fields.clientId,
                keyLocation:           tt.fields.keyLocation,
                keyLocations:          tt.fields.keyLocations,
                checksumType:          tt.fields.checksumType,
                blocks:                tt.fields.blocks,
                fileName:              tt.fields.fileName,
                volumeName:            tt.fields.volumeName,
                bucketName:            tt.fields.bucketName,
                keyName:               tt.fields.keyName,
                path:                  tt.fields.path,
                length:                tt.fields.length,
                offset:                tt.fields.offset,
                blockSize:             tt.fields.blockSize,
                chunkSize:             tt.fields.chunkSize,
                bytesPerChecksum:      tt.fields.bytesPerChecksum,
                blockIndex:            tt.fields.blockIndex,
                closed:                tt.fields.closed,
                streamBufferFlushSize: tt.fields.streamBufferFlushSize,
                streamBufferMaxSize:   tt.fields.streamBufferMaxSize,
                currentWritingBytes:   tt.fields.currentWritingBytes,
                excludeList:           tt.fields.excludeList,
                mux:                   tt.fields.mux,
            }
            k.AddDataNodesToExcludeList(tt.args.writer)
        })
    }
}

func TestKeyOutputStream_Close(t *testing.T) {
    type fields struct {
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
        streamBufferFlushSize uint64
        streamBufferMaxSize   uint64
        currentWritingBytes   uint64
        excludeList           *hdds.ExcludeListProto
        mux                   sync.Mutex
    }
    tests := []struct {
        name    string
        fields  fields
        wantErr bool
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            k := &KeyOutputStream{
                keyInfo:               tt.fields.keyInfo,
                omClient:              tt.fields.omClient,
                xceiverManager:        tt.fields.xceiverManager,
                clientId:              tt.fields.clientId,
                keyLocation:           tt.fields.keyLocation,
                keyLocations:          tt.fields.keyLocations,
                checksumType:          tt.fields.checksumType,
                blocks:                tt.fields.blocks,
                fileName:              tt.fields.fileName,
                volumeName:            tt.fields.volumeName,
                bucketName:            tt.fields.bucketName,
                keyName:               tt.fields.keyName,
                path:                  tt.fields.path,
                length:                tt.fields.length,
                offset:                tt.fields.offset,
                blockSize:             tt.fields.blockSize,
                chunkSize:             tt.fields.chunkSize,
                bytesPerChecksum:      tt.fields.bytesPerChecksum,
                blockIndex:            tt.fields.blockIndex,
                closed:                tt.fields.closed,
                streamBufferFlushSize: tt.fields.streamBufferFlushSize,
                streamBufferMaxSize:   tt.fields.streamBufferMaxSize,
                currentWritingBytes:   tt.fields.currentWritingBytes,
                excludeList:           tt.fields.excludeList,
                mux:                   tt.fields.mux,
            }
            if err := k.Close(); (err != nil) != tt.wantErr {
                t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestKeyOutputStream_Write(t *testing.T) {
    type fields struct {
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
        streamBufferFlushSize uint64
        streamBufferMaxSize   uint64
        currentWritingBytes   uint64
        excludeList           *hdds.ExcludeListProto
        mux                   sync.Mutex
    }
    type args struct {
        buff []byte
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    int
        wantErr bool
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            k := &KeyOutputStream{
                keyInfo:               tt.fields.keyInfo,
                omClient:              tt.fields.omClient,
                xceiverManager:        tt.fields.xceiverManager,
                clientId:              tt.fields.clientId,
                keyLocation:           tt.fields.keyLocation,
                keyLocations:          tt.fields.keyLocations,
                checksumType:          tt.fields.checksumType,
                blocks:                tt.fields.blocks,
                fileName:              tt.fields.fileName,
                volumeName:            tt.fields.volumeName,
                bucketName:            tt.fields.bucketName,
                keyName:               tt.fields.keyName,
                path:                  tt.fields.path,
                length:                tt.fields.length,
                offset:                tt.fields.offset,
                blockSize:             tt.fields.blockSize,
                chunkSize:             tt.fields.chunkSize,
                bytesPerChecksum:      tt.fields.bytesPerChecksum,
                blockIndex:            tt.fields.blockIndex,
                closed:                tt.fields.closed,
                streamBufferFlushSize: tt.fields.streamBufferFlushSize,
                streamBufferMaxSize:   tt.fields.streamBufferMaxSize,
                currentWritingBytes:   tt.fields.currentWritingBytes,
                excludeList:           tt.fields.excludeList,
                mux:                   tt.fields.mux,
            }
            got, err := k.Write(tt.args.buff)
            if (err != nil) != tt.wantErr {
                t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("Write() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyOutputStream_addBlockLocation(t *testing.T) {
    type fields struct {
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
        streamBufferFlushSize uint64
        streamBufferMaxSize   uint64
        currentWritingBytes   uint64
        excludeList           *hdds.ExcludeListProto
        mux                   sync.Mutex
    }
    tests := []struct {
        name    string
        fields  fields
        wantErr bool
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            k := &KeyOutputStream{
                keyInfo:               tt.fields.keyInfo,
                omClient:              tt.fields.omClient,
                xceiverManager:        tt.fields.xceiverManager,
                clientId:              tt.fields.clientId,
                keyLocation:           tt.fields.keyLocation,
                keyLocations:          tt.fields.keyLocations,
                checksumType:          tt.fields.checksumType,
                blocks:                tt.fields.blocks,
                fileName:              tt.fields.fileName,
                volumeName:            tt.fields.volumeName,
                bucketName:            tt.fields.bucketName,
                keyName:               tt.fields.keyName,
                path:                  tt.fields.path,
                length:                tt.fields.length,
                offset:                tt.fields.offset,
                blockSize:             tt.fields.blockSize,
                chunkSize:             tt.fields.chunkSize,
                bytesPerChecksum:      tt.fields.bytesPerChecksum,
                blockIndex:            tt.fields.blockIndex,
                closed:                tt.fields.closed,
                streamBufferFlushSize: tt.fields.streamBufferFlushSize,
                streamBufferMaxSize:   tt.fields.streamBufferMaxSize,
                currentWritingBytes:   tt.fields.currentWritingBytes,
                excludeList:           tt.fields.excludeList,
                mux:                   tt.fields.mux,
            }
            if err := k.addBlockLocation(); (err != nil) != tt.wantErr {
                t.Errorf("addBlockLocation() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestKeyOutputStream_allocateBlock(t *testing.T) {
    type fields struct {
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
        streamBufferFlushSize uint64
        streamBufferMaxSize   uint64
        currentWritingBytes   uint64
        excludeList           *hdds.ExcludeListProto
        mux                   sync.Mutex
    }
    tests := []struct {
        name    string
        fields  fields
        want    *BlockOutputStream
        wantErr bool
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            k := &KeyOutputStream{
                keyInfo:               tt.fields.keyInfo,
                omClient:              tt.fields.omClient,
                xceiverManager:        tt.fields.xceiverManager,
                clientId:              tt.fields.clientId,
                keyLocation:           tt.fields.keyLocation,
                keyLocations:          tt.fields.keyLocations,
                checksumType:          tt.fields.checksumType,
                blocks:                tt.fields.blocks,
                fileName:              tt.fields.fileName,
                volumeName:            tt.fields.volumeName,
                bucketName:            tt.fields.bucketName,
                keyName:               tt.fields.keyName,
                path:                  tt.fields.path,
                length:                tt.fields.length,
                offset:                tt.fields.offset,
                blockSize:             tt.fields.blockSize,
                chunkSize:             tt.fields.chunkSize,
                bytesPerChecksum:      tt.fields.bytesPerChecksum,
                blockIndex:            tt.fields.blockIndex,
                closed:                tt.fields.closed,
                streamBufferFlushSize: tt.fields.streamBufferFlushSize,
                streamBufferMaxSize:   tt.fields.streamBufferMaxSize,
                currentWritingBytes:   tt.fields.currentWritingBytes,
                excludeList:           tt.fields.excludeList,
                mux:                   tt.fields.mux,
            }
            got, err := k.allocateBlock()
            if (err != nil) != tt.wantErr {
                t.Errorf("allocateBlock() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("allocateBlock() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyOutputStream_allocateBlockIfNeed(t *testing.T) {
    type fields struct {
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
        streamBufferFlushSize uint64
        streamBufferMaxSize   uint64
        currentWritingBytes   uint64
        excludeList           *hdds.ExcludeListProto
        mux                   sync.Mutex
    }
    tests := []struct {
        name    string
        fields  fields
        wantB   *BlockOutputStream
        wantErr bool
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            k := &KeyOutputStream{
                keyInfo:               tt.fields.keyInfo,
                omClient:              tt.fields.omClient,
                xceiverManager:        tt.fields.xceiverManager,
                clientId:              tt.fields.clientId,
                keyLocation:           tt.fields.keyLocation,
                keyLocations:          tt.fields.keyLocations,
                checksumType:          tt.fields.checksumType,
                blocks:                tt.fields.blocks,
                fileName:              tt.fields.fileName,
                volumeName:            tt.fields.volumeName,
                bucketName:            tt.fields.bucketName,
                keyName:               tt.fields.keyName,
                path:                  tt.fields.path,
                length:                tt.fields.length,
                offset:                tt.fields.offset,
                blockSize:             tt.fields.blockSize,
                chunkSize:             tt.fields.chunkSize,
                bytesPerChecksum:      tt.fields.bytesPerChecksum,
                blockIndex:            tt.fields.blockIndex,
                closed:                tt.fields.closed,
                streamBufferFlushSize: tt.fields.streamBufferFlushSize,
                streamBufferMaxSize:   tt.fields.streamBufferMaxSize,
                currentWritingBytes:   tt.fields.currentWritingBytes,
                excludeList:           tt.fields.excludeList,
                mux:                   tt.fields.mux,
            }
            gotB, err := k.allocateBlockIfNeed()
            if (err != nil) != tt.wantErr {
                t.Errorf("allocateBlockIfNeed() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(gotB, tt.wantB) {
                t.Errorf("allocateBlockIfNeed() gotB = %v, want %v", gotB, tt.wantB)
            }
        })
    }
}

func TestKeyOutputStream_complete(t *testing.T) {
    type fields struct {
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
        streamBufferFlushSize uint64
        streamBufferMaxSize   uint64
        currentWritingBytes   uint64
        excludeList           *hdds.ExcludeListProto
        mux                   sync.Mutex
    }
    tests := []struct {
        name   string
        fields fields
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            k := &KeyOutputStream{
                keyInfo:               tt.fields.keyInfo,
                omClient:              tt.fields.omClient,
                xceiverManager:        tt.fields.xceiverManager,
                clientId:              tt.fields.clientId,
                keyLocation:           tt.fields.keyLocation,
                keyLocations:          tt.fields.keyLocations,
                checksumType:          tt.fields.checksumType,
                blocks:                tt.fields.blocks,
                fileName:              tt.fields.fileName,
                volumeName:            tt.fields.volumeName,
                bucketName:            tt.fields.bucketName,
                keyName:               tt.fields.keyName,
                path:                  tt.fields.path,
                length:                tt.fields.length,
                offset:                tt.fields.offset,
                blockSize:             tt.fields.blockSize,
                chunkSize:             tt.fields.chunkSize,
                bytesPerChecksum:      tt.fields.bytesPerChecksum,
                blockIndex:            tt.fields.blockIndex,
                closed:                tt.fields.closed,
                streamBufferFlushSize: tt.fields.streamBufferFlushSize,
                streamBufferMaxSize:   tt.fields.streamBufferMaxSize,
                currentWritingBytes:   tt.fields.currentWritingBytes,
                excludeList:           tt.fields.excludeList,
                mux:                   tt.fields.mux,
            }
            k.complete()
        })
    }
}

func TestKeyOutputStream_handleFlushOrClose(t *testing.T) {
    type fields struct {
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
        streamBufferFlushSize uint64
        streamBufferMaxSize   uint64
        currentWritingBytes   uint64
        excludeList           *hdds.ExcludeListProto
        mux                   sync.Mutex
    }
    tests := []struct {
        name    string
        fields  fields
        wantErr bool
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            k := &KeyOutputStream{
                keyInfo:               tt.fields.keyInfo,
                omClient:              tt.fields.omClient,
                xceiverManager:        tt.fields.xceiverManager,
                clientId:              tt.fields.clientId,
                keyLocation:           tt.fields.keyLocation,
                keyLocations:          tt.fields.keyLocations,
                checksumType:          tt.fields.checksumType,
                blocks:                tt.fields.blocks,
                fileName:              tt.fields.fileName,
                volumeName:            tt.fields.volumeName,
                bucketName:            tt.fields.bucketName,
                keyName:               tt.fields.keyName,
                path:                  tt.fields.path,
                length:                tt.fields.length,
                offset:                tt.fields.offset,
                blockSize:             tt.fields.blockSize,
                chunkSize:             tt.fields.chunkSize,
                bytesPerChecksum:      tt.fields.bytesPerChecksum,
                blockIndex:            tt.fields.blockIndex,
                closed:                tt.fields.closed,
                streamBufferFlushSize: tt.fields.streamBufferFlushSize,
                streamBufferMaxSize:   tt.fields.streamBufferMaxSize,
                currentWritingBytes:   tt.fields.currentWritingBytes,
                excludeList:           tt.fields.excludeList,
                mux:                   tt.fields.mux,
            }
            if err := k.handleFlushOrClose(); (err != nil) != tt.wantErr {
                t.Errorf("handleFlushOrClose() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestKeyOutputStream_remainDataLen(t *testing.T) {
    type fields struct {
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
        streamBufferFlushSize uint64
        streamBufferMaxSize   uint64
        currentWritingBytes   uint64
        excludeList           *hdds.ExcludeListProto
        mux                   sync.Mutex
    }
    tests := []struct {
        name   string
        fields fields
        want   uint64
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            k := &KeyOutputStream{
                keyInfo:               tt.fields.keyInfo,
                omClient:              tt.fields.omClient,
                xceiverManager:        tt.fields.xceiverManager,
                clientId:              tt.fields.clientId,
                keyLocation:           tt.fields.keyLocation,
                keyLocations:          tt.fields.keyLocations,
                checksumType:          tt.fields.checksumType,
                blocks:                tt.fields.blocks,
                fileName:              tt.fields.fileName,
                volumeName:            tt.fields.volumeName,
                bucketName:            tt.fields.bucketName,
                keyName:               tt.fields.keyName,
                path:                  tt.fields.path,
                length:                tt.fields.length,
                offset:                tt.fields.offset,
                blockSize:             tt.fields.blockSize,
                chunkSize:             tt.fields.chunkSize,
                bytesPerChecksum:      tt.fields.bytesPerChecksum,
                blockIndex:            tt.fields.blockIndex,
                closed:                tt.fields.closed,
                streamBufferFlushSize: tt.fields.streamBufferFlushSize,
                streamBufferMaxSize:   tt.fields.streamBufferMaxSize,
                currentWritingBytes:   tt.fields.currentWritingBytes,
                excludeList:           tt.fields.excludeList,
                mux:                   tt.fields.mux,
            }
            if got := k.remainDataLen(); got != tt.want {
                t.Errorf("remainDataLen() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestNewKeyOutputStream(t *testing.T) {
    type args struct {
        keyInfo               *ozone.KeyInfo
        xceiverManager        *dnClient.XceiverClientManager
        volume                string
        bucket                string
        key                   string
        omClient              *omclient.OmClient
        clientId              *uint64
        length                int64
        checksumType          *datanode.ChecksumType
        blockSize             uint64
        chunkSize             uint64
        bytesPerChecksum      uint32
        streamBufferFlushSize uint64
        streamBufferMaxSize   uint64
    }
    tests := []struct {
        name    string
        args    args
        want    *KeyOutputStream
        wantErr bool
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got, err := NewKeyOutputStream(tt.args.keyInfo, tt.args.xceiverManager, tt.args.volume, tt.args.bucket, tt.args.key, tt.args.omClient, tt.args.clientId, tt.args.length, tt.args.checksumType, tt.args.blockSize, tt.args.chunkSize, tt.args.bytesPerChecksum, tt.args.streamBufferFlushSize, tt.args.streamBufferMaxSize); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewKeyOutputStream() = %v, want %v", got, tt.want)
            } else {
                if (err != nil) != tt.wantErr {
                    t.Errorf("NewKeyOutputStream() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
