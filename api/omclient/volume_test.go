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

func TestOmClient_CreateVolume(t *testing.T) {
    type fields struct {
        OmHost    string
        clientId  string
        IpcClient IpcClient
    }
    type args struct {
        volumeName string
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        wantErr bool
    }{
        // TODO: Add test cases.
        {
            name: "create volume",
            fields: fields{
                OmHost:    omHost,
                clientId:  clientId.String(),
                IpcClient: newMockIpcClient(),
            },
            args:    args{volumeName: "v"},
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
            if err := om.CreateVolume(tt.args.volumeName); (err != nil) != tt.wantErr {
                t.Errorf("CreateVolume() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestOmClient_InfoVolume(t *testing.T) {
    type fields struct {
        OmHost    string
        clientId  string
        IpcClient IpcClient
    }
    type args struct {
        name string
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    *common.Volume
        wantErr bool
    }{
        // TODO: Add test cases.
        {
            name: "info volume",
            fields: fields{
                OmHost:    omHost,
                clientId:  clientId.String(),
                IpcClient: newMockIpcClient(),
            },
            args: args{},
            want: &common.Volume{
                QuotaInBytes:     "0 B",
                QuotaInNamespace: -2,
                CreationTime:     "1970-01-01 08:00:00",
                ModificationTime: "1970-01-01 08:00:00",
                Acls:             []common.Acl{}},
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
            got, err := om.InfoVolume(tt.args.name)
            if (err != nil) != tt.wantErr {
                t.Errorf("InfoVolume() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("InfoVolume() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestOmClient_ListVolumes(t *testing.T) {
    type fields struct {
        OmHost    string
        clientId  string
        IpcClient IpcClient
    }
    tests := []struct {
        name    string
        fields  fields
        want    []*common.Volume
        wantErr bool
    }{
        // TODO: Add test cases.
        {
            name: "list volumes",
            fields: fields{
                OmHost:    omHost,
                clientId:  clientId.String(),
                IpcClient: newMockIpcClient(),
            },
            want:    make([]*common.Volume, 0),
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
            got, err := om.ListVolumes()
            if (err != nil) != tt.wantErr {
                t.Errorf("ListVolumes() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("ListVolumes() got = %v, want %v", got, tt.want)
            }
        })
    }
}
