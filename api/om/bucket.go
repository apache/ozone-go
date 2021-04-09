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

func (om *OmClient) CreateBucket(volume string, bucket string) error {
	isVersionEnabled := false
	storageType := ozone_proto.StorageTypeProto_DISK
	bucketInfo := ozone_proto.BucketInfo{
		BucketName:       &bucket,
		VolumeName:       &volume,
		IsVersionEnabled: &isVersionEnabled,
		StorageType:      &storageType,

	}
	req := ozone_proto.CreateBucketRequest{
		BucketInfo: &bucketInfo,
	}

	cmdType := ozone_proto.Type_CreateBucket
	wrapperRequest := ozone_proto.OMRequest{
		CmdType:             &cmdType,
		CreateBucketRequest: &req,
		ClientId:            &om.clientId,
	}

	_, err := om.submitRequest(&wrapperRequest)
	if err != nil {
		return err
	}
	return nil
}

func (om *OmClient) GetBucket(volume string, bucket string) (common.Bucket, error) {
	req := ozone_proto.InfoBucketRequest{
		VolumeName: &volume,
		BucketName: &bucket,
	}

	cmdType := ozone_proto.Type_InfoBucket
	wrapperRequest := ozone_proto.OMRequest{
		CmdType:           &cmdType,
		InfoBucketRequest: &req,
		ClientId:          &om.clientId,
	}

	resp, err := om.submitRequest(&wrapperRequest)
	if err != nil {
		return common.Bucket{}, err
	}
	b := common.Bucket{
		Name:       *resp.InfoBucketResponse.BucketInfo.BucketName,
		VolumeName: *resp.InfoBucketResponse.BucketInfo.VolumeName,
	}
	return b, nil
}

func (om *OmClient) ListBucket(volume string) ([]common.Bucket, error) {
	res := make([]common.Bucket, 0)

	req := ozone_proto.ListBucketsRequest{
		VolumeName: &volume,
		StartKey:   ptr(""),
		Count:      ptri(100),
	}

	cmdType := ozone_proto.Type_ListBuckets
	wrapperRequest := ozone_proto.OMRequest{
		CmdType:            &cmdType,
		ListBucketsRequest: &req,
		ClientId:           &om.clientId,
	}

	resp, err := om.submitRequest(&wrapperRequest)
	if err != nil {
		return res, err
	}
	for _, b := range resp.ListBucketsResponse.BucketInfo {
		cb := common.Bucket{
			Name:       *b.BucketName,
			VolumeName: *b.VolumeName,
		}
		res = append(res, cb)
	}
	return res, nil
}


