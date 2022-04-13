// Package io TODO
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
package io

import "github.com/apache/ozone-go/api/proto/hdds"

// AddDNToExcludeList ExcludeList 排除一些datannode container pipeline
func (k *KeyOutputStream) AddDNToExcludeList(dn string) {
    exist := false
    for _, dataNode := range k.excludeList.GetDatanodes() {
        if dataNode == dn {
            exist = true
            break
        }
    }
    if !exist {
        k.excludeList.Datanodes = append(k.excludeList.GetDatanodes(), dn)
    }
}

// AddPipeLineDNToExcludeList ExcludeList 排除一些datannode container pipeline
func (k *KeyOutputStream) AddPipeLineDNToExcludeList(pipeline *hdds.Pipeline) {
    for _, mem := range pipeline.GetMembers() {
        exist := false
        for _, dataNode := range k.excludeList.GetDatanodes() {
            if dataNode == mem.GetIpAddress() {
                exist = true
            }
        }
        if !exist {
            k.excludeList.Datanodes = append(k.excludeList.Datanodes, mem.GetUuid())
        }
    }
    k.AddPipelineToExcludeList(pipeline.GetId())
}

// AddContainerIdToExcludeList TODO
func (k *KeyOutputStream) AddContainerIdToExcludeList(id int64) {
    exist := false
    for _, cid := range k.excludeList.GetContainerIds() {
        if cid == id {
            exist = true
            break
        }
    }
    if !exist {
        k.excludeList.ContainerIds = append(k.excludeList.GetContainerIds(), id)
    }
}

// AddPipelineToExcludeList TODO
func (k *KeyOutputStream) AddPipelineToExcludeList(id *hdds.PipelineID) {
    exist := false
    for _, pid := range k.excludeList.GetPipelineIds() {
        if pid.GetUuid128().String() == id.GetUuid128().String() {
            exist = true
            break
        }
    }
    if !exist {
        k.excludeList.PipelineIds = append(k.excludeList.GetPipelineIds(), id)
    }

}

// GetExcludeList TODO
func (k *KeyOutputStream) GetExcludeList() *hdds.ExcludeListProto {
    return k.excludeList
}
