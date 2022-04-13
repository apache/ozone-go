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
    "os"
    dnClient "github.com/apache/ozone-go/api/datanodeclient"
    ozone_proto "github.com/apache/ozone-go/api/proto/ozone"
    "reflect"
    "sync"
    "testing"
)

func TestKeyInputStream_Close(t *testing.T) {
    type fields struct {
        KeyInfo        *ozone_proto.KeyInfo
        xceiverManager *dnClient.XceiverClientManager
        verifyChecksum bool
        blocks         []*BlockInputStream
        blockIndex     int
        offset         int64
        length         int64
        open           bool
        closed         bool
        initialized    bool
        mutex          sync.Mutex
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
            k := &KeyInputStream{
                KeyInfo:        tt.fields.KeyInfo,
                xceiverManager: tt.fields.xceiverManager,
                verifyChecksum: tt.fields.verifyChecksum,
                blocks:         tt.fields.blocks,
                blockIndex:     tt.fields.blockIndex,
                offset:         tt.fields.offset,
                length:         tt.fields.length,
                open:           tt.fields.open,
                closed:         tt.fields.closed,
                initialized:    tt.fields.initialized,
                mutex:          tt.fields.mutex,
            }
            if err := k.Close(); (err != nil) != tt.wantErr {
                t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestKeyInputStream_Len(t *testing.T) {
    type fields struct {
        KeyInfo        *ozone_proto.KeyInfo
        xceiverManager *dnClient.XceiverClientManager
        verifyChecksum bool
        blocks         []*BlockInputStream
        blockIndex     int
        offset         int64
        length         int64
        open           bool
        closed         bool
        initialized    bool
        mutex          sync.Mutex
    }
    tests := []struct {
        name   string
        fields fields
        want   int64
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            k := &KeyInputStream{
                KeyInfo:        tt.fields.KeyInfo,
                xceiverManager: tt.fields.xceiverManager,
                verifyChecksum: tt.fields.verifyChecksum,
                blocks:         tt.fields.blocks,
                blockIndex:     tt.fields.blockIndex,
                offset:         tt.fields.offset,
                length:         tt.fields.length,
                open:           tt.fields.open,
                closed:         tt.fields.closed,
                initialized:    tt.fields.initialized,
                mutex:          tt.fields.mutex,
            }
            if got := k.Len(); got != tt.want {
                t.Errorf("Len() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyInputStream_Position(t *testing.T) {
    type fields struct {
        KeyInfo        *ozone_proto.KeyInfo
        xceiverManager *dnClient.XceiverClientManager
        verifyChecksum bool
        blocks         []*BlockInputStream
        blockIndex     int
        offset         int64
        length         int64
        open           bool
        closed         bool
        initialized    bool
        mutex          sync.Mutex
    }
    tests := []struct {
        name    string
        fields  fields
        want    int64
        wantErr bool
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            k := &KeyInputStream{
                KeyInfo:        tt.fields.KeyInfo,
                xceiverManager: tt.fields.xceiverManager,
                verifyChecksum: tt.fields.verifyChecksum,
                blocks:         tt.fields.blocks,
                blockIndex:     tt.fields.blockIndex,
                offset:         tt.fields.offset,
                length:         tt.fields.length,
                open:           tt.fields.open,
                closed:         tt.fields.closed,
                initialized:    tt.fields.initialized,
                mutex:          tt.fields.mutex,
            }
            got, err := k.Position()
            if (err != nil) != tt.wantErr {
                t.Errorf("Position() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("Position() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyInputStream_Read(t *testing.T) {
    type fields struct {
        KeyInfo        *ozone_proto.KeyInfo
        xceiverManager *dnClient.XceiverClientManager
        verifyChecksum bool
        blocks         []*BlockInputStream
        blockIndex     int
        offset         int64
        length         int64
        open           bool
        closed         bool
        initialized    bool
        mutex          sync.Mutex
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
            k := &KeyInputStream{
                KeyInfo:        tt.fields.KeyInfo,
                xceiverManager: tt.fields.xceiverManager,
                verifyChecksum: tt.fields.verifyChecksum,
                blocks:         tt.fields.blocks,
                blockIndex:     tt.fields.blockIndex,
                offset:         tt.fields.offset,
                length:         tt.fields.length,
                open:           tt.fields.open,
                closed:         tt.fields.closed,
                initialized:    tt.fields.initialized,
                mutex:          tt.fields.mutex,
            }
            got, err := k.Read(tt.args.buff)
            if (err != nil) != tt.wantErr {
                t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("Read() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyInputStream_ReadAt(t *testing.T) {
    type fields struct {
        KeyInfo        *ozone_proto.KeyInfo
        xceiverManager *dnClient.XceiverClientManager
        verifyChecksum bool
        blocks         []*BlockInputStream
        blockIndex     int
        offset         int64
        length         int64
        open           bool
        closed         bool
        initialized    bool
        mutex          sync.Mutex
    }
    type args struct {
        b   []byte
        off int64
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
            k := &KeyInputStream{
                KeyInfo:        tt.fields.KeyInfo,
                xceiverManager: tt.fields.xceiverManager,
                verifyChecksum: tt.fields.verifyChecksum,
                blocks:         tt.fields.blocks,
                blockIndex:     tt.fields.blockIndex,
                offset:         tt.fields.offset,
                length:         tt.fields.length,
                open:           tt.fields.open,
                closed:         tt.fields.closed,
                initialized:    tt.fields.initialized,
                mutex:          tt.fields.mutex,
            }
            got, err := k.ReadAt(tt.args.b, tt.args.off)
            if (err != nil) != tt.wantErr {
                t.Errorf("ReadAt() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("ReadAt() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyInputStream_Readdir(t *testing.T) {
    type fields struct {
        KeyInfo        *ozone_proto.KeyInfo
        xceiverManager *dnClient.XceiverClientManager
        verifyChecksum bool
        blocks         []*BlockInputStream
        blockIndex     int
        offset         int64
        length         int64
        open           bool
        closed         bool
        initialized    bool
        mutex          sync.Mutex
    }
    type args struct {
        n int
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    []os.FileInfo
        wantErr bool
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            k := &KeyInputStream{
                KeyInfo:        tt.fields.KeyInfo,
                xceiverManager: tt.fields.xceiverManager,
                verifyChecksum: tt.fields.verifyChecksum,
                blocks:         tt.fields.blocks,
                blockIndex:     tt.fields.blockIndex,
                offset:         tt.fields.offset,
                length:         tt.fields.length,
                open:           tt.fields.open,
                closed:         tt.fields.closed,
                initialized:    tt.fields.initialized,
                mutex:          tt.fields.mutex,
            }
            got, err := k.Readdir(tt.args.n)
            if (err != nil) != tt.wantErr {
                t.Errorf("Readdir() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Readdir() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyInputStream_Readdirnames(t *testing.T) {
    type fields struct {
        KeyInfo        *ozone_proto.KeyInfo
        xceiverManager *dnClient.XceiverClientManager
        verifyChecksum bool
        blocks         []*BlockInputStream
        blockIndex     int
        offset         int64
        length         int64
        open           bool
        closed         bool
        initialized    bool
        mutex          sync.Mutex
    }
    type args struct {
        n int
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    []string
        wantErr bool
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            k := &KeyInputStream{
                KeyInfo:        tt.fields.KeyInfo,
                xceiverManager: tt.fields.xceiverManager,
                verifyChecksum: tt.fields.verifyChecksum,
                blocks:         tt.fields.blocks,
                blockIndex:     tt.fields.blockIndex,
                offset:         tt.fields.offset,
                length:         tt.fields.length,
                open:           tt.fields.open,
                closed:         tt.fields.closed,
                initialized:    tt.fields.initialized,
                mutex:          tt.fields.mutex,
            }
            got, err := k.Readdirnames(tt.args.n)
            if (err != nil) != tt.wantErr {
                t.Errorf("Readdirnames() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Readdirnames() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyInputStream_Reset(t *testing.T) {
    type fields struct {
        KeyInfo        *ozone_proto.KeyInfo
        xceiverManager *dnClient.XceiverClientManager
        verifyChecksum bool
        blocks         []*BlockInputStream
        blockIndex     int
        offset         int64
        length         int64
        open           bool
        closed         bool
        initialized    bool
        mutex          sync.Mutex
    }
    tests := []struct {
        name   string
        fields fields
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            k := &KeyInputStream{
                KeyInfo:        tt.fields.KeyInfo,
                xceiverManager: tt.fields.xceiverManager,
                verifyChecksum: tt.fields.verifyChecksum,
                blocks:         tt.fields.blocks,
                blockIndex:     tt.fields.blockIndex,
                offset:         tt.fields.offset,
                length:         tt.fields.length,
                open:           tt.fields.open,
                closed:         tt.fields.closed,
                initialized:    tt.fields.initialized,
                mutex:          tt.fields.mutex,
            }
            k.Reset()
        })
    }
}

func TestKeyInputStream_Seek(t *testing.T) {
    type fields struct {
        KeyInfo        *ozone_proto.KeyInfo
        xceiverManager *dnClient.XceiverClientManager
        verifyChecksum bool
        blocks         []*BlockInputStream
        blockIndex     int
        offset         int64
        length         int64
        open           bool
        closed         bool
        initialized    bool
        mutex          sync.Mutex
    }
    type args struct {
        offset int64
        whence int
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    int64
        wantErr bool
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            k := &KeyInputStream{
                KeyInfo:        tt.fields.KeyInfo,
                xceiverManager: tt.fields.xceiverManager,
                verifyChecksum: tt.fields.verifyChecksum,
                blocks:         tt.fields.blocks,
                blockIndex:     tt.fields.blockIndex,
                offset:         tt.fields.offset,
                length:         tt.fields.length,
                open:           tt.fields.open,
                closed:         tt.fields.closed,
                initialized:    tt.fields.initialized,
                mutex:          tt.fields.mutex,
            }
            got, err := k.Seek(tt.args.offset, tt.args.whence)
            if (err != nil) != tt.wantErr {
                t.Errorf("Seek() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("Seek() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyInputStream_initialize(t *testing.T) {
    type fields struct {
        KeyInfo        *ozone_proto.KeyInfo
        xceiverManager *dnClient.XceiverClientManager
        verifyChecksum bool
        blocks         []*BlockInputStream
        blockIndex     int
        offset         int64
        length         int64
        open           bool
        closed         bool
        initialized    bool
        mutex          sync.Mutex
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
            k := &KeyInputStream{
                KeyInfo:        tt.fields.KeyInfo,
                xceiverManager: tt.fields.xceiverManager,
                verifyChecksum: tt.fields.verifyChecksum,
                blocks:         tt.fields.blocks,
                blockIndex:     tt.fields.blockIndex,
                offset:         tt.fields.offset,
                length:         tt.fields.length,
                open:           tt.fields.open,
                closed:         tt.fields.closed,
                initialized:    tt.fields.initialized,
                mutex:          tt.fields.mutex,
            }
            if err := k.initialize(); (err != nil) != tt.wantErr {
                t.Errorf("initialize() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestKeyInputStream_readdir(t *testing.T) {
    type fields struct {
        KeyInfo        *ozone_proto.KeyInfo
        xceiverManager *dnClient.XceiverClientManager
        verifyChecksum bool
        blocks         []*BlockInputStream
        blockIndex     int
        offset         int64
        length         int64
        open           bool
        closed         bool
        initialized    bool
        mutex          sync.Mutex
    }
    type args struct {
        n int32
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    []os.FileInfo
        wantErr bool
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            k := &KeyInputStream{
                KeyInfo:        tt.fields.KeyInfo,
                xceiverManager: tt.fields.xceiverManager,
                verifyChecksum: tt.fields.verifyChecksum,
                blocks:         tt.fields.blocks,
                blockIndex:     tt.fields.blockIndex,
                offset:         tt.fields.offset,
                length:         tt.fields.length,
                open:           tt.fields.open,
                closed:         tt.fields.closed,
                initialized:    tt.fields.initialized,
                mutex:          tt.fields.mutex,
            }
            got, err := k.readdir(tt.args.n)
            if (err != nil) != tt.wantErr {
                t.Errorf("readdir() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("readdir() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyInputStream_remainDataLen(t *testing.T) {
    type fields struct {
        KeyInfo        *ozone_proto.KeyInfo
        xceiverManager *dnClient.XceiverClientManager
        verifyChecksum bool
        blocks         []*BlockInputStream
        blockIndex     int
        offset         int64
        length         int64
        open           bool
        closed         bool
        initialized    bool
        mutex          sync.Mutex
    }
    tests := []struct {
        name   string
        fields fields
        want   int64
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            k := &KeyInputStream{
                KeyInfo:        tt.fields.KeyInfo,
                xceiverManager: tt.fields.xceiverManager,
                verifyChecksum: tt.fields.verifyChecksum,
                blocks:         tt.fields.blocks,
                blockIndex:     tt.fields.blockIndex,
                offset:         tt.fields.offset,
                length:         tt.fields.length,
                open:           tt.fields.open,
                closed:         tt.fields.closed,
                initialized:    tt.fields.initialized,
                mutex:          tt.fields.mutex,
            }
            if got := k.remainDataLen(); got != tt.want {
                t.Errorf("remainDataLen() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestNewKeyInputStream(t *testing.T) {
    type args struct {
        keyInfo        *ozone_proto.KeyInfo
        xceiverManager *dnClient.XceiverClientManager
        verifyChecksum bool
    }
    tests := []struct {
        name string
        args args
        want *KeyInputStream
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewKeyInputStream(tt.args.keyInfo, tt.args.xceiverManager, tt.args.verifyChecksum); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewKeyInputStream() = %v, want %v", got, tt.want)
			}
		})
	}
}
