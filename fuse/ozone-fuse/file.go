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

package main

import (
	"github.com/apache/ozone-go/api"
	"github.com/apache/ozone-go/api/datanode"
	ozone_proto "github.com/apache/ozone-go/api/proto/ozone"
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"time"
)

type OzoneFile struct {
	ozoneClient *api.OzoneClient
	key         *ozone_proto.KeyInfo
}

func CreateOzoneFile(ozoneClient *api.OzoneClient, key *ozone_proto.KeyInfo) nodefs.File {
	return &OzoneFile{
		ozoneClient: ozoneClient,
		key:         key,
	}
}

func (f *OzoneFile) SetInode(*nodefs.Inode) {
}

func (f *OzoneFile) InnerFile() nodefs.File {
	return nil
}

func (f *OzoneFile) String() string {
	return "OzoneFile"
}

func (f *OzoneFile) Read(buf []byte, off int64) (fuse.ReadResult, fuse.Status) {
	len := uint64(len(buf))

	currentOffset := int64(0)
	for _, location := range f.key.KeyLocationList[0].KeyLocations {
		dataLengthInBlock := int64(*location.Length)
		if currentOffset+dataLengthInBlock < off {
			//we don't need this block
			currentOffset += dataLengthInBlock
		} else {
			blockOffset := uint64(off - currentOffset)

			pipeline := location.Pipeline

			dnBlockId := api.ConvertBlockId(location.BlockID)
			dnClient, err := datanode.CreateDatanodeClient(pipeline)
			defer dnClient.Close()
			chunks, err := dnClient.GetBlock(dnBlockId)
			if err != nil {
				return fuse.ReadResultData([]byte{}), fuse.ENODATA
			}

			currentBlockOffset := uint64(0)
			for _, chunk := range chunks {
				if currentBlockOffset >= blockOffset {
					data, err := dnClient.ReadChunk(dnBlockId, chunk)
					if err != nil {
						return fuse.ReadResultData([]byte{}), fuse.ENODATA
					}
					if len > chunk.Len {
						len = chunk.Len
					}
					return fuse.ReadResultData(data[0:len]), 0
				}
				currentBlockOffset += chunk.Len
			}
		}
	}
	return fuse.ReadResultData([]byte{}), 0
}

func (f *OzoneFile) Write(data []byte, off int64) (uint32, fuse.Status) {
	return 0, fuse.ENOSYS
}

func (f *OzoneFile) GetLk(owner uint64, lk *fuse.FileLock, flags uint32, out *fuse.FileLock) (code fuse.Status) {
	return fuse.ENOSYS
}

func (f *OzoneFile) SetLk(owner uint64, lk *fuse.FileLock, flags uint32) (code fuse.Status) {
	return fuse.ENOSYS
}

func (f *OzoneFile) SetLkw(owner uint64, lk *fuse.FileLock, flags uint32) (code fuse.Status) {
	return fuse.ENOSYS
}

func (f *OzoneFile) Flush() fuse.Status {
	return fuse.OK
}

func (f *OzoneFile) Release() {

}

func (f *OzoneFile) GetAttr(attr *fuse.Attr) fuse.Status {
	attr.Size = *f.key.DataSize
	attr.Mode = fuse.S_IFREG | 0644
	attr.Blksize = 512
	attr.Blocks = 1
	attr.IsDir()
	return fuse.OK
}

func (f *OzoneFile) Fsync(flags int) (code fuse.Status) {
	return fuse.OK
}

func (f *OzoneFile) Utimens(atime *time.Time, mtime *time.Time) fuse.Status {
	return fuse.ENOSYS
}

func (f *OzoneFile) Truncate(size uint64) fuse.Status {
	return fuse.ENOSYS
}

func (f *OzoneFile) Chown(uid uint32, gid uint32) fuse.Status {
	return fuse.OK
}

func (f *OzoneFile) Chmod(perms uint32) fuse.Status {
	return fuse.OK
}

func (f *OzoneFile) Allocate(off uint64, size uint64, mode uint32) (code fuse.Status) {
	return fuse.ENOSYS
}
