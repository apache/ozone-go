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
    "github.com/apache/ozone-go/api/proto/datanode"
    "github.com/apache/ozone-go/api/proto/hdds"
    "github.com/apache/ozone-go/api/proto/ratis"
    "sync"
)

// XceiverClientSpi TODO
type XceiverClientSpi interface {
    Connect() error
    IsForRead() bool
    GetPipeline() *PipelineOperator
    GetReplicationType() hdds.ReplicationType
    SetPipeline(pipeline *PipelineOperator)
    RaftSend(req *datanode.ContainerCommandRequestProto) (*ratis.RaftClientReplyProto, error)
    StandaloneSend(req *datanode.ContainerCommandRequestProto) (*datanode.ContainerCommandResponseProto, error)
    WatchForCommit(index uint64) (*XceiverClientReply, error)
    GetReplicatedMinCommitIndex() uint64
    Close()
}

// XceiverClientManager TODO
type XceiverClientManager struct {
    ClientCache       map[string]XceiverClientSpi
    TopologyAwareRead bool
    mutex             sync.Mutex
}

// XceiverClientManagerInstance TODO
var XceiverClientManagerInstance *XceiverClientManager
var once = sync.Once{}

// GetXceiverClientManagerInstance TODO
func GetXceiverClientManagerInstance() *XceiverClientManager {
    if XceiverClientManagerInstance == nil {
        once.Do(func() {
            XceiverClientManagerInstance = &XceiverClientManager{
                ClientCache:       make(map[string]XceiverClientSpi, 0),
                TopologyAwareRead: false,
            }
        })
    }
    return XceiverClientManagerInstance
}

// GetClient TODO
func (x *XceiverClientManager) GetClient(pipeline *PipelineOperator, read bool) (XceiverClientSpi, error) {
    typ := pipeline.GetPipeline().GetType()
    key := x.getPipelineCacheKey(pipeline, read)
    x.mutex.Lock()
    defer x.mutex.Unlock()
    var client XceiverClientSpi
    var err error
    if v, ok := x.ClientCache[key]; ok {
        client = v
    } else {
        switch typ {
        case hdds.ReplicationType_RATIS:
            client = NewXceiverClientRatis(pipeline, read)
        case hdds.ReplicationType_STAND_ALONE:
            client = NewXceiverClientStandalone(pipeline, read)
        default:
            return nil, fmt.Errorf("unknown replication type: %s", typ)
        }
        if client == nil {
            return nil, fmt.Errorf("create %s client nil error", typ.String())
        }
        if err := client.Connect(); err != nil {
            return nil, fmt.Errorf("%s client create connection error %s", typ, err.Error())
        }
        x.ClientCache[key] = client
    }
    return client, err
}

// RemoveClient TODO
func (x *XceiverClientManager) RemoveClient(pipeline *PipelineOperator, read bool) {
    key := x.getPipelineCacheKey(pipeline, read)
    x.mutex.Lock()
    defer x.mutex.Unlock()
    if v, ok := x.ClientCache[key]; ok {
        if v != nil {
            v.Close()
            delete(x.ClientCache, key)
        }
    }
}

func (x *XceiverClientManager) getPipelineCacheKey(pipeline *PipelineOperator, forRead bool) string {
    key := pipeline.GetPipeline().GetId().GetId() + pipeline.GetPipeline().GetType().String()
    if x.TopologyAwareRead && forRead {
        key += pipeline.GetClosestNode().GetHostName()
	}
	return key
}
