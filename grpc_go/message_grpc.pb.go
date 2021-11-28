// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc_go

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

// CreateUserClient is the client API for CreateUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CreateUserClient interface {
	Create(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type createUserClient struct {
	cc grpc.ClientConnInterface
}

func NewCreateUserClient(cc grpc.ClientConnInterface) CreateUserClient {
	return &createUserClient{cc}
}

func (c *createUserClient) Create(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/dala.CreateUser/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateUserServer is the server API for CreateUser service.
// All implementations must embed UnimplementedCreateUserServer
// for forward compatibility
type CreateUserServer interface {
	Create(context.Context, *UserRequest) (*UserResponse, error)
	mustEmbedUnimplementedCreateUserServer()
}

// UnimplementedCreateUserServer must be embedded to have forward compatible implementations.
type UnimplementedCreateUserServer struct {
}

func (UnimplementedCreateUserServer) Create(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCreateUserServer) mustEmbedUnimplementedCreateUserServer() {}

// UnsafeCreateUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CreateUserServer will
// result in compilation errors.
type UnsafeCreateUserServer interface {
	mustEmbedUnimplementedCreateUserServer()
}

func RegisterCreateUserServer(s grpc.ServiceRegistrar, srv CreateUserServer) {
	s.RegisterService(&CreateUser_ServiceDesc, srv)
}

func _CreateUser_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateUserServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dala.CreateUser/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateUserServer).Create(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CreateUser_ServiceDesc is the grpc.ServiceDesc for CreateUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CreateUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dala.CreateUser",
	HandlerType: (*CreateUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CreateUser_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/message.proto",
}