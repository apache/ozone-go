// Package api TODO
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
package api

import (
    "errors"
    "fmt"
    "github.com/apache/ozone-go/api/common"
    "github.com/apache/ozone-go/api/config"
    ozoneIo "github.com/apache/ozone-go/api/io"
    "io"
)

// ListKeys TODO
func (ozoneClient *OzoneClient) ListKeys(volume string, bucket string, prefix string, limit int32) ([]*common.Key,
    error) {
    return ozoneClient.OmClient.ListKeys(volume, bucket, prefix, limit)
}

// ListFiles TODO
func (ozoneClient *OzoneClient) ListFiles(volume, bucket, key, startKey string, limit uint64) ([]*common.OzoneFileInfo,
    error) {
    return ozoneClient.OmClient.ListFiles(volume, bucket, key, startKey, limit)
}

// InfoKey TODO
func (ozoneClient *OzoneClient) InfoKey(volume, bucket, key string) (*common.Key, error) {
    return ozoneClient.OmClient.InfoKey(volume, bucket, key)
}

// RenameKey TODO
func (ozoneClient *OzoneClient) RenameKey(volume, bucket, key, toKeyName string) (string, error) {
    return ozoneClient.OmClient.RenameKey(volume, bucket, key, toKeyName)
}

// GetKey TODO
func (ozoneClient *OzoneClient) GetKey(volume, bucket, key string, destination io.Writer) error {
    keyInfo, err := ozoneClient.OmClient.LookupKey(volume, bucket, key)
    if err != nil {
        return err
    }
    if keyInfo == nil {
        return fmt.Errorf("%s key not found", key)
    }
    if len(keyInfo.KeyLocationList) == 0 {
        return errors.New("Get key returned with zero key location version " + volume + "/" + bucket + "/" +
            key)
    }
    if len(keyInfo.KeyLocationList[0].KeyLocations) == 0 {
        return errors.New("Key location doesn't have any datanode for key " + volume + "/" + bucket + "/" + key)
    }
    verifyChecksum := config.OzoneConfig.GetBool(config.OZONE_VERIFY_CHECKSUM_KEY, config.OZONE_VERIFY_CHECKSUM_DEFAULT)
    keyInputStream := ozoneIo.NewKeyInputStream(keyInfo, ozoneClient.xceiverManager, verifyChecksum)
    var n int64
    if n, err = io.Copy(destination, keyInputStream); err != nil && err != io.EOF {
        return err
    }
    dataSize := keyInfo.GetDataSize()
    if uint64(n) != dataSize {
        return fmt.Errorf("read data error: key data size %d but read len %d", n, dataSize)
    } else {
        return nil
    }
}

// DeleteKey TODO
func (ozoneClient *OzoneClient) DeleteKey(volume, bucket, key string) (string, error) {
    return ozoneClient.OmClient.DeleteKey(volume, bucket, key)
}

// TouchzKey TODO
func (ozoneClient *OzoneClient) TouchzKey(volume, bucket, key string) (*common.Key, error) {
    return ozoneClient.OmClient.TouchzKey(volume, bucket, key)
}

// PutKey TODO
func (ozoneClient *OzoneClient) PutKey(volume string, bucket string, key string, length int64, source io.Reader) error {

    createKey, err := ozoneClient.OmClient.CreateKey(volume, bucket, key)
    if err != nil {
        return err
    }
    ozoneClient.OmClient.Id = createKey.GetID()
    keyInfo := createKey.GetKeyInfo()
    bpc, err := config.OzoneConfig.GetBytesPerChecksum()
    if err != nil {
        return err
    }
    ckt := config.OzoneConfig.GetChecksumType()
    checksumType := common.ChecksumTypeFromName(ckt).Enum()
    bs, err := config.OzoneConfig.GetBlockSize()
    if err != nil {
        return err
    }
    cs, err := config.OzoneConfig.GetChunkSize()
    if err != nil {
        return err
    }
    sbfs, err := config.OzoneConfig.GetStreamBufferFlushSize()
    if err != nil {
        return err
    }
    sbms, err := config.OzoneConfig.GetStreamBufferMaxSize()
    if err != nil {
        return err
    }
    keyOutputStream, err := ozoneIo.NewKeyOutputStream(keyInfo, ozoneClient.xceiverManager, volume, bucket, key,
        ozoneClient.OmClient, createKey.ID, length, checksumType, bs, cs, uint32(bpc), sbfs, sbms)
    if err != nil {
        return err
    }
    var n int64
    if n, err = io.Copy(keyOutputStream, source); err != nil && err != io.EOF {
        return err
    }
    if err = keyOutputStream.Close(); err != nil {
        return err
    }
    if n != length {
        return fmt.Errorf("write data error: key data size %d but write len %d", length, n)
    } else {
        return nil
    }
}
