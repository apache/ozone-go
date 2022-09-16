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
    "bytes"
    "fmt"
    "github.com/apache/ozone-go/api/common"
    "github.com/apache/ozone-go/api/errors"
    "github.com/apache/ozone-go/api/proto/datanode"
    "github.com/apache/ozone-go/api/proto/hdds"
    "github.com/apache/ozone-go/api/utils"
    "strings"

    "github.com/sirupsen/logrus"
)

// PipelinePortName TODO
type PipelinePortName string

const (
    // PipelinePortNameREST TODO
    PipelinePortNameREST PipelinePortName = "REST"
    // PipelinePortNameRATIS TODO
    PipelinePortNameRATIS PipelinePortName = "RATIS"
    // PipelinePortNameSTANDALONE TODO
    PipelinePortNameSTANDALONE PipelinePortName = "STANDALONE"
)

type checksumFunc func(data []byte) []byte

// ChecksumOperator TODO
type ChecksumOperator struct {
    Checksum *datanode.ChecksumData
}

func newChecksumOperator(checksum *datanode.ChecksumData) *ChecksumOperator {
    return &ChecksumOperator{Checksum: checksum}
}

// NewChecksumOperatorComputer TODO
func NewChecksumOperatorComputer(checksumType *datanode.ChecksumType, bytesPerChecksum uint32) *ChecksumOperator {
    return newChecksumOperator(&datanode.ChecksumData{
        Type:             checksumType,
        BytesPerChecksum: utils.PointUint32(bytesPerChecksum),
    })
}

// NewChecksumOperatorVerifier TODO
func NewChecksumOperatorVerifier(checksum *datanode.ChecksumData) *ChecksumOperator {
    return newChecksumOperator(checksum)
}

// ComputeChecksum TODO
func (operator *ChecksumOperator) ComputeChecksum(data []byte) error {
    var cf checksumFunc
    switch *operator.Checksum.Type {
    case datanode.ChecksumType_NONE:
        operator.Checksum.Checksums = make([][]byte, 0)
        return nil
    case datanode.ChecksumType_CRC32:
        cf = common.CRC32
    case datanode.ChecksumType_CRC32C:
        cf = common.CRC32c
    case datanode.ChecksumType_SHA256:
        cf = common.SHA256
    case datanode.ChecksumType_MD5:
        cf = common.MD5
    default:
        return fmt.Errorf("unknown checksum type: %s", operator.Checksum.Type.String())
    }

    l := uint32(len(data))
    bytesPerChecksum := *operator.Checksum.BytesPerChecksum
    numChecksums := (l + bytesPerChecksum - 1) / bytesPerChecksum

    checksums := make([][]byte, numChecksums)
    for i := uint32(0); i < numChecksums; i++ {
        start := i * bytesPerChecksum
        end := start + bytesPerChecksum
        if end > l {
            end = l
        }
        csm := cf(data[start:end])
        if csm == nil {
            return errors.ChecksumWriteErr
        }
        for i, j := 0, len(csm)-1; i < j; i, j = i+1, j-1 {
            csm[i], csm[j] = csm[j], csm[i]
        }
        checksums[i] = csm[4:]
    }
    operator.Checksum.Checksums = checksums
    return nil
}

// VerifyChecksum TODO
func (operator *ChecksumOperator) VerifyChecksum(data []byte, off uint64) error {
    // Range: [startIndex, endIndex)
    startIndex := off / uint64(*operator.Checksum.BytesPerChecksum)
    endIndex := startIndex + uint64(len(data))/uint64(*operator.Checksum.BytesPerChecksum)
    if uint64(len(data))%uint64(*operator.Checksum.BytesPerChecksum) > 0 {
        // but it works. :-)
        endIndex += 1
    }

    if endIndex > uint64(len(operator.Checksum.Checksums)) {
        return fmt.Errorf(
            "data to verify checksums from %v, len: %v is out off index of chunk checksums(%v), bytes per Checksum: %v",
            off, len(data), len(operator.Checksum.Checksums), *operator.Checksum.BytesPerChecksum)
    }

    var cf checksumFunc
    switch *operator.Checksum.Type {
    case datanode.ChecksumType_NONE:
        return nil
    case datanode.ChecksumType_CRC32:
        cf = common.CRC32
    case datanode.ChecksumType_CRC32C:
        cf = common.CRC32c
    case datanode.ChecksumType_SHA256:
        cf = common.SHA256
    case datanode.ChecksumType_MD5:
        cf = common.MD5
    default:
        return fmt.Errorf("unknown Checksum type: " + operator.Checksum.Type.String())
    }

    for checkIndex, byteStart := startIndex, uint32(0); checkIndex < endIndex && byteStart < uint32(len(data)); {
        byteEnd := byteStart + *operator.Checksum.BytesPerChecksum
        if byteEnd > uint32(len(data)) {
            byteEnd = uint32(len(data))
        }
        csm := cf(data[byteStart:byteEnd])
        for i, j := 0, len(csm)-1; i < j; i, j = i+1, j-1 {
            csm[i], csm[j] = csm[j], csm[i]
        }
        checksumBytes := csm[4:]
        logrus.Debug("verify check sum ", checkIndex)
        if bytes.Compare(checksumBytes, operator.Checksum.Checksums[checkIndex]) != 0 {
            return &errors.ChecksumMismatchError{
                ChecksumType: operator.Checksum.Type.String(),
                From:         checksumBytes,
                CompareTo:    operator.Checksum.Checksums[checkIndex],
            }
        }
        byteStart = byteEnd
        checkIndex++
    }
    return nil
}

// DatanodeDetailsOperator TODO
type DatanodeDetailsOperator struct {
    DatanodeDetails *hdds.DatanodeDetailsProto
}

// NewDatanodeDetailsOperator TODO
func NewDatanodeDetailsOperator(dn *hdds.DatanodeDetailsProto) *DatanodeDetailsOperator {
    return &DatanodeDetailsOperator{DatanodeDetails: dn}
}

// ID TODO
func (dn *DatanodeDetailsOperator) ID() string {
    return dn.DatanodeDetails.GetUuid()
}

// Host TODO
func (dn *DatanodeDetailsOperator) Host() string {
    return dn.DatanodeDetails.GetHostName()
}

// IP TODO
func (dn *DatanodeDetailsOperator) IP() string {
    return dn.DatanodeDetails.GetIpAddress()
}

// Port TODO
func (dn *DatanodeDetailsOperator) Port(t PipelinePortName) uint32 {
    for _, port := range dn.DatanodeDetails.Ports {
        if port.GetName() == string(t) {
            return port.GetValue()
        }
    }
    return 0
}

// ExcludeListOperator TODO
type ExcludeListOperator struct {
    datanodes    []string
    containerIds []int64
    pipelineIds  []string
}

// NewExcludeListOperator TODO
func NewExcludeListOperator() *ExcludeListOperator {
    return &ExcludeListOperator{
        datanodes:    []string{},
        containerIds: []int64{},
        pipelineIds:  []string{},
    }
}

// AddDataNode TODO
func (operator *ExcludeListOperator) AddDataNode(datanode string) {
    operator.datanodes = append(operator.datanodes, datanode)
}

// AddDataNodes TODO
func (operator *ExcludeListOperator) AddDataNodes(datanode []string) {
    operator.datanodes = append(operator.datanodes, datanode...)
}

// AddContainerId TODO
func (operator *ExcludeListOperator) AddContainerId(containerId int64) {
    operator.containerIds = append(operator.containerIds, containerId)
}

// AddPipelineId TODO
func (operator *ExcludeListOperator) AddPipelineId(pipelineId string) {
    operator.pipelineIds = append(operator.pipelineIds, pipelineId)
}

// ToExcludeList TODO
func (operator *ExcludeListOperator) ToExcludeList() *hdds.ExcludeListProto {
    pipelineIds := make([]*hdds.PipelineID, 0)
    for _, id := range operator.pipelineIds {
        pipelineIds = append(pipelineIds, &hdds.PipelineID{
            Id: utils.PointString(id),
        })
    }

    return &hdds.ExcludeListProto{
        Datanodes:    operator.datanodes,
        ContainerIds: operator.containerIds,
        PipelineIds:  pipelineIds,
    }
}

// PipelineOperator TODO
type PipelineOperator struct {
    CurrentIndex int
    Pipeline     *hdds.Pipeline
}

// NewPipelineOperator TODO
func NewPipelineOperator(pipeline *hdds.Pipeline) (*PipelineOperator, error) {
    operator := &PipelineOperator{Pipeline: pipeline, CurrentIndex: 0}

    memberLen := len(operator.Pipeline.GetMembers())
    // memberOrderLen := len(operator.Pipeline.MemberOrders)

    switch {
    case memberLen == 0:
        return nil, fmt.Errorf("members of Operator is empty. Operator: %v", pipeline.String())
        // case memberOrderLen == 0:
        //	return nil, fmt.Errorf("memberOrders of Operator is empty. Operator: %v", pipeline.String())
        // case memberOrderLen != memberLen:
        //	return nil, fmt.Errorf("members chunkSize: %v not match memberOrders chunkSize: %v, Operator: %v",
        //		memberLen, memberOrderLen, pipeline.String())
    }
    return operator, nil
}

// GetPipeline TODO
func (operator *PipelineOperator) GetPipeline() *hdds.Pipeline {
    return operator.Pipeline
}

// GetId TODO
func (operator *PipelineOperator) GetId() string {
    return operator.Pipeline.GetId().GetId()
}

// GetNodeAddr TODO
func (operator *PipelineOperator) GetNodeAddr(portName PipelinePortName, firstNode bool) string {
    var dn *hdds.DatanodeDetailsProto
    if firstNode {
        dn = operator.GetFirstNode()
    } else {
        dn = operator.GetClosestNode()
    }

    for _, port := range dn.Ports {
        if port.GetName() == string(portName) {
            return fmt.Sprintf("%v:%v", dn.GetIpAddress(), port.GetValue())
        }
    }
    return ""
}

// SetCurrentToNext TODO
func (operator *PipelineOperator) SetCurrentToNext() {
    if operator.CurrentIndex+1 >= len(operator.GetPipeline().GetMembers()) {
        operator.CurrentIndex = 0
    } else {
        operator.CurrentIndex += 1
    }
}

// GetCurrentNode TODO
func (operator *PipelineOperator) GetCurrentNode() *hdds.DatanodeDetailsProto {
    if operator.CurrentIndex >= len(operator.GetPipeline().GetMembers()) {
        operator.CurrentIndex = 0
    }
    return operator.Pipeline.GetMembers()[operator.CurrentIndex]
}

// GetFirstNode TODO
func (operator *PipelineOperator) GetFirstNode() *hdds.DatanodeDetailsProto {
    return operator.Pipeline.GetMembers()[0]
}

// GetClosestNode TODO
func (operator *PipelineOperator) GetClosestNode() *hdds.DatanodeDetailsProto {
    return operator.Pipeline.GetMembers()[operator.Pipeline.GetMemberOrders()[0]]
}

// GetLeaderNode TODO
func (operator *PipelineOperator) GetLeaderNode() *hdds.DatanodeDetailsProto {
    leaderId := operator.Pipeline.GetLeaderID()
    for idx, member := range operator.Pipeline.Members {
        if member.GetUuid() == leaderId {
            operator.CurrentIndex = idx
            return member
        }
    }
    return nil
}

// GetLeaderAddress TODO
func (operator *PipelineOperator) GetLeaderAddress() string {
    leaderId := operator.Pipeline.GetLeaderID()
    for idx, member := range operator.Pipeline.Members {
        if member.GetUuid() == leaderId {
            operator.CurrentIndex = idx
            return member.GetIpAddress()
        }
    }
    return "0.0.0.0"
}

func (operator *PipelineOperator) getRemotePort(mode string) uint32 {
    for _, mem := range operator.GetPipeline().GetMembers() {
        for _, port := range mem.GetPorts() {
            if strings.ToLower(port.GetName()) == strings.ToLower(mode) {
				return port.GetValue()
			}
		}
	}
	return 0
}
