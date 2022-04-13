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
    "github.com/apache/ozone-go/api/common"
    "github.com/apache/ozone-go/api/config"
    ozone_proto "github.com/apache/ozone-go/api/proto/ozone"
    "github.com/apache/ozone-go/api/utils"

    "github.com/google/uuid"
)

// ListVolumes TODO
func (om *OmClient) ListVolumes() ([]*common.Volume, error) {
    scope := ozone_proto.ListVolumeRequest_VOLUMES_BY_USER
    req := ozone_proto.ListVolumeRequest{
        Scope:    &scope,
        UserName: utils.PointString(config.User),
        Prefix:   utils.PointString(""),
        PrevKey:  utils.PointString(""),
        MaxKeys:  utils.PointUint32(1000),
    }
    listKeys := ozone_proto.Type_ListVolume
    clientId, err := uuid.NewRandom()
    if err != nil {
        return nil, err
    }
    wrapperRequest := &ozone_proto.OMRequest{
        CmdType:           &listKeys,
        ListVolumeRequest: &req,
        ClientId:          utils.PointString(clientId.String()),
    }

    volumes := make([]*common.Volume, 0)
    if resp, err := om.SubmitRequest(wrapperRequest); err != nil {
        return nil, err
    } else {
        for _, volProto := range resp.GetListVolumeResponse().GetVolumeInfo() {
            volumes = append(volumes, volFromProto(volProto))
        }
        return volumes, nil
    }
}

// CreateVolume TODO
func (om *OmClient) CreateVolume(volumeName string) error {
    volumeInfo := ozone_proto.VolumeInfo{
        AdminName:        utils.PointString(config.User),
        OwnerName:        utils.PointString(config.User),
        Volume:           utils.PointString(volumeName),
        QuotaInBytes:     utils.PointUint64(utils.GB.Value()),
        QuotaInNamespace: utils.PointInt64(-1),
        VolumeAcls:       CreateAcls(ALL),
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

    _, err := om.SubmitRequest(&wrapperRequest)
    return err
}

// InfoVolume TODO
func (om *OmClient) InfoVolume(volumeName string) (*common.Volume, error) {
    req := ozone_proto.InfoVolumeRequest{
        VolumeName: &volumeName,
    }
    cmdType := ozone_proto.Type_InfoVolume
    wrapperRequest := ozone_proto.OMRequest{
        CmdType:           &cmdType,
        InfoVolumeRequest: &req,
        ClientId:          &om.ClientId,
    }
    resp, err := om.SubmitRequest(&wrapperRequest)
    if err != nil {
        return nil, err
    }
    return volFromProto(resp.InfoVolumeResponse.GetVolumeInfo()), nil
}

func volFromProto(volProto *ozone_proto.VolumeInfo) *common.Volume {
    ct := utils.MillisecondFormatToString(int64(volProto.GetCreationTime()), utils.TimeFormatterSECOND)
    mt := utils.MillisecondFormatToString(int64(volProto.GetModificationTime()), utils.TimeFormatterSECOND)
    acls := volProto.GetVolumeAcls()
    aclList := make([]common.Acl, 0)
    for _, acl := range acls {
        aclInfo := NewAclInfo()
        aclInfo.FromBytes(acl.GetRights())
        a := common.Acl{
            Type:     acl.GetType().String(),
            Name:     acl.GetName(),
            AclScope: acl.GetAclScope().String(),
            AclList:  aclInfo.Strings(),
        }
        aclList = append(aclList, a)
    }
    return &common.Volume{
        VolumeName:       volProto.GetVolume(),
        Owner:            volProto.GetOwnerName(),
        QuotaInBytes:     utils.BytesToHuman(volProto.GetQuotaInBytes()),
        QuotaInNamespace: volProto.GetQuotaInNamespace(),
        UsedNamespace:    volProto.GetUsedNamespace(),
        CreationTime:     ct,
		ModificationTime: mt,
		Acls:             aclList,
	}
}
