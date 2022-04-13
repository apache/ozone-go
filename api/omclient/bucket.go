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
    ozone_proto "github.com/apache/ozone-go/api/proto/ozone"
    "github.com/apache/ozone-go/api/utils"
)

// CreateBucket TODO
func (om *OmClient) CreateBucket(volume string, bucket string) error {
    isVersionEnabled := false
    storageType := ozone_proto.StorageTypeProto_DISK
    bucketInfo := ozone_proto.BucketInfo{
        VolumeName:       &volume,
        BucketName:       &bucket,
        QuotaInNamespace: utils.PointInt64(-1),
        Acls:             CreateAcls(ALL),
        IsVersionEnabled: &isVersionEnabled,
        StorageType:      &storageType,
        BucketLayout:     ozone_proto.BucketLayoutProto_LEGACY.Enum(),
    }
    req := ozone_proto.CreateBucketRequest{
        BucketInfo: &bucketInfo,
    }

    cmdType := ozone_proto.Type_CreateBucket
    wrapperRequest := ozone_proto.OMRequest{
        CmdType:             &cmdType,
        CreateBucketRequest: &req,
        ClientId:            &om.ClientId,
    }

    _, err := om.SubmitRequest(&wrapperRequest)
    return err
}

// InfoBucket TODO
func (om *OmClient) InfoBucket(volume string, bucket string) (*common.Bucket, error) {
    req := &ozone_proto.InfoBucketRequest{
        VolumeName: &volume,
        BucketName: &bucket,
    }

    cmdType := ozone_proto.Type_InfoBucket.Enum()
    wrapperRequest := ozone_proto.OMRequest{
        CmdType:           cmdType,
        InfoBucketRequest: req,
        ClientId:          &om.ClientId,
    }

    resp, err := om.SubmitRequest(&wrapperRequest)
    if err != nil {
        return nil, err
    }
    return bucketFromProto(resp.InfoBucketResponse.GetBucketInfo()), nil
}

// ListBucket TODO
func (om *OmClient) ListBucket(volume string) ([]*common.Bucket, error) {
    res := make([]*common.Bucket, 0)

    req := ozone_proto.ListBucketsRequest{
        VolumeName: &volume,
        StartKey:   utils.PointString(""),
        Count:      utils.PointInt32(100),
    }

    cmdType := ozone_proto.Type_ListBuckets
    wrapperRequest := ozone_proto.OMRequest{
        CmdType:            &cmdType,
        ListBucketsRequest: &req,
        ClientId:           &om.ClientId,
    }

    resp, err := om.SubmitRequest(&wrapperRequest)
    if err != nil {
        return nil, err
    }
    for _, bktProto := range resp.ListBucketsResponse.GetBucketInfo() {
        res = append(res, bucketFromProto(bktProto))
    }
    return res, nil
}

func bucketFromProto(bktProto *ozone_proto.BucketInfo) *common.Bucket {
    ct := utils.MillisecondFormatToString(int64(bktProto.GetCreationTime()), utils.TimeFormatterSECOND)
    mt := utils.MillisecondFormatToString(int64(bktProto.GetModificationTime()), utils.TimeFormatterSECOND)
    return &common.Bucket{
        VolumeName:       bktProto.GetVolumeName(),
        BucketName:       bktProto.GetBucketName(),
        StorageType:      bktProto.GetStorageType().String(),
        QuotaInBytes:     utils.BytesToHuman(uint64(bktProto.GetQuotaInBytes())),
        UsedBytes:        utils.BytesToHuman(bktProto.GetUsedBytes()),
        QuotaInNamespace: bktProto.GetQuotaInNamespace(),
        UsedNamespace:    bktProto.GetUsedNamespace(),
        CreationTime:     ct,
        ModificationTime: mt,
        BucketLayout:     bktProto.GetBucketLayout().String(),
	}
}
