// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/snapshot/v1beta1/service.proto

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
	proto.RegisterFile("scalar/snapshot/v1beta1/service.proto", fileDescriptor_ed0955c396808261)
}
func init() {
	golang_proto.RegisterFile("scalar/snapshot/v1beta1/service.proto", fileDescriptor_ed0955c396808261)
}

var fileDescriptor_ed0955c396808261 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4b, 0x32, 0x41,
	0x18, 0xc7, 0x1d, 0x0f, 0x1e, 0x96, 0xf7, 0xe5, 0x85, 0xe5, 0x85, 0x40, 0x62, 0x4b, 0x2b, 0x05,
	0xc9, 0x99, 0xb4, 0x5b, 0xc7, 0xe8, 0x1a, 0x99, 0xdd, 0xba, 0xc4, 0xb8, 0x3d, 0x8c, 0x0b, 0xba,
	0xcf, 0x38, 0x33, 0x8a, 0x5e, 0xfd, 0x04, 0x81, 0xb7, 0x8e, 0x7d, 0x92, 0x8e, 0x9d, 0x42, 0xe8,
	0xd2, 0x31, 0xdc, 0x3e, 0x48, 0xb8, 0x3b, 0x6b, 0xa4, 0x2c, 0x79, 0x7b, 0x60, 0x7f, 0xff, 0xe7,
	0xf9, 0xf1, 0xdf, 0x71, 0x8e, 0xb4, 0xcf, 0x7b, 0x5c, 0x31, 0x1d, 0x72, 0xa9, 0xbb, 0x68, 0xd8,
	0xa8, 0xd1, 0x01, 0xc3, 0x1b, 0x4c, 0x83, 0x1a, 0x05, 0x3e, 0x50, 0xa9, 0xd0, 0xa0, 0xbb, 0x93,
	0x60, 0x34, 0xc5, 0xa8, 0xc5, 0x8a, 0xff, 0x05, 0x0a, 0x8c, 0x19, 0xb6, 0x9c, 0x12, 0xbc, 0xb8,
	0x2b, 0x10, 0x45, 0x0f, 0x18, 0x97, 0x01, 0xe3, 0x61, 0x88, 0x86, 0x9b, 0x00, 0x43, 0x6d, 0xbf,
	0xee, 0x67, 0xdd, 0x34, 0x63, 0x4b, 0x1c, 0x64, 0x11, 0x83, 0x21, 0xa8, 0x49, 0x02, 0x35, 0x5f,
	0xf3, 0x8e, 0x73, 0xa9, 0xc5, 0x4d, 0x22, 0xea, 0x3e, 0x12, 0xe7, 0x6f, 0x1b, 0x44, 0xa0, 0x0d,
	0xa8, 0x96, 0xc2, 0xf1, 0xc4, 0xad, 0xd3, 0x0c, 0x6b, 0xfa, 0x83, 0x6b, 0xc3, 0x60, 0x08, 0xda,
	0x14, 0xe9, 0xb6, 0xb8, 0x96, 0x18, 0x6a, 0x28, 0xd7, 0xa6, 0x6f, 0x9f, 0xb3, 0xfc, 0x61, 0x79,
	0x8f, 0xad, 0xdb, 0x2a, 0xcb, 0xdf, 0xc9, 0x65, 0xe0, 0x8c, 0xd4, 0xdc, 0x27, 0xe2, 0xfc, 0xbb,
	0x00, 0xee, 0x9b, 0x60, 0xc4, 0x0d, 0x24, 0x7a, 0x2c, 0xf3, 0xde, 0x1a, 0x99, 0x0a, 0x9e, 0x6c,
	0x1f, 0xb0, 0x8a, 0xc7, 0xb1, 0x62, 0xa5, 0x5c, 0xda, 0x50, 0xbc, 0x5f, 0x25, 0x56, 0x92, 0xcd,
	0x19, 0x71, 0xfe, 0x5c, 0x2f, 0x0b, 0x4e, 0x2b, 0x9d, 0x12, 0xa7, 0xd0, 0xe2, 0x8a, 0xf7, 0xb5,
	0x5b, 0xc9, 0xbc, 0x9d, 0x00, 0xa9, 0x63, 0xf5, 0x57, 0xce, 0xaa, 0x55, 0x63, 0xb5, 0x92, 0xbb,
	0xd9, 0x5e, 0xfa, 0xaf, 0x65, 0x1c, 0x38, 0xbf, 0x7a, 0x59, 0x78, 0x64, 0xbe, 0xf0, 0xc8, 0xc7,
	0xc2, 0x23, 0x0f, 0x91, 0x97, 0x7b, 0x8e, 0x3c, 0x32, 0x8f, 0xbc, 0xdc, 0x7b, 0xe4, 0xe5, 0x6e,
	0x1b, 0x22, 0x30, 0xdd, 0x61, 0x87, 0xfa, 0xd8, 0xb7, 0x8b, 0x50, 0x09, 0x3b, 0xd5, 0x7d, 0x54,
	0xc0, 0xc6, 0xdf, 0x9b, 0xcd, 0x44, 0x82, 0xee, 0x14, 0xe2, 0xe7, 0x73, 0xfa, 0x15, 0x00, 0x00,
	0xff, 0xff, 0xdc, 0x4f, 0x04, 0xff, 0xfb, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
	// RegisterProxy defines a method for registering a proxy account that can act
	// in a validator account's stead.
	RegisterProxy(ctx context.Context, in *RegisterProxyRequest, opts ...grpc.CallOption) (*RegisterProxyResponse, error)
	// DeactivateProxy defines a method for deregistering a proxy account.
	DeactivateProxy(ctx context.Context, in *DeactivateProxyRequest, opts ...grpc.CallOption) (*DeactivateProxyResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) RegisterProxy(ctx context.Context, in *RegisterProxyRequest, opts ...grpc.CallOption) (*RegisterProxyResponse, error) {
	out := new(RegisterProxyResponse)
	err := c.cc.Invoke(ctx, "/scalar.snapshot.v1beta1.MsgService/RegisterProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) DeactivateProxy(ctx context.Context, in *DeactivateProxyRequest, opts ...grpc.CallOption) (*DeactivateProxyResponse, error) {
	out := new(DeactivateProxyResponse)
	err := c.cc.Invoke(ctx, "/scalar.snapshot.v1beta1.MsgService/DeactivateProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	// RegisterProxy defines a method for registering a proxy account that can act
	// in a validator account's stead.
	RegisterProxy(context.Context, *RegisterProxyRequest) (*RegisterProxyResponse, error)
	// DeactivateProxy defines a method for deregistering a proxy account.
	DeactivateProxy(context.Context, *DeactivateProxyRequest) (*DeactivateProxyResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) RegisterProxy(ctx context.Context, req *RegisterProxyRequest) (*RegisterProxyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterProxy not implemented")
}
func (*UnimplementedMsgServiceServer) DeactivateProxy(ctx context.Context, req *DeactivateProxyRequest) (*DeactivateProxyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateProxy not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_RegisterProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterProxyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).RegisterProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scalar.snapshot.v1beta1.MsgService/RegisterProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).RegisterProxy(ctx, req.(*RegisterProxyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_DeactivateProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeactivateProxyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).DeactivateProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scalar.snapshot.v1beta1.MsgService/DeactivateProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).DeactivateProxy(ctx, req.(*DeactivateProxyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scalar.snapshot.v1beta1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterProxy",
			Handler:    _MsgService_RegisterProxy_Handler,
		},
		{
			MethodName: "DeactivateProxy",
			Handler:    _MsgService_DeactivateProxy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scalar/snapshot/v1beta1/service.proto",
}

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	Params(ctx context.Context, in *ParamsRequest, opts ...grpc.CallOption) (*ParamsResponse, error)
}

type queryServiceClient struct {
	cc grpc1.ClientConn
}

func NewQueryServiceClient(cc grpc1.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) Params(ctx context.Context, in *ParamsRequest, opts ...grpc.CallOption) (*ParamsResponse, error) {
	out := new(ParamsResponse)
	err := c.cc.Invoke(ctx, "/scalar.snapshot.v1beta1.QueryService/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	Params(context.Context, *ParamsRequest) (*ParamsResponse, error)
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) Params(ctx context.Context, req *ParamsRequest) (*ParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterQueryServiceServer(s grpc1.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scalar.snapshot.v1beta1.QueryService/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).Params(ctx, req.(*ParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scalar.snapshot.v1beta1.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _QueryService_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scalar/snapshot/v1beta1/service.proto",
}
