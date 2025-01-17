// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/protocol/v1beta1/service.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("scalar/protocol/v1beta1/service.proto", fileDescriptor_8e6380763574b912)
}
func init() {
	golang_proto.RegisterFile("scalar/protocol/v1beta1/service.proto", fileDescriptor_8e6380763574b912)
}

var fileDescriptor_8e6380763574b912 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x4b, 0xe3, 0x40,
	0x18, 0x87, 0x3b, 0xbb, 0x74, 0x61, 0x73, 0x58, 0xd8, 0x50, 0xd8, 0x25, 0x2c, 0xa1, 0x64, 0x59,
	0x58, 0x8b, 0xcd, 0x90, 0xb6, 0x22, 0xf4, 0xa6, 0x3d, 0x8b, 0xff, 0xf0, 0xe2, 0xa5, 0x4c, 0x93,
	0x21, 0x0d, 0xd4, 0x4c, 0x3a, 0x33, 0x29, 0xed, 0x55, 0x3c, 0x8b, 0xe0, 0x77, 0xf0, 0x0b, 0xf8,
	0x01, 0xf4, 0xe8, 0xcd, 0x82, 0x17, 0x8f, 0xd2, 0xf8, 0x41, 0x24, 0x93, 0x89, 0xd2, 0xd6, 0xa1,
	0xe9, 0x6d, 0x32, 0xef, 0xf3, 0x7b, 0xdf, 0x27, 0x6f, 0x88, 0xf6, 0x8f, 0xb9, 0x68, 0x80, 0x28,
	0x8c, 0x28, 0xe1, 0xc4, 0x25, 0x03, 0x38, 0x72, 0x7a, 0x98, 0x23, 0x07, 0x32, 0x4c, 0x47, 0x81,
	0x8b, 0x6d, 0x51, 0xd0, 0x7f, 0x65, 0x98, 0x9d, 0x63, 0xb6, 0xc4, 0x8c, 0x8a, 0x4f, 0x7c, 0x22,
	0x6e, 0x61, 0x7a, 0xca, 0x00, 0xe3, 0x8f, 0x4f, 0x88, 0x3f, 0xc0, 0x10, 0x45, 0x01, 0x44, 0x61,
	0x48, 0x38, 0xe2, 0x01, 0x09, 0x99, 0xac, 0xfe, 0x55, 0xcd, 0x1c, 0xc6, 0x98, 0x4e, 0x24, 0x54,
	0x55, 0x41, 0x7c, 0x9c, 0x11, 0x8d, 0xc7, 0xb2, 0xf6, 0x75, 0x8f, 0xf9, 0xfa, 0x0d, 0xd0, 0x7e,
	0x74, 0x28, 0x46, 0x1c, 0x1f, 0x48, 0x56, 0xb7, 0x6d, 0x85, 0xaf, 0x3d, 0x0f, 0x1e, 0xe1, 0x61,
	0x8c, 0x19, 0x37, 0x60, 0x61, 0x9e, 0x45, 0x24, 0x64, 0xd8, 0x6a, 0x9e, 0x3f, 0xbd, 0x5e, 0x7f,
	0xa9, 0x5b, 0xff, 0xa1, 0x4a, 0xd3, 0x15, 0xc1, 0x6e, 0x7e, 0xdf, 0x06, 0x35, 0x21, 0x7a, 0x12,
	0x79, 0xc5, 0x44, 0xe7, 0xc1, 0xd5, 0xa2, 0x8b, 0x7c, 0x61, 0xd1, 0x58, 0x04, 0xe7, 0x44, 0x6f,
	0x81, 0xf6, 0x73, 0xc7, 0xf3, 0x8e, 0xe3, 0x28, 0x22, 0x94, 0x63, 0xaf, 0xd3, 0x47, 0x41, 0xa8,
	0x3b, 0xca, 0xd9, 0x4b, 0x6c, 0xae, 0xdb, 0x58, 0x27, 0x22, 0x8d, 0xb7, 0x85, 0xb1, 0x63, 0x6d,
	0x2a, 0x8d, 0x91, 0xe7, 0x75, 0x59, 0x1e, 0xee, 0xba, 0x69, 0x3a, 0xb5, 0xbe, 0x03, 0x5a, 0x25,
	0xdb, 0xc2, 0x82, 0x78, 0x6b, 0xc5, 0xd2, 0x3e, 0x77, 0xdf, 0x5a, 0x33, 0x25, 0xf5, 0xdb, 0x42,
	0xbf, 0x65, 0xc1, 0x55, 0x0b, 0x5f, 0x7e, 0x83, 0xc6, 0x25, 0xd0, 0xca, 0x87, 0xe9, 0x3f, 0xa0,
	0x5f, 0x00, 0xed, 0x7b, 0xfe, 0x2d, 0x99, 0xbe, 0xa1, 0x54, 0x79, 0x67, 0x72, 0xeb, 0x5a, 0x11,
	0x54, 0xaa, 0x56, 0x85, 0xaa, 0xa1, 0xff, 0x56, 0xa9, 0xee, 0xee, 0x3f, 0xcc, 0x4c, 0x30, 0x9d,
	0x99, 0xe0, 0x65, 0x66, 0x82, 0xab, 0xc4, 0x2c, 0xdd, 0x27, 0x26, 0x98, 0x26, 0x66, 0xe9, 0x39,
	0x31, 0x4b, 0xa7, 0x8e, 0x1f, 0xf0, 0x7e, 0xdc, 0xb3, 0x5d, 0x72, 0x26, 0x3b, 0x10, 0xea, 0xcb,
	0x53, 0xdd, 0x25, 0x14, 0xc3, 0xf1, 0x47, 0x4b, 0x3e, 0x89, 0x30, 0xeb, 0x7d, 0x13, 0xcf, 0xcd,
	0xb7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0x53, 0x7f, 0x63, 0x77, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// Create protocol
	CreateProtocol(ctx context.Context, in *CreateProtocolRequest, opts ...grpc.CallOption) (*CreateProtocolResponse, error)
	UpdateProtocol(ctx context.Context, in *UpdateProtocolRequest, opts ...grpc.CallOption) (*UpdateProtocolResponse, error)
	// Add DestinationChain into protocol
	AddSupportedChain(ctx context.Context, in *AddSupportedChainRequest, opts ...grpc.CallOption) (*AddSupportedChainResponse, error)
	// Delete DestinationChain from protocol
	UpdateSupportedChain(ctx context.Context, in *UpdateSupportedChainRequest, opts ...grpc.CallOption) (*UpdateSupportedChainResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateProtocol(ctx context.Context, in *CreateProtocolRequest, opts ...grpc.CallOption) (*CreateProtocolResponse, error) {
	out := new(CreateProtocolResponse)
	err := c.cc.Invoke(ctx, "/scalar.protocol.v1beta1.Msg/CreateProtocol", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateProtocol(ctx context.Context, in *UpdateProtocolRequest, opts ...grpc.CallOption) (*UpdateProtocolResponse, error) {
	out := new(UpdateProtocolResponse)
	err := c.cc.Invoke(ctx, "/scalar.protocol.v1beta1.Msg/UpdateProtocol", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddSupportedChain(ctx context.Context, in *AddSupportedChainRequest, opts ...grpc.CallOption) (*AddSupportedChainResponse, error) {
	out := new(AddSupportedChainResponse)
	err := c.cc.Invoke(ctx, "/scalar.protocol.v1beta1.Msg/AddSupportedChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateSupportedChain(ctx context.Context, in *UpdateSupportedChainRequest, opts ...grpc.CallOption) (*UpdateSupportedChainResponse, error) {
	out := new(UpdateSupportedChainResponse)
	err := c.cc.Invoke(ctx, "/scalar.protocol.v1beta1.Msg/UpdateSupportedChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Create protocol
	CreateProtocol(context.Context, *CreateProtocolRequest) (*CreateProtocolResponse, error)
	UpdateProtocol(context.Context, *UpdateProtocolRequest) (*UpdateProtocolResponse, error)
	// Add DestinationChain into protocol
	AddSupportedChain(context.Context, *AddSupportedChainRequest) (*AddSupportedChainResponse, error)
	// Delete DestinationChain from protocol
	UpdateSupportedChain(context.Context, *UpdateSupportedChainRequest) (*UpdateSupportedChainResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateProtocol(ctx context.Context, req *CreateProtocolRequest) (*CreateProtocolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProtocol not implemented")
}
func (*UnimplementedMsgServer) UpdateProtocol(ctx context.Context, req *UpdateProtocolRequest) (*UpdateProtocolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProtocol not implemented")
}
func (*UnimplementedMsgServer) AddSupportedChain(ctx context.Context, req *AddSupportedChainRequest) (*AddSupportedChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSupportedChain not implemented")
}
func (*UnimplementedMsgServer) UpdateSupportedChain(ctx context.Context, req *UpdateSupportedChainRequest) (*UpdateSupportedChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSupportedChain not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateProtocol_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProtocolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateProtocol(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scalar.protocol.v1beta1.Msg/CreateProtocol",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateProtocol(ctx, req.(*CreateProtocolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateProtocol_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProtocolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateProtocol(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scalar.protocol.v1beta1.Msg/UpdateProtocol",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateProtocol(ctx, req.(*UpdateProtocolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddSupportedChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSupportedChainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddSupportedChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scalar.protocol.v1beta1.Msg/AddSupportedChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddSupportedChain(ctx, req.(*AddSupportedChainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateSupportedChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSupportedChainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateSupportedChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scalar.protocol.v1beta1.Msg/UpdateSupportedChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateSupportedChain(ctx, req.(*UpdateSupportedChainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scalar.protocol.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProtocol",
			Handler:    _Msg_CreateProtocol_Handler,
		},
		{
			MethodName: "UpdateProtocol",
			Handler:    _Msg_UpdateProtocol_Handler,
		},
		{
			MethodName: "AddSupportedChain",
			Handler:    _Msg_AddSupportedChain_Handler,
		},
		{
			MethodName: "UpdateSupportedChain",
			Handler:    _Msg_UpdateSupportedChain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scalar/protocol/v1beta1/service.proto",
}

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// GetProtocols returns all Protocol
	Protocols(ctx context.Context, in *ProtocolsRequest, opts ...grpc.CallOption) (*ProtocolsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Protocols(ctx context.Context, in *ProtocolsRequest, opts ...grpc.CallOption) (*ProtocolsResponse, error) {
	out := new(ProtocolsResponse)
	err := c.cc.Invoke(ctx, "/scalar.protocol.v1beta1.Query/Protocols", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// GetProtocols returns all Protocol
	Protocols(context.Context, *ProtocolsRequest) (*ProtocolsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Protocols(ctx context.Context, req *ProtocolsRequest) (*ProtocolsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Protocols not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Protocols_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProtocolsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Protocols(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scalar.protocol.v1beta1.Query/Protocols",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Protocols(ctx, req.(*ProtocolsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scalar.protocol.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Protocols",
			Handler:    _Query_Protocols_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scalar/protocol/v1beta1/service.proto",
}
