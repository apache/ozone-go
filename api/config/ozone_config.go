// Package config TODO
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
package config

import (
    "fmt"
    "strings"

    "github.com/Mengqi777/xmlconfig"
    "github.com/sirupsen/logrus"
    "github.com/apache/ozone-go/api/utils"
)

var (
    // ConfFilePath TODO
    ConfFilePath string
    // OmAddress TODO
    OmAddress string
    // OmAddresses TODO
    OmAddresses = []string{"localhost:9870"}
    // LogLevel TODO
    LogLevel = "INFO"
    // User TODO
    User = "guest"
)

// OzoneConfiguration TODO
type OzoneConfiguration struct {
    config      *xmlconfig.XmlConfig
    ConfigPath  string
    Data        string
    OmNodeIndex int
}

// OzoneConfig TODO
var OzoneConfig = OzoneConfiguration{
    config:      xmlconfig.NewXmlConfig(),
    OmNodeIndex: -1,
}

// InitLog TODO
func InitLog() error {
    level, err := logrus.ParseLevel(LogLevel)
    if err != nil {
        return err
    }
    logrus.SetLevel(level)
    return nil
}

// String TODO
func (o *OzoneConfiguration) String() string {
    return fmt.Sprintf("{"+
        "\"ConfigPath\":\"%s\","+
        "\"Data\":\"%s\","+
        "\"OmNodeIndex\":\"%d\","+
        "\"config\":\"%s\",}", o.ConfigPath, o.Data, o.OmNodeIndex, o.config.String())
}

// LoadFromFile TODO
func (o *OzoneConfiguration) LoadFromFile() error {
    if o.config == nil {
        o.config = new(xmlconfig.XmlConfig)
    }
    return o.config.ReadXmlFile(o.ConfigPath)
}

// NextOmNode TODO
func (o *OzoneConfiguration) NextOmNode() string {
    o.OmNodeIndex = (o.OmNodeIndex + 1) % len(OmAddresses)
    return OmAddresses[o.OmNodeIndex]
}

// GetOmAddresses TODO
func (o *OzoneConfiguration) GetOmAddresses() []string {
    serviceIds := o.config.GetTrimmedStrings(OZONE_OM_SERVICE_IDS_KEY, ",")
    if len(serviceIds) == 0 {
        panic(fmt.Errorf("%s not exist", OZONE_OM_SERVICE_IDS_KEY))
    }
    addresses := make([]string, 0)
    for _, serviceId := range serviceIds {
        nodes := o.config.GetStrings(strings.Join([]string{OZONE_OM_NODES_KEY, serviceId}, "."), ",")
        for _, node := range nodes {
            key := strings.Join([]string{OZONE_OM_ADDRESS_KEY, serviceId, node}, ".")
            if addr := o.config.GetTrimmedString(key, ""); len(addr) > 0 {
                addresses = append(addresses, addr)
            }
        }
    }
    logrus.Debug("om addresses: ", addresses)
    return addresses
}

// ParseStorageSize TODO
func (o *OzoneConfiguration) ParseStorageSize(value string) (uint64, error) {
    ss := &utils.StorageSize{
        Uint:  0,
        Value: 0,
    }
    err := ss.Parse(value)
    return ss.Size(), err
}

// GetChunkSize TODO
func (o *OzoneConfiguration) GetChunkSize() (uint64, error) {
    cs := o.config.GetString(OZONE_SCM_CHUNK_SIZE, OZONE_SCM_CHUNK_SIZE_DEFAULT)
    return o.ParseStorageSize(cs)
}

// GetBytesPerChecksum TODO
func (o *OzoneConfiguration) GetBytesPerChecksum() (uint64, error) {
    cs := o.config.GetString(OZONE_CLIENT_BYTES_PER_CHECKSUM, OZONE_CLIENT_BYTES_PER_CHECKSUM_DEFAULT)
    return o.ParseStorageSize(cs)
}

// GetBlockSize TODO
func (o *OzoneConfiguration) GetBlockSize() (uint64, error) {
    bs := o.config.GetString(OZONE_SCM_BLOCK_SIZE, OZONE_SCM_CHUNK_SIZE_DEFAULT)
    return o.ParseStorageSize(bs)
}

// GetStreamBufferFlushSize TODO
func (o *OzoneConfiguration) GetStreamBufferFlushSize() (uint64, error) {
    bs := o.config.GetString(OZONE_SCM_BLOCK_SIZE, OZONE_SCM_CHUNK_SIZE_DEFAULT)
    return o.ParseStorageSize(bs)
}

// GetStreamBufferMaxSize TODO
func (o *OzoneConfiguration) GetStreamBufferMaxSize() (uint64, error) {
    bs := o.config.GetString(OZONE_SCM_BLOCK_SIZE, OZONE_SCM_CHUNK_SIZE_DEFAULT)
    return o.ParseStorageSize(bs)
}

// GetReplicationType TODO
func (o *OzoneConfiguration) GetReplicationType() string {
    return o.config.GetString(OZONE_REPLICATION_TYPE_KEY, OZONE_REPLICATION_TYPE_DEFAULT)
}

// GetReplicationFactor TODO
func (o *OzoneConfiguration) GetReplicationFactor() (int8, error) {
    return o.config.GetInt8(OZONE_REPLICATION_KEY, OZONE_REPLICATION_DEFAULT)
}

// Set TODO
func (o *OzoneConfiguration) Set(key, value string) {
    o.config.SetString(key, value)
}

// GetString TODO
func (o *OzoneConfiguration) GetString(key, defaultString string) string {
    return o.config.GetString(key, defaultString)
}

// GetBool TODO
func (o *OzoneConfiguration) GetBool(key string, defaultBool bool) bool {
    return o.config.GetBool(key, defaultBool)
}

// GetChecksumType TODO
func (o *OzoneConfiguration) GetChecksumType() string {
    return o.config.GetString(OZONE_CLIENT_CHECKSUM_TYPE, OZONE_CLIENT_CHECKSUM_TYPE_DEFAULT)
}
