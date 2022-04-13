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
    "reflect"
    "testing"
)

func TestOmClient_CreateBucket(t *testing.T) {
    type fields struct {
        OmHost    string
        clientId  string
        IpcClient IpcClient
    }
    type args struct {
        volume string
        bucket string
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        wantErr bool
    }{
        // TODO: Add test cases.
        {
            name: "create bucket",
            fields: fields{
                OmHost:    omHost,
                clientId:  clientId.String(),
                IpcClient: newMockIpcClient(),
            },
            args: args{
                volume: "v",
                bucket: "b",
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
            if err := om.CreateBucket(tt.args.volume, tt.args.bucket); (err != nil) != tt.wantErr {
                t.Errorf("CreateBucket() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestOmClient_InfoBucket(t *testing.T) {
    type fields struct {
        OmHost    string
        clientId  string
        IpcClient IpcClient
    }
    type args struct {
        volume string
        bucket string
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    *common.Bucket
        wantErr bool
    }{
        // TODO: Add test cases.
        {
            name: "info bucket",
            fields: fields{
                OmHost:    omHost,
                clientId:  clientId.String(),
                IpcClient: newMockIpcClient(),
            },
            args: args{
                volume: "v",
                bucket: "b",
            },
            want: &common.Bucket{
                BucketName:       "b",
                VolumeName:       "v",
                StorageType:      "DISK",
                QuotaInBytes:     "16.00 EB",
                UsedBytes:        "0 B",
                QuotaInNamespace: -2,
                UsedNamespace:    0,
                CreationTime:     "1970-01-01 08:00:00",
                ModificationTime: "1970-01-01 08:00:00",
                BucketLayout:     "LEGACY",
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
            got, err := om.InfoBucket(tt.args.volume, tt.args.bucket)
            if (err != nil) != tt.wantErr {
                t.Errorf("InfoBucket() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("InfoBucket() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestOmClient_ListBucket(t *testing.T) {
    type fields struct {
        OmHost    string
        clientId  string
        IpcClient IpcClient
    }
    type args struct {
        volume string
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    []*common.Bucket
        wantErr bool
    }{
        // TODO: Add test cases.
        {
            name: "list buckets",
            fields: fields{
                OmHost:    omHost,
                clientId:  clientId.String(),
                IpcClient: newMockIpcClient(),
            },
            args:    args{volume: "v"},
            want:    make([]*common.Bucket, 0),
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
            got, err := om.ListBucket(tt.args.volume)
            if (err != nil) != tt.wantErr {
                t.Errorf("ListBucket() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("ListBucket() got = %v, want %v", got, tt.want)
			}
		})
	}
}
