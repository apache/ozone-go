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
    "io"
    "github.com/apache/ozone-go/api/common"
    "github.com/apache/ozone-go/api/config"
    dnproto "github.com/apache/ozone-go/api/proto/datanode"
    "github.com/apache/ozone-go/api/proto/hdds"
    ozoneproto "github.com/apache/ozone-go/api/proto/ozone"
    "github.com/apache/ozone-go/api/utils"
    "time"
)

// InfoKey TODO
func (om *OmClient) InfoKey(volume string, bucket string, key string) (*common.Key, error) {
    if keyProto, err := om.LookupKey(volume, bucket, key); err == nil {
        if keyProto == nil {
            return nil, nil
        }
        return KeyFromProto(keyProto), nil
    } else {
        return nil, err
    }
}

// TouchzKey TODO
func (om *OmClient) TouchzKey(volume string, bucket string, key string) (*common.Key, error) {
    if keyProto, err := om.CreateKey(volume, bucket, key); err == nil {
        if keyProto == nil {
            return nil, nil
        }
        return KeyFromProto(keyProto.GetKeyInfo()), nil
    } else {
        return nil, err
    }
}

// LookupKey TODO
func (om *OmClient) LookupKey(volume string, bucket string, key string) (*ozoneproto.KeyInfo, error) {
    keyArgs := &ozoneproto.KeyArgs{
        VolumeName: &volume,
        BucketName: &bucket,
        KeyName:    &key,
    }
    req := ozoneproto.LookupKeyRequest{
        KeyArgs: keyArgs,
    }

    requestType := ozoneproto.Type_LookupKey
    wrapperRequest := ozoneproto.OMRequest{
        CmdType:          &requestType,
        LookupKeyRequest: &req,
        ClientId:         &om.ClientId,
    }

    resp, err := om.SubmitRequest(&wrapperRequest)
    if resp.GetStatus() == ozoneproto.Status_KEY_NOT_FOUND {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    return resp.GetLookupKeyResponse().GetKeyInfo(), nil
}

// RenameKey TODO
func (om *OmClient) RenameKey(volume, bucket, key, toKeyName string) (string, error) {
    req := ozoneproto.RenameKeyRequest{
        KeyArgs: &ozoneproto.KeyArgs{
            VolumeName: &volume,
            BucketName: &bucket,
            KeyName:    &key,
        },
        ToKeyName: &toKeyName,
    }
    cmdType := ozoneproto.Type_RenameKey
    wrapperRequest := ozoneproto.OMRequest{
        CmdType:          &cmdType,
        RenameKeyRequest: &req,
        ClientId:         &om.ClientId,
    }
    resp, err := om.SubmitRequest(&wrapperRequest)
    if err != nil {
        return "", err
    }
    k := resp.RenameKeysResponse.GetUnRenamedKeys()
    ret := ""
    for _, keysMap := range k {
        ret = keysMap.String() + "\t"
    }
    return ret, nil
}

// CreateKey TODO
func (om *OmClient) CreateKey(volume string, bucket string, key string) (*ozoneproto.CreateKeyResponse, error) {
    timeNow := time.Now().UnixNano() / 1e6
    req := ozoneproto.CreateKeyRequest{
        KeyArgs: &ozoneproto.KeyArgs{
            VolumeName:       &volume,
            BucketName:       &bucket,
            KeyName:          &key,
            DataSize:         utils.PointUint64(0),
            Type:             GetReplicationType(),
            Factor:           GetReplication(),
            IsMultipartKey:   utils.PointBool(false),
            Acls:             CreateAcls(ALL),
            ModificationTime: utils.PointUint64(uint64(timeNow)),
            Recursive:        utils.PointBool(false),
        },
        ClientID: utils.PointUint64(0),
    }

    wrapperRequest := ozoneproto.OMRequest{
        CmdType:  ozoneproto.Type_CreateKey.Enum(),
        TraceID:  utils.PointString("trace"),
        ClientId: &om.ClientId,
        UserInfo: &ozoneproto.UserInfo{
            UserName: utils.PointString(config.User),
        },
        Version:          utils.PointUint32(1),
        CreateKeyRequest: &req,
    }
    resp, err := om.SubmitRequest(&wrapperRequest)
    if err != nil {
        return resp.CreateKeyResponse, err
    }
    return resp.CreateKeyResponse, nil
}

// ConvertBlockId TODO
func ConvertBlockId(bid *hdds.BlockID) *dnproto.DatanodeBlockID {
    id := dnproto.DatanodeBlockID{
        ContainerID:           bid.ContainerBlockID.ContainerID,
        LocalID:               bid.ContainerBlockID.LocalID,
        BlockCommitSequenceId: bid.BlockCommitSequenceId,
    }
    return &id
}

// PutKey TODO
func (om *OmClient) PutKey(volume string, bucket string, key string, length int64, source io.Reader) error {
    panic("")
}

// GetKey TODO
func (om *OmClient) GetKey(volume, bucket, key string, destination io.Writer) error {
    panic("")
}

// CommitKey TODO
func (om *OmClient) CommitKey(volume string, bucket string, key string, id *uint64,
    keyLocations []*ozoneproto.KeyLocation, dataSize uint64) (*ozoneproto.OMResponse, error) {
    factor := hdds.ReplicationFactor_THREE.Enum()
    typ := hdds.ReplicationType_RATIS.Enum()
    req := &ozoneproto.CommitKeyRequest{
        KeyArgs: &ozoneproto.KeyArgs{
            VolumeName:   &volume,
            BucketName:   &bucket,
            KeyName:      &key,
            DataSize:     &dataSize,
            Type:         typ,
            Factor:       factor,
            KeyLocations: keyLocations,
        },
        ClientID: id,
    }

    cmdType := ozoneproto.Type_CommitKey.Enum()
    wrapperRequest := ozoneproto.OMRequest{
        CmdType:          cmdType,
        CommitKeyRequest: req,
        ClientId:         &om.ClientId,
    }
    return om.SubmitRequest(&wrapperRequest)
}

// ListKeys TODO
func (om *OmClient) ListKeys(volume string, bucket string, prefix string, limit int32) ([]*common.Key, error) {
    req := ozoneproto.ListKeysRequest{
        VolumeName: &volume,
        BucketName: &bucket,
        Prefix:     utils.PointString(prefix),
        Count:      utils.PointInt32(limit),
        StartKey:   utils.PointString(""),
    }

    listKeys := ozoneproto.Type_ListKeys
    wrapperRequest := ozoneproto.OMRequest{
        CmdType:         &listKeys,
        ListKeysRequest: &req,
        ClientId:        &om.ClientId,
    }

    resp, err := om.SubmitRequest(&wrapperRequest)
    if err != nil {
        return nil, err
    }
    keys := make([]*common.Key, 0)
    for _, r := range resp.GetListKeysResponse().GetKeyInfo() {
        keys = append(keys, KeyFromProto(r))
    }
    return keys, nil
}

// KeyFromProto TODO
func KeyFromProto(keyProto *ozoneproto.KeyInfo) *common.Key {
    rt := keyProto.GetType().String()
    rf := keyProto.GetFactor().Number()
    size := keyProto.GetDataSize()
    ct := utils.MillisecondFormatToString(int64(keyProto.GetCreationTime()), utils.TimeFormatterSECOND)
    mt := utils.MillisecondFormatToString(int64(keyProto.GetModificationTime()), utils.TimeFormatterSECOND)
    key := &common.Key{
        VolumeName:        keyProto.GetVolumeName(),
        BucketName:        keyProto.GetBucketName(),
        Name:              keyProto.GetKeyName(),
        DataSize:          size,
        DataSizeFriendly:  utils.BytesToHuman(size),
        CreationTime:      ct,
        ModificationTime:  mt,
        ReplicationType:   common.ReplicationType(rt),
        ReplicationFactor: common.ReplicationFactor(rf),
    }
    return key
}

// DeleteKey TODO
func (om *OmClient) DeleteKey(volume, bucket, key string) (string, error) {
    req := ozoneproto.DeleteKeyRequest{
        KeyArgs: &ozoneproto.KeyArgs{
            VolumeName: &volume,
            BucketName: &bucket,
            KeyName:    &key,
        },
    }

    listKeys := ozoneproto.Type_DeleteKey
    wrapperRequest := ozoneproto.OMRequest{
        CmdType:          &listKeys,
        DeleteKeyRequest: &req,
        ClientId:         &om.ClientId,
    }

    resp, err := om.SubmitRequest(&wrapperRequest)
    if err != nil {
        return resp.GetDeleteKeyResponse().String(), err
	}
	return resp.GetDeleteKeyResponse().GetKeyInfo().GetKeyName(), nil
}
