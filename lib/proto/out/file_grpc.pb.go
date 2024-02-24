// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: file.proto

package proto

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

// GetHelloClient is the client API for GetHello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetHelloClient interface {
	HelloWorld(ctx context.Context, in *HelloWorldResponse, opts ...grpc.CallOption) (*HelloWorldResponse, error)
}

type getHelloClient struct {
	cc grpc.ClientConnInterface
}

func NewGetHelloClient(cc grpc.ClientConnInterface) GetHelloClient {
	return &getHelloClient{cc}
}

func (c *getHelloClient) HelloWorld(ctx context.Context, in *HelloWorldResponse, opts ...grpc.CallOption) (*HelloWorldResponse, error) {
	out := new(HelloWorldResponse)
	err := c.cc.Invoke(ctx, "/file.GetHello/HelloWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetHelloServer is the server API for GetHello service.
// All implementations must embed UnimplementedGetHelloServer
// for forward compatibility
type GetHelloServer interface {
	HelloWorld(context.Context, *HelloWorldResponse) (*HelloWorldResponse, error)
	mustEmbedUnimplementedGetHelloServer()
}

// UnimplementedGetHelloServer must be embedded to have forward compatible implementations.
type UnimplementedGetHelloServer struct {
}

func (UnimplementedGetHelloServer) HelloWorld(context.Context, *HelloWorldResponse) (*HelloWorldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWorld not implemented")
}
func (UnimplementedGetHelloServer) mustEmbedUnimplementedGetHelloServer() {}

// UnsafeGetHelloServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetHelloServer will
// result in compilation errors.
type UnsafeGetHelloServer interface {
	mustEmbedUnimplementedGetHelloServer()
}

func RegisterGetHelloServer(s grpc.ServiceRegistrar, srv GetHelloServer) {
	s.RegisterService(&GetHello_ServiceDesc, srv)
}

func _GetHello_HelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWorldResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetHelloServer).HelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.GetHello/HelloWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetHelloServer).HelloWorld(ctx, req.(*HelloWorldResponse))
	}
	return interceptor(ctx, in, info, handler)
}

// GetHello_ServiceDesc is the grpc.ServiceDesc for GetHello service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetHello_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "file.GetHello",
	HandlerType: (*GetHelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloWorld",
			Handler:    _GetHello_HelloWorld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "file.proto",
}
