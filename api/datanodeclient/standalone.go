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
    "io"
    "github.com/apache/ozone-go/api/config"
    "github.com/apache/ozone-go/api/errors"
    "github.com/apache/ozone-go/api/proto/datanode"
    "github.com/apache/ozone-go/api/proto/hdds"
    "github.com/apache/ozone-go/api/proto/ratis"
    "strconv"
    "sync"

    log "github.com/sirupsen/logrus"
    "google.golang.org/grpc"
)

// XceiverClientStandalone TODO
type XceiverClientStandalone struct {
    pipeline *PipelineOperator
    client   datanode.XceiverClientProtocolService_SendClient
    forRead  bool
    close    bool
    mutex    sync.Mutex
    response chan *datanode.ContainerCommandResponseProto
}

// IsForRead TODO
func (x *XceiverClientStandalone) IsForRead() bool {
    // TODO implement me
    return x.forRead
}

// NewXceiverClientStandalone TODO
func NewXceiverClientStandalone(pipeline *PipelineOperator, forRead bool) XceiverClientSpi {
    spi := &XceiverClientStandalone{pipeline: pipeline, forRead: forRead, close: false}
    p := spi.init()
    log.Debug("standalone pipeline: ", p)
    return spi
}

// init TODO must add a init function, otherwise ,unit test not pass, can not mock XceiverClientSpi, why ?
func (x *XceiverClientStandalone) init() string {
    return x.pipeline.GetPipeline().String()
}

// GetReplicationType TODO
func (x *XceiverClientStandalone) GetReplicationType() hdds.ReplicationType {
    // TODO implement me
    return x.pipeline.GetPipeline().GetType()
}

// WatchForCommit TODO
func (x *XceiverClientStandalone) WatchForCommit(index uint64) (*XceiverClientReply, error) {
    // TODO implement me
    panic("implement me")
}

// GetReplicatedMinCommitIndex TODO
func (x *XceiverClientStandalone) GetReplicatedMinCommitIndex() uint64 {
    // TODO implement me
    panic("implement me")
}

// Connect TODO
func (x *XceiverClientStandalone) Connect() error {
    remotePort := x.pipeline.getRemotePort(config.STANDALONE)
    address := x.pipeline.GetCurrentNode().GetIpAddress() + ":" + strconv.Itoa(int(remotePort))
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        return err
    }
    x.response = make(chan *datanode.ContainerCommandResponseProto)
    maxSizeOption := grpc.MaxCallRecvMsgSize(32 * 10e6)
    client, err := datanode.NewXceiverClientProtocolServiceClient(conn).Send(context.Background(), maxSizeOption)
    if err != nil {
        return err
    }
    log.Info("Connected to the standalone " + address)
    x.client = client
    go x.StandaloneReceive()
    x.close = false
    return nil
}

// GetPipeline TODO
func (x *XceiverClientStandalone) GetPipeline() *PipelineOperator {
    return x.pipeline
}

// SetPipeline TODO
func (x *XceiverClientStandalone) SetPipeline(pipeline *PipelineOperator) {
    x.pipeline = pipeline
}

// RaftSend TODO
func (x *XceiverClientStandalone) RaftSend(req *datanode.ContainerCommandRequestProto) (*ratis.RaftClientReplyProto,
    error) {
    panic("")
}

// StandaloneSend TODO
func (x *XceiverClientStandalone) StandaloneSend(req *datanode.ContainerCommandRequestProto) (
    *datanode.ContainerCommandResponseProto, error) {
    err := (x.client).Send(req)

    if err != nil {
        return nil, err
    }
    reply := <-x.response
    if reply.GetMessage() == errors.StandaloneReceiveErr.Error() {
        return nil, errors.StandaloneReceiveErr
    }
    return reply, err
}

// Close TODO
func (x *XceiverClientStandalone) Close() {
    x.mutex.Lock()
    defer x.mutex.Unlock()
    if x.close {
        return
    }
    if x.client != nil {
        if err := (x.client).CloseSend(); err != nil {
            log.Errorln("stand alone client close error ", err)
        }
        x.client = nil
    }
    x.close = true
}

// StandaloneReceive TODO
func (x *XceiverClientStandalone) StandaloneReceive() {
    for {
        if x.close {
            break
        }
        reply, err := (x.client).Recv()
        if err == io.EOF {
            return
        }
        if err != nil {
            msg := errors.StandaloneReceiveErr.Error()
            log.Warn(msg, err)
            x.response <- &datanode.ContainerCommandResponseProto{Message: &msg}
			return
		}
		x.response <- reply
	}
}
