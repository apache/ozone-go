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
    "context"
    "encoding/binary"
    "encoding/hex"
    "fmt"
    "io"
    "github.com/apache/ozone-go/api/common"
    "github.com/apache/ozone-go/api/config"
    "github.com/apache/ozone-go/api/errors"
    "github.com/apache/ozone-go/api/proto/datanode"
    "github.com/apache/ozone-go/api/proto/hdds"
    "github.com/apache/ozone-go/api/proto/ratis"
    "strconv"
    "strings"
    "sync"

    "github.com/gogo/protobuf/proto"
    log "github.com/sirupsen/logrus"
    "google.golang.org/grpc"
)

// XceiverClientRatis TODO
type XceiverClientRatis struct {
    pipeline      *PipelineOperator
    client        ratis.RaftClientProtocolService_UnorderedClient
    reply         chan *ratis.RaftClientReplyProto
    close         bool
    seqNum        uint64
    groupId       []byte
    commitInfoMap map[string]uint64
    mutex         sync.Mutex
    forRead       bool
}

// XceiverClientReply TODO
type XceiverClientReply struct {
    LogIndex  uint64
    Reply     *datanode.ContainerCommandResponseProto
    Datanodes []string
}

// GetLogIndex TODO
func (x *XceiverClientReply) GetLogIndex() uint64 {
    return x.LogIndex
}

// NewXceiverClientRatis TODO
func NewXceiverClientRatis(pipeline *PipelineOperator, forRead bool) XceiverClientSpi {
    groupId, err := hex.DecodeString(strings.ReplaceAll(pipeline.GetId(), "-", ""))
    if err != nil {
        log.Errorf("Decode pipe line id error %s", err.Error())
        return nil
    }
    return &XceiverClientRatis{
        pipeline:      pipeline,
        groupId:       groupId,
        close:         false,
        forRead:       forRead,
        reply:         make(chan *ratis.RaftClientReplyProto),
        commitInfoMap: make(map[string]uint64),
    }
}

// IsForRead TODO
func (x *XceiverClientRatis) IsForRead() bool {
    return x.forRead
}

// GetPipeline TODO
func (x *XceiverClientRatis) GetPipeline() *PipelineOperator {
    return x.pipeline
}

// SetPipeline TODO
func (x *XceiverClientRatis) SetPipeline(pipeline *PipelineOperator) {
    x.pipeline = pipeline
}

func (x *XceiverClientRatis) updateCommitInfosMap(commitInfoProtos []*ratis.CommitInfoProto) {
    for _, commitInfoProto := range commitInfoProtos {
        address := commitInfoProto.GetServer().GetAddress()
        x.commitInfoMap[address] = commitInfoProto.GetCommitIndex()
    }
}

// GetReplicationType TODO
func (x *XceiverClientRatis) GetReplicationType() hdds.ReplicationType {
    // TODO implement me
    return x.pipeline.GetPipeline().GetType()
}

// WatchForCommit TODO
func (x *XceiverClientRatis) WatchForCommit(index uint64) (*XceiverClientReply, error) {
    // TODO implement me
    minIndex := x.GetReplicatedMinCommitIndex()
    if minIndex >= index {
        return &XceiverClientReply{LogIndex: minIndex}, nil
    }

    // init watch 0,ratis.ReplicationLevel_MAJORITY
    // first watch index,ratis.ReplicationLevel_ALL_COMMITTED
    // last watch index,ratis.ReplicationLevel_ALL_COMMITTED
    // exception index,ratis.ReplicationLevel_MAJORITY_COMMITTED
    reply, err := x.SendWatch(index, ratis.ReplicationLevel_ALL_COMMITTED)
    xceiverClientReply := &XceiverClientReply{
        LogIndex:  0,
        Reply:     &datanode.ContainerCommandResponseProto{},
        Datanodes: make([]string, 0),
    }
    if err != nil {
        log.Warn(fmt.Sprintf("3 way commit failed on pipeline %s", x.pipeline.GetPipeline().String()), err)
        if strings.Contains(reply.GetMessage().String(), "GroupMismatchException") {
            return nil, err
        }
        re, er := x.SendWatch(index, ratis.ReplicationLevel_MAJORITY_COMMITTED)
        if er != nil {
            return nil, er
        }
        // since 3 way commit has failed, the updated map from now on  will
        // only store entries for those datanodes which have had successful
        // replication.
        for _, infoProto := range re.GetCommitInfos() {
            if infoProto.GetCommitIndex() < index {
                address := infoProto.GetServer().GetAddress()
                delete(x.commitInfoMap, address)
                xceiverClientReply.Datanodes = append(xceiverClientReply.Datanodes, address)
                log.Info(fmt.Sprintf("Could not commit index %d on pipeline %s to all the nodes. "+
                    "Server %s has failed. Committed by majority.", index, x.pipeline.GetPipeline().String(), address))
            }
        }
    }
    x.updateCommitInfosMap(reply.GetCommitInfos())
    xceiverClientReply.LogIndex = index
    return xceiverClientReply, nil
}

// SendWatch TODO
func (x *XceiverClientRatis) SendWatch(index uint64, replicationLevel ratis.ReplicationLevel) (
    *ratis.RaftClientReplyProto, error) {
    // init watch 0,ratis.ReplicationLevel_MAJORITY
    // first watch index,ratis.ReplicationLevel_ALL_COMMITTED
    // last watch index,ratis.ReplicationLevel_ALL_COMMITTED
    // exception index,ratis.ReplicationLevel_MAJORITY_COMMITTED

    Watch := &ratis.RaftClientRequestProto_Watch{Watch: &ratis.WatchRequestTypeProto{
        Index:       index,
        Replication: replicationLevel,
    }}

    callId := common.OId.CallIdGetAndIncrement()
    clientId := []byte(common.OId.GetClientId())
    x.seqNum++

    raftClientReq := &ratis.RaftClientRequestProto{
        RpcRequest: &ratis.RaftRpcRequestProto{
            RequestorId: clientId,
            ReplyId:     clientId,
            RaftGroupId: &ratis.RaftGroupIdProto{
                Id: x.groupId,
            },
            CallId: callId,
            SlidingWindowEntry: &ratis.SlidingWindowEntry{
                SeqNum:  x.seqNum,
                IsFirst: true,
            },
        },
        Message: nil,
        Type:    Watch,
    }

    if err := (x.client).Send(raftClientReq); err != nil {
        return nil, err
    } else {
        reply := <-x.reply
        return reply, nil
    }
}

// GetReplicatedMinCommitIndex TODO
func (x *XceiverClientRatis) GetReplicatedMinCommitIndex() uint64 {
    minIndex := uint64(0)
    for _, u := range x.commitInfoMap {
        if u > minIndex {
            minIndex = u
        }
    }
    return minIndex
}

// Connect TODO
func (x *XceiverClientRatis) Connect() error {
    x.close = false
    remotePort := x.pipeline.getRemotePort(config.RATIS)
    ip := x.pipeline.GetLeaderAddress()
    address := ip + ":" + strconv.Itoa(int(remotePort))
    log.Debug("Connecting to the ratis " + address)
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Error(err)
        return err
    }
    // var cancel context.CancelFunc
    // pool.ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
    maxSizeOption := grpc.MaxCallRecvMsgSize(32 * 10e6)
    client, err := ratis.NewRaftClientProtocolServiceClient(conn).Unordered(context.Background(), maxSizeOption)
    if err != nil {
        log.Error(err)
        return err
    }
    log.Debug("Connected to the ratis " + address)
    x.client = client
    go x.RaftReceive()
    return nil
}

// RaftReceive TODO
func (x *XceiverClientRatis) RaftReceive() {
    for {
        if x.close {
            break
        }
        reply, err := (x.client).Recv()
        if err == io.EOF {
            log.Warn("ratis eof")
            return
        }
        if err != nil {
            log.Warn(errors.RatisReceiveErr.Error(), err)
            x.reply <- reply
            return
        }
        x.reply <- reply
    }
}

// RaftSend TODO
func (x *XceiverClientRatis) RaftSend(req *datanode.ContainerCommandRequestProto) (*ratis.RaftClientReplyProto, error) {
    callId := common.OId.CallIdGetAndIncrement()
    clientId := []byte(common.OId.GetClientId())
    x.seqNum++
    raftRpcReq := &ratis.RaftRpcRequestProto{
        RequestorId: clientId,
        ReplyId:     clientId,
        RaftGroupId: &ratis.RaftGroupIdProto{
            Id: x.groupId,
        },
        CallId: callId,
        SlidingWindowEntry: &ratis.SlidingWindowEntry{
            SeqNum:  x.seqNum,
            IsFirst: false,
        },
    }
    message := toMessage(req)
    var raftClientReq = &ratis.RaftClientRequestProto{RpcRequest: raftRpcReq,
        Message: message}
    if IsReadOnly(req.GetCmdType().String()) {
        raftClientReq.Type = &ratis.RaftClientRequestProto_Read{Read: &ratis.ReadRequestTypeProto{}}
    } else {
        raftClientReq.Type = &ratis.RaftClientRequestProto_Write{Write: &ratis.WriteRequestTypeProto{}}
    }

    err := (x.client).Send(raftClientReq)
    if err != nil {
        log.Errorf("raft send request [%s] error %v", raftClientReq.String(), err)
        return nil, err
    }
    reply := <-x.reply
    if reply.RpcReply == nil {
        return nil, errors.RatisReceiveErr
    }
    return reply, nil
}

// StandaloneSend TODO
func (x *XceiverClientRatis) StandaloneSend(req *datanode.ContainerCommandRequestProto) (
    *datanode.ContainerCommandResponseProto, error) {
    panic("")
}

// Close TODO
func (x *XceiverClientRatis) Close() {
    x.mutex.Lock()
    defer x.mutex.Unlock()
    if x.close {
        return
    }
    x.close = true
    if x.client != nil {
        if err := (x.client).CloseSend(); err != nil {
            log.Errorln("close ratis client error:", err)
        }
        x.client = nil
    }
}

func toMessage(req *datanode.ContainerCommandRequestProto) *ratis.ClientMessageEntryProto {
    data := make([]byte, 0)
    cmdType := req.GetCmdType().String()
    if cmdType == datanode.Type_WriteChunk.String() {
        data = req.WriteChunk.GetData()
        req.WriteChunk.Data = nil
    } else if cmdType == datanode.Type_PutSmallFile.String() {
        data = req.PutSmallFile.GetData()
        req.PutSmallFile.Data = nil
    }
    bytes, err := proto.Marshal(req)
    if err != nil {
        log.Error(fmt.Sprintf("proto marshal error when send request %s", req.String()), err)
    }
    lengthHeader := make([]byte, 4)
    binary.BigEndian.PutUint32(lengthHeader, uint32(len(bytes)))
    content := append(lengthHeader, bytes...)
	message := &ratis.ClientMessageEntryProto{
		Content: append(content, data...),
	}
	return message
}
