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
	"github.com/apache/ozone-go/api/common"
	ozone_proto "github.com/apache/ozone-go/api/proto/ozone"
)

func (om *OmClient) ListVolumes() ([]common.Volume, error) {
	scope := ozone_proto.ListVolumeRequest_VOLUMES_BY_USER
	req := ozone_proto.ListVolumeRequest{
		Scope:    &scope,
		UserName: ptr("hadoop"),
		Prefix:   ptr(""),
	}

	listKeys := ozone_proto.Type_ListVolume
	clientId := "goClient"
	wrapperRequest := ozone_proto.OMRequest{
		CmdType:           &listKeys,
		ListVolumeRequest: &req,
		ClientId:          &clientId,
	}

	volumes := make([]common.Volume, 0)
	resp, err := om.submitRequest(&wrapperRequest)
	if err != nil {
		return nil, err
	}
	for _, volProto := range resp.GetListVolumeResponse().GetVolumeInfo() {
		volumes = append(volumes, common.Volume{Name: *volProto.Volume})
	}
	return volumes, nil
}

func (om *OmClient) CreateVolume(name string) error {
	onegig := uint64(1024 * 1024 * 1024)
	volumeInfo := ozone_proto.VolumeInfo{
		AdminName:    ptr("hadoop"),
		OwnerName:    ptr("hadoop"),
		Volume:       ptr(name),
		QuotaInBytes: &onegig,
	}
	req := ozone_proto.CreateVolumeRequest{
		VolumeInfo: &volumeInfo,
	}

	cmdType := ozone_proto.Type_CreateVolume
	clientId := "goClient"
	wrapperRequest := ozone_proto.OMRequest{
		CmdType:             &cmdType,
		CreateVolumeRequest: &req,
		ClientId:            &clientId,
	}

	_, err := om.submitRequest(&wrapperRequest)
	if err != nil {
		return err
	}
	return nil
}

func (om *OmClient) GetVolume(name string) (common.Volume, error) {
	req := ozone_proto.InfoVolumeRequest{
		VolumeName: &name,
	}

	cmdType := ozone_proto.Type_InfoVolume
	wrapperRequest := ozone_proto.OMRequest{
		CmdType:           &cmdType,
		InfoVolumeRequest: &req,
		ClientId:          &om.clientId,
	}

	resp, err := om.submitRequest(&wrapperRequest)
	if err != nil {
		return common.Volume{}, err
	}

	vol := common.Volume{}
	vol.Name = *resp.InfoVolumeResponse.VolumeInfo.Volume
	vol.Owner = *resp.InfoVolumeResponse.VolumeInfo.OwnerName

	return vol, nil
}
