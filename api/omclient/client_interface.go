// Package omclient TODO
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
package omclient

import (
    "io"
    "github.com/apache/ozone-go/api/common"
    ozone_proto "github.com/apache/ozone-go/api/proto/ozone"
)

// IpcClient TODO
type IpcClient interface {
    UpdateServerAddress(serverAddress string)
    SubmitRequest(request *ozone_proto.OMRequest) (*ozone_proto.OMResponse, error)
}

// OzoneManagerClient TODO
type OzoneManagerClient interface {
    CreateVolume(name string) error
    InfoVolume(name string) (*common.Volume, error)
    ListVolumes() ([]*common.Volume, error)

    CreateBucket(volume string, bucket string) error
    InfoBucket(volume string, bucket string) (*common.Bucket, error)
    ListBucket(volume string) ([]*common.Bucket, error)

    // InfoKey TODO
    // CreateKey(volume string, bucket string, key string) (*ozone_proto.CreateKeyResponse, error)
    InfoKey(volume string, bucket string, key string) (*common.Key, error)
    ListKeys(volume, bucket, prefix string, limit int32) ([]*common.Key, error)
    RenameKey(volume, bucket, key, toKeyName string) (string, error)
    DeleteKey(volume, bucket, key string) (string, error)
    GetKey(volume, bucket, key string, destination io.Writer) error
    PutKey(volume string, bucket string, key string, length int64, source io.Reader) error
}

// FsClient TODO
type FsClient interface {
    LookupFile(volume, bucket, key string) (*common.OzoneFileInfo, error)
    DeleteFile(volume, bucket, key string, skipTrash bool) (string, error)
    ListFiles(volume, bucket, key, startKey string, limit uint64) ([]*common.OzoneFileInfo, error)
    Mkdir(volume, bucket, key string) error
    RenameFile(volume, bucket, fromKey, toKey string) error
    CreateFile(volume, bucket, key string) error
    GetFile(volume, bucket, key string, destination io.Writer) error
    PutFile(volume string, bucket string, key string, length int64, source io.Reader) error
}
