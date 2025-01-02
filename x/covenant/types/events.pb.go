// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/covenant/v1beta1/events.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	exported "github.com/scalarorg/scalar-core/x/covenant/exported"
	github_com_scalarorg_scalar_core_x_multisig_exported "github.com/scalarorg/scalar-core/x/multisig/exported"
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

type SigningStarted struct {
	Module           string                                                                    `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	SigID            uint64                                                                    `protobuf:"varint,2,opt,name=sig_id,json=sigId,proto3" json:"sig_id,omitempty"`
	KeyID            github_com_scalarorg_scalar_core_x_multisig_exported.KeyID                `protobuf:"bytes,3,opt,name=key_id,json=keyId,proto3,casttype=github.com/scalarorg/scalar-core/x/multisig/exported.KeyID" json:"key_id,omitempty"`
	PubKeys          map[string]github_com_scalarorg_scalar_core_x_multisig_exported.PublicKey `protobuf:"bytes,4,rep,name=pub_keys,json=pubKeys,proto3,castvalue=github.com/scalarorg/scalar-core/x/multisig/exported.PublicKey" json:"pub_keys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Psbt             Psbt                                                                      `protobuf:"bytes,5,opt,name=psbt,proto3,casttype=Psbt" json:"psbt,omitempty"`
	RequestingModule string                                                                    `protobuf:"bytes,6,opt,name=requesting_module,json=requestingModule,proto3" json:"requesting_module,omitempty"`
}

func (m *SigningStarted) Reset()         { *m = SigningStarted{} }
func (m *SigningStarted) String() string { return proto.CompactTextString(m) }
func (*SigningStarted) ProtoMessage()    {}
func (*SigningStarted) Descriptor() ([]byte, []int) {
	return fileDescriptor_1539e454966a6a42, []int{0}
}
func (m *SigningStarted) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SigningStarted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SigningStarted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigningStarted.Merge(m, src)
}
func (m *SigningStarted) XXX_Size() int {
	return m.Size()
}
func (m *SigningStarted) XXX_DiscardUnknown() {
	xxx_messageInfo_SigningStarted.DiscardUnknown(m)
}

var xxx_messageInfo_SigningStarted proto.InternalMessageInfo

type SigningCompleted struct {
	Module string `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	SigID  uint64 `protobuf:"varint,2,opt,name=sig_id,json=sigId,proto3" json:"sig_id,omitempty"`
}

func (m *SigningCompleted) Reset()         { *m = SigningCompleted{} }
func (m *SigningCompleted) String() string { return proto.CompactTextString(m) }
func (*SigningCompleted) ProtoMessage()    {}
func (*SigningCompleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_1539e454966a6a42, []int{1}
}
func (m *SigningCompleted) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SigningCompleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SigningCompleted.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SigningCompleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigningCompleted.Merge(m, src)
}
func (m *SigningCompleted) XXX_Size() int {
	return m.Size()
}
func (m *SigningCompleted) XXX_DiscardUnknown() {
	xxx_messageInfo_SigningCompleted.DiscardUnknown(m)
}

var xxx_messageInfo_SigningCompleted proto.InternalMessageInfo

type SigningExpired struct {
	Module string `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	SigID  uint64 `protobuf:"varint,2,opt,name=sig_id,json=sigId,proto3" json:"sig_id,omitempty"`
}

func (m *SigningExpired) Reset()         { *m = SigningExpired{} }
func (m *SigningExpired) String() string { return proto.CompactTextString(m) }
func (*SigningExpired) ProtoMessage()    {}
func (*SigningExpired) Descriptor() ([]byte, []int) {
	return fileDescriptor_1539e454966a6a42, []int{2}
}
func (m *SigningExpired) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SigningExpired) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SigningExpired.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SigningExpired) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SigningExpired.Merge(m, src)
}
func (m *SigningExpired) XXX_Size() int {
	return m.Size()
}
func (m *SigningExpired) XXX_DiscardUnknown() {
	xxx_messageInfo_SigningExpired.DiscardUnknown(m)
}

var xxx_messageInfo_SigningExpired proto.InternalMessageInfo

type TapScriptSigSubmitted struct {
	Module       string                                        `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	SigID        uint64                                        `protobuf:"varint,2,opt,name=sig_id,json=sigId,proto3" json:"sig_id,omitempty"`
	Participant  github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,3,opt,name=participant,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"participant,omitempty"`
	TapScriptSig *exported.TapScriptSig                        `protobuf:"bytes,4,opt,name=tap_script_sig,json=tapScriptSig,proto3" json:"tap_script_sig,omitempty"`
}

func (m *TapScriptSigSubmitted) Reset()         { *m = TapScriptSigSubmitted{} }
func (m *TapScriptSigSubmitted) String() string { return proto.CompactTextString(m) }
func (*TapScriptSigSubmitted) ProtoMessage()    {}
func (*TapScriptSigSubmitted) Descriptor() ([]byte, []int) {
	return fileDescriptor_1539e454966a6a42, []int{3}
}
func (m *TapScriptSigSubmitted) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TapScriptSigSubmitted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TapScriptSigSubmitted.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TapScriptSigSubmitted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TapScriptSigSubmitted.Merge(m, src)
}
func (m *TapScriptSigSubmitted) XXX_Size() int {
	return m.Size()
}
func (m *TapScriptSigSubmitted) XXX_DiscardUnknown() {
	xxx_messageInfo_TapScriptSigSubmitted.DiscardUnknown(m)
}

var xxx_messageInfo_TapScriptSigSubmitted proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SigningStarted)(nil), "scalar.covenant.v1beta1.SigningStarted")
	proto.RegisterMapType((map[string]github_com_scalarorg_scalar_core_x_multisig_exported.PublicKey)(nil), "scalar.covenant.v1beta1.SigningStarted.PubKeysEntry")
	proto.RegisterType((*SigningCompleted)(nil), "scalar.covenant.v1beta1.SigningCompleted")
	proto.RegisterType((*SigningExpired)(nil), "scalar.covenant.v1beta1.SigningExpired")
	proto.RegisterType((*TapScriptSigSubmitted)(nil), "scalar.covenant.v1beta1.TapScriptSigSubmitted")
}

func init() {
	proto.RegisterFile("scalar/covenant/v1beta1/events.proto", fileDescriptor_1539e454966a6a42)
}

var fileDescriptor_1539e454966a6a42 = []byte{
	// 574 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x4f, 0x8b, 0xd3, 0x4e,
	0x18, 0xc7, 0x3b, 0xdb, 0xb4, 0xbf, 0xdd, 0xd9, 0xb2, 0xf4, 0x17, 0x56, 0x2d, 0x45, 0x92, 0x50,
	0x3d, 0x14, 0xb4, 0x09, 0x5d, 0x3d, 0x48, 0x0f, 0x82, 0x71, 0x17, 0xa9, 0x55, 0x2c, 0xc9, 0xe2,
	0xc1, 0x83, 0x25, 0x7f, 0x86, 0x71, 0x68, 0x92, 0x19, 0x33, 0x93, 0xd2, 0xbc, 0x05, 0x4f, 0x82,
	0x17, 0xaf, 0xbe, 0x9b, 0xbd, 0x08, 0x7b, 0xf4, 0x54, 0xb5, 0x7d, 0x17, 0xc5, 0x83, 0x24, 0xcd,
	0xd6, 0xfa, 0x67, 0x41, 0xec, 0x29, 0xf3, 0x4c, 0x3e, 0xf3, 0xf0, 0x7d, 0x9e, 0xe7, 0xfb, 0xc0,
	0x9b, 0xdc, 0x73, 0x02, 0x27, 0x36, 0x3c, 0x3a, 0x41, 0x91, 0x13, 0x09, 0x63, 0xd2, 0x75, 0x91,
	0x70, 0xba, 0x06, 0x9a, 0xa0, 0x48, 0x70, 0x9d, 0xc5, 0x54, 0x50, 0xf9, 0xda, 0x8a, 0xd2, 0x2f,
	0x28, 0xbd, 0xa0, 0x9a, 0x87, 0x98, 0x62, 0x9a, 0x33, 0x46, 0x76, 0x5a, 0xe1, 0xcd, 0x1b, 0x97,
	0x25, 0x15, 0x29, 0x43, 0x45, 0xce, 0xe6, 0xed, 0x5f, 0x21, 0x34, 0x65, 0x34, 0x16, 0xc8, 0xff,
	0x13, 0xdd, 0xfa, 0x58, 0x86, 0x07, 0x36, 0xc1, 0x11, 0x89, 0xb0, 0x2d, 0x9c, 0x0c, 0x93, 0xaf,
	0xc2, 0x6a, 0x48, 0xfd, 0x24, 0x40, 0x0d, 0xa0, 0x81, 0xf6, 0x9e, 0x55, 0x44, 0xb2, 0x06, 0xab,
	0x9c, 0xe0, 0x11, 0xf1, 0x1b, 0x3b, 0x1a, 0x68, 0x4b, 0xe6, 0xde, 0x7c, 0xa6, 0x56, 0x6c, 0x82,
	0xfb, 0xc7, 0x56, 0x85, 0x13, 0xdc, 0xf7, 0xe5, 0x97, 0xb0, 0x3a, 0x46, 0x69, 0x46, 0x94, 0xb3,
	0x97, 0xe6, 0xa3, 0x8c, 0x18, 0xa0, 0xb4, 0x7f, 0xbc, 0x9c, 0xa9, 0x3d, 0x4c, 0xc4, 0xab, 0xc4,
	0xd5, 0x3d, 0x1a, 0x1a, 0x2b, 0x89, 0x34, 0xc6, 0xc5, 0xa9, 0xe3, 0xd1, 0x18, 0x19, 0x53, 0x23,
	0x4c, 0x02, 0x41, 0x38, 0xc1, 0x6b, 0xcd, 0x7a, 0xfe, 0xda, 0xaa, 0x8c, 0x51, 0xda, 0xf7, 0xe5,
	0x77, 0x00, 0xee, 0xb2, 0xc4, 0x1d, 0x8d, 0x51, 0xca, 0x1b, 0x92, 0x56, 0x6e, 0xef, 0x1f, 0xdd,
	0xd5, 0x2f, 0x69, 0xa1, 0xfe, 0x73, 0x55, 0xfa, 0x30, 0x71, 0x07, 0x28, 0xe5, 0x27, 0x91, 0x88,
	0x53, 0xd3, 0x7c, 0xf3, 0x59, 0xbd, 0xff, 0x4f, 0x7a, 0x86, 0x89, 0x1b, 0x10, 0x6f, 0x80, 0x52,
	0xeb, 0x3f, 0xb6, 0xca, 0x28, 0x5f, 0x87, 0x12, 0xe3, 0xae, 0x68, 0x54, 0x34, 0xd0, 0xae, 0x99,
	0xbb, 0xcb, 0x99, 0x2a, 0x0d, 0xb9, 0x2b, 0xac, 0xfc, 0x56, 0xbe, 0x05, 0xff, 0x8f, 0xd1, 0xeb,
	0x04, 0x71, 0x41, 0x22, 0x3c, 0x2a, 0x1a, 0x5b, 0xcd, 0x1b, 0x5b, 0xff, 0xf1, 0xe3, 0x69, 0x7e,
	0xdf, 0xec, 0xc1, 0xda, 0xa6, 0x4e, 0xb9, 0x0e, 0xcb, 0x63, 0x94, 0x16, 0x73, 0xc8, 0x8e, 0xf2,
	0x21, 0xac, 0x4c, 0x9c, 0x20, 0x41, 0xf9, 0x0c, 0x6a, 0xd6, 0x2a, 0xe8, 0xed, 0xdc, 0x03, 0x3d,
	0xe9, 0xfd, 0x07, 0x15, 0xb4, 0x9e, 0xc0, 0x7a, 0x51, 0xf8, 0x43, 0x1a, 0xb2, 0x00, 0x6d, 0x35,
	0xd0, 0xd6, 0xe3, 0xb5, 0x39, 0x4e, 0xa6, 0x8c, 0xc4, 0x5b, 0xe5, 0xfa, 0x06, 0xe0, 0x95, 0x53,
	0x87, 0xd9, 0x5e, 0x4c, 0x98, 0xb0, 0x09, 0xb6, 0x13, 0x37, 0x24, 0x62, 0x3b, 0xc3, 0xd9, 0x70,
	0x9f, 0x39, 0xb1, 0x20, 0x1e, 0x61, 0x4e, 0x24, 0x72, 0xd7, 0xd5, 0xcc, 0xee, 0x72, 0xa6, 0x76,
	0x36, 0x86, 0xeb, 0x51, 0x1e, 0x52, 0x5e, 0x7c, 0x3a, 0xdc, 0x1f, 0x17, 0x0b, 0xf0, 0xdc, 0x09,
	0x1e, 0xf8, 0x7e, 0x8c, 0x38, 0xb7, 0x36, 0xb3, 0xc8, 0xa7, 0xf0, 0x40, 0x38, 0x6c, 0xc4, 0x73,
	0xa1, 0x23, 0x4e, 0x70, 0x43, 0xd2, 0x40, 0x7b, 0xff, 0x48, 0xff, 0xcd, 0x6a, 0x6b, 0x57, 0x5c,
	0x78, 0x6e, 0xb3, 0x3e, 0xab, 0x26, 0x36, 0x22, 0xf3, 0xd9, 0xd9, 0x57, 0xa5, 0x74, 0x36, 0x57,
	0xc0, 0xf9, 0x5c, 0x01, 0x5f, 0xe6, 0x0a, 0x78, 0xbb, 0x50, 0x4a, 0xe7, 0x0b, 0xa5, 0xf4, 0x69,
	0xa1, 0x94, 0x5e, 0x74, 0xff, 0xc2, 0x8c, 0xeb, 0x85, 0xce, 0xe5, 0xbb, 0xd5, 0x7c, 0x81, 0xef,
	0x7c, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x6a, 0x8d, 0x39, 0x6a, 0x04, 0x00, 0x00,
}

func (m *SigningStarted) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SigningStarted) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SigningStarted) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RequestingModule) > 0 {
		i -= len(m.RequestingModule)
		copy(dAtA[i:], m.RequestingModule)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.RequestingModule)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Psbt) > 0 {
		i -= len(m.Psbt)
		copy(dAtA[i:], m.Psbt)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Psbt)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PubKeys) > 0 {
		keysForPubKeys := make([]string, 0, len(m.PubKeys))
		for k := range m.PubKeys {
			keysForPubKeys = append(keysForPubKeys, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForPubKeys)
		for iNdEx := len(keysForPubKeys) - 1; iNdEx >= 0; iNdEx-- {
			v := m.PubKeys[string(keysForPubKeys[iNdEx])]
			baseI := i
			if len(v) > 0 {
				i -= len(v)
				copy(dAtA[i:], v)
				i = encodeVarintEvents(dAtA, i, uint64(len(v)))
				i--
				dAtA[i] = 0x12
			}
			i -= len(keysForPubKeys[iNdEx])
			copy(dAtA[i:], keysForPubKeys[iNdEx])
			i = encodeVarintEvents(dAtA, i, uint64(len(keysForPubKeys[iNdEx])))
			i--
			dAtA[i] = 0xa
			i = encodeVarintEvents(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.KeyID) > 0 {
		i -= len(m.KeyID)
		copy(dAtA[i:], m.KeyID)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.KeyID)))
		i--
		dAtA[i] = 0x1a
	}
	if m.SigID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.SigID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Module) > 0 {
		i -= len(m.Module)
		copy(dAtA[i:], m.Module)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Module)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SigningCompleted) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SigningCompleted) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SigningCompleted) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SigID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.SigID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Module) > 0 {
		i -= len(m.Module)
		copy(dAtA[i:], m.Module)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Module)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SigningExpired) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SigningExpired) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SigningExpired) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SigID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.SigID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Module) > 0 {
		i -= len(m.Module)
		copy(dAtA[i:], m.Module)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Module)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TapScriptSigSubmitted) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TapScriptSigSubmitted) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TapScriptSigSubmitted) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TapScriptSig != nil {
		{
			size, err := m.TapScriptSig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Participant) > 0 {
		i -= len(m.Participant)
		copy(dAtA[i:], m.Participant)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Participant)))
		i--
		dAtA[i] = 0x1a
	}
	if m.SigID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.SigID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Module) > 0 {
		i -= len(m.Module)
		copy(dAtA[i:], m.Module)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Module)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SigningStarted) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Module)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.SigID != 0 {
		n += 1 + sovEvents(uint64(m.SigID))
	}
	l = len(m.KeyID)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.PubKeys) > 0 {
		for k, v := range m.PubKeys {
			_ = k
			_ = v
			l = 0
			if len(v) > 0 {
				l = 1 + len(v) + sovEvents(uint64(len(v)))
			}
			mapEntrySize := 1 + len(k) + sovEvents(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovEvents(uint64(mapEntrySize))
		}
	}
	l = len(m.Psbt)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.RequestingModule)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *SigningCompleted) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Module)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.SigID != 0 {
		n += 1 + sovEvents(uint64(m.SigID))
	}
	return n
}

func (m *SigningExpired) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Module)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.SigID != 0 {
		n += 1 + sovEvents(uint64(m.SigID))
	}
	return n
}

func (m *TapScriptSigSubmitted) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Module)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.SigID != 0 {
		n += 1 + sovEvents(uint64(m.SigID))
	}
	l = len(m.Participant)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.TapScriptSig != nil {
		l = m.TapScriptSig.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SigningStarted) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: SigningStarted: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SigningStarted: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Module", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Module = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigID", wireType)
			}
			m.SigID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SigID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyID = github_com_scalarorg_scalar_core_x_multisig_exported.KeyID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKeys", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PubKeys == nil {
				m.PubKeys = make(map[string]github_com_scalarorg_scalar_core_x_multisig_exported.PublicKey)
			}
			var mapkey string
			mapvalue := []byte{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEvents
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
							return ErrIntOverflowEvents
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
						return ErrInvalidLengthEvents
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthEvents
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapbyteLen uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowEvents
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapbyteLen |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intMapbyteLen := int(mapbyteLen)
					if intMapbyteLen < 0 {
						return ErrInvalidLengthEvents
					}
					postbytesIndex := iNdEx + intMapbyteLen
					if postbytesIndex < 0 {
						return ErrInvalidLengthEvents
					}
					if postbytesIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = make([]byte, mapbyteLen)
					copy(mapvalue, dAtA[iNdEx:postbytesIndex])
					iNdEx = postbytesIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipEvents(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthEvents
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.PubKeys[mapkey] = ((github_com_scalarorg_scalar_core_x_multisig_exported.PublicKey)(mapvalue))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Psbt", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Psbt = append(m.Psbt[:0], dAtA[iNdEx:postIndex]...)
			if m.Psbt == nil {
				m.Psbt = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestingModule", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestingModule = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *SigningCompleted) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: SigningCompleted: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SigningCompleted: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Module", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Module = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigID", wireType)
			}
			m.SigID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SigID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *SigningExpired) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: SigningExpired: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SigningExpired: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Module", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Module = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigID", wireType)
			}
			m.SigID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SigID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *TapScriptSigSubmitted) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: TapScriptSigSubmitted: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TapScriptSigSubmitted: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Module", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Module = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigID", wireType)
			}
			m.SigID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SigID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Participant", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Participant = append(m.Participant[:0], dAtA[iNdEx:postIndex]...)
			if m.Participant == nil {
				m.Participant = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TapScriptSig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TapScriptSig == nil {
				m.TapScriptSig = &exported.TapScriptSig{}
			}
			if err := m.TapScriptSig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)