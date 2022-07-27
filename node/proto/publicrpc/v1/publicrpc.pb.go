// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: proto/publicrpc/v1/publicrpc.proto

package publicrpcv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	v1 "wormhole/node/proto/gossip/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ChainID int32

const (
	ChainID_CHAIN_ID_UNSPECIFIED ChainID = 0
	ChainID_CHAIN_ID_SOLANA      ChainID = 1
	ChainID_CHAIN_ID_ETHEREUM    ChainID = 2
	ChainID_CHAIN_ID_TERRA       ChainID = 3
	ChainID_CHAIN_ID_BSC         ChainID = 4
	ChainID_CHAIN_ID_POLYGON     ChainID = 5
	ChainID_CHAIN_ID_AVALANCHE   ChainID = 6
	ChainID_CHAIN_ID_OASIS       ChainID = 7
	ChainID_CHAIN_ID_ALGORAND    ChainID = 8
	ChainID_CHAIN_ID_AURORA      ChainID = 9
	ChainID_CHAIN_ID_FANTOM      ChainID = 10
	ChainID_CHAIN_ID_KARURA      ChainID = 11
	ChainID_CHAIN_ID_ACALA       ChainID = 12
	ChainID_CHAIN_ID_KLAYTN      ChainID = 13
	ChainID_CHAIN_ID_CELO        ChainID = 14
	ChainID_CHAIN_ID_NEAR        ChainID = 15
	ChainID_CHAIN_ID_MOONBEAM    ChainID = 16
	ChainID_CHAIN_ID_NEON        ChainID = 17
	ChainID_CHAIN_ID_TERRA2      ChainID = 18
	ChainID_CHAIN_ID_INJECTIVE   ChainID = 19
	// Special case - Eth has two testnets. CHAIN_ID_ETHEREUM is Goerli,
	// but we also want to connect to Ropsten, so we add a separate chain.
	ChainID_CHAIN_ID_ETHEREUM_ROPSTEN ChainID = 10001
)

// Enum value maps for ChainID.
var (
	ChainID_name = map[int32]string{
		0:     "CHAIN_ID_UNSPECIFIED",
		1:     "CHAIN_ID_SOLANA",
		2:     "CHAIN_ID_ETHEREUM",
		3:     "CHAIN_ID_TERRA",
		4:     "CHAIN_ID_BSC",
		5:     "CHAIN_ID_POLYGON",
		6:     "CHAIN_ID_AVALANCHE",
		7:     "CHAIN_ID_OASIS",
		8:     "CHAIN_ID_ALGORAND",
		9:     "CHAIN_ID_AURORA",
		10:    "CHAIN_ID_FANTOM",
		11:    "CHAIN_ID_KARURA",
		12:    "CHAIN_ID_ACALA",
		13:    "CHAIN_ID_KLAYTN",
		14:    "CHAIN_ID_CELO",
		15:    "CHAIN_ID_NEAR",
		16:    "CHAIN_ID_MOONBEAM",
		17:    "CHAIN_ID_NEON",
		18:    "CHAIN_ID_TERRA2",
		19:    "CHAIN_ID_INJECTIVE",
		10001: "CHAIN_ID_ETHEREUM_ROPSTEN",
	}
	ChainID_value = map[string]int32{
		"CHAIN_ID_UNSPECIFIED":      0,
		"CHAIN_ID_SOLANA":           1,
		"CHAIN_ID_ETHEREUM":         2,
		"CHAIN_ID_TERRA":            3,
		"CHAIN_ID_BSC":              4,
		"CHAIN_ID_POLYGON":          5,
		"CHAIN_ID_AVALANCHE":        6,
		"CHAIN_ID_OASIS":            7,
		"CHAIN_ID_ALGORAND":         8,
		"CHAIN_ID_AURORA":           9,
		"CHAIN_ID_FANTOM":           10,
		"CHAIN_ID_KARURA":           11,
		"CHAIN_ID_ACALA":            12,
		"CHAIN_ID_KLAYTN":           13,
		"CHAIN_ID_CELO":             14,
		"CHAIN_ID_NEAR":             15,
		"CHAIN_ID_MOONBEAM":         16,
		"CHAIN_ID_NEON":             17,
		"CHAIN_ID_TERRA2":           18,
		"CHAIN_ID_INJECTIVE":        19,
		"CHAIN_ID_ETHEREUM_ROPSTEN": 10001,
	}
)

func (x ChainID) Enum() *ChainID {
	p := new(ChainID)
	*p = x
	return p
}

func (x ChainID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChainID) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_publicrpc_v1_publicrpc_proto_enumTypes[0].Descriptor()
}

func (ChainID) Type() protoreflect.EnumType {
	return &file_proto_publicrpc_v1_publicrpc_proto_enumTypes[0]
}

func (x ChainID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChainID.Descriptor instead.
func (ChainID) EnumDescriptor() ([]byte, []int) {
	return file_proto_publicrpc_v1_publicrpc_proto_rawDescGZIP(), []int{0}
}

// MessageID is a VAA's globally unique identifier (see data availability design document).
type MessageID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Emitter chain ID.
	EmitterChain ChainID `protobuf:"varint,1,opt,name=emitter_chain,json=emitterChain,proto3,enum=publicrpc.v1.ChainID" json:"emitter_chain,omitempty"`
	// Hex-encoded (without leading 0x) emitter address.
	EmitterAddress string `protobuf:"bytes,2,opt,name=emitter_address,json=emitterAddress,proto3" json:"emitter_address,omitempty"`
	// Sequence number for (emitter_chain, emitter_address).
	Sequence uint64 `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (x *MessageID) Reset() {
	*x = MessageID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_publicrpc_v1_publicrpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageID) ProtoMessage() {}

func (x *MessageID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_publicrpc_v1_publicrpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageID.ProtoReflect.Descriptor instead.
func (*MessageID) Descriptor() ([]byte, []int) {
	return file_proto_publicrpc_v1_publicrpc_proto_rawDescGZIP(), []int{0}
}

func (x *MessageID) GetEmitterChain() ChainID {
	if x != nil {
		return x.EmitterChain
	}
	return ChainID_CHAIN_ID_UNSPECIFIED
}

func (x *MessageID) GetEmitterAddress() string {
	if x != nil {
		return x.EmitterAddress
	}
	return ""
}

func (x *MessageID) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

type GetSignedVAARequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId *MessageID `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (x *GetSignedVAARequest) Reset() {
	*x = GetSignedVAARequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_publicrpc_v1_publicrpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSignedVAARequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSignedVAARequest) ProtoMessage() {}

func (x *GetSignedVAARequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_publicrpc_v1_publicrpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSignedVAARequest.ProtoReflect.Descriptor instead.
func (*GetSignedVAARequest) Descriptor() ([]byte, []int) {
	return file_proto_publicrpc_v1_publicrpc_proto_rawDescGZIP(), []int{1}
}

func (x *GetSignedVAARequest) GetMessageId() *MessageID {
	if x != nil {
		return x.MessageId
	}
	return nil
}

type GetSignedVAAResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VaaBytes []byte `protobuf:"bytes,1,opt,name=vaa_bytes,json=vaaBytes,proto3" json:"vaa_bytes,omitempty"`
}

func (x *GetSignedVAAResponse) Reset() {
	*x = GetSignedVAAResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_publicrpc_v1_publicrpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSignedVAAResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSignedVAAResponse) ProtoMessage() {}

func (x *GetSignedVAAResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_publicrpc_v1_publicrpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSignedVAAResponse.ProtoReflect.Descriptor instead.
func (*GetSignedVAAResponse) Descriptor() ([]byte, []int) {
	return file_proto_publicrpc_v1_publicrpc_proto_rawDescGZIP(), []int{2}
}

func (x *GetSignedVAAResponse) GetVaaBytes() []byte {
	if x != nil {
		return x.VaaBytes
	}
	return nil
}

type GetLastHeartbeatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLastHeartbeatsRequest) Reset() {
	*x = GetLastHeartbeatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_publicrpc_v1_publicrpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLastHeartbeatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastHeartbeatsRequest) ProtoMessage() {}

func (x *GetLastHeartbeatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_publicrpc_v1_publicrpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastHeartbeatsRequest.ProtoReflect.Descriptor instead.
func (*GetLastHeartbeatsRequest) Descriptor() ([]byte, []int) {
	return file_proto_publicrpc_v1_publicrpc_proto_rawDescGZIP(), []int{3}
}

type GetLastHeartbeatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*GetLastHeartbeatsResponse_Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *GetLastHeartbeatsResponse) Reset() {
	*x = GetLastHeartbeatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_publicrpc_v1_publicrpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLastHeartbeatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastHeartbeatsResponse) ProtoMessage() {}

func (x *GetLastHeartbeatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_publicrpc_v1_publicrpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastHeartbeatsResponse.ProtoReflect.Descriptor instead.
func (*GetLastHeartbeatsResponse) Descriptor() ([]byte, []int) {
	return file_proto_publicrpc_v1_publicrpc_proto_rawDescGZIP(), []int{4}
}

func (x *GetLastHeartbeatsResponse) GetEntries() []*GetLastHeartbeatsResponse_Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type GetCurrentGuardianSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCurrentGuardianSetRequest) Reset() {
	*x = GetCurrentGuardianSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_publicrpc_v1_publicrpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentGuardianSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentGuardianSetRequest) ProtoMessage() {}

func (x *GetCurrentGuardianSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_publicrpc_v1_publicrpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentGuardianSetRequest.ProtoReflect.Descriptor instead.
func (*GetCurrentGuardianSetRequest) Descriptor() ([]byte, []int) {
	return file_proto_publicrpc_v1_publicrpc_proto_rawDescGZIP(), []int{5}
}

type GetCurrentGuardianSetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GuardianSet *GuardianSet `protobuf:"bytes,1,opt,name=guardian_set,json=guardianSet,proto3" json:"guardian_set,omitempty"`
}

func (x *GetCurrentGuardianSetResponse) Reset() {
	*x = GetCurrentGuardianSetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_publicrpc_v1_publicrpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentGuardianSetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentGuardianSetResponse) ProtoMessage() {}

func (x *GetCurrentGuardianSetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_publicrpc_v1_publicrpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentGuardianSetResponse.ProtoReflect.Descriptor instead.
func (*GetCurrentGuardianSetResponse) Descriptor() ([]byte, []int) {
	return file_proto_publicrpc_v1_publicrpc_proto_rawDescGZIP(), []int{6}
}

func (x *GetCurrentGuardianSetResponse) GetGuardianSet() *GuardianSet {
	if x != nil {
		return x.GuardianSet
	}
	return nil
}

type GuardianSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Guardian set index
	Index uint32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	// List of guardian addresses as human-readable hex-encoded (leading 0x) addresses.
	Addresses []string `protobuf:"bytes,2,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (x *GuardianSet) Reset() {
	*x = GuardianSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_publicrpc_v1_publicrpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuardianSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuardianSet) ProtoMessage() {}

func (x *GuardianSet) ProtoReflect() protoreflect.Message {
	mi := &file_proto_publicrpc_v1_publicrpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuardianSet.ProtoReflect.Descriptor instead.
func (*GuardianSet) Descriptor() ([]byte, []int) {
	return file_proto_publicrpc_v1_publicrpc_proto_rawDescGZIP(), []int{7}
}

func (x *GuardianSet) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *GuardianSet) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

type GetLastHeartbeatsResponse_Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Verified, hex-encoded (with leading 0x) guardian address. This is the guardian address
	// which signed this heartbeat. The GuardianAddr field inside the heartbeat
	// is NOT verified - remote nodes can put arbitrary data in it.
	VerifiedGuardianAddr string `protobuf:"bytes,1,opt,name=verified_guardian_addr,json=verifiedGuardianAddr,proto3" json:"verified_guardian_addr,omitempty"`
	// Base58-encoded libp2p node address that sent this heartbeat, used to
	// distinguish between multiple nodes running for the same guardian.
	P2PNodeAddr string `protobuf:"bytes,2,opt,name=p2p_node_addr,json=p2pNodeAddr,proto3" json:"p2p_node_addr,omitempty"`
	// Raw heartbeat received from the network. Data is only as trusted
	// as the guardian node that sent it - none of the fields are verified.
	RawHeartbeat *v1.Heartbeat `protobuf:"bytes,3,opt,name=raw_heartbeat,json=rawHeartbeat,proto3" json:"raw_heartbeat,omitempty"`
}

func (x *GetLastHeartbeatsResponse_Entry) Reset() {
	*x = GetLastHeartbeatsResponse_Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_publicrpc_v1_publicrpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLastHeartbeatsResponse_Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastHeartbeatsResponse_Entry) ProtoMessage() {}

func (x *GetLastHeartbeatsResponse_Entry) ProtoReflect() protoreflect.Message {
	mi := &file_proto_publicrpc_v1_publicrpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastHeartbeatsResponse_Entry.ProtoReflect.Descriptor instead.
func (*GetLastHeartbeatsResponse_Entry) Descriptor() ([]byte, []int) {
	return file_proto_publicrpc_v1_publicrpc_proto_rawDescGZIP(), []int{4, 0}
}

func (x *GetLastHeartbeatsResponse_Entry) GetVerifiedGuardianAddr() string {
	if x != nil {
		return x.VerifiedGuardianAddr
	}
	return ""
}

func (x *GetLastHeartbeatsResponse_Entry) GetP2PNodeAddr() string {
	if x != nil {
		return x.P2PNodeAddr
	}
	return ""
}

func (x *GetLastHeartbeatsResponse_Entry) GetRawHeartbeat() *v1.Heartbeat {
	if x != nil {
		return x.RawHeartbeat
	}
	return nil
}

var File_proto_publicrpc_v1_publicrpc_proto protoreflect.FileDescriptor

var file_proto_publicrpc_v1_publicrpc_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x72, 0x70,
	0x63, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x72, 0x70, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x72, 0x70, 0x63, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c,
	0x01, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x3a, 0x0a, 0x0d,
	0x65, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x72, 0x70, 0x63, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x52, 0x0c, 0x65, 0x6d, 0x69, 0x74,
	0x74, 0x65, 0x72, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6d, 0x69, 0x74,
	0x74, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x65, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x4d, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x56, 0x41, 0x41, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x44, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x56, 0x41, 0x41, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x61, 0x61, 0x5f, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x76, 0x61, 0x61, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x22, 0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x83, 0x02,
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x07, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c,
	0x61, 0x73, 0x74, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x1a, 0x9c, 0x01, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x34,
	0x0a, 0x16, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x67, 0x75, 0x61, 0x72, 0x64,
	0x69, 0x61, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x47, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e,
	0x41, 0x64, 0x64, 0x72, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x32, 0x70, 0x5f, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x32, 0x70,
	0x4e, 0x6f, 0x64, 0x65, 0x41, 0x64, 0x64, 0x72, 0x12, 0x39, 0x0a, 0x0d, 0x72, 0x61, 0x77, 0x5f,
	0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x0c, 0x72, 0x61, 0x77, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x22, 0x1e, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x47, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x5d, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x47, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x67, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e,
	0x5f, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x75, 0x61, 0x72, 0x64, 0x69,
	0x61, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x0b, 0x67, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x53,
	0x65, 0x74, 0x22, 0x41, 0x0a, 0x0b, 0x47, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x53, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x2a, 0xd3, 0x03, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x43,
	0x48, 0x41, 0x49, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x53, 0x4f, 0x4c, 0x41, 0x4e, 0x41, 0x10, 0x01,
	0x12, 0x15, 0x0a, 0x11, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x45, 0x54, 0x48,
	0x45, 0x52, 0x45, 0x55, 0x4d, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x48, 0x41, 0x49, 0x4e,
	0x5f, 0x49, 0x44, 0x5f, 0x54, 0x45, 0x52, 0x52, 0x41, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x43,
	0x48, 0x41, 0x49, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x42, 0x53, 0x43, 0x10, 0x04, 0x12, 0x14, 0x0a,
	0x10, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x50, 0x4f, 0x4c, 0x59, 0x47, 0x4f,
	0x4e, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x49, 0x44, 0x5f,
	0x41, 0x56, 0x41, 0x4c, 0x41, 0x4e, 0x43, 0x48, 0x45, 0x10, 0x06, 0x12, 0x12, 0x0a, 0x0e, 0x43,
	0x48, 0x41, 0x49, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x4f, 0x41, 0x53, 0x49, 0x53, 0x10, 0x07, 0x12,
	0x15, 0x0a, 0x11, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x41, 0x4c, 0x47, 0x4f,
	0x52, 0x41, 0x4e, 0x44, 0x10, 0x08, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f,
	0x49, 0x44, 0x5f, 0x41, 0x55, 0x52, 0x4f, 0x52, 0x41, 0x10, 0x09, 0x12, 0x13, 0x0a, 0x0f, 0x43,
	0x48, 0x41, 0x49, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x46, 0x41, 0x4e, 0x54, 0x4f, 0x4d, 0x10, 0x0a,
	0x12, 0x13, 0x0a, 0x0f, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x4b, 0x41, 0x52,
	0x55, 0x52, 0x41, 0x10, 0x0b, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x49,
	0x44, 0x5f, 0x41, 0x43, 0x41, 0x4c, 0x41, 0x10, 0x0c, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x48, 0x41,
	0x49, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x4b, 0x4c, 0x41, 0x59, 0x54, 0x4e, 0x10, 0x0d, 0x12, 0x11,
	0x0a, 0x0d, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x43, 0x45, 0x4c, 0x4f, 0x10,
	0x0e, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x4e, 0x45,
	0x41, 0x52, 0x10, 0x0f, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x49, 0x44,
	0x5f, 0x4d, 0x4f, 0x4f, 0x4e, 0x42, 0x45, 0x41, 0x4d, 0x10, 0x10, 0x12, 0x11, 0x0a, 0x0d, 0x43,
	0x48, 0x41, 0x49, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x4e, 0x45, 0x4f, 0x4e, 0x10, 0x11, 0x12, 0x13,
	0x0a, 0x0f, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x54, 0x45, 0x52, 0x52, 0x41,
	0x32, 0x10, 0x12, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x49, 0x44, 0x5f,
	0x49, 0x4e, 0x4a, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x13, 0x12, 0x1e, 0x0a, 0x19, 0x43,
	0x48, 0x41, 0x49, 0x4e, 0x5f, 0x49, 0x44, 0x5f, 0x45, 0x54, 0x48, 0x45, 0x52, 0x45, 0x55, 0x4d,
	0x5f, 0x52, 0x4f, 0x50, 0x53, 0x54, 0x45, 0x4e, 0x10, 0x91, 0x4e, 0x32, 0xe2, 0x03, 0x0a, 0x10,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x50, 0x43, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x7c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x48, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x73, 0x12, 0x26, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x72, 0x70,
	0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x61, 0x73, 0x74, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e,
	0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x73, 0x12, 0xbb,
	0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x56, 0x41, 0x41, 0x12,
	0x21, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x56, 0x41, 0x41, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x72, 0x70, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x56, 0x41, 0x41, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x5e, 0x12, 0x5c,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x61, 0x2f, 0x7b,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x2e, 0x65, 0x6d, 0x69, 0x74, 0x74,
	0x65, 0x72, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x7d, 0x2f, 0x7b, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x2e, 0x65, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x2f, 0x7b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x2e, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x7d, 0x12, 0x91, 0x01, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x47, 0x75, 0x61, 0x72, 0x64,
	0x69, 0x61, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x2a, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x72,
	0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x47, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x72, 0x70, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x47, 0x75, 0x61, 0x72,
	0x64, 0x69, 0x61, 0x6e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x75, 0x61,
	0x72, 0x64, 0x69, 0x61, 0x6e, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x42, 0x2e, 0x5a, 0x2c, 0x77, 0x6f, 0x72, 0x6d, 0x68, 0x6f, 0x6c, 0x65, 0x2f, 0x6e, 0x6f, 0x64,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x72, 0x70,
	0x63, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x72, 0x70, 0x63, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_publicrpc_v1_publicrpc_proto_rawDescOnce sync.Once
	file_proto_publicrpc_v1_publicrpc_proto_rawDescData = file_proto_publicrpc_v1_publicrpc_proto_rawDesc
)

func file_proto_publicrpc_v1_publicrpc_proto_rawDescGZIP() []byte {
	file_proto_publicrpc_v1_publicrpc_proto_rawDescOnce.Do(func() {
		file_proto_publicrpc_v1_publicrpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_publicrpc_v1_publicrpc_proto_rawDescData)
	})
	return file_proto_publicrpc_v1_publicrpc_proto_rawDescData
}

var file_proto_publicrpc_v1_publicrpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_publicrpc_v1_publicrpc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_publicrpc_v1_publicrpc_proto_goTypes = []interface{}{
	(ChainID)(0),                            // 0: publicrpc.v1.ChainID
	(*MessageID)(nil),                       // 1: publicrpc.v1.MessageID
	(*GetSignedVAARequest)(nil),             // 2: publicrpc.v1.GetSignedVAARequest
	(*GetSignedVAAResponse)(nil),            // 3: publicrpc.v1.GetSignedVAAResponse
	(*GetLastHeartbeatsRequest)(nil),        // 4: publicrpc.v1.GetLastHeartbeatsRequest
	(*GetLastHeartbeatsResponse)(nil),       // 5: publicrpc.v1.GetLastHeartbeatsResponse
	(*GetCurrentGuardianSetRequest)(nil),    // 6: publicrpc.v1.GetCurrentGuardianSetRequest
	(*GetCurrentGuardianSetResponse)(nil),   // 7: publicrpc.v1.GetCurrentGuardianSetResponse
	(*GuardianSet)(nil),                     // 8: publicrpc.v1.GuardianSet
	(*GetLastHeartbeatsResponse_Entry)(nil), // 9: publicrpc.v1.GetLastHeartbeatsResponse.Entry
	(*v1.Heartbeat)(nil),                    // 10: gossip.v1.Heartbeat
}
var file_proto_publicrpc_v1_publicrpc_proto_depIdxs = []int32{
	0,  // 0: publicrpc.v1.MessageID.emitter_chain:type_name -> publicrpc.v1.ChainID
	1,  // 1: publicrpc.v1.GetSignedVAARequest.message_id:type_name -> publicrpc.v1.MessageID
	9,  // 2: publicrpc.v1.GetLastHeartbeatsResponse.entries:type_name -> publicrpc.v1.GetLastHeartbeatsResponse.Entry
	8,  // 3: publicrpc.v1.GetCurrentGuardianSetResponse.guardian_set:type_name -> publicrpc.v1.GuardianSet
	10, // 4: publicrpc.v1.GetLastHeartbeatsResponse.Entry.raw_heartbeat:type_name -> gossip.v1.Heartbeat
	4,  // 5: publicrpc.v1.PublicRPCService.GetLastHeartbeats:input_type -> publicrpc.v1.GetLastHeartbeatsRequest
	2,  // 6: publicrpc.v1.PublicRPCService.GetSignedVAA:input_type -> publicrpc.v1.GetSignedVAARequest
	6,  // 7: publicrpc.v1.PublicRPCService.GetCurrentGuardianSet:input_type -> publicrpc.v1.GetCurrentGuardianSetRequest
	5,  // 8: publicrpc.v1.PublicRPCService.GetLastHeartbeats:output_type -> publicrpc.v1.GetLastHeartbeatsResponse
	3,  // 9: publicrpc.v1.PublicRPCService.GetSignedVAA:output_type -> publicrpc.v1.GetSignedVAAResponse
	7,  // 10: publicrpc.v1.PublicRPCService.GetCurrentGuardianSet:output_type -> publicrpc.v1.GetCurrentGuardianSetResponse
	8,  // [8:11] is the sub-list for method output_type
	5,  // [5:8] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_publicrpc_v1_publicrpc_proto_init() }
func file_proto_publicrpc_v1_publicrpc_proto_init() {
	if File_proto_publicrpc_v1_publicrpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_publicrpc_v1_publicrpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageID); i {
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
		file_proto_publicrpc_v1_publicrpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSignedVAARequest); i {
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
		file_proto_publicrpc_v1_publicrpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSignedVAAResponse); i {
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
		file_proto_publicrpc_v1_publicrpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLastHeartbeatsRequest); i {
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
		file_proto_publicrpc_v1_publicrpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLastHeartbeatsResponse); i {
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
		file_proto_publicrpc_v1_publicrpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentGuardianSetRequest); i {
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
		file_proto_publicrpc_v1_publicrpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentGuardianSetResponse); i {
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
		file_proto_publicrpc_v1_publicrpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuardianSet); i {
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
		file_proto_publicrpc_v1_publicrpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLastHeartbeatsResponse_Entry); i {
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
			RawDescriptor: file_proto_publicrpc_v1_publicrpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_publicrpc_v1_publicrpc_proto_goTypes,
		DependencyIndexes: file_proto_publicrpc_v1_publicrpc_proto_depIdxs,
		EnumInfos:         file_proto_publicrpc_v1_publicrpc_proto_enumTypes,
		MessageInfos:      file_proto_publicrpc_v1_publicrpc_proto_msgTypes,
	}.Build()
	File_proto_publicrpc_v1_publicrpc_proto = out.File
	file_proto_publicrpc_v1_publicrpc_proto_rawDesc = nil
	file_proto_publicrpc_v1_publicrpc_proto_goTypes = nil
	file_proto_publicrpc_v1_publicrpc_proto_depIdxs = nil
}
