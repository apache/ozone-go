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
package datanodeclient

import (
    "context"
    "io"
    "github.com/apache/ozone-go/api/proto/datanode"
    "github.com/apache/ozone-go/api/proto/hdds"
    "github.com/apache/ozone-go/api/utils"
    "reflect"
    "testing"

    "github.com/agiledragon/gomonkey/v2"
    "google.golang.org/grpc"
)

func TestNewXceiverClientStandalone(t *testing.T) {
    p, _ := NewPipelineOperator(NewPipeline())
    type args struct {
        pipeline *PipelineOperator
        read     bool
    }
    tests := []struct {
        name string
        args args
        want XceiverClientSpi
    }{
        // TODO: Add test cases.
        {
            name: "new standalone client",
            args: args{
                pipeline: p,
                read:     false,
            },
            want: &XceiverClientStandalone{pipeline: p, forRead: false, close: false},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewXceiverClientStandalone(tt.args.pipeline, tt.args.read); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewXceiverClientStandalone() = %v, want %v", got, tt.want)
            }
        })
    }
}

type MockXceiverClientProtocolServiceClient struct {
}

func NewMockXceiverClientProtocolServiceClient(cc grpc.ClientConnInterface) datanode.XceiverClientProtocolServiceClient {
    return &MockXceiverClientProtocolServiceClient{}
}

func (c *MockXceiverClientProtocolServiceClient) Send(ctx context.Context, opts ...grpc.CallOption) (datanode.XceiverClientProtocolService_SendClient, error) {
    return NewMockXceiverClientProtocolService_SendClient(), nil
}

type MockXceiverClientProtocolService_SendClient struct {
    grpc.ClientStream
    Err error
}

func NewMockXceiverClientProtocolService_SendClient() datanode.XceiverClientProtocolService_SendClient {
    return &MockXceiverClientProtocolService_SendClient{}
}

func (x *MockXceiverClientProtocolService_SendClient) Send(p *datanode.ContainerCommandRequestProto) error {
    return x.Err
}

func (x *MockXceiverClientProtocolService_SendClient) Recv() (*datanode.ContainerCommandResponseProto, error) {
    return nil, io.EOF
}

func TestXceiverClientStandalone_Close(t *testing.T) {
    c := NewMockXceiverClientProtocolService_SendClient()
    patch := gomonkey.ApplyMethod(reflect.TypeOf(c), "CloseSend", func(*MockXceiverClientProtocolService_SendClient) error {
        return nil
    })
    defer patch.Reset()
    type fields struct {
        pipeline *PipelineOperator
        client   datanode.XceiverClientProtocolService_SendClient
        forRead  bool
        close    bool
        response chan *datanode.ContainerCommandResponseProto
    }
    tests := []struct {
        name   string
        fields fields
    }{
        // TODO: Add test cases.
        {
            name: "close standalone",
            fields: fields{
                pipeline: nil,
                client:   nil,
                forRead:  false,
                close:    true,
                response: nil,
            },
        },
        {
            name: "close standalone",
            fields: fields{
                pipeline: nil,
                client:   c,
                forRead:  false,
                close:    false,
                response: nil,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            x := &XceiverClientStandalone{
                pipeline: tt.fields.pipeline,
                client:   tt.fields.client,
                forRead:  tt.fields.forRead,
                close:    tt.fields.close,
                response: tt.fields.response,
            }
            x.Close()
        })
    }
}

func TestXceiverClientStandalone_Connect(t *testing.T) {
    op, _ := NewPipelineOperator(NewPipeline())
    patch := gomonkey.ApplyFunc(grpc.DialContext, func(ctx context.Context, target string, opts ...grpc.DialOption) (conn *grpc.ClientConn, err error) {
        return &grpc.ClientConn{}, nil
    })
    defer patch.Reset()

    patch1 := gomonkey.ApplyFunc(datanode.NewXceiverClientProtocolServiceClient, func(cc grpc.ClientConnInterface) datanode.XceiverClientProtocolServiceClient {
        return NewMockXceiverClientProtocolServiceClient(cc)
    })
    defer patch1.Reset()

    type fields struct {
        pipeline *PipelineOperator
        client   datanode.XceiverClientProtocolService_SendClient
        forRead  bool
        close    bool
        response chan *datanode.ContainerCommandResponseProto
    }
    tests := []struct {
        name    string
        fields  fields
        wantErr bool
    }{
        // TODO: Add test cases.
        {
            name: "connect",
            fields: fields{
                pipeline: op,
                client:   nil,
                forRead:  false,
                close:    false,
                response: nil,
            },
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            x := &XceiverClientStandalone{
                pipeline: tt.fields.pipeline,
                client:   tt.fields.client,
                forRead:  tt.fields.forRead,
                close:    tt.fields.close,
                response: tt.fields.response,
            }
            if err := x.Connect(); (err != nil) != tt.wantErr {
                t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestXceiverClientStandalone_GetPipeline(t *testing.T) {
    p, _ := NewPipelineOperator(NewPipeline())
    type fields struct {
        pipeline *PipelineOperator
        client   datanode.XceiverClientProtocolService_SendClient
        forRead  bool
        close    bool
        response chan *datanode.ContainerCommandResponseProto
    }
    tests := []struct {
        name   string
        fields fields
        want   *PipelineOperator
    }{
        // TODO: Add test cases.
        {
            name: "get pipeline",
            fields: fields{
                pipeline: p,
                client:   nil,
                forRead:  false,
                close:    false,
                response: nil,
            },
            want: p,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            x := &XceiverClientStandalone{
                pipeline: tt.fields.pipeline,
                client:   tt.fields.client,
                forRead:  tt.fields.forRead,
                close:    tt.fields.close,
                response: tt.fields.response,
            }
            if got := x.GetPipeline(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("GetPipeline() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestXceiverClientStandalone_GetReplicationType(t *testing.T) {
    p, _ := NewPipelineOperator(NewPipeline())
    type fields struct {
        pipeline *PipelineOperator
        client   datanode.XceiverClientProtocolService_SendClient
        forRead  bool
        close    bool
        response chan *datanode.ContainerCommandResponseProto
    }
    tests := []struct {
        name   string
        fields fields
        want   hdds.ReplicationType
    }{
        // TODO: Add test cases.
        {
            name: "get replication type",
            fields: fields{
                pipeline: p,
                client:   nil,
                forRead:  false,
                close:    false,
                response: nil,
            },
            want: hdds.ReplicationType_STAND_ALONE,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            x := &XceiverClientStandalone{
                pipeline: tt.fields.pipeline,
                client:   tt.fields.client,
                forRead:  tt.fields.forRead,
                close:    tt.fields.close,
                response: tt.fields.response,
            }
            if got := x.GetReplicationType(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("GetReplicationType() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestXceiverClientStandalone_IsForRead(t *testing.T) {
    type fields struct {
        pipeline *PipelineOperator
        client   datanode.XceiverClientProtocolService_SendClient
        forRead  bool
        close    bool
        response chan *datanode.ContainerCommandResponseProto
    }
    tests := []struct {
        name   string
        fields fields
        want   bool
    }{
        // TODO: Add test cases.
        {
            name: "is for read",
            fields: fields{
                pipeline: nil,
                client:   nil,
                forRead:  false,
                close:    false,
                response: nil,
            },
            want: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            x := &XceiverClientStandalone{
                pipeline: tt.fields.pipeline,
                client:   tt.fields.client,
                forRead:  tt.fields.forRead,
                close:    tt.fields.close,
                response: tt.fields.response,
            }
            if got := x.IsForRead(); got != tt.want {
                t.Errorf("IsForRead() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestXceiverClientStandalone_SetPipeline(t *testing.T) {
    p, _ := NewPipelineOperator(NewPipeline())
    type fields struct {
        pipeline *PipelineOperator
        client   datanode.XceiverClientProtocolService_SendClient
        forRead  bool
        close    bool
        response chan *datanode.ContainerCommandResponseProto
    }
    type args struct {
        pipeline *PipelineOperator
    }
    tests := []struct {
        name   string
        fields fields
        args   args
    }{
        // TODO: Add test cases.
        {
            name: "set pipeline",
            fields: fields{
                pipeline: nil,
                client:   nil,
                forRead:  false,
                close:    false,
                response: nil,
            },
            args: args{p},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            x := &XceiverClientStandalone{
                pipeline: tt.fields.pipeline,
                client:   tt.fields.client,
                forRead:  tt.fields.forRead,
                close:    tt.fields.close,
                response: tt.fields.response,
            }
            x.SetPipeline(tt.args.pipeline)
        })
    }
}

func TestXceiverClientStandalone_StandaloneReceive(t *testing.T) {
    c := NewMockXceiverClientProtocolService_SendClient()
    patch := gomonkey.ApplyMethod(reflect.TypeOf(c), "Recv", func(*MockXceiverClientProtocolService_SendClient) (*datanode.ContainerCommandResponseProto, error) {
        return &datanode.ContainerCommandResponseProto{}, nil
    })
    defer patch.Reset()

    c1 := NewMockXceiverClientProtocolService_SendClient()
    patch1 := gomonkey.ApplyMethod(reflect.TypeOf(c1), "Recv", func(*MockXceiverClientProtocolService_SendClient) (*datanode.ContainerCommandResponseProto, error) {
        return nil, io.EOF
    })
    defer patch1.Reset()
    type fields struct {
        pipeline *PipelineOperator
        client   datanode.XceiverClientProtocolService_SendClient
        forRead  bool
        close    bool
        response chan *datanode.ContainerCommandResponseProto
    }
    tests := []struct {
        name   string
        fields fields
    }{
        // TODO: Add test cases.
        {
            name: "standalone receive",
            fields: fields{
                pipeline: nil,
                client:   c,
                forRead:  false,
                close:    true,
                response: make(chan *datanode.ContainerCommandResponseProto, 1),
            },
        },
        {
            name: "standalone receive error",
            fields: fields{
                pipeline: nil,
                client:   c1,
                forRead:  false,
                close:    false,
                response: make(chan *datanode.ContainerCommandResponseProto, 1),
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            x := &XceiverClientStandalone{
                pipeline: tt.fields.pipeline,
                client:   tt.fields.client,
                forRead:  tt.fields.forRead,
                close:    tt.fields.close,
                response: tt.fields.response,
            }
            x.StandaloneReceive()
        })
        close(tt.fields.response)
    }
}

func TestXceiverClientStandalone_StandaloneSend(t *testing.T) {
    c := NewMockXceiverClientProtocolService_SendClient()
    patch := gomonkey.ApplyMethod(reflect.TypeOf(c), "Send", func(_ *MockXceiverClientProtocolService_SendClient, p *datanode.ContainerCommandRequestProto) error {
        return nil
    })
    defer patch.Reset()

    responseProto := &datanode.ContainerCommandResponseProto{
        Message: utils.PointString(""),
    }
    resp := make(chan *datanode.ContainerCommandResponseProto, 1)
    resp <- responseProto
    defer patch.Reset()
    type fields struct {
        pipeline *PipelineOperator
        client   datanode.XceiverClientProtocolService_SendClient
        forRead  bool
        close    bool
        response chan *datanode.ContainerCommandResponseProto
    }
    type args struct {
        req *datanode.ContainerCommandRequestProto
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    *datanode.ContainerCommandResponseProto
        wantErr bool
    }{
        // TODO: Add test cases.
        {
            name: "standalone send",
            fields: fields{
                pipeline: nil,
                client:   c,
                forRead:  false,
                close:    false,
                response: resp,
            },
            args:    args{},
            want:    &datanode.ContainerCommandResponseProto{Message: utils.PointString("")},
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            x := &XceiverClientStandalone{
                pipeline: tt.fields.pipeline,
                client:   tt.fields.client,
                forRead:  tt.fields.forRead,
                close:    tt.fields.close,
                response: tt.fields.response,
            }
            got, err := x.StandaloneSend(tt.args.req)
            if (err != nil) != tt.wantErr {
                t.Errorf("StandaloneSend() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("StandaloneSend() got = %v, want %v", got, tt.want)
            }
        })
    }
    close(resp)
}

func TestXceiverClientStandalone_StandaloneSend_Eror(t *testing.T) {
    c1 := NewMockXceiverClientProtocolService_SendClient()
    patch1 := gomonkey.ApplyMethod(reflect.TypeOf(c1), "Send", func(_ *MockXceiverClientProtocolService_SendClient, p *datanode.ContainerCommandRequestProto) error {
        return io.EOF
    })
    defer patch1.Reset()
    type fields struct {
        pipeline *PipelineOperator
        client   datanode.XceiverClientProtocolService_SendClient
        forRead  bool
        close    bool
        response chan *datanode.ContainerCommandResponseProto
    }
    type args struct {
        req *datanode.ContainerCommandRequestProto
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    *datanode.ContainerCommandResponseProto
        wantErr bool
    }{
        // TODO: Add test cases.
        // TODO: Add test cases.
        {
            name: "standalone send error",
            fields: fields{
                pipeline: nil,
                client:   c1,
                forRead:  false,
                close:    false,
                response: nil,
            },
            args:    args{},
            want:    nil,
            wantErr: true,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            x := &XceiverClientStandalone{
                pipeline: tt.fields.pipeline,
                client:   tt.fields.client,
                forRead:  tt.fields.forRead,
                close:    tt.fields.close,
                response: tt.fields.response,
            }
            got, err := x.StandaloneSend(tt.args.req)
            if (err != nil) != tt.wantErr {
                t.Errorf("StandaloneSend() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("StandaloneSend() got = %v, want %v", got, tt.want)
			}
		})
	}
}
