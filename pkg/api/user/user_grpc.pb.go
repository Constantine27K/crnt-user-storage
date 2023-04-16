// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/user/user.proto

package user

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

// UserRegistryClient is the client API for UserRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserRegistryClient interface {
	CreateUser(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error)
	UpdateUser(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserUpdateResponse, error)
	UpdateUserContacts(ctx context.Context, in *UserContactsUpdateRequest, opts ...grpc.CallOption) (*UserContactsUpdateResponse, error)
	GetUsers(ctx context.Context, in *UserGetRequest, opts ...grpc.CallOption) (*UserGetResponse, error)
	GetUserByID(ctx context.Context, in *UserGetByIDRequest, opts ...grpc.CallOption) (*UserGetByIDResponse, error)
}

type userRegistryClient struct {
	cc grpc.ClientConnInterface
}

func NewUserRegistryClient(cc grpc.ClientConnInterface) UserRegistryClient {
	return &userRegistryClient{cc}
}

func (c *userRegistryClient) CreateUser(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserCreateResponse, error) {
	out := new(UserCreateResponse)
	err := c.cc.Invoke(ctx, "/github.constantine27k.crnt_user_service.api.user.UserRegistry/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRegistryClient) UpdateUser(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserUpdateResponse, error) {
	out := new(UserUpdateResponse)
	err := c.cc.Invoke(ctx, "/github.constantine27k.crnt_user_service.api.user.UserRegistry/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRegistryClient) UpdateUserContacts(ctx context.Context, in *UserContactsUpdateRequest, opts ...grpc.CallOption) (*UserContactsUpdateResponse, error) {
	out := new(UserContactsUpdateResponse)
	err := c.cc.Invoke(ctx, "/github.constantine27k.crnt_user_service.api.user.UserRegistry/UpdateUserContacts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRegistryClient) GetUsers(ctx context.Context, in *UserGetRequest, opts ...grpc.CallOption) (*UserGetResponse, error) {
	out := new(UserGetResponse)
	err := c.cc.Invoke(ctx, "/github.constantine27k.crnt_user_service.api.user.UserRegistry/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRegistryClient) GetUserByID(ctx context.Context, in *UserGetByIDRequest, opts ...grpc.CallOption) (*UserGetByIDResponse, error) {
	out := new(UserGetByIDResponse)
	err := c.cc.Invoke(ctx, "/github.constantine27k.crnt_user_service.api.user.UserRegistry/GetUserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserRegistryServer is the server API for UserRegistry service.
// All implementations should embed UnimplementedUserRegistryServer
// for forward compatibility
type UserRegistryServer interface {
	CreateUser(context.Context, *UserCreateRequest) (*UserCreateResponse, error)
	UpdateUser(context.Context, *UserUpdateRequest) (*UserUpdateResponse, error)
	UpdateUserContacts(context.Context, *UserContactsUpdateRequest) (*UserContactsUpdateResponse, error)
	GetUsers(context.Context, *UserGetRequest) (*UserGetResponse, error)
	GetUserByID(context.Context, *UserGetByIDRequest) (*UserGetByIDResponse, error)
}

// UnimplementedUserRegistryServer should be embedded to have forward compatible implementations.
type UnimplementedUserRegistryServer struct {
}

func (UnimplementedUserRegistryServer) CreateUser(context.Context, *UserCreateRequest) (*UserCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserRegistryServer) UpdateUser(context.Context, *UserUpdateRequest) (*UserUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserRegistryServer) UpdateUserContacts(context.Context, *UserContactsUpdateRequest) (*UserContactsUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserContacts not implemented")
}
func (UnimplementedUserRegistryServer) GetUsers(context.Context, *UserGetRequest) (*UserGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedUserRegistryServer) GetUserByID(context.Context, *UserGetByIDRequest) (*UserGetByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByID not implemented")
}

// UnsafeUserRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserRegistryServer will
// result in compilation errors.
type UnsafeUserRegistryServer interface {
	mustEmbedUnimplementedUserRegistryServer()
}

func RegisterUserRegistryServer(s grpc.ServiceRegistrar, srv UserRegistryServer) {
	s.RegisterService(&UserRegistry_ServiceDesc, srv)
}

func _UserRegistry_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegistryServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.constantine27k.crnt_user_service.api.user.UserRegistry/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegistryServer).CreateUser(ctx, req.(*UserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRegistry_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegistryServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.constantine27k.crnt_user_service.api.user.UserRegistry/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegistryServer).UpdateUser(ctx, req.(*UserUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRegistry_UpdateUserContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserContactsUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegistryServer).UpdateUserContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.constantine27k.crnt_user_service.api.user.UserRegistry/UpdateUserContacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegistryServer).UpdateUserContacts(ctx, req.(*UserContactsUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRegistry_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegistryServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.constantine27k.crnt_user_service.api.user.UserRegistry/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegistryServer).GetUsers(ctx, req.(*UserGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRegistry_GetUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGetByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegistryServer).GetUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.constantine27k.crnt_user_service.api.user.UserRegistry/GetUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegistryServer).GetUserByID(ctx, req.(*UserGetByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserRegistry_ServiceDesc is the grpc.ServiceDesc for UserRegistry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserRegistry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.constantine27k.crnt_user_service.api.user.UserRegistry",
	HandlerType: (*UserRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserRegistry_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserRegistry_UpdateUser_Handler,
		},
		{
			MethodName: "UpdateUserContacts",
			Handler:    _UserRegistry_UpdateUserContacts_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _UserRegistry_GetUsers_Handler,
		},
		{
			MethodName: "GetUserByID",
			Handler:    _UserRegistry_GetUserByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user/user.proto",
}
