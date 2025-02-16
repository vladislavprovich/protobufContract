// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: userinfo/userinfo.proto

package userinfo

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
	UserInfoService_GetUserByID_FullMethodName    = "/userinfo.UserInfoService/GetUserByID"
	UserInfoService_GetUserByEmail_FullMethodName = "/userinfo.UserInfoService/GetUserByEmail"
)

// UserInfoServiceClient is the client API for UserInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserInfoServiceClient interface {
	GetUserByID(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetUserByEmail(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserInfoServiceClient(cc grpc.ClientConnInterface) UserInfoServiceClient {
	return &userInfoServiceClient{cc}
}

func (c *userInfoServiceClient) GetUserByID(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserInfoService_GetUserByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoServiceClient) GetUserByEmail(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserInfoService_GetUserByEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInfoServiceServer is the server API for UserInfoService service.
// All implementations must embed UnimplementedUserInfoServiceServer
// for forward compatibility.
type UserInfoServiceServer interface {
	GetUserByID(context.Context, *GetUserRequest) (*UserResponse, error)
	GetUserByEmail(context.Context, *GetUserRequest) (*UserResponse, error)
	mustEmbedUnimplementedUserInfoServiceServer()
}

// UnimplementedUserInfoServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserInfoServiceServer struct{}

func (UnimplementedUserInfoServiceServer) GetUserByID(context.Context, *GetUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByID not implemented")
}
func (UnimplementedUserInfoServiceServer) GetUserByEmail(context.Context, *GetUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByEmail not implemented")
}
func (UnimplementedUserInfoServiceServer) mustEmbedUnimplementedUserInfoServiceServer() {}
func (UnimplementedUserInfoServiceServer) testEmbeddedByValue()                         {}

// UnsafeUserInfoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserInfoServiceServer will
// result in compilation errors.
type UnsafeUserInfoServiceServer interface {
	mustEmbedUnimplementedUserInfoServiceServer()
}

func RegisterUserInfoServiceServer(s grpc.ServiceRegistrar, srv UserInfoServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserInfoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserInfoService_ServiceDesc, srv)
}

func _UserInfoService_GetUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServiceServer).GetUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserInfoService_GetUserByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServiceServer).GetUserByID(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInfoService_GetUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServiceServer).GetUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserInfoService_GetUserByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServiceServer).GetUserByEmail(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserInfoService_ServiceDesc is the grpc.ServiceDesc for UserInfoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserInfoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userinfo.UserInfoService",
	HandlerType: (*UserInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserByID",
			Handler:    _UserInfoService_GetUserByID_Handler,
		},
		{
			MethodName: "GetUserByEmail",
			Handler:    _UserInfoService_GetUserByEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userinfo/userinfo.proto",
}
