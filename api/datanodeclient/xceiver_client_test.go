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
    "github.com/apache/ozone-go/api/proto/hdds"
    "reflect"
    "testing"

    "github.com/agiledragon/gomonkey/v2"
)

func TestGetXceiverClientManagerInstance(t *testing.T) {
    manager := GetXceiverClientManagerInstance()
    tests := []struct {
        name string
        want *XceiverClientManager
    }{
        // TODO: Add test cases.
        {
            name: "new manager",
            want: manager,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := GetXceiverClientManagerInstance(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("GetXceiverClientManagerInstance() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestXceiverClientManager_GetClient_Ratis(t *testing.T) {
    ratisPipeline := NewPipeline()
    ratisPipeline.Type = hdds.ReplicationType_RATIS.Enum()
    ratisOp, _ := NewPipelineOperator(ratisPipeline)
    errPipeline := NewPipeline()
    errPipeline.Type = hdds.ReplicationType_CHAINED.Enum()
    errOp, _ := NewPipelineOperator(errPipeline)
    patch := gomonkey.ApplyFunc(NewXceiverClientRatis, func(pipeline *PipelineOperator, forRead bool) XceiverClientSpi {
        return NewMockXceiverClient()
    })
    defer patch.Reset()
    type fields struct {
        ClientCache       map[string]XceiverClientSpi
        TopologyAwareRead bool
    }
    type args struct {
        pipeline *PipelineOperator
        read     bool
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    XceiverClientSpi
        wantErr bool
    }{
        // TODO: Add test cases.
        {
            name: "get ratis client",
            fields: fields{
                ClientCache:       make(map[string]XceiverClientSpi, 0),
                TopologyAwareRead: false,
            },
            args: args{
                pipeline: ratisOp,
                read:     false,
            },
            want:    NewMockXceiverClient(),
            wantErr: false,
        },
        {
            name: "get error client",
            fields: fields{
                ClientCache:       nil,
                TopologyAwareRead: false,
            },
            args: args{
                pipeline: errOp,
                read:     false,
            },
            want:    nil,
            wantErr: true,
        },
        {
            name: "get cache client",
            fields: fields{
                ClientCache:       map[string]XceiverClientSpi{"idRATIS": NewMockXceiverClient()},
                TopologyAwareRead: false,
            },
            args: args{
                pipeline: ratisOp,
                read:     false,
            },
            want:    NewMockXceiverClient(),
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            x := &XceiverClientManager{
                ClientCache:       tt.fields.ClientCache,
                TopologyAwareRead: tt.fields.TopologyAwareRead,
            }
            got, err := x.GetClient(tt.args.pipeline, tt.args.read)
            if (err != nil) != tt.wantErr {
                t.Errorf("GetClient() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("GetClient() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestXceiverClientManager_GetClient_Standalone(t *testing.T) {

    standalonePipeline := NewPipeline()
    standalonePipeline.Type = hdds.ReplicationType_STAND_ALONE.Enum()
    standaloneOp, _ := NewPipelineOperator(standalonePipeline)
    errPipeline := NewPipeline()
    errPipeline.Type = hdds.ReplicationType_CHAINED.Enum()
    errOp, _ := NewPipelineOperator(errPipeline)
    patch := gomonkey.ApplyFunc(NewXceiverClientStandalone, func(pipeline *PipelineOperator, forRead bool) XceiverClientSpi {
        return NewMockXceiverClient()
    })
    defer patch.Reset()
    type fields struct {
        ClientCache       map[string]XceiverClientSpi
        TopologyAwareRead bool
    }
    type args struct {
        pipeline *PipelineOperator
        read     bool
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    XceiverClientSpi
        wantErr bool
    }{
        {
            name: "get standalone client",
            fields: fields{
                ClientCache:       make(map[string]XceiverClientSpi, 0),
                TopologyAwareRead: false,
            },
            args: args{
                pipeline: standaloneOp,
                read:     false,
            },
            want:    NewMockXceiverClient(),
            wantErr: false,
        },
        {
            name: "get error client",
            fields: fields{
                ClientCache:       nil,
                TopologyAwareRead: false,
            },
            args: args{
                pipeline: errOp,
                read:     false,
            },
            want:    nil,
            wantErr: true,
        },
        {
            name: "get cache client",
            fields: fields{
                ClientCache:       map[string]XceiverClientSpi{"idSTAND_ALONE": NewMockXceiverClient()},
                TopologyAwareRead: false,
            },
            args: args{
                pipeline: standaloneOp,
                read:     false,
            },
            want:    NewMockXceiverClient(),
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            x := &XceiverClientManager{
                ClientCache:       tt.fields.ClientCache,
                TopologyAwareRead: tt.fields.TopologyAwareRead,
            }
            got, err := x.GetClient(tt.args.pipeline, tt.args.read)
            if (err != nil) != tt.wantErr {
                t.Errorf("GetClient() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("GetClient() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestXceiverClientManager_RemoveClient(t *testing.T) {
    ratisPipeline := NewPipeline()
    ratisPipeline.Type = hdds.ReplicationType_RATIS.Enum()
    ratisOp, _ := NewPipelineOperator(ratisPipeline)
    type fields struct {
        ClientCache       map[string]XceiverClientSpi
        TopologyAwareRead bool
    }
    type args struct {
        pipeline *PipelineOperator
        read     bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
    }{
        {
            name: "remove client",
            fields: fields{
                ClientCache:       map[string]XceiverClientSpi{"idRATIS": NewMockXceiverClient()},
                TopologyAwareRead: false,
            },
            args: args{
                pipeline: ratisOp,
                read:     false,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            x := &XceiverClientManager{
                ClientCache:       tt.fields.ClientCache,
                TopologyAwareRead: tt.fields.TopologyAwareRead,
            }
            x.RemoveClient(tt.args.pipeline, tt.args.read)
        })
    }
}

func TestXceiverClientManager_getPipelineCacheKey(t *testing.T) {
    ratisPipeline := NewPipeline()
    ratisPipeline.Type = hdds.ReplicationType_RATIS.Enum()
    ratisOp, _ := NewPipelineOperator(ratisPipeline)
    type fields struct {
        ClientCache       map[string]XceiverClientSpi
        TopologyAwareRead bool
    }
    type args struct {
        pipeline *PipelineOperator
        forRead  bool
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   string
    }{
        // TODO: Add test cases.
        {
            name: "cache key",
            fields: fields{
                ClientCache:       map[string]XceiverClientSpi{"idRATIS": NewMockXceiverClient()},
                TopologyAwareRead: false,
            },
            args: args{
                pipeline: ratisOp,
                forRead:  false,
            },
            want: "idRATIS",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            x := &XceiverClientManager{
                ClientCache:       tt.fields.ClientCache,
                TopologyAwareRead: tt.fields.TopologyAwareRead,
            }
            if got := x.getPipelineCacheKey(tt.args.pipeline, tt.args.forRead); got != tt.want {
                t.Errorf("getPipelineCacheKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
