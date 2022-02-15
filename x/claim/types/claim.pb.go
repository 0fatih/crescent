// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: squad/claim/v1beta1/claim.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ActionType defines the type of action that a recipient must execute in order to receive a claimable amount.
type ActionType int32

const (
	// ACTION_TYPE_UNSPECIFIED specifies an unknown action type
	ActionTypeUnspecified ActionType = 0
	// ACTION_TYPE_DEPOSIT specifies deposit action type
	ActionTypeDeposit ActionType = 1
	// ACTION_TYPE_SWAP specifies swap action type
	ActionTypeSwap ActionType = 2
	// ACTION_TYPE_FARMING specifies farming (stake) action type
	ActionTypeFarming ActionType = 3
)

var ActionType_name = map[int32]string{
	0: "ACTION_TYPE_UNSPECIFIED",
	1: "ACTION_TYPE_DEPOSIT",
	2: "ACTION_TYPE_SWAP",
	3: "ACTION_TYPE_FARMING",
}

var ActionType_value = map[string]int32{
	"ACTION_TYPE_UNSPECIFIED": 0,
	"ACTION_TYPE_DEPOSIT":     1,
	"ACTION_TYPE_SWAP":        2,
	"ACTION_TYPE_FARMING":     3,
}

func (x ActionType) String() string {
	return proto.EnumName(ActionType_name, int32(x))
}

func (ActionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_84886eaa62c7639a, []int{0}
}

// Airdrop defines airdrop information.
type Airdrop struct {
	// airdrop_id specifies index of the airdrop
	AirdropId uint64 `protobuf:"varint,1,opt,name=airdrop_id,json=airdropId,proto3" json:"airdrop_id,omitempty"`
	// source_address defines the bech32-encoded source address
	// where the source of coins from
	SourceAddress string `protobuf:"bytes,2,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	// source_coins specifies the airdrop coins
	SourceCoins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=source_coins,json=sourceCoins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"source_coins"`
	// termination_address defines the bech32-encoded termination address
	// where the remaining source coins are sent to
	TerminationAddress string `protobuf:"bytes,4,opt,name=termination_address,json=terminationAddress,proto3" json:"termination_address,omitempty"`
	// start_time specifies the start time of the airdrop
	StartTime time.Time `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time"`
	// end_time specifies the start time of the airdrop
	EndTime time.Time `protobuf:"bytes,6,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time"`
}

func (m *Airdrop) Reset()         { *m = Airdrop{} }
func (m *Airdrop) String() string { return proto.CompactTextString(m) }
func (*Airdrop) ProtoMessage()    {}
func (*Airdrop) Descriptor() ([]byte, []int) {
	return fileDescriptor_84886eaa62c7639a, []int{0}
}
func (m *Airdrop) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Airdrop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Airdrop.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Airdrop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Airdrop.Merge(m, src)
}
func (m *Airdrop) XXX_Size() int {
	return m.Size()
}
func (m *Airdrop) XXX_DiscardUnknown() {
	xxx_messageInfo_Airdrop.DiscardUnknown(m)
}

var xxx_messageInfo_Airdrop proto.InternalMessageInfo

// ClaimRecord defines claim record that corresponds to the airdrop.
type ClaimRecord struct {
	// airdrop_id specifies airdrop id
	AirdropId uint64 `protobuf:"varint,1,opt,name=airdrop_id,json=airdropId,proto3" json:"airdrop_id,omitempty"`
	// recipient specifies the bech32-encoded address that is eligible to claim airdrop
	Recipient string `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// initial_claimable_coins specifies the initial claimable coins
	InitialClaimableCoins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=initial_claimable_coins,json=initialClaimableCoins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"initial_claimable_coins"`
	// claimable_coins specifies the unclaimed claimable coins
	ClaimableCoins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=claimable_coins,json=claimableCoins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"claimable_coins"`
	// actions specifies a list of actions
	Actions []Action `protobuf:"bytes,5,rep,name=actions,proto3" json:"actions"`
}

func (m *ClaimRecord) Reset()         { *m = ClaimRecord{} }
func (m *ClaimRecord) String() string { return proto.CompactTextString(m) }
func (*ClaimRecord) ProtoMessage()    {}
func (*ClaimRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_84886eaa62c7639a, []int{1}
}
func (m *ClaimRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimRecord.Merge(m, src)
}
func (m *ClaimRecord) XXX_Size() int {
	return m.Size()
}
func (m *ClaimRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimRecord proto.InternalMessageInfo

// Action defines an action type and its claimed status.
type Action struct {
	// action_type specifies the action type
	ActionType ActionType `protobuf:"varint,1,opt,name=action_type,json=actionType,proto3,enum=squad.claim.v1beta1.ActionType" json:"action_type,omitempty"`
	// claimed specifies the status of an action
	Claimed bool `protobuf:"varint,2,opt,name=claimed,proto3" json:"claimed,omitempty"`
}

func (m *Action) Reset()         { *m = Action{} }
func (m *Action) String() string { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()    {}
func (*Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_84886eaa62c7639a, []int{2}
}
func (m *Action) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Action.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Action.Merge(m, src)
}
func (m *Action) XXX_Size() int {
	return m.Size()
}
func (m *Action) XXX_DiscardUnknown() {
	xxx_messageInfo_Action.DiscardUnknown(m)
}

var xxx_messageInfo_Action proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("squad.claim.v1beta1.ActionType", ActionType_name, ActionType_value)
	proto.RegisterType((*Airdrop)(nil), "squad.claim.v1beta1.Airdrop")
	proto.RegisterType((*ClaimRecord)(nil), "squad.claim.v1beta1.ClaimRecord")
	proto.RegisterType((*Action)(nil), "squad.claim.v1beta1.Action")
}

func init() { proto.RegisterFile("squad/claim/v1beta1/claim.proto", fileDescriptor_84886eaa62c7639a) }

var fileDescriptor_84886eaa62c7639a = []byte{
	// 657 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xb6, 0x9b, 0xb4, 0x69, 0x2e, 0x10, 0xc2, 0x95, 0xaa, 0xa9, 0x01, 0xc7, 0xaa, 0x84, 0x14,
	0x21, 0xd5, 0xa6, 0x45, 0x62, 0x61, 0x80, 0x24, 0x4d, 0x51, 0x24, 0x68, 0x23, 0x27, 0x15, 0x82,
	0xc5, 0x3a, 0xdb, 0xd7, 0x70, 0x22, 0xf1, 0x19, 0xdf, 0x05, 0xe8, 0xdc, 0x05, 0x75, 0xea, 0x1f,
	0xe8, 0xc4, 0xc6, 0x2f, 0xe9, 0xd8, 0x81, 0x81, 0x89, 0x42, 0xfb, 0x1f, 0x98, 0xd1, 0xdd, 0xd9,
	0x4d, 0x29, 0x08, 0x84, 0x04, 0x53, 0xee, 0xbd, 0xf7, 0xbd, 0xf7, 0xbd, 0x7c, 0x9f, 0xef, 0x40,
	0x8d, 0xbd, 0x1a, 0xa3, 0xd0, 0x09, 0x86, 0x88, 0x8c, 0x9c, 0xd7, 0x2b, 0x3e, 0xe6, 0x68, 0x45,
	0x45, 0x76, 0x9c, 0x50, 0x4e, 0xe1, 0x9c, 0x04, 0xd8, 0x2a, 0x95, 0x02, 0x8c, 0x6b, 0x03, 0x3a,
	0xa0, 0xb2, 0xee, 0x88, 0x93, 0x82, 0x1a, 0xb5, 0x01, 0xa5, 0x83, 0x21, 0x76, 0x64, 0xe4, 0x8f,
	0xb7, 0x1d, 0x4e, 0x46, 0x98, 0x71, 0x34, 0x8a, 0x53, 0x80, 0x19, 0x50, 0x36, 0xa2, 0xcc, 0xf1,
	0x11, 0xc3, 0x13, 0x32, 0x4a, 0x22, 0x55, 0x5f, 0xfa, 0x36, 0x05, 0x0a, 0x0d, 0x92, 0x84, 0x09,
	0x8d, 0xe1, 0x4d, 0x00, 0x90, 0x3a, 0x7a, 0x24, 0xac, 0xea, 0x96, 0x5e, 0xcf, 0xbb, 0xc5, 0x34,
	0xd3, 0x09, 0xe1, 0x2d, 0x50, 0x66, 0x74, 0x9c, 0x04, 0xd8, 0x43, 0x61, 0x98, 0x60, 0xc6, 0xaa,
	0x53, 0x96, 0x5e, 0x2f, 0xba, 0x97, 0x55, 0xb6, 0xa1, 0x92, 0x30, 0x02, 0x97, 0x52, 0x98, 0xa0,
	0x61, 0xd5, 0x9c, 0x95, 0xab, 0x97, 0x56, 0x17, 0x6d, 0xb5, 0x88, 0x2d, 0x16, 0xc9, 0xfe, 0x94,
	0xdd, 0xa2, 0x24, 0x6a, 0xde, 0x39, 0xfc, 0x5c, 0xd3, 0x3e, 0x1c, 0xd7, 0xea, 0x03, 0xc2, 0x5f,
	0x8c, 0x7d, 0x3b, 0xa0, 0x23, 0x27, 0xdd, 0x5a, 0xfd, 0x2c, 0xb3, 0xf0, 0xa5, 0xc3, 0x77, 0x62,
	0xcc, 0x64, 0x03, 0x73, 0x4b, 0x8a, 0x40, 0x06, 0xd0, 0x01, 0x73, 0x1c, 0x27, 0x23, 0x12, 0x21,
	0x4e, 0x68, 0x74, 0xb6, 0x5b, 0x5e, 0xee, 0x06, 0xcf, 0x95, 0xb2, 0x05, 0x5b, 0x00, 0x30, 0x8e,
	0x12, 0xee, 0x09, 0xad, 0xaa, 0xd3, 0x96, 0x5e, 0x2f, 0xad, 0x1a, 0xb6, 0x12, 0xd2, 0xce, 0x84,
	0xb4, 0xfb, 0x99, 0x90, 0xcd, 0x59, 0xb1, 0xdf, 0xfe, 0x71, 0x4d, 0x77, 0x8b, 0xb2, 0x4f, 0x54,
	0xe0, 0x03, 0x30, 0x8b, 0xa3, 0x50, 0x8d, 0x98, 0xf9, 0x8b, 0x11, 0x05, 0x1c, 0x85, 0x22, 0xbf,
	0xb4, 0x9b, 0x03, 0xa5, 0x96, 0x70, 0xd8, 0xc5, 0x01, 0x4d, 0xc2, 0x3f, 0x89, 0x7f, 0x03, 0x14,
	0x13, 0x1c, 0x90, 0x98, 0xe0, 0x88, 0xa7, 0xba, 0x4f, 0x12, 0x70, 0x57, 0x07, 0x0b, 0x24, 0x22,
	0x9c, 0xa0, 0xa1, 0x27, 0x3f, 0x1b, 0xe4, 0x0f, 0xff, 0xa3, 0xfe, 0xf3, 0x29, 0x57, 0x2b, 0xa3,
	0x52, 0x4e, 0x70, 0x70, 0xe5, 0x22, 0x79, 0xfe, 0xdf, 0x93, 0x97, 0x83, 0x1f, 0x59, 0xef, 0x83,
	0x02, 0x0a, 0x84, 0xbf, 0xac, 0x3a, 0x2d, 0xd9, 0xae, 0xdb, 0xbf, 0xb8, 0x3f, 0x76, 0x43, 0x62,
	0x9a, 0x79, 0xc1, 0xe7, 0x66, 0x1d, 0x4b, 0x21, 0x98, 0x51, 0x05, 0xf8, 0x10, 0x94, 0x54, 0xd2,
	0x13, 0x64, 0xd2, 0x80, 0xf2, 0x6a, 0xed, 0x37, 0xa3, 0xfa, 0x3b, 0x31, 0x76, 0x01, 0x3a, 0x3b,
	0xc3, 0x2a, 0x28, 0x48, 0x1c, 0x0e, 0xa5, 0x41, 0xb3, 0x6e, 0x16, 0xde, 0xfe, 0xa8, 0x03, 0x30,
	0x69, 0x82, 0xf7, 0xc0, 0x42, 0xa3, 0xd5, 0xef, 0x6c, 0x6e, 0x78, 0xfd, 0x67, 0xdd, 0xb6, 0xb7,
	0xb5, 0xd1, 0xeb, 0xb6, 0x5b, 0x9d, 0xf5, 0x4e, 0x7b, 0xad, 0xa2, 0x19, 0x8b, 0x7b, 0x07, 0xd6,
	0xfc, 0x04, 0xbc, 0x15, 0xb1, 0x18, 0x07, 0x64, 0x9b, 0xe0, 0x10, 0xda, 0x60, 0xee, 0x7c, 0xdf,
	0x5a, 0xbb, 0xbb, 0xd9, 0xeb, 0xf4, 0x2b, 0xba, 0x31, 0xbf, 0x77, 0x60, 0x5d, 0x9d, 0xf4, 0xac,
	0xe1, 0x98, 0x32, 0xc2, 0x61, 0x1d, 0x54, 0xce, 0xe3, 0x7b, 0x4f, 0x1b, 0xdd, 0xca, 0x94, 0x01,
	0xf7, 0x0e, 0xac, 0xf2, 0x04, 0xdc, 0x7b, 0x83, 0xe2, 0x8b, 0x93, 0xd7, 0x1b, 0xee, 0x93, 0xce,
	0xc6, 0xa3, 0x4a, 0xee, 0xe2, 0xe4, 0x75, 0x24, 0x6e, 0xd3, 0xc0, 0xc8, 0xbf, 0x7b, 0x6f, 0x6a,
	0xcd, 0xc7, 0x87, 0x5f, 0x4d, 0xed, 0xf0, 0xc4, 0xd4, 0x8f, 0x4e, 0x4c, 0xfd, 0xcb, 0x89, 0xa9,
	0xef, 0x9f, 0x9a, 0xda, 0xd1, 0xa9, 0xa9, 0x7d, 0x3a, 0x35, 0xb5, 0xe7, 0xf6, 0x4f, 0x8e, 0x0a,
	0x29, 0x97, 0x87, 0xc8, 0x67, 0x8e, 0x7a, 0x01, 0xdf, 0xa6, 0x6f, 0xa0, 0x74, 0xd7, 0x9f, 0x91,
	0xf7, 0xe6, 0xee, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x34, 0xce, 0x13, 0x1f, 0x05, 0x00,
	0x00,
}

func (m *Airdrop) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Airdrop) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Airdrop) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintClaim(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x32
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintClaim(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x2a
	if len(m.TerminationAddress) > 0 {
		i -= len(m.TerminationAddress)
		copy(dAtA[i:], m.TerminationAddress)
		i = encodeVarintClaim(dAtA, i, uint64(len(m.TerminationAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SourceCoins) > 0 {
		for iNdEx := len(m.SourceCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SourceCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintClaim(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.SourceAddress) > 0 {
		i -= len(m.SourceAddress)
		copy(dAtA[i:], m.SourceAddress)
		i = encodeVarintClaim(dAtA, i, uint64(len(m.SourceAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.AirdropId != 0 {
		i = encodeVarintClaim(dAtA, i, uint64(m.AirdropId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ClaimRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Actions) > 0 {
		for iNdEx := len(m.Actions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Actions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintClaim(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ClaimableCoins) > 0 {
		for iNdEx := len(m.ClaimableCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClaimableCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintClaim(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.InitialClaimableCoins) > 0 {
		for iNdEx := len(m.InitialClaimableCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitialClaimableCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintClaim(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintClaim(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x12
	}
	if m.AirdropId != 0 {
		i = encodeVarintClaim(dAtA, i, uint64(m.AirdropId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Action) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Action) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Action) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Claimed {
		i--
		if m.Claimed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.ActionType != 0 {
		i = encodeVarintClaim(dAtA, i, uint64(m.ActionType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintClaim(dAtA []byte, offset int, v uint64) int {
	offset -= sovClaim(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Airdrop) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AirdropId != 0 {
		n += 1 + sovClaim(uint64(m.AirdropId))
	}
	l = len(m.SourceAddress)
	if l > 0 {
		n += 1 + l + sovClaim(uint64(l))
	}
	if len(m.SourceCoins) > 0 {
		for _, e := range m.SourceCoins {
			l = e.Size()
			n += 1 + l + sovClaim(uint64(l))
		}
	}
	l = len(m.TerminationAddress)
	if l > 0 {
		n += 1 + l + sovClaim(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovClaim(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovClaim(uint64(l))
	return n
}

func (m *ClaimRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AirdropId != 0 {
		n += 1 + sovClaim(uint64(m.AirdropId))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovClaim(uint64(l))
	}
	if len(m.InitialClaimableCoins) > 0 {
		for _, e := range m.InitialClaimableCoins {
			l = e.Size()
			n += 1 + l + sovClaim(uint64(l))
		}
	}
	if len(m.ClaimableCoins) > 0 {
		for _, e := range m.ClaimableCoins {
			l = e.Size()
			n += 1 + l + sovClaim(uint64(l))
		}
	}
	if len(m.Actions) > 0 {
		for _, e := range m.Actions {
			l = e.Size()
			n += 1 + l + sovClaim(uint64(l))
		}
	}
	return n
}

func (m *Action) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ActionType != 0 {
		n += 1 + sovClaim(uint64(m.ActionType))
	}
	if m.Claimed {
		n += 2
	}
	return n
}

func sovClaim(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaim(x uint64) (n int) {
	return sovClaim(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Airdrop) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Airdrop: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Airdrop: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropId", wireType)
			}
			m.AirdropId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AirdropId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceCoins = append(m.SourceCoins, types.Coin{})
			if err := m.SourceCoins[len(m.SourceCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TerminationAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TerminationAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClaim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClaimRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClaimRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropId", wireType)
			}
			m.AirdropId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AirdropId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialClaimableCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialClaimableCoins = append(m.InitialClaimableCoins, types.Coin{})
			if err := m.InitialClaimableCoins[len(m.InitialClaimableCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimableCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimableCoins = append(m.ClaimableCoins, types.Coin{})
			if err := m.ClaimableCoins[len(m.ClaimableCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Actions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Actions = append(m.Actions, Action{})
			if err := m.Actions[len(m.Actions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClaim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Action) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Action: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Action: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionType", wireType)
			}
			m.ActionType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ActionType |= ActionType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claimed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Claimed = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipClaim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipClaim(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClaim
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthClaim
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClaim
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClaim
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClaim        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClaim          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClaim = fmt.Errorf("proto: unexpected end of group")
)
