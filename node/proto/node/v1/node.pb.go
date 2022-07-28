// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: proto/node/v1/node.proto

package nodev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	v1 "vebridge/node/proto/gossip/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InjectGovernanceVAARequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Index of the current guardian set.
	CurrentSetIndex uint32 `protobuf:"varint,1,opt,name=current_set_index,json=currentSetIndex,proto3" json:"current_set_index,omitempty"`
	// List of governance VAA messages to inject.
	Messages []*GovernanceMessage `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *InjectGovernanceVAARequest) Reset() {
	*x = InjectGovernanceVAARequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_node_v1_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InjectGovernanceVAARequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InjectGovernanceVAARequest) ProtoMessage() {}

func (x *InjectGovernanceVAARequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_node_v1_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InjectGovernanceVAARequest.ProtoReflect.Descriptor instead.
func (*InjectGovernanceVAARequest) Descriptor() ([]byte, []int) {
	return file_proto_node_v1_node_proto_rawDescGZIP(), []int{0}
}

func (x *InjectGovernanceVAARequest) GetCurrentSetIndex() uint32 {
	if x != nil {
		return x.CurrentSetIndex
	}
	return 0
}

func (x *InjectGovernanceVAARequest) GetMessages() []*GovernanceMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

type GovernanceMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is a uint32 to match the on-chain timestamp representation. This becomes a problem in 2106 (sorry).
	Timestamp uint32 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are assignable to Payload:
	//	*GovernanceMessage_GuardianSet
	//	*GovernanceMessage_ContractUpgrade
	Payload isGovernanceMessage_Payload `protobuf_oneof:"payload"`
}

func (x *GovernanceMessage) Reset() {
	*x = GovernanceMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_node_v1_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GovernanceMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GovernanceMessage) ProtoMessage() {}

func (x *GovernanceMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_node_v1_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GovernanceMessage.ProtoReflect.Descriptor instead.
func (*GovernanceMessage) Descriptor() ([]byte, []int) {
	return file_proto_node_v1_node_proto_rawDescGZIP(), []int{1}
}

func (x *GovernanceMessage) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (m *GovernanceMessage) GetPayload() isGovernanceMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *GovernanceMessage) GetGuardianSet() *GuardianSetUpdate {
	if x, ok := x.GetPayload().(*GovernanceMessage_GuardianSet); ok {
		return x.GuardianSet
	}
	return nil
}

func (x *GovernanceMessage) GetContractUpgrade() *ContractUpgrade {
	if x, ok := x.GetPayload().(*GovernanceMessage_ContractUpgrade); ok {
		return x.ContractUpgrade
	}
	return nil
}

type isGovernanceMessage_Payload interface {
	isGovernanceMessage_Payload()
}

type GovernanceMessage_GuardianSet struct {
	GuardianSet *GuardianSetUpdate `protobuf:"bytes,3,opt,name=guardian_set,json=guardianSet,proto3,oneof"`
}

type GovernanceMessage_ContractUpgrade struct {
	ContractUpgrade *ContractUpgrade `protobuf:"bytes,4,opt,name=contract_upgrade,json=contractUpgrade,proto3,oneof"`
}

func (*GovernanceMessage_GuardianSet) isGovernanceMessage_Payload() {}

func (*GovernanceMessage_ContractUpgrade) isGovernanceMessage_Payload() {}

type InjectGovernanceVAAResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Canonical digests of the submitted VAAs.
	Digests [][]byte `protobuf:"bytes,1,rep,name=digests,proto3" json:"digests,omitempty"`
}

func (x *InjectGovernanceVAAResponse) Reset() {
	*x = InjectGovernanceVAAResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_node_v1_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InjectGovernanceVAAResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InjectGovernanceVAAResponse) ProtoMessage() {}

func (x *InjectGovernanceVAAResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_node_v1_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InjectGovernanceVAAResponse.ProtoReflect.Descriptor instead.
func (*InjectGovernanceVAAResponse) Descriptor() ([]byte, []int) {
	return file_proto_node_v1_node_proto_rawDescGZIP(), []int{2}
}

func (x *InjectGovernanceVAAResponse) GetDigests() [][]byte {
	if x != nil {
		return x.Digests
	}
	return nil
}

// GuardianSet represents a new guardian set to be submitted to and signed by the node.
// During the genesis procedure, this data structure will be assembled using off-chain collaborative tooling
// like GitHub using a human-readable encoding, so readability is a concern.
type GuardianSetUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guardians []*GuardianSetUpdate_Guardian `protobuf:"bytes,3,rep,name=guardians,proto3" json:"guardians,omitempty"`
}

func (x *GuardianSetUpdate) Reset() {
	*x = GuardianSetUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_node_v1_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuardianSetUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuardianSetUpdate) ProtoMessage() {}

func (x *GuardianSetUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_node_v1_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuardianSetUpdate.ProtoReflect.Descriptor instead.
func (*GuardianSetUpdate) Descriptor() ([]byte, []int) {
	return file_proto_node_v1_node_proto_rawDescGZIP(), []int{3}
}

func (x *GuardianSetUpdate) GetGuardians() []*GuardianSetUpdate_Guardian {
	if x != nil {
		return x.Guardians
	}
	return nil
}

// GuardianKey specifies the on-disk format for a node's guardian key.
type GuardianKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// data is the binary representation of the secp256k1 private key.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Whether this key is deterministically generated and unsuitable for production mode.
	UnsafeDeterministicKey bool `protobuf:"varint,2,opt,name=unsafe_deterministic_key,json=unsafeDeterministicKey,proto3" json:"unsafe_deterministic_key,omitempty"`
}

func (x *GuardianKey) Reset() {
	*x = GuardianKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_node_v1_node_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuardianKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuardianKey) ProtoMessage() {}

func (x *GuardianKey) ProtoReflect() protoreflect.Message {
	mi := &file_proto_node_v1_node_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuardianKey.ProtoReflect.Descriptor instead.
func (*GuardianKey) Descriptor() ([]byte, []int) {
	return file_proto_node_v1_node_proto_rawDescGZIP(), []int{4}
}

func (x *GuardianKey) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GuardianKey) GetUnsafeDeterministicKey() bool {
	if x != nil {
		return x.UnsafeDeterministicKey
	}
	return false
}

// ContractUpgrade represents a Wormhole contract update to be submitted to and signed by the node.
type ContractUpgrade struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the chain where the Wormhole contract should be updated (uint8).
	ChainId uint32 `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// Address of the new program/contract.
	NewContract []byte `protobuf:"bytes,2,opt,name=new_contract,json=newContract,proto3" json:"new_contract,omitempty"`
}

func (x *ContractUpgrade) Reset() {
	*x = ContractUpgrade{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_node_v1_node_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractUpgrade) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractUpgrade) ProtoMessage() {}

func (x *ContractUpgrade) ProtoReflect() protoreflect.Message {
	mi := &file_proto_node_v1_node_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractUpgrade.ProtoReflect.Descriptor instead.
func (*ContractUpgrade) Descriptor() ([]byte, []int) {
	return file_proto_node_v1_node_proto_rawDescGZIP(), []int{5}
}

func (x *ContractUpgrade) GetChainId() uint32 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *ContractUpgrade) GetNewContract() []byte {
	if x != nil {
		return x.NewContract
	}
	return nil
}

type SendObservationRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObservationRequest *v1.ObservationRequest `protobuf:"bytes,1,opt,name=observation_request,json=observationRequest,proto3" json:"observation_request,omitempty"`
}

func (x *SendObservationRequestRequest) Reset() {
	*x = SendObservationRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_node_v1_node_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendObservationRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendObservationRequestRequest) ProtoMessage() {}

func (x *SendObservationRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_node_v1_node_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendObservationRequestRequest.ProtoReflect.Descriptor instead.
func (*SendObservationRequestRequest) Descriptor() ([]byte, []int) {
	return file_proto_node_v1_node_proto_rawDescGZIP(), []int{6}
}

func (x *SendObservationRequestRequest) GetObservationRequest() *v1.ObservationRequest {
	if x != nil {
		return x.ObservationRequest
	}
	return nil
}

type SendObservationRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendObservationRequestResponse) Reset() {
	*x = SendObservationRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_node_v1_node_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendObservationRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendObservationRequestResponse) ProtoMessage() {}

func (x *SendObservationRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_node_v1_node_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendObservationRequestResponse.ProtoReflect.Descriptor instead.
func (*SendObservationRequestResponse) Descriptor() ([]byte, []int) {
	return file_proto_node_v1_node_proto_rawDescGZIP(), []int{7}
}

// List of guardian set members.
type GuardianSetUpdate_Guardian struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Guardian key pubkey. Stored as hex string with 0x prefix for human readability -
	// this is the canonical Ethereum representation.
	Pubkey string `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	// Optional descriptive name. Not stored on any chain, purely informational.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GuardianSetUpdate_Guardian) Reset() {
	*x = GuardianSetUpdate_Guardian{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_node_v1_node_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuardianSetUpdate_Guardian) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuardianSetUpdate_Guardian) ProtoMessage() {}

func (x *GuardianSetUpdate_Guardian) ProtoReflect() protoreflect.Message {
	mi := &file_proto_node_v1_node_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuardianSetUpdate_Guardian.ProtoReflect.Descriptor instead.
func (*GuardianSetUpdate_Guardian) Descriptor() ([]byte, []int) {
	return file_proto_node_v1_node_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GuardianSetUpdate_Guardian) GetPubkey() string {
	if x != nil {
		return x.Pubkey
	}
	return ""
}

func (x *GuardianSetUpdate_Guardian) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_proto_node_v1_node_proto protoreflect.FileDescriptor

var file_proto_node_v1_node_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x73, 0x73, 0x69,
	0x70, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x80, 0x01, 0x0a, 0x1a, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x6f, 0x76, 0x65,
	0x72, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x56, 0x41, 0x41, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2a, 0x0a, 0x11, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x36, 0x0a, 0x08,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x22, 0xc4, 0x01, 0x0a, 0x11, 0x47, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3f, 0x0a, 0x0c, 0x67, 0x75, 0x61, 0x72,
	0x64, 0x69, 0x61, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61,
	0x6e, 0x53, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x67, 0x75,
	0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x45, 0x0a, 0x10, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x48, 0x00, 0x52,
	0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x37, 0x0a, 0x1b, 0x49,
	0x6e, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x56,
	0x41, 0x41, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x07, 0x64, 0x69, 0x67,
	0x65, 0x73, 0x74, 0x73, 0x22, 0x8e, 0x01, 0x0a, 0x11, 0x47, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61,
	0x6e, 0x53, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x41, 0x0a, 0x09, 0x67, 0x75,
	0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e,
	0x53, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x75, 0x61, 0x72, 0x64, 0x69,
	0x61, 0x6e, 0x52, 0x09, 0x67, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x73, 0x1a, 0x36, 0x0a,
	0x08, 0x47, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5b, 0x0a, 0x0b, 0x47, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61,
	0x6e, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x18, 0x75, 0x6e, 0x73, 0x61,
	0x66, 0x65, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x75, 0x6e, 0x73, 0x61,
	0x66, 0x65, 0x44, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x22, 0x4f, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x55, 0x70,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x22, 0x6f, 0x0a, 0x1d, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x13, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x12, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x20, 0x0a, 0x1e, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xe4, 0x01, 0x0a, 0x15, 0x4e, 0x6f, 0x64, 0x65, 0x50,
	0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x60, 0x0a, 0x13, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x6f, 0x76, 0x65, 0x72, 0x6e,
	0x61, 0x6e, 0x63, 0x65, 0x56, 0x41, 0x41, 0x12, 0x23, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x56, 0x41, 0x41, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x6f, 0x76,
	0x65, 0x72, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x56, 0x41, 0x41, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x69, 0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x24, 0x5a,
	0x22, 0x76, 0x65, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6e, 0x6f, 0x64,
	0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_node_v1_node_proto_rawDescOnce sync.Once
	file_proto_node_v1_node_proto_rawDescData = file_proto_node_v1_node_proto_rawDesc
)

func file_proto_node_v1_node_proto_rawDescGZIP() []byte {
	file_proto_node_v1_node_proto_rawDescOnce.Do(func() {
		file_proto_node_v1_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_node_v1_node_proto_rawDescData)
	})
	return file_proto_node_v1_node_proto_rawDescData
}

var file_proto_node_v1_node_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_node_v1_node_proto_goTypes = []interface{}{
	(*InjectGovernanceVAARequest)(nil),     // 0: node.v1.InjectGovernanceVAARequest
	(*GovernanceMessage)(nil),              // 1: node.v1.GovernanceMessage
	(*InjectGovernanceVAAResponse)(nil),    // 2: node.v1.InjectGovernanceVAAResponse
	(*GuardianSetUpdate)(nil),              // 3: node.v1.GuardianSetUpdate
	(*GuardianKey)(nil),                    // 4: node.v1.GuardianKey
	(*ContractUpgrade)(nil),                // 5: node.v1.ContractUpgrade
	(*SendObservationRequestRequest)(nil),  // 6: node.v1.SendObservationRequestRequest
	(*SendObservationRequestResponse)(nil), // 7: node.v1.SendObservationRequestResponse
	(*GuardianSetUpdate_Guardian)(nil),     // 8: node.v1.GuardianSetUpdate.Guardian
	(*v1.ObservationRequest)(nil),          // 9: gossip.v1.ObservationRequest
}
var file_proto_node_v1_node_proto_depIdxs = []int32{
	1, // 0: node.v1.InjectGovernanceVAARequest.messages:type_name -> node.v1.GovernanceMessage
	3, // 1: node.v1.GovernanceMessage.guardian_set:type_name -> node.v1.GuardianSetUpdate
	5, // 2: node.v1.GovernanceMessage.contract_upgrade:type_name -> node.v1.ContractUpgrade
	8, // 3: node.v1.GuardianSetUpdate.guardians:type_name -> node.v1.GuardianSetUpdate.Guardian
	9, // 4: node.v1.SendObservationRequestRequest.observation_request:type_name -> gossip.v1.ObservationRequest
	0, // 5: node.v1.NodePrivilegedService.InjectGovernanceVAA:input_type -> node.v1.InjectGovernanceVAARequest
	6, // 6: node.v1.NodePrivilegedService.SendObservationRequest:input_type -> node.v1.SendObservationRequestRequest
	2, // 7: node.v1.NodePrivilegedService.InjectGovernanceVAA:output_type -> node.v1.InjectGovernanceVAAResponse
	7, // 8: node.v1.NodePrivilegedService.SendObservationRequest:output_type -> node.v1.SendObservationRequestResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_node_v1_node_proto_init() }
func file_proto_node_v1_node_proto_init() {
	if File_proto_node_v1_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_node_v1_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InjectGovernanceVAARequest); i {
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
		file_proto_node_v1_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GovernanceMessage); i {
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
		file_proto_node_v1_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InjectGovernanceVAAResponse); i {
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
		file_proto_node_v1_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuardianSetUpdate); i {
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
		file_proto_node_v1_node_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuardianKey); i {
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
		file_proto_node_v1_node_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractUpgrade); i {
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
		file_proto_node_v1_node_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendObservationRequestRequest); i {
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
		file_proto_node_v1_node_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendObservationRequestResponse); i {
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
		file_proto_node_v1_node_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuardianSetUpdate_Guardian); i {
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
	file_proto_node_v1_node_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*GovernanceMessage_GuardianSet)(nil),
		(*GovernanceMessage_ContractUpgrade)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_node_v1_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_node_v1_node_proto_goTypes,
		DependencyIndexes: file_proto_node_v1_node_proto_depIdxs,
		MessageInfos:      file_proto_node_v1_node_proto_msgTypes,
	}.Build()
	File_proto_node_v1_node_proto = out.File
	file_proto_node_v1_node_proto_rawDesc = nil
	file_proto_node_v1_node_proto_goTypes = nil
	file_proto_node_v1_node_proto_depIdxs = nil
}
