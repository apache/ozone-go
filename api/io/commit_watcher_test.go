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
    dnClient "github.com/apache/ozone-go/api/datanodeclient"
    "github.com/apache/ozone-go/api/proto/hdds"
    "github.com/apache/ozone-go/api/utils"
    "reflect"
    "testing"

    "github.com/agiledragon/gomonkey/v2"
)

func TestCommitWatcher_UpdateCommitInfoMap(t *testing.T) {
    type fields struct {
        xceiverClient              dnClient.XceiverClientSpi
        CommitIndex2flushedDataMap map[uint64]int
    }
    type args struct {
        commitIndex       uint64
        flushedChunkIndex int
    }
    tests := []struct {
        name   string
        fields fields
        args   args
    }{
        // TODO: Add test cases.
        {
            name: "UpdateCommitInfoMap",
            fields: fields{
                xceiverClient:              nil,
                CommitIndex2flushedDataMap: map[uint64]int{},
            },
            args: args{
                commitIndex:       0,
                flushedChunkIndex: 0,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            commitWatcher := &CommitWatcher{
                xceiverClient:              tt.fields.xceiverClient,
                CommitIndex2flushedDataMap: tt.fields.CommitIndex2flushedDataMap,
            }
            commitWatcher.UpdateCommitInfoMap(tt.args.commitIndex, tt.args.flushedChunkIndex)
        })
    }
}

func NewPipeline() *hdds.Pipeline {
    return &hdds.Pipeline{
        Members: []*hdds.DatanodeDetailsProto{{
            Uuid:                   utils.PointString("LeaderID"),
            IpAddress:              utils.PointString("0.0.0.0"),
            HostName:               utils.PointString("localhost"),
            Ports:                  []*hdds.Port{{Name: utils.PointString("ratis"), Value: utils.PointUint32(9070)}},
            CertSerialId:           nil,
            NetworkName:            nil,
            NetworkLocation:        nil,
            PersistedOpState:       nil,
            PersistedOpStateExpiry: nil,
            Uuid128:                nil,
        }},
        State:  nil,
        Type:   nil,
        Factor: nil,
        Id: &hdds.PipelineID{
            Id: utils.PointString("id"),
            Uuid128: &hdds.UUID{
                MostSigBits:  nil,
                LeastSigBits: nil,
            },
        },
        LeaderID:          utils.PointString("LeaderID"),
        MemberOrders:      []uint32{0},
        CreationTimeStamp: nil,
        SuggestedLeaderID: nil,
        LeaderID128:       nil,
    }
}
func TestCommitWatcher_WatchOnFirstIndex(t *testing.T) {
    patch := gomonkey.ApplyFunc(dnClient.NewXceiverClientRatis, func(pipeline *dnClient.PipelineOperator, forRead bool) dnClient.XceiverClientSpi {
        return &dnClient.XceiverClientRatis{}
    })
    defer patch.Reset()
    c := dnClient.NewXceiverClientRatis(nil, false)
    patch1 := gomonkey.ApplyMethod(reflect.TypeOf(c), "WatchForCommit", func(_ *dnClient.XceiverClientRatis, commitIndex uint64) (*dnClient.XceiverClientReply, error) {
        return &dnClient.XceiverClientReply{LogIndex: 1}, nil
    })
    defer patch1.Reset()
    type fields struct {
        xceiverClient              dnClient.XceiverClientSpi
        CommitIndex2flushedDataMap map[uint64]int
    }
    tests := []struct {
        name    string
        fields  fields
        want    *dnClient.XceiverClientReply
        wantErr bool
    }{
        // TODO: Add test cases.
        {
            name: "WatchOnFirstIndex",
            fields: fields{
                CommitIndex2flushedDataMap: map[uint64]int{},
            },
            want:    &dnClient.XceiverClientReply{LogIndex: 0},
            wantErr: false,
        },
        {
            name: "WatchOnFirstIndex",
            fields: fields{
                CommitIndex2flushedDataMap: map[uint64]int{0: 0},
            },
            want:    &dnClient.XceiverClientReply{LogIndex: 1},
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            commitWatcher := &CommitWatcher{
                xceiverClient:              c,
                CommitIndex2flushedDataMap: tt.fields.CommitIndex2flushedDataMap,
            }
            got, err := commitWatcher.WatchOnFirstIndex()
            if (err != nil) != tt.wantErr {
                t.Errorf("WatchOnFirstIndex() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("WatchOnFirstIndex() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestCommitWatcher_WatchOnLastIndex(t *testing.T) {
    patch := gomonkey.ApplyFunc(dnClient.NewXceiverClientRatis, func(pipeline *dnClient.PipelineOperator, forRead bool) dnClient.XceiverClientSpi {
        return &dnClient.XceiverClientRatis{}
    })
    defer patch.Reset()
    c := dnClient.NewXceiverClientRatis(nil, false)
    patch1 := gomonkey.ApplyMethod(reflect.TypeOf(c), "WatchForCommit", func(_ *dnClient.XceiverClientRatis, commitIndex uint64) (*dnClient.XceiverClientReply, error) {
        return &dnClient.XceiverClientReply{LogIndex: 1}, nil
    })
    defer patch1.Reset()
    type fields struct {
        xceiverClient              dnClient.XceiverClientSpi
        CommitIndex2flushedDataMap map[uint64]int
    }
    tests := []struct {
        name    string
        fields  fields
        want    *dnClient.XceiverClientReply
        wantErr bool
    }{
        // TODO: Add test cases.
        {
            name: "WatchOnLastIndex",
            fields: fields{
                CommitIndex2flushedDataMap: map[uint64]int{},
            },
            want:    &dnClient.XceiverClientReply{LogIndex: 0},
            wantErr: false,
        },
        {
            name: "WatchOnLastIndex",
            fields: fields{
                CommitIndex2flushedDataMap: map[uint64]int{1: 1},
            },
            want:    &dnClient.XceiverClientReply{LogIndex: 1},
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            commitWatcher := &CommitWatcher{
                xceiverClient:              c,
                CommitIndex2flushedDataMap: tt.fields.CommitIndex2flushedDataMap,
            }
            got, err := commitWatcher.WatchOnLastIndex()
            if (err != nil) != tt.wantErr {
                t.Errorf("WatchOnLastIndex() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("WatchOnLastIndex() got = %v, want %v", got, tt.want)
			}
		})
	}
}
