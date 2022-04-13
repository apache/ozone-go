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
    "fmt"
    "io"
    "github.com/apache/ozone-go/api/common"
    "github.com/apache/ozone-go/api/config"
    ozoneproto "github.com/apache/ozone-go/api/proto/ozone"
    "github.com/apache/ozone-go/api/utils"
    "strconv"
    "strings"
    "time"
)

// Mkdir TODO
func (om *OmClient) Mkdir(volume, bucket, key string) error {
    req := &ozoneproto.CreateDirectoryRequest{
        KeyArgs: &ozoneproto.KeyArgs{
            VolumeName: &volume,
            BucketName: &bucket,
            KeyName:    &key,
        },
    }
    cmdType := ozoneproto.Type_CreateDirectory.Enum()
    wrapperRequest := ozoneproto.OMRequest{
        CmdType:                cmdType,
        CreateDirectoryRequest: req,
        ClientId:               &om.ClientId,
    }
    _, err := om.SubmitRequest(&wrapperRequest)
    return err
}

// RenameFile TODO
func (om *OmClient) RenameFile(volume, bucket, fromKey, toKey string) error {
    ret, err := om.RenameKey(volume, bucket, fromKey, toKey)
    if ret == "" {
        if err == nil {
            return err
        }
    } else if len(ret) > 0 {
        if err == nil {
            return fmt.Errorf("rename file %s not complete", ret)
        } else {
            return fmt.Errorf("rename file %s not complete, error: %s", ret, err.Error())
        }
    }
    return nil
}

// DeleteFile TODO
func (om *OmClient) DeleteFile(volume, bucket, key string, skipTrash bool) (string, error) {
    if skipTrash {
        return om.DeleteKey(volume, bucket, key)
    } else {
        timeNow := time.Now().UnixNano() / 1e6
        trashPath := "user/" + config.User + "/.Trash/" + key + strconv.Itoa(int(timeNow))
        return trashPath, om.RenameFile(volume, bucket, key, trashPath)
    }
}

// CreateFile TODO
func (om *OmClient) CreateFile(volume, bucket, key string) error {
    _, err := om.CreateEmptyFile(volume, bucket, key, false, false)
    return err
}

// CreateEmptyFile TODO
func (om *OmClient) CreateEmptyFile(volume, bucket, key string, recursive, overwrite bool) (
    *ozoneproto.CreateKeyResponse, error) {
    timeNow := time.Now().UnixNano() / 1e6
    req := &ozoneproto.CreateFileRequest{
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
        IsRecursive: utils.PointBool(recursive),
        IsOverwrite: utils.PointBool(overwrite),
        ClientID:    utils.PointUint64(0),
    }
    cmdType := ozoneproto.Type_CreateFile.Enum()
    wrapperRequest := ozoneproto.OMRequest{
        CmdType:           cmdType,
        CreateFileRequest: req,
        ClientId:          &om.ClientId,
    }
    resp, err := om.SubmitRequest(&wrapperRequest)
    if err != nil {
        return nil, err
    }
    return resp.GetCreateKeyResponse(), nil
}

// PutFile TODO
func (om *OmClient) PutFile(volume string, bucket string, key string, length int64, source io.Reader) error {
    panic("")
}

// GetFile TODO
func (om *OmClient) GetFile(volume, bucket, key string, destination io.Writer) error {
    panic("")
}

// LookupFile TODO
func (om *OmClient) LookupFile(volume, bucket, key string) (*common.OzoneFileInfo, error) {
    req := &ozoneproto.LookupFileRequest{KeyArgs: &ozoneproto.KeyArgs{
        VolumeName: utils.PointString(volume),
        BucketName: utils.PointString(bucket),
        KeyName:    utils.PointString(key),
    }}
    cmdType := ozoneproto.Type_LookupFile.Enum()
    wrapperRequest := ozoneproto.OMRequest{
        CmdType:           cmdType,
        LookupFileRequest: req,
        ClientId:          &om.ClientId,
    }
    resp, err := om.SubmitRequest(&wrapperRequest)
    if err != nil {
        return nil, err
    }
    bs, err := config.OzoneConfig.GetBlockSize()
    if err != nil {
        return nil, err
    }
    ofs := &ozoneproto.OzoneFileStatusProto{
        KeyInfo:     resp.GetLookupFileResponse().GetKeyInfo(),
        BlockSize:   utils.PointUint64(bs),
        IsDirectory: utils.PointBool(false),
    }
    return FileFromProto(ofs), nil
}

// ListFiles TODO
func (om *OmClient) ListFiles(volume, bucket, key, startKey string, limit uint64) ([]*common.OzoneFileInfo, error) {
    req := ozoneproto.ListStatusRequest{
        KeyArgs: &ozoneproto.KeyArgs{
            VolumeName: utils.PointString(volume),
            BucketName: utils.PointString(bucket),
            KeyName:    utils.PointString(key),
        },
        StartKey:   utils.PointString(startKey),
        NumEntries: utils.PointUint64(limit),
        Recursive:  utils.PointBool(false),
    }

    listKeys := ozoneproto.Type_ListStatus
    wrapperRequest := ozoneproto.OMRequest{
        CmdType:           &listKeys,
        ListStatusRequest: &req,
        ClientId:          &om.ClientId,
    }
    resp, err := om.SubmitRequest(&wrapperRequest)
    list := resp.GetListStatusResponse().GetStatuses()
    var res = make([]*common.OzoneFileInfo, 0)
    for _, proto := range list {
        res = append(res, FileFromProto(proto))
    }
    return res, err
}

// FileFromProto TODO
func FileFromProto(proto *ozoneproto.OzoneFileStatusProto) *common.OzoneFileInfo {

    keyProto := proto.GetKeyInfo()
    replica := int(keyProto.GetFactor().Number())
    size := keyProto.GetDataSize()
    mode := "-"
    if proto.GetIsDirectory() {
        mode = "d"
    }
    owner := "guest"
    group := "guest"
    for _, info := range keyProto.GetAcls() {
        if info.GetType() == ozoneproto.OzoneAclInfo_USER {
            owner = info.GetName()
            aclInfo := NewAclInfo()
            aclInfo.FromBytes(info.GetRights())
            mode = mode + aclInfo.String()
        }
        if info.GetType() == ozoneproto.OzoneAclInfo_GROUP {
            group = info.GetName()
            aclInfo := NewAclInfo()
            aclInfo.FromBytes(info.GetRights())
            mode = mode + aclInfo.String()
        }
    }
    key := &common.OzoneFileInfo{
        Name: keyProto.GetKeyName(),
        Path: strings.Join([]string{"", keyProto.GetVolumeName(), keyProto.GetBucketName(), keyProto.GetKeyName()},
            "/"),
        Replica:     replica,
        Size:        size,
        ReplicaSize: size * uint64(replica),
        Mode:        mode,
        Owner:       owner,
        Group:       group,
		ModifyTime:  keyProto.GetModificationTime(),
	}
	return key
}
