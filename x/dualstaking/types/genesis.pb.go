// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/dualstaking/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/lavanet/lava/x/fixationstore/types"
	_ "github.com/lavanet/lava/x/timerstore/types"
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

// GenesisState defines the dualstaking module's genesis state.
type GenesisState struct {
	Params              Params             `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	DelegationsFS       types.GenesisState `protobuf:"bytes,2,opt,name=delegationsFS,proto3" json:"delegationsFS"`
	DelegatorsFS        types.GenesisState `protobuf:"bytes,3,opt,name=delegatorsFS,proto3" json:"delegatorsFS"`
	DelegatorRewardList []DelegatorReward  `protobuf:"bytes,5,rep,name=delegator_reward_list,json=delegatorRewardList,proto3" json:"delegator_reward_list"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5bca863c53f218f, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetDelegationsFS() types.GenesisState {
	if m != nil {
		return m.DelegationsFS
	}
	return types.GenesisState{}
}

func (m *GenesisState) GetDelegatorsFS() types.GenesisState {
	if m != nil {
		return m.DelegatorsFS
	}
	return types.GenesisState{}
}

func (m *GenesisState) GetDelegatorRewardList() []DelegatorReward {
	if m != nil {
		return m.DelegatorRewardList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "lavanet.lava.dualstaking.GenesisState")
}

func init() {
	proto.RegisterFile("lavanet/lava/dualstaking/genesis.proto", fileDescriptor_d5bca863c53f218f)
}

var fileDescriptor_d5bca863c53f218f = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x4d, 0x4b, 0xfb, 0x30,
	0x1c, 0xc7, 0xdb, 0x3d, 0xf1, 0xa7, 0xdb, 0x1f, 0xa4, 0x2a, 0x94, 0x1d, 0xe2, 0x50, 0x94, 0x0d,
	0x21, 0x85, 0x79, 0xf7, 0x30, 0x7c, 0x00, 0xf1, 0x20, 0x9b, 0x27, 0x2f, 0x23, 0x5b, 0x63, 0x0c,
	0x76, 0xcd, 0x48, 0x7e, 0xd3, 0xf9, 0x2e, 0x7c, 0x59, 0x3b, 0xee, 0xe8, 0x49, 0xa4, 0x7d, 0x23,
	0xd2, 0x34, 0x8e, 0x45, 0xe8, 0xc5, 0x53, 0xd2, 0xf2, 0xf9, 0x7e, 0x92, 0x6f, 0x7e, 0xde, 0x49,
	0x4c, 0x5e, 0x48, 0x42, 0x21, 0xcc, 0xd7, 0x30, 0x5a, 0x90, 0x58, 0x01, 0x79, 0xe6, 0x09, 0x0b,
	0x19, 0x4d, 0xa8, 0xe2, 0x0a, 0xcf, 0xa5, 0x00, 0xe1, 0x07, 0x86, 0xc3, 0xf9, 0x8a, 0xb7, 0xb8,
	0xf6, 0x1e, 0x13, 0x4c, 0x68, 0x28, 0xcc, 0x77, 0x05, 0xdf, 0x3e, 0x2e, 0xf5, 0xce, 0x89, 0x24,
	0x33, 0xa3, 0x6d, 0xf7, 0x2c, 0xec, 0x91, 0x2f, 0x09, 0x70, 0x91, 0x28, 0x10, 0x92, 0x6e, 0xbe,
	0x0c, 0x7a, 0x64, 0xa1, 0xc0, 0x67, 0x54, 0x16, 0x9c, 0xde, 0x1a, 0x28, 0x2c, 0x3d, 0x36, 0xa2,
	0x31, 0x65, 0x04, 0x84, 0x1c, 0x4b, 0xfa, 0x4a, 0x64, 0x54, 0x04, 0x0e, 0xb3, 0x8a, 0xd7, 0xba,
	0x2e, 0x9a, 0x8e, 0x80, 0x00, 0xf5, 0xcf, 0xbd, 0x46, 0x71, 0xc3, 0xc0, 0xed, 0xb8, 0xdd, 0x66,
	0xbf, 0x83, 0xcb, 0x9a, 0xe3, 0x3b, 0xcd, 0x0d, 0x6a, 0xab, 0xcf, 0x03, 0x67, 0x68, 0x52, 0xfe,
	0xbd, 0xf7, 0xdf, 0x1c, 0x95, 0x17, 0xb9, 0x1a, 0x05, 0x15, 0xad, 0xe9, 0xda, 0x1a, 0xab, 0x29,
	0xde, 0xbe, 0x80, 0xd1, 0xd9, 0x12, 0x7f, 0xe8, 0xb5, 0x36, 0x05, 0x72, 0x69, 0xf5, 0x4f, 0x52,
	0xcb, 0xe1, 0x4f, 0xbd, 0xfd, 0xdf, 0x8f, 0x32, 0x8e, 0xb9, 0x82, 0xa0, 0xde, 0xa9, 0x76, 0x9b,
	0xfd, 0x5e, 0x79, 0xf1, 0x8b, 0x9f, 0xd8, 0x50, 0xa7, 0x8c, 0x7d, 0x37, 0xb2, 0x7f, 0xdf, 0x72,
	0x05, 0x37, 0xb5, 0x7f, 0xb5, 0x9d, 0xfa, 0xe0, 0x72, 0x95, 0x22, 0x77, 0x9d, 0x22, 0xf7, 0x2b,
	0x45, 0xee, 0x7b, 0x86, 0x9c, 0x75, 0x86, 0x9c, 0x8f, 0x0c, 0x39, 0x0f, 0xa7, 0x8c, 0xc3, 0xd3,
	0x62, 0x82, 0xa7, 0x62, 0x66, 0xcf, 0x6e, 0x69, 0x4d, 0x0f, 0xde, 0xe6, 0x54, 0x4d, 0x1a, 0x7a,
	0x66, 0x67, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xd4, 0x46, 0x25, 0xb5, 0x02, 0x00, 0x00,
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
	if len(m.DelegatorRewardList) > 0 {
		for iNdEx := len(m.DelegatorRewardList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegatorRewardList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size, err := m.DelegatorsFS.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.DelegationsFS.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
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
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.DelegationsFS.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.DelegatorsFS.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.DelegatorRewardList) > 0 {
		for _, e := range m.DelegatorRewardList {
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
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationsFS", wireType)
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
			if err := m.DelegationsFS.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorsFS", wireType)
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
			if err := m.DelegatorsFS.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorRewardList", wireType)
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
			m.DelegatorRewardList = append(m.DelegatorRewardList, DelegatorReward{})
			if err := m.DelegatorRewardList[len(m.DelegatorRewardList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
