// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: proto/template/template.proto

package template

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TestRPC_TestConnection_FullMethodName = "/template.TestRPC/TestConnection"
)

// TestRPCClient is the client API for TestRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestRPCClient interface {
	TestConnection(ctx context.Context, in *TestConnectionRequest, opts ...grpc.CallOption) (*TestConnectionResponse, error)
}

type testRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewTestRPCClient(cc grpc.ClientConnInterface) TestRPCClient {
	return &testRPCClient{cc}
}

func (c *testRPCClient) TestConnection(ctx context.Context, in *TestConnectionRequest, opts ...grpc.CallOption) (*TestConnectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TestConnectionResponse)
	err := c.cc.Invoke(ctx, TestRPC_TestConnection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestRPCServer is the server API for TestRPC service.
// All implementations must embed UnimplementedTestRPCServer
// for forward compatibility.
type TestRPCServer interface {
	TestConnection(context.Context, *TestConnectionRequest) (*TestConnectionResponse, error)
	mustEmbedUnimplementedTestRPCServer()
}

// UnimplementedTestRPCServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTestRPCServer struct{}

func (UnimplementedTestRPCServer) TestConnection(context.Context, *TestConnectionRequest) (*TestConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestConnection not implemented")
}
func (UnimplementedTestRPCServer) mustEmbedUnimplementedTestRPCServer() {}
func (UnimplementedTestRPCServer) testEmbeddedByValue()                 {}

// UnsafeTestRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestRPCServer will
// result in compilation errors.
type UnsafeTestRPCServer interface {
	mustEmbedUnimplementedTestRPCServer()
}

func RegisterTestRPCServer(s grpc.ServiceRegistrar, srv TestRPCServer) {
	// If the following call pancis, it indicates UnimplementedTestRPCServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TestRPC_ServiceDesc, srv)
}

func _TestRPC_TestConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestRPCServer).TestConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestRPC_TestConnection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestRPCServer).TestConnection(ctx, req.(*TestConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TestRPC_ServiceDesc is the grpc.ServiceDesc for TestRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "template.TestRPC",
	HandlerType: (*TestRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestConnection",
			Handler:    _TestRPC_TestConnection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/template/template.proto",
}
