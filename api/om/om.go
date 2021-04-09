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
package om

import (
	"errors"
	"github.com/apache/ozone-go/api/common"
	"github.com/apache/ozone-go/api/proto/hdds"
	ozone_proto "github.com/apache/ozone-go/api/proto/ozone"
	"github.com/hortonworks/gohadoop"
	hadoop_ipc_client "github.com/hortonworks/gohadoop/hadoop_common/ipc/client"
	uuid "github.com/nu7hatch/gouuid"
	"net"
	"strconv"
	"sync"
)

var OM_PROTOCOL = "org.apache.hadoop.ozone.om.protocol.OzoneManagerProtocol"

type OmClient struct {
	OmHost   string
	client   *hadoop_ipc_client.Client
	clientId string
	mu       sync.Mutex
}

func CreateOmClient(omhost string) OmClient {
	clientId, _ := uuid.NewV4()
	ugi, _ := gohadoop.CreateSimpleUGIProto()
	c := &hadoop_ipc_client.Client{
		ClientId:      clientId,
		Ugi:           ugi,
		ServerAddress: net.JoinHostPort(omhost, strconv.Itoa(9862))}

	return OmClient{
		OmHost: omhost,
		client: c,
	}
}

func (om *OmClient) GetKey(volume string, bucket string, key string) (*ozone_proto.KeyInfo, error) {

	keyArgs := &ozone_proto.KeyArgs{
		VolumeName: &volume,
		BucketName: &bucket,
		KeyName:    &key,
	}
	req := ozone_proto.LookupKeyRequest{
		KeyArgs: keyArgs,
	}

	requestType := ozone_proto.Type_LookupKey
	wrapperRequest := ozone_proto.OMRequest{
		CmdType:          &requestType,
		LookupKeyRequest: &req,
		ClientId:         &om.clientId,
	}

	resp, err := om.submitRequest(&wrapperRequest)
	if err != nil {
		return nil, err
	}
	keyProto := resp.GetLookupKeyResponse().GetKeyInfo()
	println(keyProto)
	return keyProto, nil
}

func (om *OmClient) ListKeys(volume string, bucket string) ([]*ozone_proto.KeyInfo, error) {
	return om.ListKeysPrefix(volume, bucket, "")

}

func (om *OmClient) AllocateBlock(volume string, bucket string, key string, clientID *uint64) (*ozone_proto.AllocateBlockResponse,error) {
	req := ozone_proto.AllocateBlockRequest{
		KeyArgs: &ozone_proto.KeyArgs{
			VolumeName: &volume,
			BucketName: &bucket,
			KeyName:    &key,
		},
		ClientID: clientID,
	}

	msgType := ozone_proto.Type_AllocateBlock
	wrapperRequest := ozone_proto.OMRequest{
		CmdType:              &msgType,
		AllocateBlockRequest: &req,
		ClientId:             &om.clientId,
	}
	resp, err := om.submitRequest(&wrapperRequest)
	if err != nil {
		return nil, err
	}
	return resp.AllocateBlockResponse, nil
}

func (om *OmClient) CreateKey(volume string, bucket string, key string) (*ozone_proto.CreateKeyResponse, error) {
	req := ozone_proto.CreateKeyRequest{
		KeyArgs: &ozone_proto.KeyArgs{
			VolumeName: &volume,
			BucketName: &bucket,
			KeyName:    &key,
		},
	}

	createKeys := ozone_proto.Type_CreateKey
	wrapperRequest := ozone_proto.OMRequest{
		CmdType:          &createKeys,
		CreateKeyRequest: &req,
		ClientId:         &om.clientId,
	}
	resp, err := om.submitRequest(&wrapperRequest)
	if err != nil {
		return nil, err
	}
	return resp.CreateKeyResponse, nil
}

func (om *OmClient) CommitKey(volume string, bucket string, key string, id *uint64, keyLocations []*ozone_proto.KeyLocation, size uint64) (common.Key, error) {
	one := hdds.ReplicationFactor_ONE
	standalone := hdds.ReplicationType_STAND_ALONE
	req := ozone_proto.CommitKeyRequest{
		KeyArgs: &ozone_proto.KeyArgs{
			VolumeName:   &volume,
			BucketName:   &bucket,
			KeyName:      &key,
			KeyLocations: keyLocations,
			DataSize:     &size,
			Factor:       &one,
			Type:         &standalone,
		},
		ClientID: id,
	}

	messageType := ozone_proto.Type_CommitKey
	wrapperRequest := ozone_proto.OMRequest{
		CmdType:          &messageType,
		CommitKeyRequest: &req,
		ClientId:         &om.clientId,
	}
	_, err := om.submitRequest(&wrapperRequest)
	if err != nil {
		return common.Key{}, err
	}

	return common.Key{}, nil
}

func (om *OmClient) ListKeysPrefix(volume string, bucket string, prefix string) ([]*ozone_proto.KeyInfo, error) {

	req := ozone_proto.ListKeysRequest{
		VolumeName: &volume,
		BucketName: &bucket,
		Prefix:     ptr(prefix),
		Count:      ptri(1000),
	}

	listKeys := ozone_proto.Type_ListKeys
	wrapperRequest := ozone_proto.OMRequest{
		CmdType:         &listKeys,
		ListKeysRequest: &req,
		ClientId:        &om.clientId,
	}

	resp, err := om.submitRequest(&wrapperRequest)
	if err != nil {
		return nil, err
	}

	return resp.GetListKeysResponse().GetKeyInfo(), nil
}

func getRpcPort(ports []*hdds.Port) uint32 {
	for _, port := range ports {
		if port.GetName() == "RATIS" {
			return port.GetValue()
		}
	}
	return 0
}

func ptri(i int32) *int32 {
	return &i
}

func ptr(s string) *string {
	return &s
}

func (om *OmClient) submitRequest(request *ozone_proto.OMRequest, ) (*ozone_proto.OMResponse, error) {
	wrapperResponse := ozone_proto.OMResponse{}
	om.mu.Lock()
	err := om.client.Call(gohadoop.GetCalleeRPCRequestHeaderProto(&OM_PROTOCOL), request, &wrapperResponse)
	om.mu.Unlock()
	if err != nil {
		return nil, err
	}
	if *wrapperResponse.Status != ozone_proto.Status_OK {
		return nil, errors.New("Error on calling OM " + wrapperResponse.Status.String() + " " + *wrapperResponse.Message)
	}
	return &wrapperResponse, nil
}

