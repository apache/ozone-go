// Package common TODO
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
package common

import (
    "encoding/json"
    "fmt"
    "os"
    "github.com/apache/ozone-go/api/config"
    "github.com/apache/ozone-go/api/proto/hdds"
    "github.com/apache/ozone-go/api/proto/ozone"
    "github.com/apache/ozone-go/api/utils"
    "path"
    "strings"

    log "github.com/sirupsen/logrus"
)

var sep = string(os.PathSeparator)

// ReplicationType TODO
type ReplicationType string

// ReplicationFactor TODO
type ReplicationFactor int8

const (
    // RATIS TODO
    RATIS ReplicationType = "RATIS"
    // STANDALONE TODO
    STANDALONE = "STANDALONE"
    // CHAINED TODO
    CHAINED = "CHAINED"
)

const (
    // ONE TODO
    ONE ReplicationFactor = 1
    // THREE TODO
    THREE ReplicationFactor = 3
)

// FriendlyFileInfo TODO
type FriendlyFileInfo interface {
    FriendlyFileInfoString() string
}

// Acl client acl struct
type Acl struct {
    Type     string   `json:"type"`
    Name     string   `json:"name"`
    AclScope string   `json:"acl_scope"`
    AclList  []string `json:"acl_list"`
}

// Volume client volume struct
type Volume struct {
    VolumeName       string `json:"volume_name"`
    Owner            string `json:"owner"`
    QuotaInBytes     string `json:"quota_in_bytes"`
    QuotaInNamespace int64  `json:"quota_in_namespace"`
    UsedNamespace    uint64 `json:"used_namespace"`
    CreationTime     string `json:"creation_time"`
    ModificationTime string `json:"modification_time"`
    Acls             []Acl  `json:"acls"`
}

// String json volume string
func (v Volume) String() string {
    return toJsonString(v)
}

// FriendlyFileInfoString friendly volume fs string
func (v Volume) FriendlyFileInfoString() string {
    mode := "drwx-rw-x"
    return fmt.Sprintf("%s \t %s \t %d \t %s \t %s \t %s \t %s", mode,
        v.QuotaInBytes, v.UsedNamespace, v.Owner, v.CreationTime, v.ModificationTime, sep+v.VolumeName)
}

// Bucket client bucket struct
type Bucket struct {
    VolumeName       string `json:"volume_name"`
    BucketName       string `json:"bucket_name"`
    StorageType      string `json:"storage_type"`
    QuotaInBytes     string `json:"quota_in_bytes"`
    UsedBytes        string `json:"used_bytes"`
    QuotaInNamespace int64  `json:"quota_in_namespace"`
    UsedNamespace    uint64 `json:"used_namespace"`
    CreationTime     string `json:"creation_time"`
    ModificationTime string `json:"modification_time"`
    BucketLayout     string `json:"bucket_layout"`
}

// String json bucket string
func (b Bucket) String() string {
    return toJsonString(b)
}

// FriendlyFileInfoString friendly bucket fs string
func (b Bucket) FriendlyFileInfoString() string {
    mode := "drwx-rw-x"
    return fmt.Sprintf("%s \t %s \t %d \t %s \t %s \t%s \t%s",
        mode, b.UsedBytes, b.UsedNamespace, b.CreationTime, b.ModificationTime, b.BucketLayout,
        sep+b.VolumeName+sep+b.BucketName)
}

// Key client key struct
type Key struct {
    VolumeName        string            `json:"volume_name"`
    BucketName        string            `json:"bucket_name"`
    Name              string            `json:"name"`
    DataSize          uint64            `json:"data_size"`
    DataSizeFriendly  string            `json:"data_size_friendly"`
    CreationTime      string            `json:"creation_time"`
    ModificationTime  string            `json:"modification_time"`
    ReplicationFactor ReplicationFactor `json:"replication_factor"`
    ReplicationType   ReplicationType   `json:"replication_type"`
}

// String json key string
func (k Key) String() string {
    return toJsonString(k)
}

// FriendlyFileInfoString friendly key fs string
func (k Key) FriendlyFileInfoString() string {
    mode := "orwx-rw-x"
    return fmt.Sprintf("%s \t %d \t %d \t %s \t %s \t%s \t%s",
        mode, k.ReplicationFactor, k.DataSize, k.DataSizeFriendly, k.CreationTime, k.ModificationTime,
        sep+k.VolumeName+sep+k.BucketName+sep+k.Name)
}

// OzoneFileInfo client ozone file info
type OzoneFileInfo struct {
    Name        string
    Path        string
    Mode        string
    Owner       string
    Group       string
    Replica     int
    Size        uint64
    ReplicaSize uint64
    ModifyTime  uint64
}

// String pretty print
func (o *OzoneFileInfo) String() string {
    mTime := utils.MillisecondFormatToString(int64(o.ModifyTime), utils.TimeFormatterSECOND)
    return fmt.Sprintf("%s \t%d \t%d \t %d \t %s \t%s \t%s \t%s",
        o.Mode, o.Replica, o.Size, o.ReplicaSize, o.Owner, o.Group, mTime, o.Path)
}

// FriendlyFileInfoString convert size to human read
func (o *OzoneFileInfo) FriendlyFileInfoString() string {
    size := utils.BytesToHuman(o.Size)
    replicaSize := utils.BytesToHuman(o.Size * uint64(o.Replica))
    mTime := utils.MillisecondFormatToString(int64(o.ModifyTime), utils.TimeFormatterSECOND)
    return fmt.Sprintf("%s \t%d \t%s \t %s \t %s \t%s \t%s \t%s",
        o.Mode, o.Replica, size, replicaSize, o.Owner, o.Group, mTime, o.Path)
}

// OzoneObjectAddress object convert struct
type OzoneObjectAddress struct {
    Om       string `json:"om"`
    Volume   string `json:"volume"`
    Bucket   string `json:"bucket"`
    Key      string `json:"key"`
    FilePath string `json:"file_path"`
}

// OzoneObjectAddressFromFsPath file system path like:
// /vol/bucket/key key(relative path)
func OzoneObjectAddressFromFsPath(originPath string) OzoneObjectAddress {
    om := ""
    volumeName := ""
    bucketName := ""
    keyName := ""
    if !strings.HasPrefix(originPath, sep) {
        originPath = sep + "user" + sep + config.User + sep + originPath
    }
    cleanPath := path.Clean(originPath)
    arr := strings.SplitN(cleanPath, sep, 4)
    if len(arr) == 4 {
        volumeName = arr[1]
        bucketName = arr[2]
        keyName = arr[3]
    } else if len(arr) == 3 {
        volumeName = arr[1]
        bucketName = arr[2]
    } else if len(arr) == 2 {
        volumeName = arr[1]
    }
    return OzoneObjectAddress{
        Om:       om,
        FilePath: cleanPath,
        Volume:   volumeName,
        Bucket:   bucketName,
        Key:      keyName,
    }
}

// OzoneObjectAddressFromPath like: vol/bucket/key
func OzoneObjectAddressFromPath(originPath string) OzoneObjectAddress {
    return OzoneObjectAddressFromFsPath(originPath)
}

// IsEmpty TODO
func (o *OzoneObjectAddress) IsEmpty() bool {
    return len(o.Om) == 0 && len(o.Key) == 0 && len(o.Bucket) == 0 && len(o.Volume) == 0
}

// String TODO
func (o *OzoneObjectAddress) String() string {
    return toJsonString(o)
}

// IsRootPath TODO
func (o *OzoneObjectAddress) IsRootPath() bool {
    return o.FilePath == sep
}

// FileSystemPath TODO
func (o *OzoneObjectAddress) FileSystemPath() string {
    return o.FilePath
}

func toJsonString(i interface{}) string {
    out, err := json.MarshalIndent(i, "", "   ")
    if err != nil {
        log.Error("format to json string error: ", err.Error())
        return "{}"
    }
    return string(out)
}

// OmKeyLocationInfo TODO
type OmKeyLocationInfo struct {
    BlockID *hdds.BlockID
    // the id of this subkey in all the subkeys.
    Length        uint64
    Offset        uint64
    CreateVersion uint64
    Pipeline      *hdds.Pipeline
    // PartNumber is set for Multipart upload Keys.
    PartNumber int32
}

// KeyArgs TODO
type KeyArgs struct {
    VolumeName                string
    BucketName                string
    KeyName                   string
    dataSize                  uint64
    Replication               *hdds.ReplicationType
    LocationInfoList          []*OmKeyLocationInfo
    IsMultipartKey            bool
    MultipartUploadID         string
    MultipartUploadPartNumber int32
    Metadata                  map[string]string
    RefreshPipeline           bool
    SortDatanodesInPipeline   bool
    Acls                      []*ozone.OzoneAclInfo
    LatestVersionLocation     bool
	Recursive                 bool
	HeadOp                    bool
}
