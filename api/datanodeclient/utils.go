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
    "github.com/apache/ozone-go/api/proto/datanode"
    "github.com/apache/ozone-go/api/proto/ozone"
)

var readTypeList = []string{datanode.Type_ReadContainer.String(),
    datanode.Type_ReadChunk.String(),
    datanode.Type_ListBlock.String(),
    datanode.Type_GetBlock.String(),
    datanode.Type_GetSmallFile.String(),
    datanode.Type_ListContainer.String(),
    datanode.Type_ListChunk.String(),
    datanode.Type_GetCommittedBlockLength.String(),
    ozone.Type_InfoVolume.String(),
    ozone.Type_InfoBucket.String(),
    ozone.Type_LookupKey.String(),
    ozone.Type_CheckVolumeAccess.String(),
    ozone.Type_ListVolume.String(),
    ozone.Type_ListBuckets.String(),
    ozone.Type_ListKeys.String(),
    ozone.Type_ListTrash.String(),
    ozone.Type_ServiceList.String(),
    ozone.Type_ListMultipartUploads.String(),
    ozone.Type_ListMultiPartUploadParts.String(),
    ozone.Type_GetFileStatus.String(),
    ozone.Type_ListStatus.String(),
    ozone.Type_GetAcl.String(),
    ozone.Type_DBUpdates.String(),
    ozone.Type_FinalizeUpgrade.String(),
    ozone.Type_PrepareStatus.String()}

var writeTypeList = []string{datanode.Type_CloseContainer.String(),
    datanode.Type_WriteChunk.String(),
    datanode.Type_UpdateContainer.String(),
    datanode.Type_CompactChunk.String(),
    datanode.Type_CreateContainer.String(),
    datanode.Type_DeleteChunk.String(),
    datanode.Type_DeleteContainer.String(),
    datanode.Type_DeleteBlock.String(),
    datanode.Type_PutBlock.String(),
    datanode.Type_PutSmallFile.String(),
    ozone.Type_CreateKey.String(),
    ozone.Type_CreateKey.String(),
    ozone.Type_CreateFile.String(),
    ozone.Type_CreateBucket.String(),
    ozone.Type_CreateVolume.String(),
    ozone.Type_SetVolumeProperty.String(),
    ozone.Type_DeleteVolume.String(),
    ozone.Type_SetBucketProperty.String(),
    ozone.Type_DeleteBucket.String(),
    ozone.Type_RenameKeys.String(),
    ozone.Type_RenameKey.String(),
    ozone.Type_DeleteKeys.String(),
    ozone.Type_DeleteKey.String(),
    ozone.Type_DeleteOpenKeys.String(),
    ozone.Type_CommitKey.String(),
    ozone.Type_InitiateMultiPartUpload.String(),
    ozone.Type_CommitMultiPartUpload.String(),
    ozone.Type_CompleteMultiPartUpload.String(),
    ozone.Type_AbortMultiPartUpload.String(),
    ozone.Type_RemoveAcl.String(),
    ozone.Type_SetAcl.String(),
    ozone.Type_AddAcl.String(),
    ozone.Type_PurgeKeys.String(),
    ozone.Type_PurgePaths.String(),
    ozone.Type_DeleteOpenKeys.String()}

// IsReadOnly TODO
func IsReadOnly(cmdType string) bool {
    for _, readTye := range readTypeList {
        if readTye == cmdType {
            return true
        }
    }
    for _, writeTye := range writeTypeList {
        if writeTye == cmdType {
            return false
        }
    }
    return true
}
