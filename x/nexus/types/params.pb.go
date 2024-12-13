// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/nexus/v1beta1/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	utils "github.com/scalarorg/scalar-core/utils"
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

// Params represent the genesis parameters for the module
type Params struct {
	ChainActivationThreshold              utils.Threshold                               `protobuf:"bytes,1,opt,name=chain_activation_threshold,json=chainActivationThreshold,proto3" json:"chain_activation_threshold"`
	ChainMaintainerMissingVoteThreshold   utils.Threshold                               `protobuf:"bytes,2,opt,name=chain_maintainer_missing_vote_threshold,json=chainMaintainerMissingVoteThreshold,proto3" json:"chain_maintainer_missing_vote_threshold"`
	ChainMaintainerIncorrectVoteThreshold utils.Threshold                               `protobuf:"bytes,3,opt,name=chain_maintainer_incorrect_vote_threshold,json=chainMaintainerIncorrectVoteThreshold,proto3" json:"chain_maintainer_incorrect_vote_threshold"`
	ChainMaintainerCheckWindow            int32                                         `protobuf:"varint,4,opt,name=chain_maintainer_check_window,json=chainMaintainerCheckWindow,proto3" json:"chain_maintainer_check_window,omitempty"`
	Gateway                               github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,5,opt,name=gateway,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"gateway,omitempty"`
	EndBlockerLimit                       uint64                                        `protobuf:"varint,6,opt,name=end_blocker_limit,json=endBlockerLimit,proto3" json:"end_blocker_limit,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd4883104db874c0, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "scalar.nexus.v1beta1.Params")
}

func init() { proto.RegisterFile("scalar/nexus/v1beta1/params.proto", fileDescriptor_cd4883104db874c0) }

var fileDescriptor_cd4883104db874c0 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x8a, 0xd3, 0x40,
	0x1c, 0xc7, 0x33, 0xda, 0xad, 0x30, 0x0a, 0x62, 0xd8, 0x43, 0x28, 0x98, 0xd6, 0x7f, 0x18, 0x85,
	0x4e, 0xa8, 0x3e, 0x41, 0xe2, 0x49, 0xdc, 0x05, 0x09, 0xa2, 0xe0, 0x25, 0x4c, 0x26, 0x43, 0x32,
	0x34, 0x99, 0x5f, 0x99, 0x99, 0xb6, 0xbb, 0x78, 0xf0, 0x15, 0x7c, 0x0b, 0x5f, 0xa5, 0xc7, 0x3d,
	0x7a, 0x5a, 0xb4, 0x7d, 0x0b, 0x4f, 0xd2, 0xc9, 0x9f, 0x5d, 0xbb, 0xa7, 0x3d, 0x75, 0x98, 0x7e,
	0xf8, 0x7c, 0xbe, 0x84, 0xc1, 0x4f, 0x34, 0xa3, 0x15, 0x55, 0xa1, 0xe4, 0x67, 0x4b, 0x1d, 0xae,
	0x66, 0x19, 0x37, 0x74, 0x16, 0x2e, 0xa8, 0xa2, 0xb5, 0x26, 0x0b, 0x05, 0x06, 0xdc, 0xe3, 0x06,
	0x21, 0x16, 0x21, 0x2d, 0x32, 0x3a, 0x2e, 0xa0, 0x00, 0x0b, 0x84, 0xfb, 0x53, 0xc3, 0x8e, 0x9e,
	0xb7, 0xba, 0xa5, 0x11, 0xd5, 0x95, 0xce, 0x94, 0x8a, 0xeb, 0x12, 0xaa, 0xbc, 0xa1, 0x9e, 0xfe,
	0x1c, 0xe0, 0xe1, 0x47, 0x9b, 0x70, 0x19, 0x1e, 0xb1, 0x92, 0x0a, 0x99, 0x52, 0x66, 0xc4, 0x8a,
	0x1a, 0x01, 0x32, 0xed, 0x71, 0x0f, 0x4d, 0x50, 0x70, 0xff, 0xcd, 0x98, 0xb4, 0x0b, 0xac, 0xb5,
	0x5b, 0x40, 0x3e, 0x75, 0x58, 0x3c, 0xd8, 0x5c, 0x8e, 0x9d, 0xc4, 0xb3, 0xa2, 0xa8, 0xf7, 0xf4,
	0xff, 0xbb, 0xdf, 0xf0, 0xcb, 0x26, 0x52, 0x53, 0x21, 0x0d, 0x15, 0x92, 0xab, 0xb4, 0x16, 0x5a,
	0x0b, 0x59, 0xa4, 0x2b, 0x30, 0xfc, 0x5a, 0xf1, 0xce, 0x6d, 0x8a, 0xcf, 0xac, 0xf5, 0xb4, 0x97,
	0x9e, 0x36, 0xce, 0xcf, 0x60, 0xf8, 0x55, 0xfc, 0x3b, 0x7e, 0x75, 0x23, 0x2e, 0x24, 0x03, 0xa5,
	0x38, 0x33, 0x87, 0xf9, 0xbb, 0xb7, 0xc9, 0xbf, 0x38, 0xc8, 0xbf, 0xef, 0xac, 0xff, 0x0f, 0x88,
	0xf0, 0xe3, 0x1b, 0x03, 0x58, 0xc9, 0xd9, 0x3c, 0x5d, 0x0b, 0x99, 0xc3, 0xda, 0x1b, 0x4c, 0x50,
	0x70, 0x94, 0x8c, 0x0e, 0x6c, 0xef, 0xf6, 0xc8, 0x17, 0x4b, 0xb8, 0x1f, 0xf0, 0xbd, 0x82, 0x1a,
	0xbe, 0xa6, 0xe7, 0xde, 0xd1, 0x04, 0x05, 0x0f, 0xe2, 0xd9, 0xdf, 0xcb, 0xf1, 0xb4, 0x10, 0xa6,
	0x5c, 0x66, 0x84, 0x41, 0x1d, 0x32, 0xd0, 0x35, 0xe8, 0xf6, 0x67, 0xaa, 0xf3, 0x79, 0x68, 0xce,
	0x17, 0x5c, 0x93, 0x88, 0xb1, 0x28, 0xcf, 0x15, 0xd7, 0x3a, 0xe9, 0x0c, 0xee, 0x6b, 0xfc, 0x88,
	0xcb, 0x3c, 0xcd, 0x2a, 0x60, 0x73, 0xae, 0xd2, 0x4a, 0xd4, 0xc2, 0x78, 0xc3, 0x09, 0x0a, 0x06,
	0xc9, 0x43, 0x2e, 0xf3, 0xb8, 0xb9, 0x3f, 0xd9, 0x5f, 0xc7, 0x27, 0x9b, 0x3f, 0xbe, 0xb3, 0xd9,
	0xfa, 0xe8, 0x62, 0xeb, 0xa3, 0xdf, 0x5b, 0x1f, 0xfd, 0xd8, 0xf9, 0xce, 0xc5, 0xce, 0x77, 0x7e,
	0xed, 0x7c, 0xe7, 0x2b, 0xb9, 0xb6, 0xa0, 0xf9, 0x62, 0xa0, 0x8a, 0xf6, 0x34, 0x65, 0xa0, 0x78,
	0x78, 0xd6, 0x3e, 0x6c, 0xbb, 0x26, 0x1b, 0xda, 0xe7, 0xf7, 0xf6, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xf6, 0x83, 0x20, 0x75, 0xf5, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EndBlockerLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.EndBlockerLimit))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Gateway) > 0 {
		i -= len(m.Gateway)
		copy(dAtA[i:], m.Gateway)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Gateway)))
		i--
		dAtA[i] = 0x2a
	}
	if m.ChainMaintainerCheckWindow != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ChainMaintainerCheckWindow))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.ChainMaintainerIncorrectVoteThreshold.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.ChainMaintainerMissingVoteThreshold.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ChainActivationThreshold.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ChainActivationThreshold.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.ChainMaintainerMissingVoteThreshold.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.ChainMaintainerIncorrectVoteThreshold.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.ChainMaintainerCheckWindow != 0 {
		n += 1 + sovParams(uint64(m.ChainMaintainerCheckWindow))
	}
	l = len(m.Gateway)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.EndBlockerLimit != 0 {
		n += 1 + sovParams(uint64(m.EndBlockerLimit))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainActivationThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ChainActivationThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainMaintainerMissingVoteThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ChainMaintainerMissingVoteThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainMaintainerIncorrectVoteThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ChainMaintainerIncorrectVoteThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainMaintainerCheckWindow", wireType)
			}
			m.ChainMaintainerCheckWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainMaintainerCheckWindow |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gateway", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Gateway = append(m.Gateway[:0], dAtA[iNdEx:postIndex]...)
			if m.Gateway == nil {
				m.Gateway = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndBlockerLimit", wireType)
			}
			m.EndBlockerLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndBlockerLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)