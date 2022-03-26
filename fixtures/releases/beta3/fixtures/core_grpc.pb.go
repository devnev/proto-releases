// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: fixtures/core.proto

package fixtures

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestServiceClient interface {
	EmptyMethodReleased(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	MethodReleased(ctx context.Context, in *MessageNotAnnotated, opts ...grpc.CallOption) (*MessageNotAnnotated, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) EmptyMethodReleased(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/com.github.devnev.proto_releases.fixtures.releases.beta3.fixtures.TestService/EmptyMethodReleased", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) MethodReleased(ctx context.Context, in *MessageNotAnnotated, opts ...grpc.CallOption) (*MessageNotAnnotated, error) {
	out := new(MessageNotAnnotated)
	err := c.cc.Invoke(ctx, "/com.github.devnev.proto_releases.fixtures.releases.beta3.fixtures.TestService/MethodReleased", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
// All implementations must embed UnimplementedTestServiceServer
// for forward compatibility
type TestServiceServer interface {
	EmptyMethodReleased(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	MethodReleased(context.Context, *MessageNotAnnotated) (*MessageNotAnnotated, error)
	mustEmbedUnimplementedTestServiceServer()
}

// UnimplementedTestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (UnimplementedTestServiceServer) EmptyMethodReleased(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmptyMethodReleased not implemented")
}
func (UnimplementedTestServiceServer) MethodReleased(context.Context, *MessageNotAnnotated) (*MessageNotAnnotated, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MethodReleased not implemented")
}
func (UnimplementedTestServiceServer) mustEmbedUnimplementedTestServiceServer() {}

// UnsafeTestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServiceServer will
// result in compilation errors.
type UnsafeTestServiceServer interface {
	mustEmbedUnimplementedTestServiceServer()
}

func RegisterTestServiceServer(s grpc.ServiceRegistrar, srv TestServiceServer) {
	s.RegisterService(&TestService_ServiceDesc, srv)
}

func _TestService_EmptyMethodReleased_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).EmptyMethodReleased(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.devnev.proto_releases.fixtures.releases.beta3.fixtures.TestService/EmptyMethodReleased",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).EmptyMethodReleased(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_MethodReleased_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageNotAnnotated)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).MethodReleased(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.github.devnev.proto_releases.fixtures.releases.beta3.fixtures.TestService/MethodReleased",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).MethodReleased(ctx, req.(*MessageNotAnnotated))
	}
	return interceptor(ctx, in, info, handler)
}

// TestService_ServiceDesc is the grpc.ServiceDesc for TestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.github.devnev.proto_releases.fixtures.releases.beta3.fixtures.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EmptyMethodReleased",
			Handler:    _TestService_EmptyMethodReleased_Handler,
		},
		{
			MethodName: "MethodReleased",
			Handler:    _TestService_MethodReleased_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fixtures/core.proto",
}
