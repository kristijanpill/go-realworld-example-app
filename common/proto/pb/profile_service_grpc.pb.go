// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.0--rc1
// source: profile_service.proto

package pb

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

// ProfileServiceClient is the client API for ProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileServiceClient interface {
	GetProfileByUsername(ctx context.Context, in *ProfileUsernameRequest, opts ...grpc.CallOption) (*ProfileResponse, error)
	FollowUserByUsername(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*ProfileResponse, error)
	UnfollowUserByUsername(ctx context.Context, in *UnfollowRequest, opts ...grpc.CallOption) (*ProfileResponse, error)
	GetProfileById(ctx context.Context, in *ProfileIdRequest, opts ...grpc.CallOption) (*ProfileResponse, error)
	CreateProfile(ctx context.Context, in *CreateProfileRequest, opts ...grpc.CallOption) (*ProfileInfo, error)
	UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*ProfileInfo, error)
	GetProfileIdByUsername(ctx context.Context, in *ProfileIdUsernameRequest, opts ...grpc.CallOption) (*ProfileIdResponse, error)
	GetFollowedProfileIds(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FollowedIds, error)
}

type profileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileServiceClient(cc grpc.ClientConnInterface) ProfileServiceClient {
	return &profileServiceClient{cc}
}

func (c *profileServiceClient) GetProfileByUsername(ctx context.Context, in *ProfileUsernameRequest, opts ...grpc.CallOption) (*ProfileResponse, error) {
	out := new(ProfileResponse)
	err := c.cc.Invoke(ctx, "/profile.ProfileService/GetProfileByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) FollowUserByUsername(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*ProfileResponse, error) {
	out := new(ProfileResponse)
	err := c.cc.Invoke(ctx, "/profile.ProfileService/FollowUserByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) UnfollowUserByUsername(ctx context.Context, in *UnfollowRequest, opts ...grpc.CallOption) (*ProfileResponse, error) {
	out := new(ProfileResponse)
	err := c.cc.Invoke(ctx, "/profile.ProfileService/UnfollowUserByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) GetProfileById(ctx context.Context, in *ProfileIdRequest, opts ...grpc.CallOption) (*ProfileResponse, error) {
	out := new(ProfileResponse)
	err := c.cc.Invoke(ctx, "/profile.ProfileService/GetProfileById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) CreateProfile(ctx context.Context, in *CreateProfileRequest, opts ...grpc.CallOption) (*ProfileInfo, error) {
	out := new(ProfileInfo)
	err := c.cc.Invoke(ctx, "/profile.ProfileService/CreateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*ProfileInfo, error) {
	out := new(ProfileInfo)
	err := c.cc.Invoke(ctx, "/profile.ProfileService/UpdateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) GetProfileIdByUsername(ctx context.Context, in *ProfileIdUsernameRequest, opts ...grpc.CallOption) (*ProfileIdResponse, error) {
	out := new(ProfileIdResponse)
	err := c.cc.Invoke(ctx, "/profile.ProfileService/GetProfileIdByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) GetFollowedProfileIds(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FollowedIds, error) {
	out := new(FollowedIds)
	err := c.cc.Invoke(ctx, "/profile.ProfileService/GetFollowedProfileIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServiceServer is the server API for ProfileService service.
// All implementations must embed UnimplementedProfileServiceServer
// for forward compatibility
type ProfileServiceServer interface {
	GetProfileByUsername(context.Context, *ProfileUsernameRequest) (*ProfileResponse, error)
	FollowUserByUsername(context.Context, *FollowRequest) (*ProfileResponse, error)
	UnfollowUserByUsername(context.Context, *UnfollowRequest) (*ProfileResponse, error)
	GetProfileById(context.Context, *ProfileIdRequest) (*ProfileResponse, error)
	CreateProfile(context.Context, *CreateProfileRequest) (*ProfileInfo, error)
	UpdateProfile(context.Context, *UpdateProfileRequest) (*ProfileInfo, error)
	GetProfileIdByUsername(context.Context, *ProfileIdUsernameRequest) (*ProfileIdResponse, error)
	GetFollowedProfileIds(context.Context, *emptypb.Empty) (*FollowedIds, error)
	mustEmbedUnimplementedProfileServiceServer()
}

// UnimplementedProfileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProfileServiceServer struct {
}

func (UnimplementedProfileServiceServer) GetProfileByUsername(context.Context, *ProfileUsernameRequest) (*ProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileByUsername not implemented")
}
func (UnimplementedProfileServiceServer) FollowUserByUsername(context.Context, *FollowRequest) (*ProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowUserByUsername not implemented")
}
func (UnimplementedProfileServiceServer) UnfollowUserByUsername(context.Context, *UnfollowRequest) (*ProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnfollowUserByUsername not implemented")
}
func (UnimplementedProfileServiceServer) GetProfileById(context.Context, *ProfileIdRequest) (*ProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileById not implemented")
}
func (UnimplementedProfileServiceServer) CreateProfile(context.Context, *CreateProfileRequest) (*ProfileInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProfile not implemented")
}
func (UnimplementedProfileServiceServer) UpdateProfile(context.Context, *UpdateProfileRequest) (*ProfileInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedProfileServiceServer) GetProfileIdByUsername(context.Context, *ProfileIdUsernameRequest) (*ProfileIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileIdByUsername not implemented")
}
func (UnimplementedProfileServiceServer) GetFollowedProfileIds(context.Context, *emptypb.Empty) (*FollowedIds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowedProfileIds not implemented")
}
func (UnimplementedProfileServiceServer) mustEmbedUnimplementedProfileServiceServer() {}

// UnsafeProfileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileServiceServer will
// result in compilation errors.
type UnsafeProfileServiceServer interface {
	mustEmbedUnimplementedProfileServiceServer()
}

func RegisterProfileServiceServer(s grpc.ServiceRegistrar, srv ProfileServiceServer) {
	s.RegisterService(&ProfileService_ServiceDesc, srv)
}

func _ProfileService_GetProfileByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetProfileByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileService/GetProfileByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetProfileByUsername(ctx, req.(*ProfileUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_FollowUserByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).FollowUserByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileService/FollowUserByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).FollowUserByUsername(ctx, req.(*FollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_UnfollowUserByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnfollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).UnfollowUserByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileService/UnfollowUserByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).UnfollowUserByUsername(ctx, req.(*UnfollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_GetProfileById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetProfileById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileService/GetProfileById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetProfileById(ctx, req.(*ProfileIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_CreateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).CreateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileService/CreateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).CreateProfile(ctx, req.(*CreateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileService/UpdateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).UpdateProfile(ctx, req.(*UpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_GetProfileIdByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileIdUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetProfileIdByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileService/GetProfileIdByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetProfileIdByUsername(ctx, req.(*ProfileIdUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_GetFollowedProfileIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetFollowedProfileIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileService/GetFollowedProfileIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetFollowedProfileIds(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfileService_ServiceDesc is the grpc.ServiceDesc for ProfileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "profile.ProfileService",
	HandlerType: (*ProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfileByUsername",
			Handler:    _ProfileService_GetProfileByUsername_Handler,
		},
		{
			MethodName: "FollowUserByUsername",
			Handler:    _ProfileService_FollowUserByUsername_Handler,
		},
		{
			MethodName: "UnfollowUserByUsername",
			Handler:    _ProfileService_UnfollowUserByUsername_Handler,
		},
		{
			MethodName: "GetProfileById",
			Handler:    _ProfileService_GetProfileById_Handler,
		},
		{
			MethodName: "CreateProfile",
			Handler:    _ProfileService_CreateProfile_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _ProfileService_UpdateProfile_Handler,
		},
		{
			MethodName: "GetProfileIdByUsername",
			Handler:    _ProfileService_GetProfileIdByUsername_Handler,
		},
		{
			MethodName: "GetFollowedProfileIds",
			Handler:    _ProfileService_GetFollowedProfileIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile_service.proto",
}
