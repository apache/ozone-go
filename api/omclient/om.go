// Package omclient TODO
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
package omclient

import (
    "errors"
    "fmt"
    "github.com/apache/ozone-go/api/config"
    hadoop_ipc_client "github.com/apache/ozone-go/api/hadoop/ipc/client"
    "github.com/apache/ozone-go/api/proto/hdds"
    ozone_proto "github.com/apache/ozone-go/api/proto/ozone"
    "strings"

    uuid "github.com/nu7hatch/gouuid"
    log "github.com/sirupsen/logrus"
)

// OM_PROTOCOL TODO
var OM_PROTOCOL = "org.apache.hadoop.ozone.om.protocol.OzoneManagerProtocol"

// HadoopIpcClient TODO
type HadoopIpcClient struct {
    Client *hadoop_ipc_client.Client
}

// OmClient TODO
type OmClient struct {
    OmHost    string
    ClientId  string
    IpcClient IpcClient
    Id        uint64
}

// CreateOmClient TODO
func CreateOmClient(omHost string) (*OmClient, error) {
    clientId, err := uuid.NewV4()
    if err != nil {
        return nil, err
    }
    ugi, err := hadoop_ipc_client.CreateSimpleUGIProto()
    if err != nil {
        return nil, err
    }
    ipcClient := &HadoopIpcClient{Client: &hadoop_ipc_client.Client{
        ClientId:      clientId,
        Ugi:           ugi,
        ServerAddress: omHost,
        TCPNoDelay:    false,
    }}
    return &OmClient{
        OmHost:    omHost,
        ClientId:  clientId.String(),
        IpcClient: ipcClient,
    }, nil
}

// Equal TODO
func (om *OmClient) Equal(i interface{}) bool {
    if o, ok := i.(OmClient); ok {
        return o.OmHost == om.OmHost && o.ClientId == om.ClientId
    }
    return false
}

// AllocateBlock TODO
func (om *OmClient) AllocateBlock(volume string, bucket string, key string, clientID *uint64,
    excludeList *hdds.ExcludeListProto) (
    *ozone_proto.AllocateBlockResponse, error) {
    req := ozone_proto.AllocateBlockRequest{
        KeyArgs: &ozone_proto.KeyArgs{
            VolumeName: &volume,
            BucketName: &bucket,
            KeyName:    &key,
        },
        ClientID:    clientID,
        ExcludeList: excludeList,
    }

    msgType := ozone_proto.Type_AllocateBlock
    wrapperRequest := ozone_proto.OMRequest{
        CmdType:              &msgType,
        AllocateBlockRequest: &req,
        ClientId:             &om.ClientId,
    }

    resp, err := om.SubmitRequest(&wrapperRequest)
    if err != nil {
        return nil, err
    }
    return resp.GetAllocateBlockResponse(), nil
}

func getRpcPort(ports []*hdds.Port) uint32 {
    for _, port := range ports {
        if port.GetName() == "RATIS" {
            return port.GetValue()
        }
    }
    return 0
}

// SubmitRequest TODO
func (om *OmClient) SubmitRequest(request *ozone_proto.OMRequest) (resp *ozone_proto.OMResponse, err error) {
    maxRetry := len(config.OmAddresses)
    retry := 0
    for retry < maxRetry {
        if resp, err = om.IpcClient.SubmitRequest(request); err == nil {
            return resp, nil
        } else {
            retry++
            if strings.Contains(err.Error(), "OMNotLeaderException") {
                om.IpcClient.UpdateServerAddress(config.OzoneConfig.NextOmNode())
            } else {
                if resp.GetStatus() == ozone_proto.Status_KEY_NOT_FOUND {
                    return resp, nil
                }
                log.Debug(fmt.Sprintf("om submit request call %s error ", request.String()), err)
            }
        }
    }
    return nil, err
}

// UpdateServerAddress TODO
func (ipcClient *HadoopIpcClient) UpdateServerAddress(address string) {
    ipcClient.Client.ServerAddress = address
}

// SubmitRequest TODO
func (ipcClient *HadoopIpcClient) SubmitRequest(request *ozone_proto.OMRequest) (*ozone_proto.OMResponse, error) {
    wrapperResponse := &ozone_proto.OMResponse{}
    if err := ipcClient.Client.Call(hadoop_ipc_client.GetCalleeRPCRequestHeaderProto(&OM_PROTOCOL), request,
        wrapperResponse); err != nil {
        return wrapperResponse, err
    }
    if wrapperResponse.GetStatus() != ozone_proto.Status_OK {
        return wrapperResponse, errors.New("Error on calling OM " + wrapperResponse.Status.String() + " " +
			*wrapperResponse.Message)
	}
	return wrapperResponse, nil
}
