// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: stable3/example.proto

package examples

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ExampleServiceClient is the client API for ExampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExampleServiceClient interface {
	ExampleReleasedMethod(ctx context.Context, in *Example, opts ...grpc.CallOption) (*Example, error)
}

type exampleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleServiceClient(cc grpc.ClientConnInterface) ExampleServiceClient {
	return &exampleServiceClient{cc}
}

func (c *exampleServiceClient) ExampleReleasedMethod(ctx context.Context, in *Example, opts ...grpc.CallOption) (*Example, error) {
	out := new(Example)
	err := c.cc.Invoke(ctx, "/ExampleService/ExampleReleasedMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServiceServer is the server API for ExampleService service.
// All implementations must embed UnimplementedExampleServiceServer
// for forward compatibility
type ExampleServiceServer interface {
	ExampleReleasedMethod(context.Context, *Example) (*Example, error)
	mustEmbedUnimplementedExampleServiceServer()
}

// UnimplementedExampleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExampleServiceServer struct {
}

func (UnimplementedExampleServiceServer) ExampleReleasedMethod(context.Context, *Example) (*Example, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExampleReleasedMethod not implemented")
}
func (UnimplementedExampleServiceServer) mustEmbedUnimplementedExampleServiceServer() {}

// UnsafeExampleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExampleServiceServer will
// result in compilation errors.
type UnsafeExampleServiceServer interface {
	mustEmbedUnimplementedExampleServiceServer()
}

func RegisterExampleServiceServer(s grpc.ServiceRegistrar, srv ExampleServiceServer) {
	s.RegisterService(&ExampleService_ServiceDesc, srv)
}

func _ExampleService_ExampleReleasedMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Example)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).ExampleReleasedMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExampleService/ExampleReleasedMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).ExampleReleasedMethod(ctx, req.(*Example))
	}
	return interceptor(ctx, in, info, handler)
}

// ExampleService_ServiceDesc is the grpc.ServiceDesc for ExampleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExampleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ExampleService",
	HandlerType: (*ExampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExampleReleasedMethod",
			Handler:    _ExampleService_ExampleReleasedMethod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stable3/example.proto",
}
