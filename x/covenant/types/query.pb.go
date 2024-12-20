// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/covenant/v1beta1/query.proto

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

type CustodiansRequest struct {
	Name   string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Pubkey []byte          `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Status CustodianStatus `protobuf:"varint,3,opt,name=status,proto3,enum=scalar.covenant.v1beta1.CustodianStatus" json:"status,omitempty"`
}

func (m *CustodiansRequest) Reset()         { *m = CustodiansRequest{} }
func (m *CustodiansRequest) String() string { return proto.CompactTextString(m) }
func (*CustodiansRequest) ProtoMessage()    {}
func (*CustodiansRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b8372e44a779179, []int{0}
}
func (m *CustodiansRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CustodiansRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CustodiansRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CustodiansRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustodiansRequest.Merge(m, src)
}
func (m *CustodiansRequest) XXX_Size() int {
	return m.Size()
}
func (m *CustodiansRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CustodiansRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CustodiansRequest proto.InternalMessageInfo

func (m *CustodiansRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CustodiansRequest) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *CustodiansRequest) GetStatus() CustodianStatus {
	if m != nil {
		return m.Status
	}
	return Unspecified
}

type CustodiansResponse struct {
	Custodians []*Custodian `protobuf:"bytes,1,rep,name=custodians,proto3" json:"custodians,omitempty"`
}

func (m *CustodiansResponse) Reset()         { *m = CustodiansResponse{} }
func (m *CustodiansResponse) String() string { return proto.CompactTextString(m) }
func (*CustodiansResponse) ProtoMessage()    {}
func (*CustodiansResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b8372e44a779179, []int{1}
}
func (m *CustodiansResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CustodiansResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CustodiansResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CustodiansResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustodiansResponse.Merge(m, src)
}
func (m *CustodiansResponse) XXX_Size() int {
	return m.Size()
}
func (m *CustodiansResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CustodiansResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CustodiansResponse proto.InternalMessageInfo

func (m *CustodiansResponse) GetCustodians() []*Custodian {
	if m != nil {
		return m.Custodians
	}
	return nil
}

type CustodianGroupsRequest struct {
	Uid             string          `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name            string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	GroupPubkey     []byte          `protobuf:"bytes,3,opt,name=group_pubkey,json=groupPubkey,proto3" json:"group_pubkey,omitempty"`
	CustodianPubkey []byte          `protobuf:"bytes,4,opt,name=custodian_pubkey,json=custodianPubkey,proto3" json:"custodian_pubkey,omitempty"`
	Status          CustodianStatus `protobuf:"varint,5,opt,name=status,proto3,enum=scalar.covenant.v1beta1.CustodianStatus" json:"status,omitempty"`
}

func (m *CustodianGroupsRequest) Reset()         { *m = CustodianGroupsRequest{} }
func (m *CustodianGroupsRequest) String() string { return proto.CompactTextString(m) }
func (*CustodianGroupsRequest) ProtoMessage()    {}
func (*CustodianGroupsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b8372e44a779179, []int{2}
}
func (m *CustodianGroupsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CustodianGroupsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CustodianGroupsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CustodianGroupsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustodianGroupsRequest.Merge(m, src)
}
func (m *CustodianGroupsRequest) XXX_Size() int {
	return m.Size()
}
func (m *CustodianGroupsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CustodianGroupsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CustodianGroupsRequest proto.InternalMessageInfo

func (m *CustodianGroupsRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *CustodianGroupsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CustodianGroupsRequest) GetGroupPubkey() []byte {
	if m != nil {
		return m.GroupPubkey
	}
	return nil
}

func (m *CustodianGroupsRequest) GetCustodianPubkey() []byte {
	if m != nil {
		return m.CustodianPubkey
	}
	return nil
}

func (m *CustodianGroupsRequest) GetStatus() CustodianStatus {
	if m != nil {
		return m.Status
	}
	return Unspecified
}

type CustodianGroupsResponse struct {
	Groups []*CustodianGroup `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (m *CustodianGroupsResponse) Reset()         { *m = CustodianGroupsResponse{} }
func (m *CustodianGroupsResponse) String() string { return proto.CompactTextString(m) }
func (*CustodianGroupsResponse) ProtoMessage()    {}
func (*CustodianGroupsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b8372e44a779179, []int{3}
}
func (m *CustodianGroupsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CustodianGroupsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CustodianGroupsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CustodianGroupsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustodianGroupsResponse.Merge(m, src)
}
func (m *CustodianGroupsResponse) XXX_Size() int {
	return m.Size()
}
func (m *CustodianGroupsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CustodianGroupsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CustodianGroupsResponse proto.InternalMessageInfo

func (m *CustodianGroupsResponse) GetGroups() []*CustodianGroup {
	if m != nil {
		return m.Groups
	}
	return nil
}

func init() {
	proto.RegisterType((*CustodiansRequest)(nil), "scalar.covenant.v1beta1.CustodiansRequest")
	proto.RegisterType((*CustodiansResponse)(nil), "scalar.covenant.v1beta1.CustodiansResponse")
	proto.RegisterType((*CustodianGroupsRequest)(nil), "scalar.covenant.v1beta1.CustodianGroupsRequest")
	proto.RegisterType((*CustodianGroupsResponse)(nil), "scalar.covenant.v1beta1.CustodianGroupsResponse")
}

func init() {
	proto.RegisterFile("scalar/covenant/v1beta1/query.proto", fileDescriptor_7b8372e44a779179)
}

var fileDescriptor_7b8372e44a779179 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0x19, 0xca, 0x6d, 0x72, 0x0f, 0xe4, 0x5e, 0x6e, 0x73, 0x03, 0x0d, 0x8b, 0xa6, 0xd6,
	0x85, 0x75, 0x61, 0x1b, 0xf0, 0x01, 0x34, 0xb8, 0x70, 0xe1, 0x86, 0xd4, 0x8d, 0x61, 0x63, 0xa6,
	0x65, 0x52, 0x89, 0xd2, 0x29, 0x9d, 0x19, 0x22, 0x4b, 0xdf, 0xc0, 0xc7, 0x72, 0x65, 0x58, 0xba,
	0x34, 0xf0, 0x22, 0x86, 0xe9, 0x30, 0x12, 0x0d, 0x21, 0x71, 0x77, 0x7a, 0xfa, 0x9d, 0x7f, 0xce,
	0xff, 0xe7, 0xc0, 0x21, 0x4b, 0xf0, 0x03, 0x2e, 0xc2, 0x84, 0xce, 0x48, 0x86, 0x33, 0x1e, 0xce,
	0xba, 0x31, 0xe1, 0xb8, 0x1b, 0x4e, 0x05, 0x29, 0xe6, 0x41, 0x5e, 0x50, 0x4e, 0xad, 0x76, 0x09,
	0x05, 0x1b, 0x28, 0x50, 0x50, 0xe7, 0x7f, 0x4a, 0x53, 0x2a, 0x99, 0x70, 0x5d, 0x95, 0x78, 0x67,
	0xa7, 0x26, 0x9f, 0xe7, 0x84, 0x95, 0x90, 0xf7, 0x84, 0xe0, 0xdf, 0x85, 0x60, 0x9c, 0x8e, 0xc6,
	0x38, 0x63, 0x11, 0x99, 0x0a, 0xc2, 0xb8, 0x65, 0x41, 0x2d, 0xc3, 0x13, 0x62, 0x23, 0x17, 0xf9,
	0xbf, 0x23, 0x59, 0x5b, 0x2d, 0x30, 0x73, 0x11, 0xdf, 0x93, 0xb9, 0x5d, 0x75, 0x91, 0xdf, 0x88,
	0xd4, 0x97, 0x75, 0x0e, 0x26, 0xe3, 0x98, 0x0b, 0x66, 0x1b, 0x2e, 0xf2, 0xff, 0xf4, 0xfc, 0x60,
	0xc7, 0x9a, 0x81, 0x7e, 0xe7, 0x5a, 0xf2, 0x91, 0x9a, 0xf3, 0x6e, 0xc0, 0xda, 0x5e, 0x81, 0xe5,
	0x34, 0x63, 0xc4, 0xea, 0x03, 0x24, 0xba, 0x6b, 0x23, 0xd7, 0xf0, 0xeb, 0x3d, 0x6f, 0xbf, 0x76,
	0xb4, 0x35, 0xe5, 0xbd, 0x22, 0x68, 0xe9, 0x3f, 0x97, 0x05, 0x15, 0xb9, 0xb6, 0xd8, 0x04, 0x43,
	0x8c, 0x47, 0xca, 0xe1, 0xba, 0xd4, 0xa6, 0xab, 0x5b, 0xa6, 0x0f, 0xa0, 0x91, 0xae, 0xc7, 0x6e,
	0x95, 0x75, 0x43, 0x5a, 0xaf, 0xcb, 0xde, 0xa0, 0xf4, 0x7f, 0x0c, 0x4d, 0xfd, 0xe2, 0x06, 0xab,
	0x49, 0xec, 0xaf, 0xee, 0x0f, 0xbe, 0x46, 0xf5, 0xeb, 0x87, 0x51, 0x0d, 0xa1, 0xfd, 0xcd, 0x8f,
	0xca, 0xeb, 0x0c, 0x4c, 0xb9, 0xd6, 0x26, 0xab, 0xa3, 0xfd, 0xe2, 0x52, 0x21, 0x52, 0x63, 0xfd,
	0xab, 0x97, 0xa5, 0x83, 0x16, 0x4b, 0x07, 0xbd, 0x2f, 0x1d, 0xf4, 0xbc, 0x72, 0x2a, 0x8b, 0x95,
	0x53, 0x79, 0x5b, 0x39, 0x95, 0x61, 0x37, 0x1d, 0xf3, 0x3b, 0x11, 0x07, 0x09, 0x9d, 0x84, 0xa5,
	0x28, 0x2d, 0x52, 0x55, 0x9d, 0x24, 0xb4, 0x20, 0xe1, 0xe3, 0xe7, 0x95, 0xc9, 0xeb, 0x8a, 0x4d,
	0x79, 0x5e, 0xa7, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xed, 0x33, 0x36, 0x99, 0xd9, 0x02, 0x00,
	0x00,
}

func (m *CustodiansRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CustodiansRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CustodiansRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Pubkey) > 0 {
		i -= len(m.Pubkey)
		copy(dAtA[i:], m.Pubkey)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Pubkey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CustodiansResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CustodiansResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CustodiansResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Custodians) > 0 {
		for iNdEx := len(m.Custodians) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Custodians[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CustodianGroupsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CustodianGroupsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CustodianGroupsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	if len(m.CustodianPubkey) > 0 {
		i -= len(m.CustodianPubkey)
		copy(dAtA[i:], m.CustodianPubkey)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.CustodianPubkey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.GroupPubkey) > 0 {
		i -= len(m.GroupPubkey)
		copy(dAtA[i:], m.GroupPubkey)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.GroupPubkey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Uid) > 0 {
		i -= len(m.Uid)
		copy(dAtA[i:], m.Uid)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Uid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CustodianGroupsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CustodianGroupsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CustodianGroupsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Groups) > 0 {
		for iNdEx := len(m.Groups) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Groups[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CustodiansRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Pubkey)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovQuery(uint64(m.Status))
	}
	return n
}

func (m *CustodiansResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Custodians) > 0 {
		for _, e := range m.Custodians {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *CustodianGroupsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Uid)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.GroupPubkey)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.CustodianPubkey)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovQuery(uint64(m.Status))
	}
	return n
}

func (m *CustodianGroupsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Groups) > 0 {
		for _, e := range m.Groups {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CustodiansRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: CustodiansRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CustodiansRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pubkey = append(m.Pubkey[:0], dAtA[iNdEx:postIndex]...)
			if m.Pubkey == nil {
				m.Pubkey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= CustodianStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *CustodiansResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: CustodiansResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CustodiansResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Custodians", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Custodians = append(m.Custodians, &Custodian{})
			if err := m.Custodians[len(m.Custodians)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *CustodianGroupsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: CustodianGroupsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CustodianGroupsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupPubkey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupPubkey = append(m.GroupPubkey[:0], dAtA[iNdEx:postIndex]...)
			if m.GroupPubkey == nil {
				m.GroupPubkey = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustodianPubkey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CustodianPubkey = append(m.CustodianPubkey[:0], dAtA[iNdEx:postIndex]...)
			if m.CustodianPubkey == nil {
				m.CustodianPubkey = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= CustodianStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *CustodianGroupsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: CustodianGroupsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CustodianGroupsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Groups", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Groups = append(m.Groups, &CustodianGroup{})
			if err := m.Groups[len(m.Groups)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
