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
    "github.com/apache/ozone-go/api/common"
    "github.com/apache/ozone-go/api/config"
    dnproto "github.com/apache/ozone-go/api/proto/datanode"
    "github.com/apache/ozone-go/api/proto/hdds"
    ozone_proto "github.com/apache/ozone-go/api/proto/ozone"
    "github.com/apache/ozone-go/api/utils"
    "reflect"
    "testing"
)

func mockAcls() []*ozone_proto.OzoneAclInfo {
    aclInfo := NewAclInfo()
    aclInfo.AddRight(ALL)
    rights := aclInfo.ToBytes()
    return []*ozone_proto.OzoneAclInfo{
        createOzoneAclInfo(ozone_proto.OzoneAclInfo_USER.Enum(), "test", rights, ozone_proto.OzoneAclInfo_ACCESS.Enum()),
        createOzoneAclInfo(ozone_proto.OzoneAclInfo_GROUP.Enum(), "test", rights, ozone_proto.OzoneAclInfo_ACCESS.Enum()),
    }
}

func TestConvertBlockId(t *testing.T) {
    type args struct {
        bid *hdds.BlockID
    }
    tests := []struct {
        name string
        args args
        want *dnproto.DatanodeBlockID
    }{
        // TODO: Add test cases.
        {
            name: "convert_block_id",
            args: args{bid: &hdds.BlockID{
                ContainerBlockID: &hdds.ContainerBlockID{
                    ContainerID: utils.PointInt64(1),
                    LocalID:     utils.PointInt64(2),
                },
                BlockCommitSequenceId: utils.PointUint64(3),
            }},
            want: &dnproto.DatanodeBlockID{
                ContainerID:           utils.PointInt64(1),
                LocalID:               utils.PointInt64(2),
                BlockCommitSequenceId: utils.PointUint64(3),
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := ConvertBlockId(tt.args.bid); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("ConvertBlockId() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestKeyFromProto(t *testing.T) {
    type args struct {
        keyProto *ozone_proto.KeyInfo
    }
    tests := []struct {
        name string
        args args
        want *common.Key
    }{
        // TODO: Add test cases.
        {
            name: "key from proto",
            args: args{keyProto: &ozone_proto.KeyInfo{
                VolumeName: utils.PointString("v"),
                BucketName: utils.PointString("b"),
                KeyName:    utils.PointString("k"),
                Type:       hdds.ReplicationType_RATIS.Enum(),
                Factor:     hdds.ReplicationFactor_THREE.Enum(),
                DataSize:   utils.PointUint64(1),
            }},
            want: &common.Key{
                VolumeName: "v", BucketName: "b", Name: "k",
                ReplicationType:   common.RATIS,
                ReplicationFactor: common.THREE,
                DataSize:          1,
                DataSizeFriendly:  "1.00 B",
                CreationTime:      "1970-01-01 08:00:00",
                ModificationTime:  "1970-01-01 08:00:00",
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := KeyFromProto(tt.args.keyProto); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("KeyFromProto() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestOmClient_CommitKey(t *testing.T) {
    type fields struct {
        OmHost    string
        clientId  string
        IpcClient IpcClient
    }
    type args struct {
        volume       string
        bucket       string
        key          string
        id           *uint64
        keyLocations []*ozone_proto.KeyLocation
        size         uint64
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    *ozone_proto.CommitKeyResponse
        wantErr bool
    }{
        // TODO: Add test cases.
        {
            name: "commit key",
            fields: fields{
                OmHost:    omHost,
                clientId:  clientId.String(),
                IpcClient: newMockIpcClient(),
            },
            args:    args{},
            want:    &ozone_proto.CommitKeyResponse{},
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            om := &OmClient{
                OmHost:    tt.fields.OmHost,
                ClientId:  tt.fields.clientId,
                IpcClient: tt.fields.IpcClient,
            }
            got, err := om.CommitKey(tt.args.volume, tt.args.bucket, tt.args.key, tt.args.id, tt.args.keyLocations, tt.args.size)
            if (err != nil) != tt.wantErr {
                t.Errorf("CommitKey() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got.GetCommitKeyResponse(), tt.want) {
                t.Errorf("CommitKey() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestOmClient_CreateKey(t *testing.T) {
    config.OzoneConfig.Set(config.OZONE_REPLICATION_KEY, "3")
    config.OzoneConfig.Set(config.OZONE_REPLICATION_TYPE_KEY, "ratis")
    type fields struct {
        OmHost    string
        clientId  string
        IpcClient IpcClient
    }
    type args struct {
        volume string
        bucket string
        key    string
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    *ozone_proto.CreateKeyResponse
        wantErr bool
    }{
        // TODO: Add test cases.
        {
            name: "create key",
            fields: fields{
                OmHost:    "localhost:9870",
                clientId:  "1",
                IpcClient: newMockIpcClient(),
            },
            args: args{
                volume: "v",
                bucket: "b",
                key:    "k",
            },
            want: &ozone_proto.CreateKeyResponse{
                KeyInfo: &ozone_proto.KeyInfo{
                    VolumeName: utils.PointString("v"),
                    BucketName: utils.PointString("b"),
                    KeyName:    utils.PointString("k"),
                    Type:       hdds.ReplicationType_RATIS.Enum(),
                    Factor:     hdds.ReplicationFactor_THREE.Enum(),
                }},
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            om := &OmClient{
                OmHost:    tt.fields.OmHost,
                ClientId:  tt.fields.clientId,
                IpcClient: tt.fields.IpcClient,
            }
            got, err := om.CreateKey(tt.args.volume, tt.args.bucket, tt.args.key)
            if (err != nil) != tt.wantErr {
                t.Errorf("CreateKey() error = %v, wantErr %v", err, tt.wantErr)
                return
            }

            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("CreateKey() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestOmClient_DeleteKey(t *testing.T) {
    type fields struct {
        OmHost    string
        clientId  string
        IpcClient IpcClient
    }
    type args struct {
        volume string
        bucket string
        key    string
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    string
        wantErr bool
    }{
        // TODO: Add test cases.
        {
            name: "delete key",
            fields: fields{
                OmHost:    omHost,
                clientId:  clientId.String(),
                IpcClient: newMockIpcClient(),
            },
            args: args{
                volume: "v",
                bucket: "b",
                key:    "k",
            },
            want:    "k",
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            om := &OmClient{
                OmHost:    tt.fields.OmHost,
                ClientId:  tt.fields.clientId,
                IpcClient: tt.fields.IpcClient,
            }
            got, err := om.DeleteKey(tt.args.volume, tt.args.bucket, tt.args.key)
            if (err != nil) != tt.wantErr {
                t.Errorf("DeleteKey() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("DeleteKey() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestOmClient_InfoKey(t *testing.T) {
    config.OzoneConfig.Set(config.OZONE_REPLICATION_TYPE_KEY, "ratis")
    type fields struct {
        OmHost    string
        clientId  string
        IpcClient IpcClient
    }
    type args struct {
        volume string
        bucket string
        key    string
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    *common.Key
        wantErr bool
    }{
        // TODO: Add test cases.
        {
            name: "info key",
            fields: fields{
                OmHost:    omHost,
                clientId:  clientId.String(),
                IpcClient: newMockIpcClient(),
            },
            args: args{
                volume: "v",
                bucket: "b",
                key:    "k",
            },
            want: &common.Key{
                Name:              "k",
                BucketName:        "b",
                VolumeName:        "v",
                DataSize:          0,
                ReplicationType:   common.RATIS,
                ReplicationFactor: common.THREE,
                DataSizeFriendly:  "0 B",
                CreationTime:      "1970-01-01 08:00:00",
                ModificationTime:  "1970-01-01 08:00:00",
            },
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            om := &OmClient{
                OmHost:    tt.fields.OmHost,
                ClientId:  tt.fields.clientId,
                IpcClient: tt.fields.IpcClient,
            }
            got, err := om.InfoKey(tt.args.volume, tt.args.bucket, tt.args.key)
            if (err != nil) != tt.wantErr {
                t.Errorf("InfoKey() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("InfoKey() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestOmClient_ListKeys(t *testing.T) {
    type fields struct {
        OmHost    string
        clientId  string
        IpcClient IpcClient
    }
    type args struct {
        volume string
        bucket string
        limit  int32
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    []*common.Key
        wantErr bool
    }{
        // TODO: Add test cases.
        {
            name: "list keys",
            fields: fields{
                OmHost:    "",
                clientId:  "",
                IpcClient: newMockIpcClient(),
            },
            args: args{
                volume: "v",
                bucket: "b",
                limit:  3,
            },
            want:    make([]*common.Key, 0),
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            om := &OmClient{
                OmHost:    tt.fields.OmHost,
                ClientId:  tt.fields.clientId,
                IpcClient: tt.fields.IpcClient,
            }
            got, err := om.ListKeys(tt.args.volume, tt.args.bucket, "", tt.args.limit)
            if (err != nil) != tt.wantErr {
                t.Errorf("ListKeys() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("ListKeys() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestOmClient_ListKeysPrefix(t *testing.T) {
    type fields struct {
        OmHost    string
        clientId  string
        IpcClient IpcClient
    }
    type args struct {
        volume string
        bucket string
        prefix string
        limit  int32
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    []*common.Key
        wantErr bool
    }{
        // TODO: Add test cases.
        {
            name: "list prefix keys",
            fields: fields{
                OmHost:    "",
                clientId:  "",
                IpcClient: newMockIpcClient(),
            },
            args: args{
                volume: "v",
                bucket: "b",
                prefix: "prefix",
                limit:  3,
            },
            want:    make([]*common.Key, 0),
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            om := &OmClient{
                OmHost:    tt.fields.OmHost,
                ClientId:  tt.fields.clientId,
                IpcClient: tt.fields.IpcClient,
            }
            got, err := om.ListKeys(tt.args.volume, tt.args.bucket, tt.args.prefix, tt.args.limit)
            if (err != nil) != tt.wantErr {
                t.Errorf("ListKeysPrefix() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("ListKeysPrefix() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestOmClient_LookupKey(t *testing.T) {
    type fields struct {
        OmHost    string
        clientId  string
        IpcClient IpcClient
    }
    type args struct {
        volume string
        bucket string
        key    string
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    *ozone_proto.KeyInfo
        wantErr bool
    }{
        // TODO: Add test cases.
        {
            name: "look up key",
            fields: fields{
                OmHost:    omHost,
                clientId:  clientId.String(),
                IpcClient: newMockIpcClient(),
            },
            args: args{
                volume: "v",
                bucket: "b",
                key:    "k",
            },
            want: &ozone_proto.KeyInfo{
                VolumeName: utils.PointString("v"),
                BucketName: utils.PointString("b"),
                KeyName:    utils.PointString("k"),
                DataSize:   utils.PointUint64(0),
                Type:       GetReplicationType(),
                Factor:     GetReplication(),
            },
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            om := &OmClient{
                OmHost:    tt.fields.OmHost,
                ClientId:  tt.fields.clientId,
                IpcClient: tt.fields.IpcClient,
            }
            got, err := om.LookupKey(tt.args.volume, tt.args.bucket, tt.args.key)
            if (err != nil) != tt.wantErr {
                t.Errorf("LookupKey() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("LookupKey() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestOmClient_RenameKey(t *testing.T) {
    type fields struct {
        OmHost    string
        clientId  string
        IpcClient IpcClient
    }
    type args struct {
        volume    string
        bucket    string
        key       string
        toKeyName string
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    string
        wantErr bool
    }{
        // TODO: Add test cases.
        {
            name: "rename key",
            fields: fields{
                OmHost:    omHost,
                clientId:  clientId.String(),
                IpcClient: newMockIpcClient(),
            },
            args: args{
                volume:    "v",
                bucket:    "b",
                key:       "k",
                toKeyName: "to_key",
            },
            want:    "",
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            om := &OmClient{
                OmHost:    tt.fields.OmHost,
                ClientId:  tt.fields.clientId,
                IpcClient: tt.fields.IpcClient,
            }
            got, err := om.RenameKey(tt.args.volume, tt.args.bucket, tt.args.key, tt.args.toKeyName)
            if (err != nil) != tt.wantErr {
                t.Errorf("RenameKey() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("RenameKey() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestOmClient_ListFiles(t *testing.T) {
    type fields struct {
        OmHost    string
        ClientId  string
        IpcClient IpcClient
        Id        uint64
    }
    type args struct {
        volume   string
        bucket   string
        key      string
        startKey string
        limit    uint64
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    []*common.OzoneFileInfo
        wantErr bool
    }{
        // TODO: Add test cases.
        {
            name: "list file",
            fields: fields{
                OmHost:    "localhost",
                ClientId:  "0",
                IpcClient: newMockIpcClient(),
                Id:        0,
            },
            args: args{
                volume:   "v",
                bucket:   "b",
                key:      "k",
                startKey: "",
                limit:    2,
            },
            want: []*common.OzoneFileInfo{{
                Name:        "k",
                Path:        "/v/b/k",
                Replica:     3,
                Size:        1,
                ReplicaSize: 3,
                Mode:        "-ALLALL",
                Owner:       "test",
                Group:       "test",
                ModifyTime:  0,
            }, {
                Name:        "k",
                Path:        "/v/b/k",
                Replica:     3,
                Size:        1,
                ReplicaSize: 3,
                Mode:        "dALLALL",
                Owner:       "test",
                Group:       "test",
                ModifyTime:  0,
            }},
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            om := &OmClient{
                OmHost:    tt.fields.OmHost,
                ClientId:  tt.fields.ClientId,
                IpcClient: tt.fields.IpcClient,
                Id:        tt.fields.Id,
            }
            got, err := om.ListFiles(tt.args.volume, tt.args.bucket, tt.args.key, tt.args.startKey, tt.args.limit)
            if (err != nil) != tt.wantErr {
                t.Errorf("ListFiles() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("ListFiles() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestOmClient_TouchzKey(t *testing.T) {
    type fields struct {
        OmHost    string
        ClientId  string
        IpcClient IpcClient
        Id        uint64
    }
    type args struct {
        volume string
        bucket string
        key    string
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    *common.Key
        wantErr bool
    }{
        // TODO: Add test cases.
        {
            name: "touch key",
            fields: fields{
                OmHost:    "localhost",
                ClientId:  "0",
                IpcClient: newMockIpcClient(),
                Id:        0,
            },
            args: args{
                volume: "v",
                bucket: "b",
                key:    "k",
            },
            want: &common.Key{
                VolumeName:        "v",
                BucketName:        "b",
                Name:              "k",
                DataSize:          0,
                DataSizeFriendly:  "0 B",
                ReplicationFactor: 3,
                ReplicationType:   "RATIS",
                CreationTime:      "1970-01-01 08:00:00",
                ModificationTime:  "1970-01-01 08:00:00",
            },
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            om := &OmClient{
                OmHost:    tt.fields.OmHost,
                ClientId:  tt.fields.ClientId,
                IpcClient: tt.fields.IpcClient,
                Id:        tt.fields.Id,
            }
            got, err := om.TouchzKey(tt.args.volume, tt.args.bucket, tt.args.key)
            if (err != nil) != tt.wantErr {
                t.Errorf("TouchzKey() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("TouchzKey() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func mockOzoneFileStatusProto(isDir bool) *ozone_proto.OzoneFileStatusProto {
    return &ozone_proto.OzoneFileStatusProto{
        KeyInfo: &ozone_proto.KeyInfo{
            VolumeName:         utils.PointString("v"),
            BucketName:         utils.PointString("b"),
            KeyName:            utils.PointString("k"),
            DataSize:           utils.PointUint64(1),
            Type:               hdds.ReplicationType_RATIS.Enum(),
            Factor:             hdds.ReplicationFactor_THREE.Enum(),
            KeyLocationList:    nil,
            CreationTime:       utils.PointUint64(0),
            ModificationTime:   utils.PointUint64(0),
            LatestVersion:      utils.PointUint64(0),
            Metadata:           nil,
            FileEncryptionInfo: nil,
            Acls:               mockAcls(),
            ObjectID:           nil,
            UpdateID:           nil,
            ParentID:           nil,
        },
        BlockSize:   utils.PointUint64(0),
        IsDirectory: utils.PointBool(isDir),
    }
}

func TestFileFromProto(t *testing.T) {
    type args struct {
        proto *ozone_proto.OzoneFileStatusProto
    }
    tests := []struct {
        name string
        args args
        want *common.OzoneFileInfo
    }{
        // TODO: Add test cases.
        {
            name: "file from proto",
            args: args{proto: mockOzoneFileStatusProto(false)},
            want: &common.OzoneFileInfo{
                Name:        "k",
                Path:        "/v/b/k",
                Replica:     3,
                Size:        1,
                ReplicaSize: 3,
                Mode:        "-ALLALL",
                Owner:       "test",
                Group:       "test",
                ModifyTime:  0,
            },
        },
        {
            name: "dir from proto",
            args: args{proto: mockOzoneFileStatusProto(true)},
            want: &common.OzoneFileInfo{
                Name:        "k",
                Path:        "/v/b/k",
                Replica:     3,
                Size:        1,
                ReplicaSize: 3,
                Mode:        "dALLALL",
                Owner:       "test",
                Group:       "test",
                ModifyTime:  0,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := FileFromProto(tt.args.proto); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileFromProto() = %v, want %v", got, tt.want)
			}
		})
	}
}
