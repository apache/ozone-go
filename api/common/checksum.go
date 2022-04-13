// Package common TODO
// Package checksum
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
    "crypto/md5"
    "crypto/sha256"
    "hash/crc32"
    "github.com/apache/ozone-go/api/proto/datanode"
    "github.com/apache/ozone-go/api/utils"
    "strings"
)

// CRC32 crc32 checksum
func CRC32(data []byte) []byte {
    var b = make([]byte, 8)
    utils.Uint64toBytes(b, uint64(crc32.ChecksumIEEE(data)))
    return b
}

// CRC32c crc32c checksum
func CRC32c(data []byte) []byte {
    var b = make([]byte, 8)
    t := crc32.MakeTable(crc32.Castagnoli)
    utils.Uint64toBytes(b, uint64(crc32.Checksum(data, t)))
    return b
}

// SHA256 sha256 checksum
func SHA256(data []byte) []byte {
    c := sha256.New()
    if _, err := c.Write(data); err != nil {
        return nil
    }
    return c.Sum(nil)
}

// MD5 md5 checksum
func MD5(data []byte) []byte {
    c := md5.New()
    if _, err := c.Write(data); err != nil {
        return nil
    }
    return c.Sum(nil)
}

// ChecksumTypeFromName get checksum type by name
func ChecksumTypeFromName(name string) datanode.ChecksumType {
    name = strings.ToLower(name)
    switch name {
    case "crc32":
        return datanode.ChecksumType_CRC32
    case "crc32c":
        return datanode.ChecksumType_CRC32C
    case "md5":
        return datanode.ChecksumType_MD5
    case "sha256":
        return datanode.ChecksumType_SHA256
    default:
        return datanode.ChecksumType_NONE
    }
}

// ChecksumTypeFromCode get checksum type by type code
func ChecksumTypeFromCode(typeCode int32) datanode.ChecksumType {
    name := datanode.ChecksumType_name[typeCode]
    return ChecksumTypeFromName(name)
}
