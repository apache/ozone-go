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
	"errors"
	"io"
	"github.com/apache/ozone-go/api/proto/datanode"
	"github.com/apache/ozone-go/api/proto/hdds"
	"github.com/apache/ozone-go/api/proto/ratis"
	"github.com/apache/ozone-go/api/utils"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"google.golang.org/grpc"
)

type MockRaftClientProtocolService_UnorderedClient struct {
	grpc.ClientStream
	Err error
}

func NewMockRaftClientProtocolService_UnorderedClient() ratis.RaftClientProtocolService_UnorderedClient {
	return &MockRaftClientProtocolService_UnorderedClient{}
}

func (x *MockRaftClientProtocolService_UnorderedClient) Send(m *ratis.RaftClientRequestProto) error {
	return x.Err
}

func (x *MockRaftClientProtocolService_UnorderedClient) Recv() (*ratis.RaftClientReplyProto, error) {
	return nil, io.EOF
}

func TestNewXceiverClientRatis(t *testing.T) {
	op, _ := NewPipelineOperator(NewPipeline())
	type args struct {
		pipeline *PipelineOperator
		forRead  bool
	}
	tests := []struct {
		name string
		args args
		want XceiverClientSpi
	}{
		// TODO: Add test cases.
		{
			name: "ratis client",
			args: args{
				pipeline: op,
				forRead:  false,
			},
			want: &XceiverClientRatis{
				pipeline: op,
				forRead:  false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewXceiverClientRatis(tt.args.pipeline, tt.args.forRead); got.GetPipeline() != tt.want.GetPipeline() {
				t.Errorf("NewXceiverClientRatis() = %v, want %v", got, tt.want)
			}
		})
	}
}

type MockRaftClientProtocolServiceClient struct {
	Err error
}

func NewMockRaftClientProtocolServiceClient() *MockRaftClientProtocolServiceClient {
	return &MockRaftClientProtocolServiceClient{
		Err: nil,
	}
}

func (c *MockRaftClientProtocolServiceClient) Ordered(ctx context.Context, opts ...grpc.CallOption) (ratis.RaftClientProtocolService_OrderedClient, error) {
	return nil, nil
}

func (c *MockRaftClientProtocolServiceClient) Unordered(ctx context.Context, opts ...grpc.CallOption) (ratis.RaftClientProtocolService_UnorderedClient, error) {
	return NewMockRaftClientProtocolService_UnorderedClient(), nil
}

func TestXceiverClientRatis_Connect(t *testing.T) {
	op, _ := NewPipelineOperator(NewPipeline())
	patch := gomonkey.ApplyFunc(grpc.DialContext, func(ctx context.Context, target string, opts ...grpc.DialOption) (conn *grpc.ClientConn, err error) {
		return &grpc.ClientConn{}, nil
	})
	defer patch.Reset()

	patch1 := gomonkey.ApplyFunc(ratis.NewRaftClientProtocolServiceClient, func(cc grpc.ClientConnInterface) ratis.RaftClientProtocolServiceClient {
		return NewMockRaftClientProtocolServiceClient()
	})
	defer patch1.Reset()

	type fields struct {
		pipeline      *PipelineOperator
		client        ratis.RaftClientProtocolService_UnorderedClient
		reply         chan *ratis.RaftClientReplyProto
		close         bool
		seqNum        uint64
		groupId       []byte
		commitInfoMap map[string]uint64
		forRead       bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "connect",
			fields: fields{
				pipeline:      op,
				client:        nil,
				reply:         nil,
				close:         false,
				seqNum:        0,
				groupId:       nil,
				commitInfoMap: nil,
				forRead:       false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XceiverClientRatis{pipeline: tt.fields.pipeline}
			if err := x.Connect(); (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestXceiverClientRatis_GetPipeline(t *testing.T) {
	op, _ := NewPipelineOperator(NewPipeline())
	type fields struct {
		pipeline      *PipelineOperator
		client        ratis.RaftClientProtocolService_UnorderedClient
		reply         chan *ratis.RaftClientReplyProto
		close         bool
		seqNum        uint64
		groupId       []byte
		commitInfoMap map[string]uint64
		forRead       bool
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
				pipeline:      op,
				client:        nil,
				reply:         nil,
				close:         false,
				seqNum:        0,
				groupId:       nil,
				commitInfoMap: nil,
				forRead:       false,
			},
			want: op,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XceiverClientRatis{
				pipeline:      tt.fields.pipeline,
				client:        tt.fields.client,
				reply:         tt.fields.reply,
				close:         tt.fields.close,
				seqNum:        tt.fields.seqNum,
				groupId:       tt.fields.groupId,
				commitInfoMap: tt.fields.commitInfoMap,
				forRead:       tt.fields.forRead,
			}
			if got := x.GetPipeline(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPipeline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXceiverClientRatis_GetReplicatedMinCommitIndex(t *testing.T) {
	type fields struct {
		pipeline      *PipelineOperator
		client        ratis.RaftClientProtocolService_UnorderedClient
		reply         chan *ratis.RaftClientReplyProto
		close         bool
		seqNum        uint64
		groupId       []byte
		commitInfoMap map[string]uint64
		forRead       bool
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
		{
			name: "get min commit index",
			fields: fields{
				pipeline:      nil,
				client:        nil,
				reply:         nil,
				close:         false,
				seqNum:        0,
				groupId:       nil,
				commitInfoMap: map[string]uint64{"1": 1},
				forRead:       false,
			},
			want: 1,
		},
		{
			name: "get min commit index",
			fields: fields{
				pipeline:      nil,
				client:        nil,
				reply:         nil,
				close:         false,
				seqNum:        0,
				groupId:       nil,
				commitInfoMap: map[string]uint64{"0": 0},
				forRead:       false,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XceiverClientRatis{
				pipeline:      tt.fields.pipeline,
				client:        tt.fields.client,
				reply:         tt.fields.reply,
				close:         tt.fields.close,
				seqNum:        tt.fields.seqNum,
				groupId:       tt.fields.groupId,
				commitInfoMap: tt.fields.commitInfoMap,
				forRead:       tt.fields.forRead,
			}
			if got := x.GetReplicatedMinCommitIndex(); got != tt.want {
				t.Errorf("GetReplicatedMinCommitIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXceiverClientRatis_GetReplicationType(t *testing.T) {
	op, _ := NewPipelineOperator(NewPipeline())
	type fields struct {
		pipeline      *PipelineOperator
		client        ratis.RaftClientProtocolService_UnorderedClient
		reply         chan *ratis.RaftClientReplyProto
		close         bool
		seqNum        uint64
		groupId       []byte
		commitInfoMap map[string]uint64
		forRead       bool
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
				pipeline:      op,
				client:        nil,
				reply:         nil,
				close:         false,
				seqNum:        0,
				groupId:       nil,
				commitInfoMap: nil,
				forRead:       false,
			},
			want: hdds.ReplicationType_STAND_ALONE,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XceiverClientRatis{
				pipeline:      tt.fields.pipeline,
				client:        tt.fields.client,
				reply:         tt.fields.reply,
				close:         tt.fields.close,
				seqNum:        tt.fields.seqNum,
				groupId:       tt.fields.groupId,
				commitInfoMap: tt.fields.commitInfoMap,
				forRead:       tt.fields.forRead,
			}
			if got := x.GetReplicationType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReplicationType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXceiverClientRatis_IsForRead(t *testing.T) {
	type fields struct {
		pipeline      *PipelineOperator
		client        ratis.RaftClientProtocolService_UnorderedClient
		reply         chan *ratis.RaftClientReplyProto
		close         bool
		seqNum        uint64
		groupId       []byte
		commitInfoMap map[string]uint64
		forRead       bool
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
				pipeline:      nil,
				client:        nil,
				reply:         nil,
				close:         false,
				seqNum:        0,
				groupId:       nil,
				commitInfoMap: nil,
				forRead:       false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XceiverClientRatis{
				pipeline:      tt.fields.pipeline,
				client:        tt.fields.client,
				reply:         tt.fields.reply,
				close:         tt.fields.close,
				seqNum:        tt.fields.seqNum,
				groupId:       tt.fields.groupId,
				commitInfoMap: tt.fields.commitInfoMap,
				forRead:       tt.fields.forRead,
			}
			if got := x.IsForRead(); got != tt.want {
				t.Errorf("IsForRead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXceiverClientRatis_RaftReceive(t *testing.T) {
	op, _ := NewPipelineOperator(NewPipeline())
	type fields struct {
		pipeline      *PipelineOperator
		client        ratis.RaftClientProtocolService_UnorderedClient
		reply         chan *ratis.RaftClientReplyProto
		close         bool
		seqNum        uint64
		groupId       []byte
		commitInfoMap map[string]uint64
		forRead       bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{
			name: "raft receive",
			fields: fields{
				pipeline:      op,
				client:        NewMockRaftClientProtocolService_UnorderedClient(),
				reply:         make(chan *ratis.RaftClientReplyProto),
				close:         true,
				seqNum:        0,
				groupId:       nil,
				commitInfoMap: nil,
				forRead:       false,
			},
		},
		{
			name: "raft receive eof",
			fields: fields{
				pipeline:      op,
				client:        NewMockRaftClientProtocolService_UnorderedClient(),
				reply:         make(chan *ratis.RaftClientReplyProto),
				close:         false,
				seqNum:        0,
				groupId:       nil,
				commitInfoMap: nil,
				forRead:       false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XceiverClientRatis{
				pipeline:      tt.fields.pipeline,
				client:        tt.fields.client,
				reply:         tt.fields.reply,
				close:         tt.fields.close,
				seqNum:        tt.fields.seqNum,
				groupId:       tt.fields.groupId,
				commitInfoMap: tt.fields.commitInfoMap,
				forRead:       tt.fields.forRead,
			}
			x.RaftReceive()
			close(x.reply)
		})
	}
}

func TestXceiverClientRatis_RaftSend(t *testing.T) {
	op, _ := NewPipelineOperator(NewPipeline())
	cli := NewMockRaftClientProtocolService_UnorderedClient()
	patch := gomonkey.ApplyMethod(reflect.TypeOf(cli), "Send", func(_ *MockRaftClientProtocolService_UnorderedClient, _ *ratis.RaftClientRequestProto) error {
		return nil
	})
	defer patch.Reset()
	replyProto := &ratis.RaftClientReplyProto{RpcReply: &ratis.RaftRpcReplyProto{
		RequestorId: nil,
		ReplyId:     nil,
		RaftGroupId: nil,
		CallId:      0,
		Success:     false,
	}}
	reply := make(chan *ratis.RaftClientReplyProto, 1)
	reply <- replyProto
	type fields struct {
		pipeline      *PipelineOperator
		client        ratis.RaftClientProtocolService_UnorderedClient
		reply         chan *ratis.RaftClientReplyProto
		close         bool
		seqNum        uint64
		groupId       []byte
		commitInfoMap map[string]uint64
		forRead       bool
	}
	type args struct {
		req *datanode.ContainerCommandRequestProto
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ratis.RaftClientReplyProto
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "raft send",
			fields: fields{
				pipeline:      op,
				client:        cli,
				reply:         reply,
				close:         false,
				seqNum:        0,
				groupId:       nil,
				commitInfoMap: nil,
				forRead:       false,
			},
			args: args{req: &datanode.ContainerCommandRequestProto{
				CmdType:     datanode.Type_WriteChunk.Enum(),
				TraceID:     utils.PointString("0"),
				ContainerID: utils.PointInt64(0),
				WriteChunk: &datanode.WriteChunkRequestProto{
					Data:    []byte{0},
					BlockID: NewDatanodeBlockId(),
					ChunkData: &datanode.ChunkInfo{
						ChunkName: utils.PointString("1"),
						Offset:    utils.PointUint64(0),
						Len:       utils.PointUint64(1),
						ChecksumData: &datanode.ChecksumData{
							Type:             datanode.ChecksumType_NONE.Enum(),
							BytesPerChecksum: utils.PointUint32(1),
						},
					},
				},
				DatanodeUuid: utils.PointString("0"),
			}},
			want:    replyProto,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XceiverClientRatis{
				pipeline:      tt.fields.pipeline,
				client:        tt.fields.client,
				reply:         tt.fields.reply,
				close:         tt.fields.close,
				seqNum:        tt.fields.seqNum,
				groupId:       tt.fields.groupId,
				commitInfoMap: tt.fields.commitInfoMap,
				forRead:       tt.fields.forRead,
			}
			got, err := x.RaftSend(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RaftSend() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RaftSend() got = %v, want %v", got, tt.want)
			}
		})
	}
	close(reply)
}

func TestXceiverClientRatis_SetPipeline(t *testing.T) {
	op, _ := NewPipelineOperator(NewPipeline())
	type fields struct {
		pipeline      *PipelineOperator
		client        ratis.RaftClientProtocolService_UnorderedClient
		reply         chan *ratis.RaftClientReplyProto
		close         bool
		seqNum        uint64
		groupId       []byte
		commitInfoMap map[string]uint64
		forRead       bool
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
			name:   "set pipe line",
			fields: fields{},
			args:   args{pipeline: op},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XceiverClientRatis{
				pipeline:      tt.fields.pipeline,
				client:        tt.fields.client,
				reply:         tt.fields.reply,
				close:         tt.fields.close,
				seqNum:        tt.fields.seqNum,
				groupId:       tt.fields.groupId,
				commitInfoMap: tt.fields.commitInfoMap,
				forRead:       tt.fields.forRead,
			}
			x.SetPipeline(tt.args.pipeline)
			if !reflect.DeepEqual(x.pipeline, op) {
				t.Errorf("set pipiline () got = %v, want %v", x.pipeline, op)
			}
		})
	}
}

func TestXceiverClientRatis_updateCommitInfosMap(t *testing.T) {
	type fields struct {
		pipeline      *PipelineOperator
		client        ratis.RaftClientProtocolService_UnorderedClient
		reply         chan *ratis.RaftClientReplyProto
		close         bool
		seqNum        uint64
		groupId       []byte
		commitInfoMap map[string]uint64
		forRead       bool
	}
	type args struct {
		commitInfoProtos []*ratis.CommitInfoProto
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "update commit index",
			fields: fields{
				commitInfoMap: map[string]uint64{},
			},
			args: args{commitInfoProtos: []*ratis.CommitInfoProto{
				{Server: &ratis.RaftPeerProto{Address: "1"}, CommitIndex: 1},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XceiverClientRatis{
				pipeline:      tt.fields.pipeline,
				client:        tt.fields.client,
				reply:         tt.fields.reply,
				close:         tt.fields.close,
				seqNum:        tt.fields.seqNum,
				groupId:       tt.fields.groupId,
				commitInfoMap: tt.fields.commitInfoMap,
				forRead:       tt.fields.forRead,
			}
			x.updateCommitInfosMap(tt.args.commitInfoProtos)
		})
	}
}

func TestXceiverClientReply_GetLogIndex(t *testing.T) {
	type fields struct {
		LogIndex  uint64
		Reply     *datanode.ContainerCommandResponseProto
		Datanodes []string
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
		{
			name:   "get log index",
			fields: fields{LogIndex: uint64(1)},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XceiverClientReply{
				LogIndex:  tt.fields.LogIndex,
				Reply:     tt.fields.Reply,
				Datanodes: tt.fields.Datanodes,
			}
			if got := x.GetLogIndex(); got != tt.want {
				t.Errorf("GetLogIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toMessage(t *testing.T) {
	type args struct {
		req *datanode.ContainerCommandRequestProto
	}
	tests := []struct {
		name string
		args args
		want *ratis.ClientMessageEntryProto
	}{
		// TODO: Add test cases.
		{
			name: "to message WriteChunk",
			args: args{req: &datanode.ContainerCommandRequestProto{
				CmdType:     datanode.Type_WriteChunk.Enum(),
				TraceID:     utils.PointString("0"),
				ContainerID: utils.PointInt64(0),
				WriteChunk: &datanode.WriteChunkRequestProto{
					Data:    []byte{0},
					BlockID: NewDatanodeBlockId(),
					ChunkData: &datanode.ChunkInfo{
						ChunkName: utils.PointString("1"),
						Offset:    utils.PointUint64(0),
						Len:       utils.PointUint64(1),
						ChecksumData: &datanode.ChecksumData{
							Type:             datanode.ChecksumType_NONE.Enum(),
							BytesPerChecksum: utils.PointUint32(1),
						},
					},
				},
				DatanodeUuid: utils.PointString("0"),
			}},
			want: &ratis.ClientMessageEntryProto{Content: []byte{0, 0, 0, 36, 8, 12, 18, 1, 48, 24, 0, 34, 1, 48, 138, 1, 23, 10, 6, 8, 1, 16, 2, 24, 3, 18, 13, 10, 1, 49, 16, 0, 24, 1, 42, 4, 8, 1, 16, 1, 0}},
		},
		{
			name: "to message PutSmallFile",
			args: args{req: &datanode.ContainerCommandRequestProto{
				CmdType:     datanode.Type_PutSmallFile.Enum(),
				TraceID:     utils.PointString("0"),
				ContainerID: utils.PointInt64(0),
				PutSmallFile: &datanode.PutSmallFileRequestProto{
					Block: &datanode.PutBlockRequestProto{BlockData: &datanode.BlockData{BlockID: NewDatanodeBlockId()}},
					ChunkInfo: &datanode.ChunkInfo{
						ChunkName: utils.PointString("1"),
						Offset:    utils.PointUint64(0),
						Len:       utils.PointUint64(1),
						ChecksumData: &datanode.ChecksumData{
							Type:             datanode.ChecksumType_NONE.Enum(),
							BytesPerChecksum: utils.PointUint32(1),
						},
					},
					Data: []byte{0},
				},
				DatanodeUuid: utils.PointString("0"),
			}},
			want: &ratis.ClientMessageEntryProto{Content: []byte{0, 0, 0, 40, 8, 15, 18, 1, 48, 24, 0, 34, 1, 48, 162, 1, 27, 10, 10, 10, 8, 10, 6, 8, 1, 16, 2, 24, 3, 18, 13, 10, 1, 49, 16, 0, 24, 1, 42, 4, 8, 1, 16, 1, 0}},
		},
		{
			name: "to message Other",
			args: args{req: &datanode.ContainerCommandRequestProto{
				CmdType:     datanode.Type_PutBlock.Enum(),
				TraceID:     utils.PointString("0"),
				ContainerID: utils.PointInt64(0),
				PutBlock: &datanode.PutBlockRequestProto{
					BlockData: &datanode.BlockData{BlockID: NewDatanodeBlockId()},
					Eof:       utils.PointBool(true),
				},
				DatanodeUuid: utils.PointString("0"),
			}},
			want: &ratis.ClientMessageEntryProto{Content: []byte{0, 0, 0, 24, 8, 6, 18, 1, 48, 24, 0, 34, 1, 48, 98, 12, 10, 8, 10, 6, 8, 1, 16, 2, 24, 3, 16, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toMessage(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXceiverClientRatis_Close(t *testing.T) {
	client := NewMockRaftClientProtocolService_UnorderedClient()
	patch := gomonkey.ApplyMethod(reflect.TypeOf(client), "CloseSend", func(*MockRaftClientProtocolService_UnorderedClient) error {
		return errors.New("test error")
	})
	defer patch.Reset()
	type fields struct {
		pipeline      *PipelineOperator
		client        ratis.RaftClientProtocolService_UnorderedClient
		reply         chan *ratis.RaftClientReplyProto
		close         bool
		seqNum        uint64
		groupId       []byte
		commitInfoMap map[string]uint64
		forRead       bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{
			name: "close client",
			fields: fields{
				pipeline:      nil,
				client:        nil,
				reply:         nil,
				close:         true,
				seqNum:        0,
				groupId:       nil,
				commitInfoMap: nil,
				forRead:       false,
			},
		},
		{
			name: "close err client",
			fields: fields{
				pipeline:      nil,
				client:        client,
				reply:         nil,
				close:         false,
				seqNum:        0,
				groupId:       nil,
				commitInfoMap: nil,
				forRead:       false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XceiverClientRatis{
				pipeline:      tt.fields.pipeline,
				client:        tt.fields.client,
				reply:         tt.fields.reply,
				close:         tt.fields.close,
				seqNum:        tt.fields.seqNum,
				groupId:       tt.fields.groupId,
				commitInfoMap: tt.fields.commitInfoMap,
				forRead:       tt.fields.forRead,
			}
			x.Close()
		})
	}
}

func TestXceiverClientRatis_WatchForCommit(t *testing.T) {
	op, _ := NewPipelineOperator(NewPipeline())
	cli := NewMockRaftClientProtocolService_UnorderedClient()
	patch := gomonkey.ApplyMethod(reflect.TypeOf(cli), "Send", func(_ *MockRaftClientProtocolService_UnorderedClient, _ *ratis.RaftClientRequestProto) error {
		return nil
	})
	defer patch.Reset()
	replyProto := &ratis.RaftClientReplyProto{RpcReply: &ratis.RaftRpcReplyProto{
		RequestorId: nil,
		ReplyId:     nil,
		RaftGroupId: nil,
		CallId:      0,
		Success:     false,
	}}
	reply := make(chan *ratis.RaftClientReplyProto, 1)
	reply <- replyProto
	type fields struct {
		pipeline      *PipelineOperator
		client        ratis.RaftClientProtocolService_UnorderedClient
		reply         chan *ratis.RaftClientReplyProto
		close         bool
		seqNum        uint64
		groupId       []byte
		commitInfoMap map[string]uint64
		forRead       bool
	}
	type args struct {
		index uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *XceiverClientReply
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "commit index 0",
			fields:  fields{},
			args:    args{},
			want:    &XceiverClientReply{},
			wantErr: false,
		},
		{
			name: "commit index 1",
			fields: fields{
				pipeline:      op,
				client:        cli,
				reply:         reply,
				close:         false,
				seqNum:        0,
				groupId:       nil,
				commitInfoMap: map[string]uint64{},
				forRead:       false,
			},
			args: args{index: 1},
			want: &XceiverClientReply{
				LogIndex:  1,
				Reply:     &datanode.ContainerCommandResponseProto{},
				Datanodes: []string{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XceiverClientRatis{
				pipeline:      tt.fields.pipeline,
				client:        tt.fields.client,
				reply:         tt.fields.reply,
				close:         tt.fields.close,
				seqNum:        tt.fields.seqNum,
				groupId:       tt.fields.groupId,
				commitInfoMap: tt.fields.commitInfoMap,
				forRead:       tt.fields.forRead,
			}
			got, err := x.WatchForCommit(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("WatchForCommit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WatchForCommit() got = %v, want %v", got, tt.want)
			}
		})
	}
	close(reply)
}
