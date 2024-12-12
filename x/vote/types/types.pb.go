// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/vote/v1beta1/types.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	_ "github.com/regen-network/cosmos-proto"
	_ "github.com/scalarorg/scalar-core/x/vote/exported"
	github_com_scalarorg_scalar_core_x_vote_exported "github.com/scalarorg/scalar-core/x/vote/exported"
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

// TalliedVote represents a vote for a poll with the accumulated stake of all
// validators voting for the same data
type TalliedVote struct {
	Tally       github_com_cosmos_cosmos_sdk_types.Uint                 `protobuf:"bytes,1,opt,name=tally,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"tally"`
	Data        *types.Any                                              `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	PollID      github_com_scalarorg_scalar_core_x_vote_exported.PollID `protobuf:"varint,4,opt,name=poll_id,json=pollId,proto3,customtype=github.com/scalarorg/scalar-core/x/vote/exported.PollID" json:"poll_id"`
	IsVoterLate map[string]bool                                         `protobuf:"bytes,5,rep,name=is_voter_late,json=isVoterLate,proto3" json:"is_voter_late,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (m *TalliedVote) Reset()         { *m = TalliedVote{} }
func (m *TalliedVote) String() string { return proto.CompactTextString(m) }
func (*TalliedVote) ProtoMessage()    {}
func (*TalliedVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_a635c404f0665110, []int{0}
}
func (m *TalliedVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TalliedVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *TalliedVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TalliedVote.Merge(m, src)
}
func (m *TalliedVote) XXX_Size() int {
	return m.Size()
}
func (m *TalliedVote) XXX_DiscardUnknown() {
	xxx_messageInfo_TalliedVote.DiscardUnknown(m)
}

var xxx_messageInfo_TalliedVote proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TalliedVote)(nil), "scalar.vote.v1beta1.TalliedVote")
	proto.RegisterMapType((map[string]bool)(nil), "scalar.vote.v1beta1.TalliedVote.IsVoterLateEntry")
}

func init() { proto.RegisterFile("scalar/vote/v1beta1/types.proto", fileDescriptor_a635c404f0665110) }

var fileDescriptor_a635c404f0665110 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x3f, 0x8f, 0xd3, 0x30,
	0x18, 0xc6, 0xe3, 0x26, 0x2d, 0x87, 0x0b, 0x52, 0x15, 0x3a, 0xe4, 0x3a, 0x24, 0x11, 0x03, 0x44,
	0x48, 0xb5, 0xd5, 0x63, 0x00, 0xdd, 0x80, 0x44, 0xc5, 0x0d, 0xe5, 0x8f, 0x74, 0x8a, 0xee, 0x18,
	0x58, 0x2a, 0x27, 0x31, 0xb9, 0xe8, 0x7c, 0x7d, 0x2b, 0xc7, 0xad, 0x2e, 0xdf, 0x82, 0x91, 0x95,
	0x8f, 0x80, 0xc4, 0x87, 0xa8, 0x98, 0x6e, 0x44, 0x0c, 0x15, 0xb4, 0x5f, 0x04, 0x25, 0x36, 0xa2,
	0xa0, 0x0e, 0x4c, 0x7e, 0xec, 0xf7, 0x97, 0x27, 0xef, 0xfb, 0xbc, 0x38, 0x28, 0x53, 0x26, 0x98,
	0xa4, 0x4b, 0x50, 0x9c, 0x2e, 0x47, 0x09, 0x57, 0x6c, 0x44, 0x55, 0x35, 0xe7, 0x25, 0x99, 0x4b,
	0x50, 0xe0, 0xde, 0xd3, 0x00, 0xa9, 0x01, 0x62, 0x80, 0xc1, 0x61, 0x0e, 0x90, 0x0b, 0x4e, 0x1b,
	0x24, 0x59, 0xbc, 0xa7, 0x6c, 0x56, 0x69, 0x7e, 0xd0, 0xcf, 0x21, 0x87, 0x46, 0xd2, 0x5a, 0x99,
	0xd7, 0xc3, 0x14, 0xca, 0x2b, 0x28, 0xa7, 0xba, 0xa0, 0x2f, 0xa6, 0x14, 0xed, 0x76, 0xc0, 0xaf,
	0xe7, 0x20, 0x15, 0xcf, 0xf6, 0xb5, 0x72, 0xff, 0xb3, 0x8d, 0xbb, 0x67, 0x4c, 0x88, 0x82, 0x67,
	0x6f, 0x41, 0x71, 0xf7, 0x04, 0xb7, 0x15, 0x13, 0xa2, 0xf2, 0x50, 0x88, 0xa2, 0x3b, 0x63, 0xba,
	0x5a, 0x07, 0xd6, 0xf7, 0x75, 0xf0, 0x30, 0x2f, 0xd4, 0xc5, 0x22, 0x21, 0x29, 0x5c, 0x99, 0x3f,
	0x99, 0x63, 0x58, 0x66, 0x97, 0xc6, 0xf0, 0xbc, 0x98, 0xa9, 0x58, 0x7f, 0xed, 0x9e, 0x61, 0x27,
	0x63, 0x8a, 0x79, 0x76, 0x88, 0xa2, 0xee, 0x51, 0x9f, 0xe8, 0xd9, 0xc8, 0xef, 0xd9, 0xc8, 0xf3,
	0x59, 0x35, 0x7e, 0xf4, 0xf5, 0xcb, 0xf0, 0xc1, 0x3e, 0xdf, 0x8c, 0xa7, 0xf4, 0xb4, 0x26, 0xdf,
	0x30, 0x59, 0x5e, 0x30, 0xc1, 0x65, 0xdc, 0xb8, 0xb9, 0x09, 0xbe, 0x35, 0x07, 0x21, 0xa6, 0x45,
	0xe6, 0x39, 0x21, 0x8a, 0x9c, 0xf1, 0xc4, 0xb4, 0xf7, 0x64, 0xc7, 0x46, 0x8f, 0x0e, 0x32, 0x37,
	0x6a, 0x98, 0x82, 0xe4, 0xf4, 0xfa, 0xef, 0x2c, 0xc8, 0x29, 0x08, 0x31, 0x79, 0xb1, 0x59, 0x07,
	0x1d, 0xad, 0xe2, 0x4e, 0xed, 0x3c, 0xc9, 0xdc, 0x73, 0x7c, 0xb7, 0x28, 0xa7, 0x35, 0x2c, 0xa7,
	0x82, 0x29, 0xee, 0xb5, 0x43, 0x3b, 0xea, 0x1e, 0x8d, 0xc8, 0x9e, 0x9d, 0x91, 0x9d, 0xe4, 0xc8,
	0xa4, 0xac, 0x0f, 0xf9, 0x9a, 0x29, 0x7e, 0x32, 0x53, 0xb2, 0x8a, 0xbb, 0xc5, 0x9f, 0x97, 0xc1,
	0x33, 0xdc, 0xfb, 0x17, 0x70, 0x7b, 0xd8, 0xbe, 0xe4, 0x3a, 0xe9, 0xdb, 0x71, 0x2d, 0xdd, 0x3e,
	0x6e, 0x2f, 0x99, 0x58, 0x70, 0xaf, 0x15, 0xa2, 0xe8, 0x20, 0xd6, 0x97, 0xe3, 0xd6, 0x53, 0x74,
	0xec, 0x7c, 0xfc, 0x14, 0xa0, 0x97, 0xce, 0x41, 0xab, 0x67, 0x8f, 0x5f, 0xad, 0x7e, 0xfa, 0xd6,
	0x6a, 0xe3, 0xa3, 0x9b, 0x8d, 0x8f, 0x7e, 0x6c, 0x7c, 0xf4, 0x61, 0xeb, 0x5b, 0x37, 0x5b, 0xdf,
	0xfa, 0xb6, 0xf5, 0xad, 0x77, 0xc3, 0xff, 0xcd, 0xa2, 0xd9, 0x5a, 0xd2, 0x69, 0x76, 0xf2, 0xf8,
	0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0x73, 0x55, 0xa3, 0xb5, 0x02, 0x00, 0x00,
}

func (m *TalliedVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TalliedVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TalliedVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IsVoterLate) > 0 {
		keysForIsVoterLate := make([]string, 0, len(m.IsVoterLate))
		for k := range m.IsVoterLate {
			keysForIsVoterLate = append(keysForIsVoterLate, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForIsVoterLate)
		for iNdEx := len(keysForIsVoterLate) - 1; iNdEx >= 0; iNdEx-- {
			v := m.IsVoterLate[string(keysForIsVoterLate[iNdEx])]
			baseI := i
			i--
			if v {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x10
			i -= len(keysForIsVoterLate[iNdEx])
			copy(dAtA[i:], keysForIsVoterLate[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(keysForIsVoterLate[iNdEx])))
			i--
			dAtA[i] = 0xa
			i = encodeVarintTypes(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.PollID != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.PollID))
		i--
		dAtA[i] = 0x20
	}
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Tally.Size()
		i -= size
		if _, err := m.Tally.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TalliedVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Tally.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.PollID != 0 {
		n += 1 + sovTypes(uint64(m.PollID))
	}
	if len(m.IsVoterLate) > 0 {
		for k, v := range m.IsVoterLate {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovTypes(uint64(len(k))) + 1 + 1
			n += mapEntrySize + 1 + sovTypes(uint64(mapEntrySize))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TalliedVote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: TalliedVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TalliedVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tally", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Tally.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &types.Any{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PollID", wireType)
			}
			m.PollID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PollID |= github_com_scalarorg_scalar_core_x_vote_exported.PollID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsVoterLate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.IsVoterLate == nil {
				m.IsVoterLate = make(map[string]bool)
			}
			var mapkey string
			var mapvalue bool
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthTypes
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthTypes
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapvaluetemp int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvaluetemp |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					mapvalue = bool(mapvaluetemp != 0)
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTypes(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthTypes
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.IsVoterLate[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
