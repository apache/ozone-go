/**
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package hdds

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PipelineState int32

const (
	PipelineState_PIPELINE_ALLOCATED PipelineState = 1
	PipelineState_PIPELINE_OPEN      PipelineState = 2
	PipelineState_PIPELINE_DORMANT   PipelineState = 3
	PipelineState_PIPELINE_CLOSED    PipelineState = 4
)

// Enum value maps for PipelineState.
var (
	PipelineState_name = map[int32]string{
		1: "PIPELINE_ALLOCATED",
		2: "PIPELINE_OPEN",
		3: "PIPELINE_DORMANT",
		4: "PIPELINE_CLOSED",
	}
	PipelineState_value = map[string]int32{
		"PIPELINE_ALLOCATED": 1,
		"PIPELINE_OPEN":      2,
		"PIPELINE_DORMANT":   3,
		"PIPELINE_CLOSED":    4,
	}
)

func (x PipelineState) Enum() *PipelineState {
	p := new(PipelineState)
	*p = x
	return p
}

func (x PipelineState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PipelineState) Descriptor() protoreflect.EnumDescriptor {
	return file_hdds_proto_enumTypes[0].Descriptor()
}

func (PipelineState) Type() protoreflect.EnumType {
	return &file_hdds_proto_enumTypes[0]
}

func (x PipelineState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *PipelineState) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = PipelineState(num)
	return nil
}

// Deprecated: Use PipelineState.Descriptor instead.
func (PipelineState) EnumDescriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{0}
}

//*
// Type of the node.
type NodeType int32

const (
	NodeType_OM       NodeType = 1 // Ozone Manager
	NodeType_SCM      NodeType = 2 // Storage Container Manager
	NodeType_DATANODE NodeType = 3 // DataNode
	NodeType_RECON    NodeType = 4 // Recon
)

// Enum value maps for NodeType.
var (
	NodeType_name = map[int32]string{
		1: "OM",
		2: "SCM",
		3: "DATANODE",
		4: "RECON",
	}
	NodeType_value = map[string]int32{
		"OM":       1,
		"SCM":      2,
		"DATANODE": 3,
		"RECON":    4,
	}
)

func (x NodeType) Enum() *NodeType {
	p := new(NodeType)
	*p = x
	return p
}

func (x NodeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeType) Descriptor() protoreflect.EnumDescriptor {
	return file_hdds_proto_enumTypes[1].Descriptor()
}

func (NodeType) Type() protoreflect.EnumType {
	return &file_hdds_proto_enumTypes[1]
}

func (x NodeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *NodeType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = NodeType(num)
	return nil
}

// Deprecated: Use NodeType.Descriptor instead.
func (NodeType) EnumDescriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{1}
}

//*
// Enum that represents the Node State. This is used in calls to getNodeList
// and getNodeCount.
type NodeState int32

const (
	NodeState_HEALTHY         NodeState = 1
	NodeState_STALE           NodeState = 2
	NodeState_DEAD            NodeState = 3
	NodeState_DECOMMISSIONING NodeState = 4
	NodeState_DECOMMISSIONED  NodeState = 5
)

// Enum value maps for NodeState.
var (
	NodeState_name = map[int32]string{
		1: "HEALTHY",
		2: "STALE",
		3: "DEAD",
		4: "DECOMMISSIONING",
		5: "DECOMMISSIONED",
	}
	NodeState_value = map[string]int32{
		"HEALTHY":         1,
		"STALE":           2,
		"DEAD":            3,
		"DECOMMISSIONING": 4,
		"DECOMMISSIONED":  5,
	}
)

func (x NodeState) Enum() *NodeState {
	p := new(NodeState)
	*p = x
	return p
}

func (x NodeState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeState) Descriptor() protoreflect.EnumDescriptor {
	return file_hdds_proto_enumTypes[2].Descriptor()
}

func (NodeState) Type() protoreflect.EnumType {
	return &file_hdds_proto_enumTypes[2]
}

func (x NodeState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *NodeState) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = NodeState(num)
	return nil
}

// Deprecated: Use NodeState.Descriptor instead.
func (NodeState) EnumDescriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{2}
}

type QueryScope int32

const (
	QueryScope_CLUSTER QueryScope = 1
	QueryScope_POOL    QueryScope = 2
)

// Enum value maps for QueryScope.
var (
	QueryScope_name = map[int32]string{
		1: "CLUSTER",
		2: "POOL",
	}
	QueryScope_value = map[string]int32{
		"CLUSTER": 1,
		"POOL":    2,
	}
)

func (x QueryScope) Enum() *QueryScope {
	p := new(QueryScope)
	*p = x
	return p
}

func (x QueryScope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QueryScope) Descriptor() protoreflect.EnumDescriptor {
	return file_hdds_proto_enumTypes[3].Descriptor()
}

func (QueryScope) Type() protoreflect.EnumType {
	return &file_hdds_proto_enumTypes[3]
}

func (x QueryScope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *QueryScope) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = QueryScope(num)
	return nil
}

// Deprecated: Use QueryScope.Descriptor instead.
func (QueryScope) EnumDescriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{3}
}

type LifeCycleState int32

const (
	LifeCycleState_OPEN         LifeCycleState = 1
	LifeCycleState_CLOSING      LifeCycleState = 2
	LifeCycleState_QUASI_CLOSED LifeCycleState = 3
	LifeCycleState_CLOSED       LifeCycleState = 4
	LifeCycleState_DELETING     LifeCycleState = 5
	LifeCycleState_DELETED      LifeCycleState = 6 // object is deleted.
)

// Enum value maps for LifeCycleState.
var (
	LifeCycleState_name = map[int32]string{
		1: "OPEN",
		2: "CLOSING",
		3: "QUASI_CLOSED",
		4: "CLOSED",
		5: "DELETING",
		6: "DELETED",
	}
	LifeCycleState_value = map[string]int32{
		"OPEN":         1,
		"CLOSING":      2,
		"QUASI_CLOSED": 3,
		"CLOSED":       4,
		"DELETING":     5,
		"DELETED":      6,
	}
)

func (x LifeCycleState) Enum() *LifeCycleState {
	p := new(LifeCycleState)
	*p = x
	return p
}

func (x LifeCycleState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LifeCycleState) Descriptor() protoreflect.EnumDescriptor {
	return file_hdds_proto_enumTypes[4].Descriptor()
}

func (LifeCycleState) Type() protoreflect.EnumType {
	return &file_hdds_proto_enumTypes[4]
}

func (x LifeCycleState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *LifeCycleState) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = LifeCycleState(num)
	return nil
}

// Deprecated: Use LifeCycleState.Descriptor instead.
func (LifeCycleState) EnumDescriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{4}
}

type LifeCycleEvent int32

const (
	LifeCycleEvent_FINALIZE    LifeCycleEvent = 1
	LifeCycleEvent_QUASI_CLOSE LifeCycleEvent = 2
	LifeCycleEvent_CLOSE       LifeCycleEvent = 3 // !!Event after this has not been used yet.
	LifeCycleEvent_FORCE_CLOSE LifeCycleEvent = 4
	LifeCycleEvent_DELETE      LifeCycleEvent = 5
	LifeCycleEvent_CLEANUP     LifeCycleEvent = 6
)

// Enum value maps for LifeCycleEvent.
var (
	LifeCycleEvent_name = map[int32]string{
		1: "FINALIZE",
		2: "QUASI_CLOSE",
		3: "CLOSE",
		4: "FORCE_CLOSE",
		5: "DELETE",
		6: "CLEANUP",
	}
	LifeCycleEvent_value = map[string]int32{
		"FINALIZE":    1,
		"QUASI_CLOSE": 2,
		"CLOSE":       3,
		"FORCE_CLOSE": 4,
		"DELETE":      5,
		"CLEANUP":     6,
	}
)

func (x LifeCycleEvent) Enum() *LifeCycleEvent {
	p := new(LifeCycleEvent)
	*p = x
	return p
}

func (x LifeCycleEvent) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LifeCycleEvent) Descriptor() protoreflect.EnumDescriptor {
	return file_hdds_proto_enumTypes[5].Descriptor()
}

func (LifeCycleEvent) Type() protoreflect.EnumType {
	return &file_hdds_proto_enumTypes[5]
}

func (x LifeCycleEvent) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *LifeCycleEvent) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = LifeCycleEvent(num)
	return nil
}

// Deprecated: Use LifeCycleEvent.Descriptor instead.
func (LifeCycleEvent) EnumDescriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{5}
}

type ReplicationType int32

const (
	ReplicationType_RATIS       ReplicationType = 1
	ReplicationType_STAND_ALONE ReplicationType = 2
	ReplicationType_CHAINED     ReplicationType = 3
)

// Enum value maps for ReplicationType.
var (
	ReplicationType_name = map[int32]string{
		1: "RATIS",
		2: "STAND_ALONE",
		3: "CHAINED",
	}
	ReplicationType_value = map[string]int32{
		"RATIS":       1,
		"STAND_ALONE": 2,
		"CHAINED":     3,
	}
)

func (x ReplicationType) Enum() *ReplicationType {
	p := new(ReplicationType)
	*p = x
	return p
}

func (x ReplicationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReplicationType) Descriptor() protoreflect.EnumDescriptor {
	return file_hdds_proto_enumTypes[6].Descriptor()
}

func (ReplicationType) Type() protoreflect.EnumType {
	return &file_hdds_proto_enumTypes[6]
}

func (x ReplicationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ReplicationType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ReplicationType(num)
	return nil
}

// Deprecated: Use ReplicationType.Descriptor instead.
func (ReplicationType) EnumDescriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{6}
}

type ReplicationFactor int32

const (
	ReplicationFactor_ONE   ReplicationFactor = 1
	ReplicationFactor_THREE ReplicationFactor = 3
)

// Enum value maps for ReplicationFactor.
var (
	ReplicationFactor_name = map[int32]string{
		1: "ONE",
		3: "THREE",
	}
	ReplicationFactor_value = map[string]int32{
		"ONE":   1,
		"THREE": 3,
	}
)

func (x ReplicationFactor) Enum() *ReplicationFactor {
	p := new(ReplicationFactor)
	*p = x
	return p
}

func (x ReplicationFactor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReplicationFactor) Descriptor() protoreflect.EnumDescriptor {
	return file_hdds_proto_enumTypes[7].Descriptor()
}

func (ReplicationFactor) Type() protoreflect.EnumType {
	return &file_hdds_proto_enumTypes[7]
}

func (x ReplicationFactor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ReplicationFactor) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ReplicationFactor(num)
	return nil
}

// Deprecated: Use ReplicationFactor.Descriptor instead.
func (ReplicationFactor) EnumDescriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{7}
}

type ScmOps int32

const (
	ScmOps_allocateBlock             ScmOps = 1
	ScmOps_keyBlocksInfoList         ScmOps = 2
	ScmOps_getScmInfo                ScmOps = 3
	ScmOps_deleteBlock               ScmOps = 4
	ScmOps_createReplicationPipeline ScmOps = 5
	ScmOps_allocateContainer         ScmOps = 6
	ScmOps_getContainer              ScmOps = 7
	ScmOps_getContainerWithPipeline  ScmOps = 8
	ScmOps_listContainer             ScmOps = 9
	ScmOps_deleteContainer           ScmOps = 10
	ScmOps_queryNode                 ScmOps = 11
)

// Enum value maps for ScmOps.
var (
	ScmOps_name = map[int32]string{
		1:  "allocateBlock",
		2:  "keyBlocksInfoList",
		3:  "getScmInfo",
		4:  "deleteBlock",
		5:  "createReplicationPipeline",
		6:  "allocateContainer",
		7:  "getContainer",
		8:  "getContainerWithPipeline",
		9:  "listContainer",
		10: "deleteContainer",
		11: "queryNode",
	}
	ScmOps_value = map[string]int32{
		"allocateBlock":             1,
		"keyBlocksInfoList":         2,
		"getScmInfo":                3,
		"deleteBlock":               4,
		"createReplicationPipeline": 5,
		"allocateContainer":         6,
		"getContainer":              7,
		"getContainerWithPipeline":  8,
		"listContainer":             9,
		"deleteContainer":           10,
		"queryNode":                 11,
	}
)

func (x ScmOps) Enum() *ScmOps {
	p := new(ScmOps)
	*p = x
	return p
}

func (x ScmOps) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScmOps) Descriptor() protoreflect.EnumDescriptor {
	return file_hdds_proto_enumTypes[8].Descriptor()
}

func (ScmOps) Type() protoreflect.EnumType {
	return &file_hdds_proto_enumTypes[8]
}

func (x ScmOps) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ScmOps) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ScmOps(num)
	return nil
}

// Deprecated: Use ScmOps.Descriptor instead.
func (ScmOps) EnumDescriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{8}
}

//*
// File access permissions mode.
type BlockTokenSecretProto_AccessModeProto int32

const (
	BlockTokenSecretProto_READ   BlockTokenSecretProto_AccessModeProto = 1
	BlockTokenSecretProto_WRITE  BlockTokenSecretProto_AccessModeProto = 2
	BlockTokenSecretProto_COPY   BlockTokenSecretProto_AccessModeProto = 3
	BlockTokenSecretProto_DELETE BlockTokenSecretProto_AccessModeProto = 4
)

// Enum value maps for BlockTokenSecretProto_AccessModeProto.
var (
	BlockTokenSecretProto_AccessModeProto_name = map[int32]string{
		1: "READ",
		2: "WRITE",
		3: "COPY",
		4: "DELETE",
	}
	BlockTokenSecretProto_AccessModeProto_value = map[string]int32{
		"READ":   1,
		"WRITE":  2,
		"COPY":   3,
		"DELETE": 4,
	}
)

func (x BlockTokenSecretProto_AccessModeProto) Enum() *BlockTokenSecretProto_AccessModeProto {
	p := new(BlockTokenSecretProto_AccessModeProto)
	*p = x
	return p
}

func (x BlockTokenSecretProto_AccessModeProto) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlockTokenSecretProto_AccessModeProto) Descriptor() protoreflect.EnumDescriptor {
	return file_hdds_proto_enumTypes[9].Descriptor()
}

func (BlockTokenSecretProto_AccessModeProto) Type() protoreflect.EnumType {
	return &file_hdds_proto_enumTypes[9]
}

func (x BlockTokenSecretProto_AccessModeProto) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *BlockTokenSecretProto_AccessModeProto) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = BlockTokenSecretProto_AccessModeProto(num)
	return nil
}

// Deprecated: Use BlockTokenSecretProto_AccessModeProto.Descriptor instead.
func (BlockTokenSecretProto_AccessModeProto) EnumDescriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{15, 0}
}

type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MostSigBits  *int64 `protobuf:"varint,1,req,name=mostSigBits" json:"mostSigBits,omitempty"`
	LeastSigBits *int64 `protobuf:"varint,2,req,name=leastSigBits" json:"leastSigBits,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hdds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_hdds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{0}
}

func (x *UUID) GetMostSigBits() int64 {
	if x != nil && x.MostSigBits != nil {
		return *x.MostSigBits
	}
	return 0
}

func (x *UUID) GetLeastSigBits() int64 {
	if x != nil && x.LeastSigBits != nil {
		return *x.LeastSigBits
	}
	return 0
}

type DatanodeDetailsProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// deprecated, please use uuid128 instead
	Uuid         *string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`           // UUID assigned to the Datanode.
	IpAddress    *string `protobuf:"bytes,2,req,name=ipAddress" json:"ipAddress,omitempty"` // IP address
	HostName     *string `protobuf:"bytes,3,req,name=hostName" json:"hostName,omitempty"`   // hostname
	Ports        []*Port `protobuf:"bytes,4,rep,name=ports" json:"ports,omitempty"`
	CertSerialId *string `protobuf:"bytes,5,opt,name=certSerialId" json:"certSerialId,omitempty"` // Certificate serial id.
	// network name, can be Ip address or host name, depends
	NetworkName     *string `protobuf:"bytes,6,opt,name=networkName" json:"networkName,omitempty"`
	NetworkLocation *string `protobuf:"bytes,7,opt,name=networkLocation" json:"networkLocation,omitempty"` // Network topology location
	Version         *string `protobuf:"bytes,8,opt,name=version" json:"version,omitempty"`                 // Datanode version
	SetupTime       *int64  `protobuf:"varint,9,opt,name=setupTime" json:"setupTime,omitempty"`
	Revision        *string `protobuf:"bytes,10,opt,name=revision" json:"revision,omitempty"`
	BuildDate       *string `protobuf:"bytes,11,opt,name=buildDate" json:"buildDate,omitempty"`
	// TODO(runzhiwang): when uuid is gone, specify 1 as the index of uuid128 and mark as required
	Uuid128 *UUID `protobuf:"bytes,100,opt,name=uuid128" json:"uuid128,omitempty"` // UUID with 128 bits assigned to the Datanode.
}

func (x *DatanodeDetailsProto) Reset() {
	*x = DatanodeDetailsProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hdds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatanodeDetailsProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatanodeDetailsProto) ProtoMessage() {}

func (x *DatanodeDetailsProto) ProtoReflect() protoreflect.Message {
	mi := &file_hdds_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatanodeDetailsProto.ProtoReflect.Descriptor instead.
func (*DatanodeDetailsProto) Descriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{1}
}

func (x *DatanodeDetailsProto) GetUuid() string {
	if x != nil && x.Uuid != nil {
		return *x.Uuid
	}
	return ""
}

func (x *DatanodeDetailsProto) GetIpAddress() string {
	if x != nil && x.IpAddress != nil {
		return *x.IpAddress
	}
	return ""
}

func (x *DatanodeDetailsProto) GetHostName() string {
	if x != nil && x.HostName != nil {
		return *x.HostName
	}
	return ""
}

func (x *DatanodeDetailsProto) GetPorts() []*Port {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *DatanodeDetailsProto) GetCertSerialId() string {
	if x != nil && x.CertSerialId != nil {
		return *x.CertSerialId
	}
	return ""
}

func (x *DatanodeDetailsProto) GetNetworkName() string {
	if x != nil && x.NetworkName != nil {
		return *x.NetworkName
	}
	return ""
}

func (x *DatanodeDetailsProto) GetNetworkLocation() string {
	if x != nil && x.NetworkLocation != nil {
		return *x.NetworkLocation
	}
	return ""
}

func (x *DatanodeDetailsProto) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

func (x *DatanodeDetailsProto) GetSetupTime() int64 {
	if x != nil && x.SetupTime != nil {
		return *x.SetupTime
	}
	return 0
}

func (x *DatanodeDetailsProto) GetRevision() string {
	if x != nil && x.Revision != nil {
		return *x.Revision
	}
	return ""
}

func (x *DatanodeDetailsProto) GetBuildDate() string {
	if x != nil && x.BuildDate != nil {
		return *x.BuildDate
	}
	return ""
}

func (x *DatanodeDetailsProto) GetUuid128() *UUID {
	if x != nil {
		return x.Uuid128
	}
	return nil
}

//*
//Proto message encapsulating information required to uniquely identify a
//OzoneManager.
type OzoneManagerDetailsProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      *string `protobuf:"bytes,1,req,name=uuid" json:"uuid,omitempty"`           // UUID assigned to the OzoneManager.
	IpAddress *string `protobuf:"bytes,2,req,name=ipAddress" json:"ipAddress,omitempty"` // IP address of OM.
	HostName  *string `protobuf:"bytes,3,req,name=hostName" json:"hostName,omitempty"`   // Hostname of OM.
	Ports     []*Port `protobuf:"bytes,4,rep,name=ports" json:"ports,omitempty"`
}

func (x *OzoneManagerDetailsProto) Reset() {
	*x = OzoneManagerDetailsProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hdds_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OzoneManagerDetailsProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OzoneManagerDetailsProto) ProtoMessage() {}

func (x *OzoneManagerDetailsProto) ProtoReflect() protoreflect.Message {
	mi := &file_hdds_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OzoneManagerDetailsProto.ProtoReflect.Descriptor instead.
func (*OzoneManagerDetailsProto) Descriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{2}
}

func (x *OzoneManagerDetailsProto) GetUuid() string {
	if x != nil && x.Uuid != nil {
		return *x.Uuid
	}
	return ""
}

func (x *OzoneManagerDetailsProto) GetIpAddress() string {
	if x != nil && x.IpAddress != nil {
		return *x.IpAddress
	}
	return ""
}

func (x *OzoneManagerDetailsProto) GetHostName() string {
	if x != nil && x.HostName != nil {
		return *x.HostName
	}
	return ""
}

func (x *OzoneManagerDetailsProto) GetPorts() []*Port {
	if x != nil {
		return x.Ports
	}
	return nil
}

type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value *uint32 `protobuf:"varint,2,req,name=value" json:"value,omitempty"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hdds_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_hdds_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{3}
}

func (x *Port) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Port) GetValue() uint32 {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return 0
}

type PipelineID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// deprecated, please use uuid128 instead
	Id *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// TODO(runzhiwang): when id is gone, specify 1 as the index of uuid128 and mark as required
	Uuid128 *UUID `protobuf:"bytes,100,opt,name=uuid128" json:"uuid128,omitempty"`
}

func (x *PipelineID) Reset() {
	*x = PipelineID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hdds_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineID) ProtoMessage() {}

func (x *PipelineID) ProtoReflect() protoreflect.Message {
	mi := &file_hdds_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineID.ProtoReflect.Descriptor instead.
func (*PipelineID) Descriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{4}
}

func (x *PipelineID) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *PipelineID) GetUuid128() *UUID {
	if x != nil {
		return x.Uuid128
	}
	return nil
}

type Pipeline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Members []*DatanodeDetailsProto `protobuf:"bytes,1,rep,name=members" json:"members,omitempty"`
	// TODO: remove the state and leaderID from this class
	State             *PipelineState     `protobuf:"varint,2,opt,name=state,enum=hadoop.hdds.PipelineState,def=1" json:"state,omitempty"`
	Type              *ReplicationType   `protobuf:"varint,3,opt,name=type,enum=hadoop.hdds.ReplicationType,def=2" json:"type,omitempty"`
	Factor            *ReplicationFactor `protobuf:"varint,4,opt,name=factor,enum=hadoop.hdds.ReplicationFactor,def=1" json:"factor,omitempty"`
	Id                *PipelineID        `protobuf:"bytes,5,req,name=id" json:"id,omitempty"`
	LeaderID          *string            `protobuf:"bytes,6,opt,name=leaderID" json:"leaderID,omitempty"`
	MemberOrders      []uint32           `protobuf:"varint,7,rep,name=memberOrders" json:"memberOrders,omitempty"`
	CreationTimeStamp *uint64            `protobuf:"varint,8,opt,name=creationTimeStamp" json:"creationTimeStamp,omitempty"`
	// TODO(runzhiwang): when leaderID is gone, specify 6 as the index of leaderID128
	LeaderID128 *UUID `protobuf:"bytes,100,opt,name=leaderID128" json:"leaderID128,omitempty"`
}

// Default values for Pipeline fields.
const (
	Default_Pipeline_State  = PipelineState_PIPELINE_ALLOCATED
	Default_Pipeline_Type   = ReplicationType_STAND_ALONE
	Default_Pipeline_Factor = ReplicationFactor_ONE
)

func (x *Pipeline) Reset() {
	*x = Pipeline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hdds_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pipeline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pipeline) ProtoMessage() {}

func (x *Pipeline) ProtoReflect() protoreflect.Message {
	mi := &file_hdds_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pipeline.ProtoReflect.Descriptor instead.
func (*Pipeline) Descriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{5}
}

func (x *Pipeline) GetMembers() []*DatanodeDetailsProto {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *Pipeline) GetState() PipelineState {
	if x != nil && x.State != nil {
		return *x.State
	}
	return Default_Pipeline_State
}

func (x *Pipeline) GetType() ReplicationType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return Default_Pipeline_Type
}

func (x *Pipeline) GetFactor() ReplicationFactor {
	if x != nil && x.Factor != nil {
		return *x.Factor
	}
	return Default_Pipeline_Factor
}

func (x *Pipeline) GetId() *PipelineID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Pipeline) GetLeaderID() string {
	if x != nil && x.LeaderID != nil {
		return *x.LeaderID
	}
	return ""
}

func (x *Pipeline) GetMemberOrders() []uint32 {
	if x != nil {
		return x.MemberOrders
	}
	return nil
}

func (x *Pipeline) GetCreationTimeStamp() uint64 {
	if x != nil && x.CreationTimeStamp != nil {
		return *x.CreationTimeStamp
	}
	return 0
}

func (x *Pipeline) GetLeaderID128() *UUID {
	if x != nil {
		return x.LeaderID128
	}
	return nil
}

type KeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (x *KeyValue) Reset() {
	*x = KeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hdds_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_hdds_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValue.ProtoReflect.Descriptor instead.
func (*KeyValue) Descriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{6}
}

func (x *KeyValue) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *KeyValue) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeID     *DatanodeDetailsProto `protobuf:"bytes,1,req,name=nodeID" json:"nodeID,omitempty"`
	NodeStates []NodeState           `protobuf:"varint,2,rep,name=nodeStates,enum=hadoop.hdds.NodeState" json:"nodeStates,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hdds_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_hdds_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{7}
}

func (x *Node) GetNodeID() *DatanodeDetailsProto {
	if x != nil {
		return x.NodeID
	}
	return nil
}

func (x *Node) GetNodeStates() []NodeState {
	if x != nil {
		return x.NodeStates
	}
	return nil
}

type NodePool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
}

func (x *NodePool) Reset() {
	*x = NodePool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hdds_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodePool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodePool) ProtoMessage() {}

func (x *NodePool) ProtoReflect() protoreflect.Message {
	mi := &file_hdds_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodePool.ProtoReflect.Descriptor instead.
func (*NodePool) Descriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{8}
}

func (x *NodePool) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type ContainerInfoProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerID         *int64             `protobuf:"varint,1,req,name=containerID" json:"containerID,omitempty"`
	State               *LifeCycleState    `protobuf:"varint,2,req,name=state,enum=hadoop.hdds.LifeCycleState" json:"state,omitempty"`
	PipelineID          *PipelineID        `protobuf:"bytes,3,opt,name=pipelineID" json:"pipelineID,omitempty"`
	UsedBytes           *uint64            `protobuf:"varint,4,req,name=usedBytes" json:"usedBytes,omitempty"`
	NumberOfKeys        *uint64            `protobuf:"varint,5,req,name=numberOfKeys" json:"numberOfKeys,omitempty"`
	StateEnterTime      *int64             `protobuf:"varint,6,opt,name=stateEnterTime" json:"stateEnterTime,omitempty"`
	Owner               *string            `protobuf:"bytes,7,req,name=owner" json:"owner,omitempty"`
	DeleteTransactionId *int64             `protobuf:"varint,8,opt,name=deleteTransactionId" json:"deleteTransactionId,omitempty"`
	SequenceId          *int64             `protobuf:"varint,9,opt,name=sequenceId" json:"sequenceId,omitempty"`
	ReplicationFactor   *ReplicationFactor `protobuf:"varint,10,req,name=replicationFactor,enum=hadoop.hdds.ReplicationFactor" json:"replicationFactor,omitempty"`
	ReplicationType     *ReplicationType   `protobuf:"varint,11,req,name=replicationType,enum=hadoop.hdds.ReplicationType" json:"replicationType,omitempty"`
}

func (x *ContainerInfoProto) Reset() {
	*x = ContainerInfoProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hdds_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerInfoProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerInfoProto) ProtoMessage() {}

func (x *ContainerInfoProto) ProtoReflect() protoreflect.Message {
	mi := &file_hdds_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerInfoProto.ProtoReflect.Descriptor instead.
func (*ContainerInfoProto) Descriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{9}
}

func (x *ContainerInfoProto) GetContainerID() int64 {
	if x != nil && x.ContainerID != nil {
		return *x.ContainerID
	}
	return 0
}

func (x *ContainerInfoProto) GetState() LifeCycleState {
	if x != nil && x.State != nil {
		return *x.State
	}
	return LifeCycleState_OPEN
}

func (x *ContainerInfoProto) GetPipelineID() *PipelineID {
	if x != nil {
		return x.PipelineID
	}
	return nil
}

func (x *ContainerInfoProto) GetUsedBytes() uint64 {
	if x != nil && x.UsedBytes != nil {
		return *x.UsedBytes
	}
	return 0
}

func (x *ContainerInfoProto) GetNumberOfKeys() uint64 {
	if x != nil && x.NumberOfKeys != nil {
		return *x.NumberOfKeys
	}
	return 0
}

func (x *ContainerInfoProto) GetStateEnterTime() int64 {
	if x != nil && x.StateEnterTime != nil {
		return *x.StateEnterTime
	}
	return 0
}

func (x *ContainerInfoProto) GetOwner() string {
	if x != nil && x.Owner != nil {
		return *x.Owner
	}
	return ""
}

func (x *ContainerInfoProto) GetDeleteTransactionId() int64 {
	if x != nil && x.DeleteTransactionId != nil {
		return *x.DeleteTransactionId
	}
	return 0
}

func (x *ContainerInfoProto) GetSequenceId() int64 {
	if x != nil && x.SequenceId != nil {
		return *x.SequenceId
	}
	return 0
}

func (x *ContainerInfoProto) GetReplicationFactor() ReplicationFactor {
	if x != nil && x.ReplicationFactor != nil {
		return *x.ReplicationFactor
	}
	return ReplicationFactor_ONE
}

func (x *ContainerInfoProto) GetReplicationType() ReplicationType {
	if x != nil && x.ReplicationType != nil {
		return *x.ReplicationType
	}
	return ReplicationType_RATIS
}

type ContainerWithPipeline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerInfo *ContainerInfoProto `protobuf:"bytes,1,req,name=containerInfo" json:"containerInfo,omitempty"`
	Pipeline      *Pipeline           `protobuf:"bytes,2,req,name=pipeline" json:"pipeline,omitempty"`
}

func (x *ContainerWithPipeline) Reset() {
	*x = ContainerWithPipeline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hdds_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerWithPipeline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerWithPipeline) ProtoMessage() {}

func (x *ContainerWithPipeline) ProtoReflect() protoreflect.Message {
	mi := &file_hdds_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerWithPipeline.ProtoReflect.Descriptor instead.
func (*ContainerWithPipeline) Descriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{10}
}

func (x *ContainerWithPipeline) GetContainerInfo() *ContainerInfoProto {
	if x != nil {
		return x.ContainerInfo
	}
	return nil
}

func (x *ContainerWithPipeline) GetPipeline() *Pipeline {
	if x != nil {
		return x.Pipeline
	}
	return nil
}

type GetScmInfoRequestProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceID *string `protobuf:"bytes,1,opt,name=traceID" json:"traceID,omitempty"`
}

func (x *GetScmInfoRequestProto) Reset() {
	*x = GetScmInfoRequestProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hdds_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScmInfoRequestProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScmInfoRequestProto) ProtoMessage() {}

func (x *GetScmInfoRequestProto) ProtoReflect() protoreflect.Message {
	mi := &file_hdds_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScmInfoRequestProto.ProtoReflect.Descriptor instead.
func (*GetScmInfoRequestProto) Descriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{11}
}

func (x *GetScmInfoRequestProto) GetTraceID() string {
	if x != nil && x.TraceID != nil {
		return *x.TraceID
	}
	return ""
}

type GetScmInfoResponseProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId *string `protobuf:"bytes,1,req,name=clusterId" json:"clusterId,omitempty"`
	ScmId     *string `protobuf:"bytes,2,req,name=scmId" json:"scmId,omitempty"`
}

func (x *GetScmInfoResponseProto) Reset() {
	*x = GetScmInfoResponseProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hdds_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScmInfoResponseProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScmInfoResponseProto) ProtoMessage() {}

func (x *GetScmInfoResponseProto) ProtoReflect() protoreflect.Message {
	mi := &file_hdds_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScmInfoResponseProto.ProtoReflect.Descriptor instead.
func (*GetScmInfoResponseProto) Descriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{12}
}

func (x *GetScmInfoResponseProto) GetClusterId() string {
	if x != nil && x.ClusterId != nil {
		return *x.ClusterId
	}
	return ""
}

func (x *GetScmInfoResponseProto) GetScmId() string {
	if x != nil && x.ScmId != nil {
		return *x.ScmId
	}
	return ""
}

type ExcludeListProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datanodes    []string      `protobuf:"bytes,1,rep,name=datanodes" json:"datanodes,omitempty"`
	ContainerIds []int64       `protobuf:"varint,2,rep,name=containerIds" json:"containerIds,omitempty"`
	PipelineIds  []*PipelineID `protobuf:"bytes,3,rep,name=pipelineIds" json:"pipelineIds,omitempty"`
}

func (x *ExcludeListProto) Reset() {
	*x = ExcludeListProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hdds_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExcludeListProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExcludeListProto) ProtoMessage() {}

func (x *ExcludeListProto) ProtoReflect() protoreflect.Message {
	mi := &file_hdds_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExcludeListProto.ProtoReflect.Descriptor instead.
func (*ExcludeListProto) Descriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{13}
}

func (x *ExcludeListProto) GetDatanodes() []string {
	if x != nil {
		return x.Datanodes
	}
	return nil
}

func (x *ExcludeListProto) GetContainerIds() []int64 {
	if x != nil {
		return x.ContainerIds
	}
	return nil
}

func (x *ExcludeListProto) GetPipelineIds() []*PipelineID {
	if x != nil {
		return x.PipelineIds
	}
	return nil
}

//*
// Block ID that uniquely identify a block by SCM.
type ContainerBlockID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerID *int64 `protobuf:"varint,1,req,name=containerID" json:"containerID,omitempty"`
	LocalID     *int64 `protobuf:"varint,2,req,name=localID" json:"localID,omitempty"`
}

func (x *ContainerBlockID) Reset() {
	*x = ContainerBlockID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hdds_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerBlockID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerBlockID) ProtoMessage() {}

func (x *ContainerBlockID) ProtoReflect() protoreflect.Message {
	mi := &file_hdds_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerBlockID.ProtoReflect.Descriptor instead.
func (*ContainerBlockID) Descriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{14}
}

func (x *ContainerBlockID) GetContainerID() int64 {
	if x != nil && x.ContainerID != nil {
		return *x.ContainerID
	}
	return 0
}

func (x *ContainerBlockID) GetLocalID() int64 {
	if x != nil && x.LocalID != nil {
		return *x.LocalID
	}
	return 0
}

//*
// Information for the Hdds block token.
// When adding further fields, make sure they are optional as they would
// otherwise not be backwards compatible.
type BlockTokenSecretProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId        *string                                 `protobuf:"bytes,1,req,name=ownerId" json:"ownerId,omitempty"`
	BlockId        *string                                 `protobuf:"bytes,2,req,name=blockId" json:"blockId,omitempty"`
	ExpiryDate     *uint64                                 `protobuf:"varint,3,req,name=expiryDate" json:"expiryDate,omitempty"`
	OmCertSerialId *string                                 `protobuf:"bytes,4,req,name=omCertSerialId" json:"omCertSerialId,omitempty"`
	Modes          []BlockTokenSecretProto_AccessModeProto `protobuf:"varint,5,rep,name=modes,enum=hadoop.hdds.BlockTokenSecretProto_AccessModeProto" json:"modes,omitempty"`
	MaxLength      *uint64                                 `protobuf:"varint,6,req,name=maxLength" json:"maxLength,omitempty"`
}

func (x *BlockTokenSecretProto) Reset() {
	*x = BlockTokenSecretProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hdds_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockTokenSecretProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockTokenSecretProto) ProtoMessage() {}

func (x *BlockTokenSecretProto) ProtoReflect() protoreflect.Message {
	mi := &file_hdds_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockTokenSecretProto.ProtoReflect.Descriptor instead.
func (*BlockTokenSecretProto) Descriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{15}
}

func (x *BlockTokenSecretProto) GetOwnerId() string {
	if x != nil && x.OwnerId != nil {
		return *x.OwnerId
	}
	return ""
}

func (x *BlockTokenSecretProto) GetBlockId() string {
	if x != nil && x.BlockId != nil {
		return *x.BlockId
	}
	return ""
}

func (x *BlockTokenSecretProto) GetExpiryDate() uint64 {
	if x != nil && x.ExpiryDate != nil {
		return *x.ExpiryDate
	}
	return 0
}

func (x *BlockTokenSecretProto) GetOmCertSerialId() string {
	if x != nil && x.OmCertSerialId != nil {
		return *x.OmCertSerialId
	}
	return ""
}

func (x *BlockTokenSecretProto) GetModes() []BlockTokenSecretProto_AccessModeProto {
	if x != nil {
		return x.Modes
	}
	return nil
}

func (x *BlockTokenSecretProto) GetMaxLength() uint64 {
	if x != nil && x.MaxLength != nil {
		return *x.MaxLength
	}
	return 0
}

type BlockID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerBlockID      *ContainerBlockID `protobuf:"bytes,1,req,name=containerBlockID" json:"containerBlockID,omitempty"`
	BlockCommitSequenceId *uint64           `protobuf:"varint,2,opt,name=blockCommitSequenceId,def=0" json:"blockCommitSequenceId,omitempty"`
}

// Default values for BlockID fields.
const (
	Default_BlockID_BlockCommitSequenceId = uint64(0)
)

func (x *BlockID) Reset() {
	*x = BlockID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hdds_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockID) ProtoMessage() {}

func (x *BlockID) ProtoReflect() protoreflect.Message {
	mi := &file_hdds_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockID.ProtoReflect.Descriptor instead.
func (*BlockID) Descriptor() ([]byte, []int) {
	return file_hdds_proto_rawDescGZIP(), []int{16}
}

func (x *BlockID) GetContainerBlockID() *ContainerBlockID {
	if x != nil {
		return x.ContainerBlockID
	}
	return nil
}

func (x *BlockID) GetBlockCommitSequenceId() uint64 {
	if x != nil && x.BlockCommitSequenceId != nil {
		return *x.BlockCommitSequenceId
	}
	return Default_BlockID_BlockCommitSequenceId
}

var File_hdds_proto protoreflect.FileDescriptor

var file_hdds_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x68, 0x64, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x68, 0x61,
	0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x64, 0x73, 0x22, 0x4c, 0x0a, 0x04, 0x55, 0x55, 0x49,
	0x44, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x6f, 0x73, 0x74, 0x53, 0x69, 0x67, 0x42, 0x69, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x6f, 0x73, 0x74, 0x53, 0x69, 0x67, 0x42,
	0x69, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x53, 0x69, 0x67, 0x42,
	0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x65, 0x61, 0x73, 0x74,
	0x53, 0x69, 0x67, 0x42, 0x69, 0x74, 0x73, 0x22, 0x9c, 0x03, 0x0a, 0x14, 0x44, 0x61, 0x74, 0x61,
	0x6e, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x02, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27,
	0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x64, 0x73, 0x2e, 0x50, 0x6f, 0x72, 0x74,
	0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x65, 0x72, 0x74, 0x53,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63,
	0x65, 0x72, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a,
	0x0f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x74, 0x75, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x74, 0x75, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x75, 0x75, 0x69,
	0x64, 0x31, 0x32, 0x38, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x68, 0x61, 0x64,
	0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x64, 0x73, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x07, 0x75,
	0x75, 0x69, 0x64, 0x31, 0x32, 0x38, 0x22, 0x91, 0x01, 0x0a, 0x18, 0x4f, 0x7a, 0x6f, 0x6e, 0x65,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x27, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x64, 0x73, 0x2e, 0x50,
	0x6f, 0x72, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x30, 0x0a, 0x04, 0x50, 0x6f,
	0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x49, 0x0a, 0x0a,
	0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x07, 0x75, 0x75,
	0x69, 0x64, 0x31, 0x32, 0x38, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x68, 0x61,
	0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x64, 0x73, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x07,
	0x75, 0x75, 0x69, 0x64, 0x31, 0x32, 0x38, 0x22, 0xd5, 0x03, 0x0a, 0x08, 0x50, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68,
	0x64, 0x64, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x12, 0x44, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1a, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x64, 0x73, 0x2e, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x12, 0x50, 0x49,
	0x50, 0x45, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x45, 0x44,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68,
	0x64, 0x64, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x3a, 0x0b, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x5f, 0x41, 0x4c, 0x4f, 0x4e, 0x45,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e,
	0x68, 0x64, 0x64, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x3a, 0x03, 0x4f, 0x4e, 0x45, 0x52, 0x06, 0x66, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x27, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x64, 0x73, 0x2e, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2c, 0x0a, 0x11,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x33, 0x0a, 0x0b, 0x6c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x49, 0x44, 0x31, 0x32, 0x38, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x64, 0x73, 0x2e, 0x55, 0x55,
	0x49, 0x44, 0x52, 0x0b, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x44, 0x31, 0x32, 0x38, 0x22,
	0x32, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x79, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x6e,
	0x6f, 0x64, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x68, 0x61,
	0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x64, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x6e, 0x6f,
	0x64, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x06,
	0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x12, 0x36, 0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x68, 0x61, 0x64,
	0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x64, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x22, 0x33,
	0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x27, 0x0a, 0x05, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x68, 0x61, 0x64, 0x6f,
	0x6f, 0x70, 0x2e, 0x68, 0x64, 0x64, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x22, 0x8a, 0x04, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x02, 0x28, 0x03, 0x52,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x31, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x68, 0x61,
	0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x64, 0x73, 0x2e, 0x4c, 0x69, 0x66, 0x65, 0x43, 0x79,
	0x63, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x37, 0x0a, 0x0a, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x64,
	0x73, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x52, 0x0a, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x64,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x02, 0x28, 0x04, 0x52, 0x09, 0x75, 0x73, 0x65,
	0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x4f, 0x66, 0x4b, 0x65, 0x79, 0x73, 0x18, 0x05, 0x20, 0x02, 0x28, 0x04, 0x52, 0x0c, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x07, 0x20, 0x02, 0x28,
	0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x13, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x4c, 0x0a, 0x11, 0x72, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18,
	0x0a, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68,
	0x64, 0x64, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x11, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x46, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x02, 0x28,
	0x0e, 0x32, 0x1c, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x64, 0x73, 0x2e,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x91, 0x01, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x57, 0x69,
	0x74, 0x68, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x45, 0x0a, 0x0d, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x64, 0x73, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x31, 0x0a, 0x08, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20,
	0x02, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x64,
	0x73, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x08, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x22, 0x32, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x63, 0x6d, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18,
	0x0a, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x44, 0x22, 0x4d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53,
	0x63, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x63, 0x6d, 0x49, 0x64, 0x22, 0x8f, 0x01, 0x0a, 0x10, 0x45, 0x78, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x0a, 0x09,
	0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x64, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x39,
	0x0a, 0x0b, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x64,
	0x73, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x44, 0x52, 0x0b, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x73, 0x22, 0x4e, 0x0a, 0x10, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x03, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x49, 0x44, 0x18, 0x02, 0x20, 0x02, 0x28, 0x03,
	0x52, 0x07, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x49, 0x44, 0x22, 0xb9, 0x02, 0x0a, 0x15, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x79, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x04, 0x52, 0x0a, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6f, 0x6d, 0x43, 0x65, 0x72,
	0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x04, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x0e, 0x6f, 0x6d, 0x43, 0x65, 0x72, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x12,
	0x48, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x32,
	0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x64, 0x73, 0x2e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x78,
	0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x06, 0x20, 0x02, 0x28, 0x04, 0x52, 0x09, 0x6d, 0x61,
	0x78, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x3c, 0x0a, 0x0f, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4d, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x45,
	0x41, 0x44, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x02, 0x12,
	0x08, 0x0a, 0x04, 0x43, 0x4f, 0x50, 0x59, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c,
	0x45, 0x54, 0x45, 0x10, 0x04, 0x22, 0x8d, 0x01, 0x0a, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49,
	0x44, 0x12, 0x49, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x68, 0x61,
	0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x64, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x12, 0x37, 0x0a, 0x15,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x53, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x3a, 0x01, 0x30, 0x52, 0x15,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x53, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x49, 0x64, 0x2a, 0x65, 0x0a, 0x0d, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x49, 0x50, 0x45, 0x4c, 0x49,
	0x4e, 0x45, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x11,
	0x0a, 0x0d, 0x50, 0x49, 0x50, 0x45, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x10,
	0x02, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x49, 0x50, 0x45, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x44, 0x4f,
	0x52, 0x4d, 0x41, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x49, 0x50, 0x45, 0x4c,
	0x49, 0x4e, 0x45, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x04, 0x2a, 0x34, 0x0a, 0x08,
	0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4d, 0x10, 0x01,
	0x12, 0x07, 0x0a, 0x03, 0x53, 0x43, 0x4d, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x41, 0x54,
	0x41, 0x4e, 0x4f, 0x44, 0x45, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x43, 0x4f, 0x4e,
	0x10, 0x04, 0x2a, 0x56, 0x0a, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x59, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x53, 0x54, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x41, 0x44, 0x10,
	0x03, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x45, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f,
	0x4e, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x45, 0x43, 0x4f, 0x4d, 0x4d,
	0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x45, 0x44, 0x10, 0x05, 0x2a, 0x23, 0x0a, 0x0a, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4c, 0x55, 0x53,
	0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4f, 0x4f, 0x4c, 0x10, 0x02, 0x2a,
	0x60, 0x0a, 0x0e, 0x4c, 0x69, 0x66, 0x65, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43,
	0x4c, 0x4f, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x51, 0x55, 0x41, 0x53,
	0x49, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c,
	0x4f, 0x53, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49,
	0x4e, 0x47, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10,
	0x06, 0x2a, 0x64, 0x0a, 0x0e, 0x4c, 0x69, 0x66, 0x65, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x49, 0x5a, 0x45, 0x10,
	0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x51, 0x55, 0x41, 0x53, 0x49, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45,
	0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x03, 0x12, 0x0f, 0x0a,
	0x0b, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x04, 0x12, 0x0a,
	0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4c,
	0x45, 0x41, 0x4e, 0x55, 0x50, 0x10, 0x06, 0x2a, 0x3a, 0x0a, 0x0f, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x41,
	0x54, 0x49, 0x53, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x5f, 0x41,
	0x4c, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x45,
	0x44, 0x10, 0x03, 0x2a, 0x27, 0x0a, 0x11, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x4e, 0x45, 0x10,
	0x01, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x48, 0x52, 0x45, 0x45, 0x10, 0x03, 0x2a, 0xf0, 0x01, 0x0a,
	0x06, 0x53, 0x63, 0x6d, 0x4f, 0x70, 0x73, 0x12, 0x11, 0x0a, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x6b, 0x65,
	0x79, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x10,
	0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x53, 0x63, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x10,
	0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x10, 0x04, 0x12, 0x1d, 0x0a, 0x19, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x10,
	0x05, 0x12, 0x15, 0x0a, 0x11, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x67, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x10, 0x07, 0x12, 0x1c, 0x0a, 0x18, 0x67, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x08, 0x12, 0x11, 0x0a, 0x0d, 0x6c, 0x69, 0x73, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x10, 0x09, 0x12, 0x13, 0x0a, 0x0f, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x10, 0x0a,
	0x12, 0x0d, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x10, 0x0b, 0x42,
	0x64, 0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x68, 0x61,
	0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x64, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0a, 0x48, 0x64, 0x64, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x65, 0x2d, 0x67, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x64, 0x64, 0x73, 0x88,
	0x01, 0x01, 0xa0, 0x01, 0x01,
}

var (
	file_hdds_proto_rawDescOnce sync.Once
	file_hdds_proto_rawDescData = file_hdds_proto_rawDesc
)

func file_hdds_proto_rawDescGZIP() []byte {
	file_hdds_proto_rawDescOnce.Do(func() {
		file_hdds_proto_rawDescData = protoimpl.X.CompressGZIP(file_hdds_proto_rawDescData)
	})
	return file_hdds_proto_rawDescData
}

var file_hdds_proto_enumTypes = make([]protoimpl.EnumInfo, 10)
var file_hdds_proto_msgTypes = make([]protoimpl.MessageInfo, 17)
var file_hdds_proto_goTypes = []interface{}{
	(PipelineState)(0),     // 0: hadoop.hdds.PipelineState
	(NodeType)(0),          // 1: hadoop.hdds.NodeType
	(NodeState)(0),         // 2: hadoop.hdds.NodeState
	(QueryScope)(0),        // 3: hadoop.hdds.QueryScope
	(LifeCycleState)(0),    // 4: hadoop.hdds.LifeCycleState
	(LifeCycleEvent)(0),    // 5: hadoop.hdds.LifeCycleEvent
	(ReplicationType)(0),   // 6: hadoop.hdds.ReplicationType
	(ReplicationFactor)(0), // 7: hadoop.hdds.ReplicationFactor
	(ScmOps)(0),            // 8: hadoop.hdds.ScmOps
	(BlockTokenSecretProto_AccessModeProto)(0), // 9: hadoop.hdds.BlockTokenSecretProto.AccessModeProto
	(*UUID)(nil),                     // 10: hadoop.hdds.UUID
	(*DatanodeDetailsProto)(nil),     // 11: hadoop.hdds.DatanodeDetailsProto
	(*OzoneManagerDetailsProto)(nil), // 12: hadoop.hdds.OzoneManagerDetailsProto
	(*Port)(nil),                     // 13: hadoop.hdds.Port
	(*PipelineID)(nil),               // 14: hadoop.hdds.PipelineID
	(*Pipeline)(nil),                 // 15: hadoop.hdds.Pipeline
	(*KeyValue)(nil),                 // 16: hadoop.hdds.KeyValue
	(*Node)(nil),                     // 17: hadoop.hdds.Node
	(*NodePool)(nil),                 // 18: hadoop.hdds.NodePool
	(*ContainerInfoProto)(nil),       // 19: hadoop.hdds.ContainerInfoProto
	(*ContainerWithPipeline)(nil),    // 20: hadoop.hdds.ContainerWithPipeline
	(*GetScmInfoRequestProto)(nil),   // 21: hadoop.hdds.GetScmInfoRequestProto
	(*GetScmInfoResponseProto)(nil),  // 22: hadoop.hdds.GetScmInfoResponseProto
	(*ExcludeListProto)(nil),         // 23: hadoop.hdds.ExcludeListProto
	(*ContainerBlockID)(nil),         // 24: hadoop.hdds.ContainerBlockID
	(*BlockTokenSecretProto)(nil),    // 25: hadoop.hdds.BlockTokenSecretProto
	(*BlockID)(nil),                  // 26: hadoop.hdds.BlockID
}
var file_hdds_proto_depIdxs = []int32{
	13, // 0: hadoop.hdds.DatanodeDetailsProto.ports:type_name -> hadoop.hdds.Port
	10, // 1: hadoop.hdds.DatanodeDetailsProto.uuid128:type_name -> hadoop.hdds.UUID
	13, // 2: hadoop.hdds.OzoneManagerDetailsProto.ports:type_name -> hadoop.hdds.Port
	10, // 3: hadoop.hdds.PipelineID.uuid128:type_name -> hadoop.hdds.UUID
	11, // 4: hadoop.hdds.Pipeline.members:type_name -> hadoop.hdds.DatanodeDetailsProto
	0,  // 5: hadoop.hdds.Pipeline.state:type_name -> hadoop.hdds.PipelineState
	6,  // 6: hadoop.hdds.Pipeline.type:type_name -> hadoop.hdds.ReplicationType
	7,  // 7: hadoop.hdds.Pipeline.factor:type_name -> hadoop.hdds.ReplicationFactor
	14, // 8: hadoop.hdds.Pipeline.id:type_name -> hadoop.hdds.PipelineID
	10, // 9: hadoop.hdds.Pipeline.leaderID128:type_name -> hadoop.hdds.UUID
	11, // 10: hadoop.hdds.Node.nodeID:type_name -> hadoop.hdds.DatanodeDetailsProto
	2,  // 11: hadoop.hdds.Node.nodeStates:type_name -> hadoop.hdds.NodeState
	17, // 12: hadoop.hdds.NodePool.nodes:type_name -> hadoop.hdds.Node
	4,  // 13: hadoop.hdds.ContainerInfoProto.state:type_name -> hadoop.hdds.LifeCycleState
	14, // 14: hadoop.hdds.ContainerInfoProto.pipelineID:type_name -> hadoop.hdds.PipelineID
	7,  // 15: hadoop.hdds.ContainerInfoProto.replicationFactor:type_name -> hadoop.hdds.ReplicationFactor
	6,  // 16: hadoop.hdds.ContainerInfoProto.replicationType:type_name -> hadoop.hdds.ReplicationType
	19, // 17: hadoop.hdds.ContainerWithPipeline.containerInfo:type_name -> hadoop.hdds.ContainerInfoProto
	15, // 18: hadoop.hdds.ContainerWithPipeline.pipeline:type_name -> hadoop.hdds.Pipeline
	14, // 19: hadoop.hdds.ExcludeListProto.pipelineIds:type_name -> hadoop.hdds.PipelineID
	9,  // 20: hadoop.hdds.BlockTokenSecretProto.modes:type_name -> hadoop.hdds.BlockTokenSecretProto.AccessModeProto
	24, // 21: hadoop.hdds.BlockID.containerBlockID:type_name -> hadoop.hdds.ContainerBlockID
	22, // [22:22] is the sub-list for method output_type
	22, // [22:22] is the sub-list for method input_type
	22, // [22:22] is the sub-list for extension type_name
	22, // [22:22] is the sub-list for extension extendee
	0,  // [0:22] is the sub-list for field type_name
}

func init() { file_hdds_proto_init() }
func file_hdds_proto_init() {
	if File_hdds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hdds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hdds_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatanodeDetailsProto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hdds_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OzoneManagerDetailsProto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hdds_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hdds_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hdds_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pipeline); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hdds_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hdds_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hdds_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodePool); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hdds_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerInfoProto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hdds_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerWithPipeline); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hdds_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetScmInfoRequestProto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hdds_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetScmInfoResponseProto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hdds_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExcludeListProto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hdds_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerBlockID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hdds_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockTokenSecretProto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hdds_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hdds_proto_rawDesc,
			NumEnums:      10,
			NumMessages:   17,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hdds_proto_goTypes,
		DependencyIndexes: file_hdds_proto_depIdxs,
		EnumInfos:         file_hdds_proto_enumTypes,
		MessageInfos:      file_hdds_proto_msgTypes,
	}.Build()
	File_hdds_proto = out.File
	file_hdds_proto_rawDesc = nil
	file_hdds_proto_goTypes = nil
	file_hdds_proto_depIdxs = nil
}
