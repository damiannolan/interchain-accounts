// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: testing/msgs.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgTryRegisterIBCAccount struct {
	// the port on which the packet will be sent
	SourcePort string `protobuf:"bytes,1,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty" yaml:"source_port"`
	// the channel by which the packet will be sent
	SourceChannel string `protobuf:"bytes,2,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty" yaml:"source_channel"`
	// the salt that will be transfered to counterparty chain.
	Salt []byte `protobuf:"bytes,3,opt,name=salt,proto3" json:"salt,omitempty" yaml:"source_channel"`
	// Timeout height relative to the current block height.
	// The timeout is disabled when set to 0.
	TimeoutHeight types.Height `protobuf:"bytes,4,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height" yaml:"timeout_height"`
	// Timeout timestamp (in nanoseconds) relative to the current block timestamp.
	// The timeout is disabled when set to 0.
	TimeoutTimestamp uint64                                        `protobuf:"varint,5,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty" yaml:"timeout_timestamp"`
	Sender           github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,6,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
}

func (m *MsgTryRegisterIBCAccount) Reset()         { *m = MsgTryRegisterIBCAccount{} }
func (m *MsgTryRegisterIBCAccount) String() string { return proto.CompactTextString(m) }
func (*MsgTryRegisterIBCAccount) ProtoMessage()    {}
func (*MsgTryRegisterIBCAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_a97478b5c8c3f49a, []int{0}
}
func (m *MsgTryRegisterIBCAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTryRegisterIBCAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTryRegisterIBCAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTryRegisterIBCAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTryRegisterIBCAccount.Merge(m, src)
}
func (m *MsgTryRegisterIBCAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgTryRegisterIBCAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTryRegisterIBCAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTryRegisterIBCAccount proto.InternalMessageInfo

func (m *MsgTryRegisterIBCAccount) GetSourcePort() string {
	if m != nil {
		return m.SourcePort
	}
	return ""
}

func (m *MsgTryRegisterIBCAccount) GetSourceChannel() string {
	if m != nil {
		return m.SourceChannel
	}
	return ""
}

func (m *MsgTryRegisterIBCAccount) GetSalt() []byte {
	if m != nil {
		return m.Salt
	}
	return nil
}

func (m *MsgTryRegisterIBCAccount) GetTimeoutHeight() types.Height {
	if m != nil {
		return m.TimeoutHeight
	}
	return types.Height{}
}

func (m *MsgTryRegisterIBCAccount) GetTimeoutTimestamp() uint64 {
	if m != nil {
		return m.TimeoutTimestamp
	}
	return 0
}

func (m *MsgTryRegisterIBCAccount) GetSender() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Sender
	}
	return nil
}

type MsgTryRunTxMsgSend struct {
	// the port on which the packet will be sent
	SourcePort string `protobuf:"bytes,1,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty" yaml:"source_port"`
	// the channel by which the packet will be sent
	SourceChannel string `protobuf:"bytes,2,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty" yaml:"source_channel"`
	// Timeout height relative to the current block height.
	// The timeout is disabled when set to 0.
	TimeoutHeight types.Height `protobuf:"bytes,4,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height" yaml:"timeout_height"`
	// Timeout timestamp (in nanoseconds) relative to the current block timestamp.
	// The timeout is disabled when set to 0.
	TimeoutTimestamp uint64                                        `protobuf:"varint,5,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty" yaml:"timeout_timestamp"`
	FromAddress      github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,6,opt,name=from_address,json=fromAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"from_address,omitempty" yaml:"from_address"`
	ToAddress        github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,7,opt,name=to_address,json=toAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"to_address,omitempty" yaml:"to_address"`
	Amount           github_com_cosmos_cosmos_sdk_types.Coins      `protobuf:"bytes,8,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
	Sender           github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,9,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
}

func (m *MsgTryRunTxMsgSend) Reset()         { *m = MsgTryRunTxMsgSend{} }
func (m *MsgTryRunTxMsgSend) String() string { return proto.CompactTextString(m) }
func (*MsgTryRunTxMsgSend) ProtoMessage()    {}
func (*MsgTryRunTxMsgSend) Descriptor() ([]byte, []int) {
	return fileDescriptor_a97478b5c8c3f49a, []int{1}
}
func (m *MsgTryRunTxMsgSend) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTryRunTxMsgSend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTryRunTxMsgSend.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTryRunTxMsgSend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTryRunTxMsgSend.Merge(m, src)
}
func (m *MsgTryRunTxMsgSend) XXX_Size() int {
	return m.Size()
}
func (m *MsgTryRunTxMsgSend) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTryRunTxMsgSend.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTryRunTxMsgSend proto.InternalMessageInfo

func (m *MsgTryRunTxMsgSend) GetSourcePort() string {
	if m != nil {
		return m.SourcePort
	}
	return ""
}

func (m *MsgTryRunTxMsgSend) GetSourceChannel() string {
	if m != nil {
		return m.SourceChannel
	}
	return ""
}

func (m *MsgTryRunTxMsgSend) GetTimeoutHeight() types.Height {
	if m != nil {
		return m.TimeoutHeight
	}
	return types.Height{}
}

func (m *MsgTryRunTxMsgSend) GetTimeoutTimestamp() uint64 {
	if m != nil {
		return m.TimeoutTimestamp
	}
	return 0
}

func (m *MsgTryRunTxMsgSend) GetFromAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.FromAddress
	}
	return nil
}

func (m *MsgTryRunTxMsgSend) GetToAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.ToAddress
	}
	return nil
}

func (m *MsgTryRunTxMsgSend) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *MsgTryRunTxMsgSend) GetSender() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Sender
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgTryRegisterIBCAccount)(nil), "testing.MsgTryRegisterIBCAccount")
	proto.RegisterType((*MsgTryRunTxMsgSend)(nil), "testing.MsgTryRunTxMsgSend")
}

func init() { proto.RegisterFile("testing/msgs.proto", fileDescriptor_a97478b5c8c3f49a) }

var fileDescriptor_a97478b5c8c3f49a = []byte{
	// 570 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x54, 0x3f, 0x6f, 0xd3, 0x40,
	0x14, 0x8f, 0x69, 0x9b, 0xd2, 0x4b, 0x5b, 0x51, 0xf3, 0x47, 0x6e, 0x04, 0x76, 0xe4, 0xc9, 0x8b,
	0x7d, 0xa4, 0x0c, 0x48, 0x4c, 0xc4, 0x91, 0x50, 0x33, 0x44, 0x42, 0x26, 0x13, 0x4b, 0xb0, 0x2f,
	0x87, 0x73, 0x4a, 0x7c, 0x17, 0xdd, 0x5d, 0xa2, 0xe4, 0x5b, 0xf0, 0x39, 0xf8, 0x0c, 0x0c, 0x8c,
	0x1d, 0x3b, 0x32, 0x19, 0x94, 0x7c, 0x83, 0x8c, 0x4c, 0xc8, 0xbe, 0x4b, 0x49, 0x84, 0x84, 0x40,
	0x1d, 0x10, 0x4b, 0xee, 0xde, 0x7b, 0xbf, 0xdf, 0xfb, 0x25, 0xef, 0x77, 0x79, 0xc0, 0x94, 0x58,
	0x48, 0x42, 0x53, 0x98, 0x89, 0x54, 0x04, 0x13, 0xce, 0x24, 0x33, 0x0f, 0x75, 0xae, 0xfe, 0x20,
	0x65, 0x29, 0x2b, 0x73, 0xb0, 0xb8, 0xa9, 0x72, 0xdd, 0x21, 0x09, 0x82, 0x88, 0x71, 0x0c, 0xd1,
	0x98, 0x60, 0x2a, 0xe1, 0xac, 0xa9, 0x6f, 0x1a, 0x60, 0x23, 0x26, 0x32, 0x26, 0x60, 0x12, 0x0b,
	0x0c, 0x67, 0xcd, 0x04, 0xcb, 0xb8, 0x09, 0x11, 0x23, 0x54, 0xd5, 0xdd, 0x4f, 0x7b, 0xc0, 0xea,
	0x8a, 0xb4, 0xc7, 0x17, 0x11, 0x4e, 0x89, 0x90, 0x98, 0x77, 0xc2, 0x76, 0x0b, 0x21, 0x36, 0xa5,
	0xd2, 0x7c, 0x0e, 0x6a, 0x82, 0x4d, 0x39, 0xc2, 0xfd, 0x09, 0xe3, 0xd2, 0x32, 0x1a, 0x86, 0x77,
	0x14, 0x3e, 0x5a, 0xe7, 0x8e, 0xb9, 0x88, 0xb3, 0xf1, 0x0b, 0x77, 0xab, 0xe8, 0x46, 0x40, 0x45,
	0xaf, 0x19, 0x97, 0xe6, 0x4b, 0x70, 0xaa, 0x6b, 0x68, 0x18, 0x53, 0x8a, 0xc7, 0xd6, 0x9d, 0x92,
	0x7b, 0xbe, 0xce, 0x9d, 0x87, 0x3b, 0x5c, 0x5d, 0x77, 0xa3, 0x13, 0x95, 0x68, 0xab, 0xd8, 0xf4,
	0xc1, 0xbe, 0x88, 0xc7, 0xd2, 0xda, 0x6b, 0x18, 0xde, 0xf1, 0xef, 0x78, 0x25, 0xcc, 0x7c, 0x07,
	0x4e, 0x25, 0xc9, 0x30, 0x9b, 0xca, 0xfe, 0x10, 0x93, 0x74, 0x28, 0xad, 0xfd, 0x86, 0xe1, 0xd5,
	0x2e, 0xea, 0x01, 0x49, 0x50, 0x50, 0x0c, 0x28, 0xd0, 0x63, 0x99, 0x35, 0x83, 0xcb, 0x12, 0x11,
	0x3e, 0xb9, 0xca, 0x9d, 0xca, 0xcf, 0xc6, 0xbb, 0x7c, 0x37, 0x3a, 0xd1, 0x09, 0x85, 0x36, 0x3b,
	0xe0, 0x6c, 0x83, 0x28, 0x4e, 0x21, 0xe3, 0x6c, 0x62, 0x1d, 0x34, 0x0c, 0x6f, 0x3f, 0x7c, 0xbc,
	0xce, 0x1d, 0x6b, 0xb7, 0xc9, 0x0d, 0xc4, 0x8d, 0xee, 0xe9, 0x5c, 0x6f, 0x93, 0x32, 0x3b, 0xa0,
	0x2a, 0x30, 0x1d, 0x60, 0x6e, 0x55, 0xcb, 0x5f, 0xd7, 0xfc, 0x9e, 0x3b, 0x7e, 0x4a, 0xe4, 0x70,
	0x9a, 0x04, 0x88, 0x65, 0x50, 0x5b, 0xa6, 0x0e, 0x5f, 0x0c, 0x46, 0x50, 0x2e, 0x26, 0x58, 0x04,
	0x2d, 0x84, 0x5a, 0x83, 0x01, 0xc7, 0x42, 0x44, 0xba, 0x81, 0xfb, 0xf9, 0x00, 0x98, 0xda, 0xbe,
	0x29, 0xed, 0xcd, 0xbb, 0x22, 0x7d, 0x83, 0xe9, 0xe0, 0x5f, 0x1a, 0xf7, 0x5f, 0x39, 0x31, 0x02,
	0xc7, 0xef, 0x39, 0xcb, 0xfa, 0xb1, 0x1a, 0xab, 0xf6, 0xe3, 0x72, 0x9d, 0x3b, 0xf7, 0x55, 0x97,
	0xed, 0xaa, 0xfb, 0xf7, 0x36, 0xd5, 0x0a, 0xbe, 0x0e, 0x4c, 0x0c, 0x80, 0x64, 0x37, 0x52, 0x87,
	0xa5, 0xd4, 0xab, 0x75, 0xee, 0x9c, 0xe9, 0x2f, 0xcc, 0x6e, 0x21, 0x74, 0x24, 0xd9, 0x46, 0x06,
	0x81, 0x6a, 0x9c, 0x15, 0x7f, 0x5f, 0xeb, 0x6e, 0x63, 0xcf, 0xab, 0x5d, 0x9c, 0x07, 0x8a, 0x18,
	0x14, 0x2b, 0x20, 0xd0, 0x2b, 0x20, 0x68, 0x33, 0x42, 0xc3, 0xa7, 0xc5, 0xdc, 0x3f, 0x7e, 0x75,
	0xbc, 0x3f, 0x10, 0x2b, 0x08, 0x22, 0xd2, 0xad, 0xb7, 0x9e, 0xf0, 0xd1, 0x2d, 0x9f, 0x70, 0x98,
	0x5e, 0x2d, 0x6d, 0xe3, 0x7a, 0x69, 0x1b, 0xdf, 0x96, 0xb6, 0xf1, 0x61, 0x65, 0x57, 0xae, 0x57,
	0x76, 0xe5, 0xcb, 0xca, 0xae, 0xbc, 0xed, 0xfe, 0xda, 0x90, 0x50, 0x89, 0x39, 0x1a, 0xc6, 0x84,
	0xfa, 0xb1, 0x5a, 0x52, 0x02, 0xce, 0x21, 0x49, 0xd0, 0x26, 0x84, 0x19, 0x43, 0x23, 0x51, 0x7e,
	0xfa, 0x25, 0xd8, 0x97, 0x73, 0xa5, 0x9d, 0x54, 0xcb, 0x8d, 0xf7, 0xec, 0x47, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x2e, 0x1c, 0x6e, 0xa4, 0x67, 0x05, 0x00, 0x00,
}

func (m *MsgTryRegisterIBCAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTryRegisterIBCAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTryRegisterIBCAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x32
	}
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintMsgs(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.TimeoutHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgs(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Salt) > 0 {
		i -= len(m.Salt)
		copy(dAtA[i:], m.Salt)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.Salt)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SourcePort) > 0 {
		i -= len(m.SourcePort)
		copy(dAtA[i:], m.SourcePort)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.SourcePort)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgTryRunTxMsgSend) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTryRunTxMsgSend) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTryRunTxMsgSend) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMsgs(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.ToAddress) > 0 {
		i -= len(m.ToAddress)
		copy(dAtA[i:], m.ToAddress)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.ToAddress)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0x32
	}
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintMsgs(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.TimeoutHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgs(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SourcePort) > 0 {
		i -= len(m.SourcePort)
		copy(dAtA[i:], m.SourcePort)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.SourcePort)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMsgs(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgTryRegisterIBCAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SourcePort)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	l = len(m.Salt)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	l = m.TimeoutHeight.Size()
	n += 1 + l + sovMsgs(uint64(l))
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovMsgs(uint64(m.TimeoutTimestamp))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	return n
}

func (m *MsgTryRunTxMsgSend) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SourcePort)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	l = m.TimeoutHeight.Size()
	n += 1 + l + sovMsgs(uint64(l))
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovMsgs(uint64(m.TimeoutTimestamp))
	}
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	l = len(m.ToAddress)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovMsgs(uint64(l))
		}
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	return n
}

func sovMsgs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgs(x uint64) (n int) {
	return sovMsgs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgTryRegisterIBCAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgs
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
			return fmt.Errorf("proto: MsgTryRegisterIBCAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTryRegisterIBCAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourcePort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourcePort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Salt", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Salt = append(m.Salt[:0], dAtA[iNdEx:postIndex]...)
			if m.Salt == nil {
				m.Salt = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TimeoutHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgs
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
func (m *MsgTryRunTxMsgSend) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgs
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
			return fmt.Errorf("proto: MsgTryRunTxMsgSend: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTryRunTxMsgSend: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourcePort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourcePort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TimeoutHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = append(m.FromAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.FromAddress == nil {
				m.FromAddress = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToAddress = append(m.ToAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.ToAddress == nil {
				m.ToAddress = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types1.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgs
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
func skipMsgs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgs
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
					return 0, ErrIntOverflowMsgs
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
					return 0, ErrIntOverflowMsgs
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
				return 0, ErrInvalidLengthMsgs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgs = fmt.Errorf("proto: unexpected end of group")
)
