// Package datanodeclient TODO
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
	"fmt"
	"github.com/apache/ozone-go/api/common"
	"github.com/apache/ozone-go/api/config"
	"github.com/apache/ozone-go/api/proto/datanode"
	"github.com/apache/ozone-go/api/proto/ratis"
	"github.com/apache/ozone-go/api/utils"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// GetBlock TODO
func GetBlock(spi XceiverClientSpi, blockId *datanode.DatanodeBlockID) (*datanode.GetBlockResponseProto, error) {
	dnUuid := spi.GetPipeline().GetCurrentNode().GetUuid()
	req := &datanode.ContainerCommandRequestProto{
		CmdType:      datanode.Type_GetBlock.Enum(),
		DatanodeUuid: &dnUuid,
		ContainerID:  blockId.ContainerID,
		GetBlock: &datanode.GetBlockRequestProto{
			BlockID: &datanode.DatanodeBlockID{
				ContainerID: blockId.ContainerID,
				LocalID:     blockId.LocalID,
			},
		},
	}
	response, err := spi.StandaloneSend(req)
	return response.GetGetBlock(), err
}

// PutBlock TODO
func PutBlock(spi XceiverClientSpi, blockData *datanode.BlockData, eof bool) (*ratis.RaftClientReplyProto, error) {
	dnUuid := spi.GetPipeline().GetCurrentNode().GetUuid()
	req := &datanode.ContainerCommandRequestProto{
		CmdType:      datanode.Type_PutBlock.Enum(),
		DatanodeUuid: &dnUuid,
		ContainerID:  blockData.BlockID.ContainerID,
		PutBlock: &datanode.PutBlockRequestProto{
			BlockData: blockData,
			Eof:       &eof,
		},
	}
	log.Debug("Put block request: ", req.String())
	return spi.RaftSend(req)
}

// ReadChunk TODO
func ReadChunk(spi XceiverClientSpi, chunkInfo *datanode.ChunkInfo, blockId *datanode.DatanodeBlockID) (
	*datanode.ReadChunkResponseProto, error) {
	dnUuid := spi.GetPipeline().GetCurrentNode().GetUuid()

	bpc, err := config.OzoneConfig.GetBytesPerChecksum()
	if err != nil {
		log.Error("Get bytes per checksum error:", err)
		return nil, err
	}
	readChunkInfo := &datanode.ChunkInfo{
		ChunkName: chunkInfo.ChunkName,
		Offset:    chunkInfo.Offset,
		Len:       chunkInfo.Len,
		ChecksumData: &datanode.ChecksumData{
			Type:             chunkInfo.ChecksumData.Type,
			BytesPerChecksum: utils.PointUint32(uint32(bpc)),
		},
	}
	req := &datanode.ContainerCommandRequestProto{
		CmdType:      datanode.Type_ReadChunk.Enum(),
		DatanodeUuid: &dnUuid,
		ContainerID:  blockId.ContainerID,
		ReadChunk: &datanode.ReadChunkRequestProto{
			BlockID:   blockId,
			ChunkData: readChunkInfo,
		},
	}
	response, err := spi.StandaloneSend(req)
	log.Debug("Read chunk response: ", response.GetReadChunk().GetChunkData().GetChunkName())
	return response.GetReadChunk(), err
}

// WriteChunk TODO
func WriteChunk(spi XceiverClientSpi, chunkInfo *datanode.ChunkInfo, blockId *datanode.DatanodeBlockID,
	data []byte) (*ratis.RaftClientReplyProto, error) {
	dnUuid := spi.GetPipeline().GetCurrentNode().GetUuid()
	traceId := ""
	req := &datanode.ContainerCommandRequestProto{
		CmdType:      datanode.Type_WriteChunk.Enum(),
		TraceID:      &traceId,
		PipelineID:   utils.PointString(spi.GetPipeline().GetId()),
		ContainerID:  blockId.ContainerID,
		DatanodeUuid: &dnUuid,
		EncodedToken: nil,
		WriteChunk: &datanode.WriteChunkRequestProto{
			BlockID:   blockId,
			ChunkData: chunkInfo,
			Data:      data,
		},
	}
	log.Debug(fmt.Sprintf("%s name:%s length:%d offset:%d",
		req.GetCmdType().String(),
		req.GetWriteChunk().GetChunkData().GetChunkName(),
		req.GetWriteChunk().GetChunkData().GetLen(),
		req.GetWriteChunk().GetChunkData().GetOffset()))
	return spi.RaftSend(req)
}

// WriteSmallFile TODO
func WriteSmallFile(spi XceiverClientSpi, blockId *datanode.DatanodeBlockID,
	data []byte) (*datanode.PutSmallFileResponseProto, error) {
	dnUuid := spi.GetPipeline().GetCurrentNode().GetUuid()
	traceId := ""
	blockData := &datanode.BlockData{
		BlockID: blockId,
	}
	keyValue := &datanode.KeyValue{
		Key:   utils.PointString("OverWriteRequested"),
		Value: utils.PointString("true"),
	}
	bpc, err := config.OzoneConfig.GetBytesPerChecksum()
	if err != nil {
		return nil, err
	}
	ckt := config.OzoneConfig.GetChecksumType()
	checksumType := common.ChecksumTypeFromName(ckt).Enum()
	checksumData := NewChecksumOperatorComputer(checksumType, uint32(bpc))
	if err := checksumData.ComputeChecksum(data); err != nil {
		return nil, err
	}
	chunkInfo := &datanode.ChunkInfo{
		ChunkName:    utils.PointString(strconv.FormatInt(blockId.GetLocalID(), 10) + "_chunk"),
		Offset:       utils.PointUint64(0),
		Len:          utils.PointUint64(uint64(len(data))),
		Metadata:     []*datanode.KeyValue{keyValue},
		ChecksumData: checksumData.Checksum,
	}
	req := &datanode.ContainerCommandRequestProto{
		CmdType:      datanode.Type_PutSmallFile.Enum(),
		TraceID:      &traceId,
		PipelineID:   utils.PointString(spi.GetPipeline().GetId()),
		ContainerID:  blockId.ContainerID,
		DatanodeUuid: &dnUuid,
		PutSmallFile: &datanode.PutSmallFileRequestProto{
			Block: &datanode.PutBlockRequestProto{
				BlockData: blockData,
			},
			ChunkInfo: chunkInfo,
			Data:      data,
		},
	}
	response, err := spi.StandaloneSend(req)
	return response.GetPutSmallFile(), err
}
