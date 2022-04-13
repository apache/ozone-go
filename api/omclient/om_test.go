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
    "fmt"
    hadoop_ipc_client "github.com/apache/ozone-go/api/hadoop/ipc/client"
    "github.com/apache/ozone-go/api/proto/hdds"
    ozone_proto "github.com/apache/ozone-go/api/proto/ozone"
    "github.com/apache/ozone-go/api/utils"
    "reflect"
    "testing"

    uuid "github.com/nu7hatch/gouuid"
)

var omHost = "localhost:9870"
var clientId, _ = uuid.NewV4()

// MockIpcClientTest TODO
type MockIpcClient struct {
    Mock string
}

func newMockIpcClient() *MockIpcClient {
    return &MockIpcClient{Mock: "mock"}
}

func (mockIpcClient *MockIpcClient) UpdateServerAddress(address string) {

}

// SubmitRequest TODO
func (mockIpcClient *MockIpcClient) SubmitRequest(request *ozone_proto.OMRequest) (*ozone_proto.OMResponse, error) {
    cmdType := *request.CmdType
    switch cmdType {
    case ozone_proto.Type_CreateVolume:
        return &ozone_proto.OMResponse{CreateVolumeResponse: &ozone_proto.CreateVolumeResponse{}}, nil
    case ozone_proto.Type_CreateBucket:
        return &ozone_proto.OMResponse{CreateBucketResponse: &ozone_proto.CreateBucketResponse{}}, nil
    case ozone_proto.Type_CreateKey:
        return &ozone_proto.OMResponse{CreateKeyResponse: &ozone_proto.CreateKeyResponse{
            KeyInfo: &ozone_proto.KeyInfo{
                VolumeName: request.CreateKeyRequest.KeyArgs.VolumeName,
                BucketName: request.CreateKeyRequest.KeyArgs.BucketName,
                KeyName:    request.CreateKeyRequest.KeyArgs.KeyName,
                Type:       request.CreateKeyRequest.KeyArgs.Type,
                Factor:     request.CreateKeyRequest.KeyArgs.Factor,
            }}}, nil
    case ozone_proto.Type_InfoVolume:
        return &ozone_proto.OMResponse{InfoVolumeResponse: &ozone_proto.InfoVolumeResponse{
            VolumeInfo: &ozone_proto.VolumeInfo{
                AdminName:        nil,
                OwnerName:        nil,
                Volume:           request.InfoVolumeRequest.VolumeName,
                QuotaInBytes:     nil,
                Metadata:         nil,
                VolumeAcls:       []*ozone_proto.OzoneAclInfo{},
                CreationTime:     nil,
                ObjectID:         nil,
                UpdateID:         nil,
                ModificationTime: nil,
                QuotaInNamespace: nil,
                UsedNamespace:    nil,
            }}}, nil
    case ozone_proto.Type_ListVolume:
        return &ozone_proto.OMResponse{ListVolumeResponse: &ozone_proto.ListVolumeResponse{
            VolumeInfo: make([]*ozone_proto.VolumeInfo, 0)}}, nil
    case ozone_proto.Type_ListBuckets:
        return &ozone_proto.OMResponse{ListBucketsResponse: &ozone_proto.ListBucketsResponse{
            BucketInfo: make([]*ozone_proto.BucketInfo, 0)}}, nil
    case ozone_proto.Type_InfoBucket:
        return &ozone_proto.OMResponse{InfoBucketResponse: &ozone_proto.InfoBucketResponse{
            BucketInfo: &ozone_proto.BucketInfo{
                VolumeName: request.InfoBucketRequest.VolumeName,
                BucketName: request.InfoBucketRequest.BucketName,
            }}}, nil
    case ozone_proto.Type_LookupKey:
        return &ozone_proto.OMResponse{LookupKeyResponse: &ozone_proto.LookupKeyResponse{
            KeyInfo: &ozone_proto.KeyInfo{
                VolumeName: request.LookupKeyRequest.KeyArgs.VolumeName,
                BucketName: request.LookupKeyRequest.KeyArgs.BucketName,
                KeyName:    request.LookupKeyRequest.KeyArgs.KeyName,
                Type:       GetReplicationType(),
                Factor:     GetReplication(),
                DataSize:   utils.PointUint64(0),
            }}}, nil
    case ozone_proto.Type_ListKeys:
        return &ozone_proto.OMResponse{ListKeysResponse: &ozone_proto.ListKeysResponse{
            KeyInfo: make([]*ozone_proto.KeyInfo, 0)}}, nil
    case ozone_proto.Type_AllocateBlock:
        return &ozone_proto.OMResponse{AllocateBlockResponse: &ozone_proto.AllocateBlockResponse{
            KeyLocation: &ozone_proto.KeyLocation{}}}, nil
    case ozone_proto.Type_CommitKey:
        return &ozone_proto.OMResponse{CommitKeyResponse: &ozone_proto.CommitKeyResponse{}}, nil
    case ozone_proto.Type_RenameKey:
        return &ozone_proto.OMResponse{RenameKeysResponse: &ozone_proto.RenameKeysResponse{}}, nil
    case ozone_proto.Type_DeleteKey:
        return &ozone_proto.OMResponse{DeleteKeyResponse: &ozone_proto.DeleteKeyResponse{
            KeyInfo: &ozone_proto.KeyInfo{KeyName: request.DeleteKeyRequest.KeyArgs.KeyName}}}, nil
    case ozone_proto.Type_ListStatus:
        return &ozone_proto.OMResponse{ListStatusResponse: &ozone_proto.ListStatusResponse{
            Statuses: []*ozone_proto.OzoneFileStatusProto{mockOzoneFileStatusProto(false), mockOzoneFileStatusProto(true)},
        }}, nil
    default:
        return nil, fmt.Errorf("unknown request type %v", request.CmdType.String())
    }
}

func TestCreateOmClient(t *testing.T) {
    type args struct {
        omHost string
    }
    ugi, _ := hadoop_ipc_client.CreateSimpleUGIProto()
    tests := []struct {
        name    string
        args    args
        want    *OmClient
        wantErr bool
    }{
        // TODO: Add test cases.
        {
            name: "create ozone manager Client",
            args: args{omHost: omHost},
            want: &OmClient{
                OmHost:   omHost,
                ClientId: clientId.String(),
                IpcClient: &HadoopIpcClient{Client: &hadoop_ipc_client.Client{
                    ClientId:      clientId,
                    Ugi:           ugi,
                    ServerAddress: "localhost:9870",
                    TCPNoDelay:    false,
                }},
            },
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := CreateOmClient(tt.args.omHost)
            if err != nil {
                t.Errorf("CreateOmClient() = %v, wantErr %v", err, tt.wantErr)
            }
            got.ClientId = clientId.String()
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("CreateOmClient() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestHadoopIpcClient_SubmitRequest(t *testing.T) {
    type fields struct {
        client *MockIpcClient
    }
    type args struct {
        request *ozone_proto.OMRequest
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    *ozone_proto.OMResponse
        wantErr bool
    }{
        // TODO: Add test cases.
        {
            name:    "submit request",
            fields:  fields{client: newMockIpcClient()},
            args:    args{request: &ozone_proto.OMRequest{CmdType: ozone_proto.Type.Enum(10000)}},
            want:    nil,
            wantErr: true,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            ipcClient := tt.fields.client
            _, err := ipcClient.SubmitRequest(tt.args.request)
            if (err != nil) != tt.wantErr {
                t.Errorf("SubmitRequest() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
        })
    }
}

func TestOmClient_AllocateBlock(t *testing.T) {
    type fields struct {
        OmHost    string
        ClientId  string
        IpcClient IpcClient
    }
    type args struct {
        volume   string
        bucket   string
        key      string
        clientID *uint64
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    *ozone_proto.AllocateBlockResponse
        wantErr bool
    }{
        // TODO: Add test cases.
        {
            name: "allocate block",
            fields: fields{
                OmHost:    omHost,
                ClientId:  clientId.String(),
                IpcClient: newMockIpcClient(),
            },
            args: args{
                volume:   "",
                bucket:   "",
                key:      "",
                clientID: nil,
            },
            want:    &ozone_proto.AllocateBlockResponse{KeyLocation: &ozone_proto.KeyLocation{}},
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            om := &OmClient{
                OmHost:    tt.fields.OmHost,
                ClientId:  tt.fields.ClientId,
                IpcClient: tt.fields.IpcClient,
            }
            got, err := om.AllocateBlock(tt.args.volume, tt.args.bucket, tt.args.key, tt.args.clientID, nil)

            if (err != nil) != tt.wantErr {
                t.Errorf("AllocateBlock() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("AllocateBlock() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func Test_getRpcPort(t *testing.T) {
    type args struct {
        ports []*hdds.Port
    }
    tests := []struct {
        name string
        args args
        want uint32
    }{
        // TODO: Add test cases.
        {
            name: "get rpc port",
            args: args{ports: []*hdds.Port{{
                Name:  utils.PointString("RATIS"),
                Value: utils.PointUint32(9870),
            }}},
            want: 9870,
        },
        {
            name: "get rpc port",
            args: args{ports: []*hdds.Port{{
                Name:  utils.PointString("STANDALONE"),
                Value: utils.PointUint32(9870),
            }}},
            want: 0,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := getRpcPort(tt.args.ports); got != tt.want {
                t.Errorf("getRpcPort() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestOmClient_Equal(t *testing.T) {
    type fields struct {
        OmHost    string
        ClientId  string
        IpcClient IpcClient
    }
    type args struct {
        i interface{}
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   bool
    }{
        // TODO: Add test cases.
        {
            name: "omclient equal",
            fields: fields{
                OmHost:    omHost,
                ClientId:  clientId.String(),
                IpcClient: newMockIpcClient(),
            },
            args: args{OmClient{
                OmHost:    omHost,
                ClientId:  clientId.String(),
                IpcClient: newMockIpcClient(),
            }},
            want: true,
        },
        {
            name: "omclient not equal",
            fields: fields{
                OmHost:    "",
                ClientId:  clientId.String(),
                IpcClient: newMockIpcClient(),
            },
            args: args{OmClient{
                OmHost:    omHost,
                ClientId:  clientId.String(),
                IpcClient: newMockIpcClient(),
            }},
            want: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            om := &OmClient{
                OmHost:    tt.fields.OmHost,
                ClientId:  tt.fields.ClientId,
                IpcClient: tt.fields.IpcClient,
            }
            if got := om.Equal(tt.args.i); got != tt.want {
                t.Errorf("Equal() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestOmClient_Equal1(t *testing.T) {
    type fields struct {
        OmHost    string
        ClientId  string
        IpcClient IpcClient
        Id        uint64
    }
    type args struct {
        i interface{}
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   bool
    }{
        // TODO: Add test cases.
        {
            name: "om client equal",
            fields: fields{
                OmHost:    "om",
                ClientId:  "1",
                IpcClient: nil,
                Id:        0,
            },
            args: args{fields{
                OmHost:    "om",
                ClientId:  "1",
                IpcClient: nil,
                Id:        0,
            }},
            want: false,
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
            if got := om.Equal(tt.args.i); got != tt.want {
                t.Errorf("Equal() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestHadoopIpcClient_UpdateServerAddress(t *testing.T) {
    type fields struct {
        Client *hadoop_ipc_client.Client
    }
    type args struct {
        address string
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   string
    }{
        // TODO: Add test cases.
        {
            name: "update server address",
            fields: fields{Client: &hadoop_ipc_client.Client{
                ClientId:      nil,
                Ugi:           nil,
                ServerAddress: "localhost",
                TCPNoDelay:    false,
            }},

            args: args{address: "127.0.0.1"},
            want: "127.0.0.1",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            ipcClient := &HadoopIpcClient{
                Client: tt.fields.Client,
            }
            ipcClient.UpdateServerAddress(tt.args.address)
            if got := ipcClient.Client.ServerAddress == tt.want; !got {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
