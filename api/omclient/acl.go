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
    "math"
    "github.com/apache/ozone-go/api/config"
    ozone_proto "github.com/apache/ozone-go/api/proto/ozone"
    "github.com/apache/ozone-go/api/utils"
    "strings"
)

// Acl TODO
type Acl uint

const (
    // READ TODO
    READ Acl = 1 << iota
    // WRITE TODO
    WRITE
    // CREATE TODO
    CREATE
    // LIST TODO
    LIST
    // DELETE TODO
    DELETE
    // READ_ACL TODO
    READ_ACL
    // WRITE_ACL TODO
    WRITE_ACL
    // ALL TODO
    ALL
    // NONE TODO
    NONE
)

var (
    // ACL_name TODO
    ACL_name = map[Acl]string{
        READ:      "READ",
        WRITE:     "WRITE",
        CREATE:    "CREATE",
        LIST:      "LIST",
        DELETE:    "DELETE",
        READ_ACL:  "READ_ACL",
        WRITE_ACL: "WRITE_ACL",
        ALL:       "ALL",
        NONE:      "NONE",
    }
    // ACL_value TODO
    ACL_value = map[string]Acl{
        "READ":      READ,
        "WRITE":     WRITE,
        "CREATE":    CREATE,
        "LIST":      LIST,
        "DELETE":    DELETE,
        "READ_ACL":  READ_ACL,
        "WRITE_ACL": WRITE_ACL,
        "ALL":       ALL,
        "NONE":      NONE,
    }
)

// ACL_LIST TODO
var ACL_LIST = []string{"READ", "WRITE", "CREATE", "LIST", "DELETE", "READ_ACL", "WRITE_ACL", "ALL", "NONE"}

// NewAcl TODO
func NewAcl(right uint) Acl {
    return Acl(right)
}

// NewAclFromString TODO
func NewAclFromString(right string) Acl {
    return ACL_value[right]
}

// GetRight TODO
func (acl Acl) GetRight() uint {
    return uint(acl)
}

// GetName TODO
func (acl Acl) GetName() string {
    return ACL_name[acl]
}

// AclInfo TODO
type AclInfo struct {
    Acls []Acl
}

// NewAclInfo TODO
func NewAclInfo() *AclInfo {
    return &AclInfo{
        Acls: make([]Acl, 0),
    }
}

// CreateAcls TODO
func CreateAcls(acl Acl) []*ozone_proto.OzoneAclInfo {
    aclInfo := NewAclInfo()
    aclInfo.AddRight(acl)
    rights := aclInfo.ToBytes()
    return []*ozone_proto.OzoneAclInfo{
        createOzoneAclInfo(ozone_proto.OzoneAclInfo_USER.Enum(), config.User, rights, ozone_proto.OzoneAclInfo_ACCESS.Enum()),
        createOzoneAclInfo(ozone_proto.OzoneAclInfo_GROUP.Enum(), config.User, rights,
            ozone_proto.OzoneAclInfo_ACCESS.Enum()),
    }
}

func createOzoneAclInfo(aclType *ozone_proto.OzoneAclInfo_OzoneAclType, name string, rights []byte,
    aclScope *ozone_proto.OzoneAclInfo_OzoneAclScope) *ozone_proto.OzoneAclInfo {
    return &ozone_proto.OzoneAclInfo{
        Type:     aclType,
        Name:     utils.PointString(name),
        Rights:   rights,
        AclScope: aclScope,
    }
}

// Len TODO
func (aclInfo *AclInfo) Len() int {
    return len(aclInfo.Acls)
}

// Less TODO
func (aclInfo *AclInfo) Less(i, j int) bool {
    return aclInfo.Acls[i] < aclInfo.Acls[j]
}

// Swap TODO
func (aclInfo *AclInfo) Swap(i, j int) {
    tmp := aclInfo.Acls[i]
    aclInfo.Acls[i] = aclInfo.Acls[j]
    aclInfo.Acls[j] = tmp
}

// GetRights TODO
func (aclInfo *AclInfo) GetRights() []Acl {
    return aclInfo.Acls
}

// SetRights TODO
func (aclInfo *AclInfo) SetRights(acls []Acl) {
    aclInfo.Acls = acls
}

// AddRight TODO
func (aclInfo *AclInfo) AddRight(acl Acl) {
    aclInfo.Acls = append(aclInfo.Acls, acl)
}

// Clear TODO
func (aclInfo *AclInfo) Clear() {
    aclInfo.Acls = make([]Acl, 0)
}

// FromBytes TODO
func (aclInfo *AclInfo) FromBytes(rights []byte) {
    rightNumber := uint(0)
    for i, right := range rights {
        rightNumber += uint(right) << (i * 8)
    }
    aclInfo.Acls = make([]Acl, 0)
    bitSet := utils.UintToBitSet(rightNumber)
    for i, b := range bitSet {
        if b == 1 {
            aclInfo.Acls = append(aclInfo.Acls, NewAcl(uint(1<<i)))
        }
    }
}

// ToBytes TODO
func (aclInfo *AclInfo) ToBytes() []byte {
    rightNumber := uint(0)
    for _, acl := range aclInfo.Acls {
        rightNumber += uint(acl)
    }
    data := make([]byte, 0)
    for rightNumber > 0 {
        data = append(data, byte(rightNumber&math.MaxUint8))
        rightNumber = rightNumber >> 8
    }
    return data
}

// String TODO
func (aclInfo *AclInfo) String() string {
    return strings.Join(aclInfo.Strings(), ",")
}

// Strings TODO
func (aclInfo *AclInfo) Strings() []string {
    names := make([]string, aclInfo.Len())
    for i, acl := range aclInfo.Acls {
		names[i] = acl.GetName()
	}
	return names
}
