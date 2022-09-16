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
    "errors"
    "fmt"
    "io"
    "os"
    dnClient "github.com/apache/ozone-go/api/datanodeclient"
    "github.com/apache/ozone-go/api/proto/hdds"
    ozone_proto "github.com/apache/ozone-go/api/proto/ozone"
    "sync"

    log "github.com/sirupsen/logrus"
)

// KeyInputStream TODO
// get key
type KeyInputStream struct {
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

    mutex sync.Mutex
}

// NewKeyInputStream TODO
func NewKeyInputStream(keyInfo *ozone_proto.KeyInfo, xceiverManager *dnClient.XceiverClientManager,
    verifyChecksum bool) *KeyInputStream {
    return &KeyInputStream{
        KeyInfo:        keyInfo,
        verifyChecksum: verifyChecksum,
        xceiverManager: xceiverManager,
        blocks:         make([]*BlockInputStream, 0),
        blockIndex:     0,
        offset:         0,
        length:         int64(keyInfo.GetDataSize()),
        open:           false,
        closed:         false,
        initialized:    false,
        mutex:          sync.Mutex{},
    }
}

func (k *KeyInputStream) initialize() error {
    k.mutex.Lock()
    defer k.mutex.Unlock()
    if k.initialized {
        return nil
    }
    bis := make([]*BlockInputStream, 0)
    xceiverClientManager := dnClient.GetXceiverClientManagerInstance()
    for _, list := range k.KeyInfo.GetKeyLocationList() {
        for _, location := range list.GetKeyLocations() {
            log.Debug(fmt.Sprintf("key location %v", location.String()))
            pipeline := location.GetPipeline()
            if pipeline.GetType() != hdds.ReplicationType_STAND_ALONE {
                pipeline.Type = hdds.ReplicationType_STAND_ALONE.Enum()
            }
            pipelineOperator, err := dnClient.NewPipelineOperator(pipeline)
            if err != nil {
                return err
            }
            xceiverClient, err := xceiverClientManager.GetClient(pipelineOperator, true)
            if err != nil {
                return err
            }
            bi := NewBlockInputStream(location, xceiverClientManager, xceiverClient, pipelineOperator, k.verifyChecksum)
            bis = append(bis, bi)
        }
    }
    k.blocks = bis
    k.initialized = true
    k.open = true
    k.closed = false
    return nil
}

// Read from key input stream to byte array
func (k *KeyInputStream) Read(buff []byte) (int, error) {
    if k.closed || k.offset >= k.length {
        log.Debug("key closed")
        return 0, io.EOF
    }
    if err := k.initialize(); err != nil {
        log.Debug("key initialised")
        return 0, err
    }

    remain := len(buff)
    log.Debug("key buff remain ", remain)
    readLen := 0
    for remain > 0 {
        block := k.blocks[k.blockIndex]
        n, err := block.Read(buff)
        readLen += n
        remain -= n
        k.offset += int64(n)
        if err != nil && err != io.EOF {
            return readLen, err
        }
        if !block.hasRemainData() {
            block.Close()
            k.blockIndex++
            block.blockIndex = k.blockIndex
            if k.blockIndex >= len(k.blocks) {
                break
            }
        }
    }

    log.Debug("key read len: ", readLen, " total read len ", k.offset, " key remain: ", k.remainDataLen())
    return readLen, nil
}

// ReadAt implements io.ReaderAt.
func (k *KeyInputStream) ReadAt(b []byte, off int64) (int, error) {
    if k.closed {
        return 0, io.EOF
    }

    if off < 0 {
        return 0, &os.PathError{"readat", k.KeyInfo.GetKeyName(), errors.New("negative offset")}
    }

    _, err := k.Seek(off, 0)
    if err != nil {
        return 0, err
    }

    n, err := io.ReadFull(k, b)

    // For some reason, os.File.ReadAt returns io.EOF in this case instead of
    // io.ErrUnexpectedEOF.
    if err == io.ErrUnexpectedEOF {
        err = io.EOF
    }

    return n, err
}

// Seek s to a given position
func (k *KeyInputStream) Seek(offset int64, whence int) (int64, error) {
    if k.closed {
        return 0, io.ErrClosedPipe
    }

    var off int64
    if whence == 0 {
        off = offset
    } else if whence == 1 {
        off = k.offset + offset
    } else if whence == 2 {
        off = k.length + offset
    } else {
        return k.offset, fmt.Errorf("invalid whence: %d", whence)
    }

    if off < 0 || off > k.length {
        return k.offset, fmt.Errorf("invalid resulting offset: %d", off)
    }

    if k.offset != off {
        k.offset = off
    }

    totalLen := uint64(0)
    for idx, block := range k.blocks {
        totalLen += block.Len()
        if uint64(k.offset) < totalLen {
            k.blockIndex = idx
            k.blocks[k.blockIndex].offset = k.offset - int64(totalLen-block.Len())
            if _, err := k.blocks[k.blockIndex].Seek(0, 1); err != nil {
                return 0, err
            }
            break
        }
    }

    return k.offset, nil
}

// Position TODO
// Returns current position
func (k *KeyInputStream) Position() (int64, error) {
    actualPos, err := k.Seek(0, 1)
    if err != nil {
        return 0, err
    }
    return actualPos, nil
}

// Close TODO
func (k *KeyInputStream) Close() error {
    k.Reset()
    k.closed = true
    return nil
}

// Reset TODO
func (k *KeyInputStream) Reset() {
    for _, r := range k.blocks {
        r.Reset()
    }
    k.blockIndex = -1
}

// Len TODO
func (k *KeyInputStream) Len() int64 {
    return k.length
}

// Readdir reads the contents of the directory associated with file and returns
// a slice of up to n os.FileInfo values, as would be returned by Stat, in
// directory order. Subsequent calls on the same file will yield further
// os.FileInfos.
//
// If n > 0, Readdir returns at most n os.FileInfo values. In this case, if
// Readdir returns an empty slice, it will return a non-nil error explaining
// why. At the end of a directory, the error is io.EOF.
//
// If n <= 0, Readdir returns all the os.FileInfo from the directory in a single
// slice. In this case, if Readdir succeeds (reads all the way to the end of
// the directory), it returns the slice and a nil error. If it encounters an
// error before the end of the directory, Readdir returns the os.FileInfo read
// until that point and a non-nil error.
//
// The os.FileInfo values returned will not have block location attached to
// the struct returned by Sys(). To fetch that information, make a separate
// call to Stat.
//
// Note that making multiple calls to Readdir with a smallish n (as you might do
// with the os version) is slower than just requesting everything at once.
// That's because HDFS has no mechanism for limiting the number of entries
// returned; whatever extra entries it returns are simply thrown away.
func (k *KeyInputStream) Readdir(n int) ([]os.FileInfo, error) {
    panic("")
}

func (k *KeyInputStream) readdir(n int32) ([]os.FileInfo, error) {
    panic("")
}

// Readdirnames reads and returns a slice of names from the directory f.
//
// If n > 0, Readdirnames returns at most n names. In this case, if Readdirnames
// returns an empty slice, it will return a non-nil error explaining why. At the
// end of a directory, the error is io.EOF.
//
// If n <= 0, Readdirnames returns all the names from the directory in a single
// slice. In this case, if Readdirnames succeeds (reads all the way to the end
// of the directory), it returns the slice and a nil error. If it encounters an
// error before the end of the directory, Readdirnames returns the names read
// until that point and a non-nil error.
func (k *KeyInputStream) Readdirnames(n int) ([]string, error) {
    if k.closed {
        return nil, io.ErrClosedPipe
    }

    fis, err := k.Readdir(n)
    if err != nil {
        return nil, err
    }

    names := make([]string, 0, len(fis))
    for _, fi := range fis {
        names = append(names, fi.Name())
    }

    return names, nil
}

func (k *KeyInputStream) remainDataLen() int64 {
	return k.length - k.offset
}
