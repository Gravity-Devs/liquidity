// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint/liquidity/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
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

// records the state of each pool after genesis export or import, used to check variables
type PoolRecord struct {
	Pool              Pool               `protobuf:"bytes,1,opt,name=pool,proto3" json:"pool" yaml:"pool"`
	PoolMetadata      PoolMetadata       `protobuf:"bytes,2,opt,name=pool_metadata,json=poolMetadata,proto3" json:"pool_metadata" yaml:"pool_metadata"`
	PoolBatch         PoolBatch          `protobuf:"bytes,3,opt,name=pool_batch,json=poolBatch,proto3" json:"pool_batch" yaml:"pool_batch"`
	DepositMsgStates  []DepositMsgState  `protobuf:"bytes,4,rep,name=deposit_msg_states,json=depositMsgStates,proto3" json:"deposit_msg_states" yaml:"deposit_msg_states"`
	WithdrawMsgStates []WithdrawMsgState `protobuf:"bytes,5,rep,name=withdraw_msg_states,json=withdrawMsgStates,proto3" json:"withdraw_msg_states" yaml:"withdraw_msg_states"`
	SwapMsgStates     []SwapMsgState     `protobuf:"bytes,6,rep,name=swap_msg_states,json=swapMsgStates,proto3" json:"swap_msg_states" yaml:"swap_msg_states"`
}

func (m *PoolRecord) Reset()         { *m = PoolRecord{} }
func (m *PoolRecord) String() string { return proto.CompactTextString(m) }
func (*PoolRecord) ProtoMessage()    {}
func (*PoolRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dc104913a173687, []int{0}
}
func (m *PoolRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolRecord.Merge(m, src)
}
func (m *PoolRecord) XXX_Size() int {
	return m.Size()
}
func (m *PoolRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolRecord.DiscardUnknown(m)
}

var xxx_messageInfo_PoolRecord proto.InternalMessageInfo

func (m *PoolRecord) GetPool() Pool {
	if m != nil {
		return m.Pool
	}
	return Pool{}
}

func (m *PoolRecord) GetPoolMetadata() PoolMetadata {
	if m != nil {
		return m.PoolMetadata
	}
	return PoolMetadata{}
}

func (m *PoolRecord) GetPoolBatch() PoolBatch {
	if m != nil {
		return m.PoolBatch
	}
	return PoolBatch{}
}

func (m *PoolRecord) GetDepositMsgStates() []DepositMsgState {
	if m != nil {
		return m.DepositMsgStates
	}
	return nil
}

func (m *PoolRecord) GetWithdrawMsgStates() []WithdrawMsgState {
	if m != nil {
		return m.WithdrawMsgStates
	}
	return nil
}

func (m *PoolRecord) GetSwapMsgStates() []SwapMsgState {
	if m != nil {
		return m.SwapMsgStates
	}
	return nil
}

// GenesisState defines the liquidity module's genesis state.
type GenesisState struct {
	// params defines all the parameters for the liquidity module.
	Params      Params       `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PoolRecords []PoolRecord `protobuf:"bytes,2,rep,name=pool_records,json=poolRecords,proto3" json:"pool_records" yaml:"pools"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dc104913a173687, []int{1}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PoolRecord)(nil), "tendermint.liquidity.v1beta1.PoolRecord")
	proto.RegisterType((*GenesisState)(nil), "tendermint.liquidity.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("tendermint/liquidity/v1beta1/genesis.proto", fileDescriptor_7dc104913a173687)
}

var fileDescriptor_7dc104913a173687 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0xed, 0x36, 0x8d, 0xe0, 0x92, 0x0a, 0x7a, 0x8d, 0x90, 0x1b, 0x55, 0x4e, 0x39, 0x21,
	0x11, 0x55, 0xd4, 0x56, 0xdb, 0xad, 0xa3, 0x85, 0xc4, 0x80, 0x22, 0x21, 0x77, 0x40, 0x62, 0x89,
	0x2e, 0xf1, 0xc9, 0x39, 0x29, 0xce, 0x1d, 0x7e, 0xd7, 0x98, 0x2c, 0x0c, 0x4c, 0x8c, 0x7c, 0x84,
	0x0e, 0x7c, 0x0f, 0xd6, 0x8e, 0x1d, 0x99, 0x2a, 0x94, 0x2c, 0xcc, 0x7c, 0x02, 0xe4, 0xf3, 0xe1,
	0x98, 0x80, 0x92, 0x4e, 0x3e, 0x9d, 0xff, 0xff, 0xff, 0xef, 0x3d, 0xdd, 0x7b, 0xe8, 0x58, 0xb1,
	0x49, 0xc4, 0xd2, 0x84, 0x4f, 0x94, 0x3f, 0xe6, 0xef, 0xaf, 0x78, 0xc4, 0xd5, 0xcc, 0x9f, 0x9e,
	0x0e, 0x98, 0xa2, 0xa7, 0x7e, 0xcc, 0x26, 0x0c, 0x38, 0x78, 0x32, 0x15, 0x4a, 0xe0, 0xc3, 0xa5,
	0xd6, 0x2b, 0xb5, 0x9e, 0xd1, 0xb6, 0x5f, 0xac, 0x4d, 0x5a, 0xea, 0x75, 0x56, 0xbb, 0x15, 0x8b,
	0x58, 0xe8, 0xa3, 0x9f, 0x9f, 0x8a, 0x5b, 0xf2, 0x75, 0x07, 0xa1, 0x37, 0x42, 0x8c, 0x43, 0x36,
	0x14, 0x69, 0x84, 0x5f, 0xa3, 0x9a, 0x14, 0x62, 0xec, 0xd8, 0x47, 0x76, 0xb7, 0x71, 0x46, 0xbc,
	0x75, 0x7c, 0x2f, 0xf7, 0x05, 0xfb, 0x37, 0x77, 0x1d, 0xeb, 0xd7, 0x5d, 0xa7, 0x31, 0xa3, 0xc9,
	0xf8, 0x82, 0xe4, 0x6e, 0x12, 0xea, 0x10, 0x9c, 0xa0, 0xdd, 0xfc, 0xdb, 0x4f, 0x98, 0xa2, 0x11,
	0x55, 0xd4, 0xd9, 0xd2, 0xa9, 0xc7, 0x9b, 0x53, 0x7b, 0xc6, 0x11, 0x1c, 0x9a, 0xf4, 0xd6, 0x32,
	0xbd, 0x8c, 0x23, 0x61, 0x53, 0x56, 0xb4, 0x98, 0x22, 0xa4, 0xff, 0x0f, 0xa8, 0x1a, 0x8e, 0x9c,
	0x6d, 0xcd, 0x7a, 0x7e, 0x8f, 0x0e, 0x72, 0x79, 0x70, 0x60, 0x40, 0x7b, 0x15, 0x90, 0x0e, 0x22,
	0xe1, 0x43, 0xf9, 0x47, 0x85, 0x3f, 0x22, 0x1c, 0x31, 0x29, 0x80, 0xab, 0x7e, 0x02, 0x71, 0x1f,
	0x14, 0x55, 0x0c, 0x9c, 0xda, 0xd1, 0x76, 0xb7, 0x71, 0x76, 0xb2, 0x1e, 0xf5, 0xb2, 0xf0, 0xf5,
	0x20, 0xbe, 0xcc, 0x5d, 0xc1, 0x53, 0x03, 0x3c, 0x28, 0x80, 0xff, 0xc6, 0x92, 0xf0, 0x71, 0xf4,
	0xb7, 0x07, 0xf0, 0x27, 0x1b, 0xed, 0x67, 0x5c, 0x8d, 0xa2, 0x94, 0x66, 0xd5, 0x0a, 0x76, 0x74,
	0x05, 0xde, 0xfa, 0x0a, 0xde, 0x1a, 0x63, 0x59, 0x02, 0x31, 0x25, 0xb4, 0x8b, 0x12, 0xfe, 0x13,
	0x4c, 0xc2, 0xbd, 0x6c, 0xc5, 0x05, 0x38, 0x45, 0x8f, 0x20, 0xa3, 0xb2, 0xca, 0xaf, 0x6b, 0xfe,
	0x86, 0x87, 0xbd, 0xcc, 0xa8, 0x2c, 0xd9, 0xae, 0x61, 0x3f, 0x29, 0xd8, 0x2b, 0x81, 0x24, 0xdc,
	0x85, 0x8a, 0x1a, 0xc8, 0x37, 0x1b, 0x35, 0x5f, 0x15, 0xab, 0xa1, 0x6f, 0x70, 0x80, 0xea, 0x92,
	0xa6, 0x34, 0x01, 0x33, 0xaa, 0xcf, 0x36, 0x3c, 0xb4, 0xd6, 0x06, 0xb5, 0x9c, 0x1a, 0x1a, 0x27,
	0xa6, 0x48, 0x0f, 0x50, 0x3f, 0xd5, 0xb3, 0x0f, 0xce, 0x96, 0xee, 0xa2, 0xbb, 0x79, 0x64, 0x8a,
	0x65, 0x09, 0x5a, 0xa6, 0x87, 0xe6, 0x72, 0x66, 0x80, 0x84, 0x0d, 0x59, 0x2a, 0xe0, 0xe2, 0xc1,
	0xe7, 0xeb, 0x8e, 0xf5, 0xf3, 0xba, 0x63, 0x05, 0xbd, 0x9b, 0xb9, 0x6b, 0xdf, 0xce, 0x5d, 0xfb,
	0xc7, 0xdc, 0xb5, 0xbf, 0x2c, 0x5c, 0xeb, 0x76, 0xe1, 0x5a, 0xdf, 0x17, 0xae, 0xf5, 0xee, 0x3c,
	0xe6, 0x6a, 0x74, 0x35, 0xf0, 0x86, 0x22, 0xf1, 0xe3, 0x94, 0x4e, 0xb9, 0x9a, 0x9d, 0x44, 0x6c,
	0x0a, 0x95, 0x9d, 0xfe, 0x50, 0x39, 0xab, 0x99, 0x64, 0x30, 0xa8, 0xeb, 0xf5, 0x3d, 0xff, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x59, 0xae, 0xc9, 0xf1, 0x4e, 0x04, 0x00, 0x00,
}

func (m *PoolRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SwapMsgStates) > 0 {
		for iNdEx := len(m.SwapMsgStates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SwapMsgStates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.WithdrawMsgStates) > 0 {
		for iNdEx := len(m.WithdrawMsgStates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WithdrawMsgStates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.DepositMsgStates) > 0 {
		for iNdEx := len(m.DepositMsgStates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DepositMsgStates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size, err := m.PoolBatch.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.PoolMetadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Pool.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PoolRecords) > 0 {
		for iNdEx := len(m.PoolRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PoolRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Pool.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.PoolMetadata.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.PoolBatch.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.DepositMsgStates) > 0 {
		for _, e := range m.DepositMsgStates {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.WithdrawMsgStates) > 0 {
		for _, e := range m.WithdrawMsgStates {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SwapMsgStates) > 0 {
		for _, e := range m.SwapMsgStates {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.PoolRecords) > 0 {
		for _, e := range m.PoolRecords {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PoolRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: PoolRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pool", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Pool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PoolMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolBatch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PoolBatch.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositMsgStates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DepositMsgStates = append(m.DepositMsgStates, DepositMsgState{})
			if err := m.DepositMsgStates[len(m.DepositMsgStates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawMsgStates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WithdrawMsgStates = append(m.WithdrawMsgStates, WithdrawMsgState{})
			if err := m.WithdrawMsgStates[len(m.WithdrawMsgStates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapMsgStates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SwapMsgStates = append(m.SwapMsgStates, SwapMsgState{})
			if err := m.SwapMsgStates[len(m.SwapMsgStates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolRecords = append(m.PoolRecords, PoolRecord{})
			if err := m.PoolRecords[len(m.PoolRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
