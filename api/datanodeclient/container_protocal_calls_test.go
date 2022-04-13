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
	"github.com/apache/ozone-go/api/common"
	"github.com/apache/ozone-go/api/proto/datanode"
	"github.com/apache/ozone-go/api/proto/hdds"
	"github.com/apache/ozone-go/api/proto/ratis"
	"github.com/apache/ozone-go/api/utils"
	"reflect"
	"testing"
)

type MockXceiverClient struct {
	Err              error
	Boo              bool
	PipelineOperator *PipelineOperator
	ReplicaType      hdds.ReplicationType
	Reply            *ratis.RaftClientReplyProto
	Response         *datanode.ContainerCommandResponseProto
	ClientReply      *XceiverClientReply
	u                uint64
}

func (c *MockXceiverClient) Connect() error {
	return c.Err
}
func (c *MockXceiverClient) IsForRead() bool {
	return c.Boo
}
func (c *MockXceiverClient) GetPipeline() *PipelineOperator {
	return c.PipelineOperator
}
func (c *MockXceiverClient) GetReplicationType() hdds.ReplicationType {
	return c.GetPipeline().GetPipeline().GetType()
}
func (c *MockXceiverClient) SetPipeline(pipeline *PipelineOperator) {
	c.PipelineOperator = pipeline
}
func (c *MockXceiverClient) RaftSend(req *datanode.ContainerCommandRequestProto) (*ratis.RaftClientReplyProto, error) {
	return c.Reply, c.Err
}
func (c *MockXceiverClient) StandaloneSend(req *datanode.ContainerCommandRequestProto) (*datanode.ContainerCommandResponseProto, error) {
	return c.Response, c.Err
}
func (c *MockXceiverClient) WatchForCommit(index uint64) (*XceiverClientReply, error) {
	return c.ClientReply, c.Err
}
func (c *MockXceiverClient) GetReplicatedMinCommitIndex() uint64 {
	return c.u
}
func (c *MockXceiverClient) Close() {
	return
}

func NewMockXceiverClient() *MockXceiverClient {
	return &MockXceiverClient{}
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

func NewDatanodeBlockId() *datanode.DatanodeBlockID {
	return &datanode.DatanodeBlockID{
		ContainerID:           utils.PointInt64(1),
		LocalID:               utils.PointInt64(2),
		BlockCommitSequenceId: utils.PointUint64(3),
	}
}

func TestGetBlock(t *testing.T) {
	xceiverClient := NewMockXceiverClient()
	pipeline := NewPipeline()
	xceiverClient.PipelineOperator, _ = NewPipelineOperator(pipeline)
	xceiverClient.Response = &datanode.ContainerCommandResponseProto{
		GetBlock: &datanode.GetBlockResponseProto{},
	}
	type args struct {
		spi     XceiverClientSpi
		blockId *datanode.DatanodeBlockID
	}
	tests := []struct {
		name    string
		args    args
		want    *datanode.GetBlockResponseProto
		wantErr bool
	}{
		{
			name: "get block",
			args: args{spi: xceiverClient,
				blockId: NewDatanodeBlockId(),
			},
			want:    &datanode.GetBlockResponseProto{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBlock(tt.args.spi, tt.args.blockId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBlock() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPutBlock(t *testing.T) {
	xceiverClient := NewMockXceiverClient()
	pipeline := NewPipeline()
	xceiverClient.PipelineOperator, _ = NewPipelineOperator(pipeline)
	xceiverClient.Reply = &ratis.RaftClientReplyProto{}
	type args struct {
		spi       XceiverClientSpi
		blockData *datanode.BlockData
		eof       bool
	}
	tests := []struct {
		name    string
		args    args
		want    *ratis.RaftClientReplyProto
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "put block",
			args: args{
				spi:       xceiverClient,
				blockData: &datanode.BlockData{BlockID: NewDatanodeBlockId()},
				eof:       false,
			},
			want:    &ratis.RaftClientReplyProto{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PutBlock(tt.args.spi, tt.args.blockData, tt.args.eof)
			if (err != nil) != tt.wantErr {
				t.Errorf("PutBlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PutBlock() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadChunk(t *testing.T) {
	xceiverClient := NewMockXceiverClient()
	pipeline := NewPipeline()
	xceiverClient.PipelineOperator, _ = NewPipelineOperator(pipeline)
	xceiverClient.Response = &datanode.ContainerCommandResponseProto{ReadChunk: &datanode.ReadChunkResponseProto{
		BlockID:      nil,
		ChunkData:    nil,
		ResponseData: nil,
	}}
	type args struct {
		spi       XceiverClientSpi
		chunkInfo *datanode.ChunkInfo
		blockId   *datanode.DatanodeBlockID
	}
	tests := []struct {
		name    string
		args    args
		want    *datanode.ReadChunkResponseProto
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "read chunk",
			args: args{
				spi: xceiverClient,
				chunkInfo: &datanode.ChunkInfo{
					ChunkName: utils.PointString("chunkname"),
					Offset:    utils.PointUint64(0),
					Len:       utils.PointUint64(0),
					ChecksumData: &datanode.ChecksumData{
						Type: common.ChecksumTypeFromName("none").Enum(),
					},
				},
				blockId: NewDatanodeBlockId(),
			},
			want:    &datanode.ReadChunkResponseProto{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadChunk(tt.args.spi, tt.args.chunkInfo, tt.args.blockId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadChunk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadChunk() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteChunk(t *testing.T) {
	xceiverClient := NewMockXceiverClient()
	pipeline := NewPipeline()
	xceiverClient.PipelineOperator, _ = NewPipelineOperator(pipeline)
	xceiverClient.Reply = &ratis.RaftClientReplyProto{}
	type args struct {
		spi       XceiverClientSpi
		chunkInfo *datanode.ChunkInfo
		blockId   *datanode.DatanodeBlockID
		data      []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *ratis.RaftClientReplyProto
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "write chunk",
			args: args{
				spi:     xceiverClient,
				blockId: NewDatanodeBlockId(),
			},
			want:    &ratis.RaftClientReplyProto{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WriteChunk(tt.args.spi, tt.args.chunkInfo, tt.args.blockId, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteChunk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WriteChunk() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteSmallFile(t *testing.T) {
	xceiverClient := NewMockXceiverClient()
	pipeline := NewPipeline()
	xceiverClient.PipelineOperator, _ = NewPipelineOperator(pipeline)
	xceiverClient.Response = &datanode.ContainerCommandResponseProto{PutSmallFile: &datanode.PutSmallFileResponseProto{}}
	type args struct {
		spi     XceiverClientSpi
		blockId *datanode.DatanodeBlockID
		data    []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *datanode.PutSmallFileResponseProto
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "write small file",
			args: args{
				spi:     xceiverClient,
				blockId: NewDatanodeBlockId(),
				data:    []byte{0},
			},
			want:    &datanode.PutSmallFileResponseProto{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WriteSmallFile(tt.args.spi, tt.args.blockId, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteSmallFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WriteSmallFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
