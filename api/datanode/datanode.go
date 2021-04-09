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
package datanode

import (
	"context"
	"errors"
	"fmt"
	dnapi "github.com/apache/ozone-go/api/proto/datanode"
	"github.com/apache/ozone-go/api/proto/hdds"
	"github.com/apache/ozone-go/api/proto/ratis"
	"google.golang.org/grpc"
	"io"
	"strconv"
)

type ChunkInfo struct {
	Name   string
	Offset uint64
	Len    uint64
}

type DatanodeClient struct {
	ratisClient        *ratis.RaftClientProtocolService_UnorderedClient
	ratisReceiver      chan ratis.RaftClientReplyProto
	standaloneClient   *dnapi.XceiverClientProtocolService_SendClient
	standaloneReceiver chan dnapi.ContainerCommandResponseProto

	ctx             context.Context
	datanodes       []*hdds.DatanodeDetailsProto
	currentDatanode hdds.DatanodeDetailsProto
	grpcConnection  *grpc.ClientConn
	pipelineId      *hdds.PipelineID
	memberIndex     int
}

func (dn *DatanodeClient) GetCurrentDnUUid() *string {
	uid := dn.currentDatanode.GetUuid()
	return &uid
}

func (dnClient *DatanodeClient) connectToNext() error {
	if dnClient.grpcConnection != nil {
		dnClient.grpcConnection.Close()
	}
	dnClient.memberIndex = dnClient.memberIndex + 1
	if dnClient.memberIndex == len(dnClient.datanodes) {
		dnClient.memberIndex = 0
	}
	selectedDatanode := dnClient.datanodes[dnClient.memberIndex]
	dnClient.currentDatanode = *selectedDatanode

	standalonePort := 0
	for _, port := range dnClient.currentDatanode.Ports {
		if *port.Name == "STANDALONE" {
			standalonePort = int(*port.Value)
		}
	}

	address := *dnClient.currentDatanode.IpAddress + ":" + strconv.Itoa(standalonePort)
	println("Connecting to the " + address)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}

	dnClient.ratisReceiver = make(chan ratis.RaftClientReplyProto)
	dnClient.standaloneReceiver = make(chan dnapi.ContainerCommandResponseProto)

	client, err := dnapi.NewXceiverClientProtocolServiceClient(conn).Send(dnClient.ctx)
	dnClient.standaloneClient = &client
	//
	//client, err := ratis.NewRaftClientProtocolServiceClient(conn).Unordered(dnClient.ctx)
	//if err != nil {
	//	panic(err)
	//}
	//dnClient.ratisClient = &client
	//go dnClient.RaftReceiver()
	go dnClient.StandaloneReceive()
	return nil
}

func CreateDatanodeClient(pipeline *hdds.Pipeline) (*DatanodeClient, error) {
	dnClient := &DatanodeClient{
		ctx:         context.Background(),
		pipelineId:  pipeline.Id,
		datanodes:   pipeline.Members,
		memberIndex: -1,
	}
	err := dnClient.connectToNext()
	if err != nil {
		return nil, err
	}
	return dnClient, nil
}
func (dnClient *DatanodeClient) RaftReceiver() {
	for {
		proto, err := (*dnClient.ratisClient).Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		dnClient.ratisReceiver <- *proto
	}
}

func (dnClient *DatanodeClient) CreateAndWriteChunk(id *dnapi.DatanodeBlockID, blockOffset uint64, buffer []byte, length uint64) (dnapi.ChunkInfo, error) {
	bpc := uint32(12)
	checksumType := dnapi.ChecksumType_NONE
	checksumDataProto := dnapi.ChecksumData{
		Type:             &checksumType,
		BytesPerChecksum: &bpc,
	}
	chunkName := fmt.Sprintf("chunk_%d", blockOffset)
	chunkInfoProto := dnapi.ChunkInfo{
		ChunkName:    &chunkName,
		Offset:       &blockOffset,
		Len:          &length,
		ChecksumData: &checksumDataProto,
	}
	return dnClient.WriteChunk(id, chunkInfoProto, buffer[0:length])
}

func (dnClient *DatanodeClient) WriteChunk(id *dnapi.DatanodeBlockID, info dnapi.ChunkInfo, data []byte) (dnapi.ChunkInfo, error) {

	req := dnapi.WriteChunkRequestProto{
		BlockID:   id,
		ChunkData: &info,
		Data:      data,
	}
	commandType := dnapi.Type_WriteChunk
	uuid := dnClient.currentDatanode.GetUuid()
	proto := dnapi.ContainerCommandRequestProto{
		CmdType:      &commandType,
		WriteChunk:   &req,
		ContainerID:  id.ContainerID,
		DatanodeUuid: &uuid,
	}

	_, err := dnClient.sendDatanodeCommand(proto)
	if err != nil {
		return info, err
	}
	return info, nil
}

func (dnClient *DatanodeClient) ReadChunk(id *dnapi.DatanodeBlockID, info ChunkInfo) ([]byte, error) {
	result := make([]byte, 0)

	bpc := uint32(12)
	checksumType := dnapi.ChecksumType_NONE
	checksumDataProto := dnapi.ChecksumData{
		Type:             &checksumType,
		BytesPerChecksum: &bpc,
	}
	chunkInfoProto := dnapi.ChunkInfo{
		ChunkName:    &info.Name,
		Offset:       &info.Offset,
		Len:          &info.Len,
		ChecksumData: &checksumDataProto,
	}
	req := dnapi.ReadChunkRequestProto{
		BlockID:   id,
		ChunkData: &chunkInfoProto,
	}
	commandType := dnapi.Type_ReadChunk
	uuid := dnClient.currentDatanode.GetUuid()
	proto := dnapi.ContainerCommandRequestProto{
		CmdType:      &commandType,
		ReadChunk:    &req,
		ContainerID:  id.ContainerID,
		DatanodeUuid: &uuid,
	}

	resp, err := dnClient.sendDatanodeCommand(proto)
	if err != nil {
		return result, err
	}
	if resp.GetResult() != dnapi.Result_SUCCESS {
		return nil, errors.New(resp.GetResult().String() + " " + resp.GetMessage())
	}
	return resp.GetReadChunk().Data, nil
}

func (dnClient *DatanodeClient) PutBlock(id *dnapi.DatanodeBlockID, chunks []*dnapi.ChunkInfo) error {

	flags := int64(0)
	req := dnapi.PutBlockRequestProto{
		BlockData: &dnapi.BlockData{
			BlockID:  id,
			Flags:    &flags,
			Metadata: make([]*dnapi.KeyValue, 0),
			Chunks:   chunks,
		},
	}
	commandType := dnapi.Type_PutBlock
	proto := dnapi.ContainerCommandRequestProto{
		CmdType:      &commandType,
		PutBlock:     &req,
		ContainerID:  id.ContainerID,
		DatanodeUuid: dnClient.GetCurrentDnUUid(),
	}

	_, err := dnClient.sendDatanodeCommand(proto)
	if err != nil {
		return err
	}
	return nil
}

func (dnClient *DatanodeClient) GetBlock(id *dnapi.DatanodeBlockID) ([]ChunkInfo, error) {
	result := make([]ChunkInfo, 0)

	req := dnapi.GetBlockRequestProto{
		BlockID: id,
	}
	commandType := dnapi.Type_GetBlock
	proto := dnapi.ContainerCommandRequestProto{
		CmdType:      &commandType,
		GetBlock:     &req,
		ContainerID:  id.ContainerID,
		DatanodeUuid: dnClient.GetCurrentDnUUid(),
	}

	resp, err := dnClient.sendDatanodeCommand(proto)
	if err != nil {
		return result, err
	}
	for _, chunkInfo := range resp.GetGetBlock().GetBlockData().Chunks {
		result = append(result, ChunkInfo{
			Name:   chunkInfo.GetChunkName(),
			Offset: chunkInfo.GetOffset(),
			Len:    chunkInfo.GetLen(),
		})
	}
	return result, nil
}

func (dnClient *DatanodeClient) sendDatanodeCommand(proto dnapi.ContainerCommandRequestProto) (dnapi.ContainerCommandResponseProto, error) {
	return dnClient.sendStandaloneDatanodeCommand(proto)
}

func (dn *DatanodeClient) Close() {
	(*dn.standaloneClient).CloseSend()
}
