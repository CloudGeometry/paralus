// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/rpc/config/override.proto

package config

import (
	context "context"
	v3 "github.com/RafaySystems/rcloud-base/components/common/proto/types/commonpb/v3"
	config "github.com/RafaySystems/rcloud-base/components/common/proto/types/config"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OverrideClient is the client API for Override service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OverrideClient interface {
	CreateOverride(ctx context.Context, in *config.Override, opts ...grpc.CallOption) (*config.Override, error)
	GetOverride(ctx context.Context, in *config.Override, opts ...grpc.CallOption) (*config.Override, error)
	GetOverrides(ctx context.Context, in *v3.QueryOptions, opts ...grpc.CallOption) (*config.OverrideList, error)
	UpdateOverride(ctx context.Context, in *config.Override, opts ...grpc.CallOption) (*config.Override, error)
	DeleteOverride(ctx context.Context, in *config.Override, opts ...grpc.CallOption) (*DeleteOverrideResponse, error)
	ApplyOverride(ctx context.Context, in *ApplyOverrideRequest, opts ...grpc.CallOption) (*ApplyOverrideResponse, error)
	AddDefaultOverrides(ctx context.Context, in *AddDefaultOverridesRequest, opts ...grpc.CallOption) (*AddDefaultOverridesResponse, error)
}

type overrideClient struct {
	cc grpc.ClientConnInterface
}

func NewOverrideClient(cc grpc.ClientConnInterface) OverrideClient {
	return &overrideClient{cc}
}

func (c *overrideClient) CreateOverride(ctx context.Context, in *config.Override, opts ...grpc.CallOption) (*config.Override, error) {
	out := new(config.Override)
	err := c.cc.Invoke(ctx, "/rafay.dev.config.rpc.Override/CreateOverride", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *overrideClient) GetOverride(ctx context.Context, in *config.Override, opts ...grpc.CallOption) (*config.Override, error) {
	out := new(config.Override)
	err := c.cc.Invoke(ctx, "/rafay.dev.config.rpc.Override/GetOverride", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *overrideClient) GetOverrides(ctx context.Context, in *v3.QueryOptions, opts ...grpc.CallOption) (*config.OverrideList, error) {
	out := new(config.OverrideList)
	err := c.cc.Invoke(ctx, "/rafay.dev.config.rpc.Override/GetOverrides", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *overrideClient) UpdateOverride(ctx context.Context, in *config.Override, opts ...grpc.CallOption) (*config.Override, error) {
	out := new(config.Override)
	err := c.cc.Invoke(ctx, "/rafay.dev.config.rpc.Override/UpdateOverride", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *overrideClient) DeleteOverride(ctx context.Context, in *config.Override, opts ...grpc.CallOption) (*DeleteOverrideResponse, error) {
	out := new(DeleteOverrideResponse)
	err := c.cc.Invoke(ctx, "/rafay.dev.config.rpc.Override/DeleteOverride", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *overrideClient) ApplyOverride(ctx context.Context, in *ApplyOverrideRequest, opts ...grpc.CallOption) (*ApplyOverrideResponse, error) {
	out := new(ApplyOverrideResponse)
	err := c.cc.Invoke(ctx, "/rafay.dev.config.rpc.Override/ApplyOverride", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *overrideClient) AddDefaultOverrides(ctx context.Context, in *AddDefaultOverridesRequest, opts ...grpc.CallOption) (*AddDefaultOverridesResponse, error) {
	out := new(AddDefaultOverridesResponse)
	err := c.cc.Invoke(ctx, "/rafay.dev.config.rpc.Override/AddDefaultOverrides", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OverrideServer is the server API for Override service.
// All implementations should embed UnimplementedOverrideServer
// for forward compatibility
type OverrideServer interface {
	CreateOverride(context.Context, *config.Override) (*config.Override, error)
	GetOverride(context.Context, *config.Override) (*config.Override, error)
	GetOverrides(context.Context, *v3.QueryOptions) (*config.OverrideList, error)
	UpdateOverride(context.Context, *config.Override) (*config.Override, error)
	DeleteOverride(context.Context, *config.Override) (*DeleteOverrideResponse, error)
	ApplyOverride(context.Context, *ApplyOverrideRequest) (*ApplyOverrideResponse, error)
	AddDefaultOverrides(context.Context, *AddDefaultOverridesRequest) (*AddDefaultOverridesResponse, error)
}

// UnimplementedOverrideServer should be embedded to have forward compatible implementations.
type UnimplementedOverrideServer struct {
}

func (UnimplementedOverrideServer) CreateOverride(context.Context, *config.Override) (*config.Override, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOverride not implemented")
}
func (UnimplementedOverrideServer) GetOverride(context.Context, *config.Override) (*config.Override, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOverride not implemented")
}
func (UnimplementedOverrideServer) GetOverrides(context.Context, *v3.QueryOptions) (*config.OverrideList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOverrides not implemented")
}
func (UnimplementedOverrideServer) UpdateOverride(context.Context, *config.Override) (*config.Override, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOverride not implemented")
}
func (UnimplementedOverrideServer) DeleteOverride(context.Context, *config.Override) (*DeleteOverrideResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOverride not implemented")
}
func (UnimplementedOverrideServer) ApplyOverride(context.Context, *ApplyOverrideRequest) (*ApplyOverrideResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyOverride not implemented")
}
func (UnimplementedOverrideServer) AddDefaultOverrides(context.Context, *AddDefaultOverridesRequest) (*AddDefaultOverridesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDefaultOverrides not implemented")
}

// UnsafeOverrideServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OverrideServer will
// result in compilation errors.
type UnsafeOverrideServer interface {
	mustEmbedUnimplementedOverrideServer()
}

func RegisterOverrideServer(s grpc.ServiceRegistrar, srv OverrideServer) {
	s.RegisterService(&Override_ServiceDesc, srv)
}

func _Override_CreateOverride_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(config.Override)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverrideServer).CreateOverride(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.config.rpc.Override/CreateOverride",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverrideServer).CreateOverride(ctx, req.(*config.Override))
	}
	return interceptor(ctx, in, info, handler)
}

func _Override_GetOverride_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(config.Override)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverrideServer).GetOverride(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.config.rpc.Override/GetOverride",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverrideServer).GetOverride(ctx, req.(*config.Override))
	}
	return interceptor(ctx, in, info, handler)
}

func _Override_GetOverrides_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.QueryOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverrideServer).GetOverrides(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.config.rpc.Override/GetOverrides",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverrideServer).GetOverrides(ctx, req.(*v3.QueryOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _Override_UpdateOverride_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(config.Override)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverrideServer).UpdateOverride(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.config.rpc.Override/UpdateOverride",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverrideServer).UpdateOverride(ctx, req.(*config.Override))
	}
	return interceptor(ctx, in, info, handler)
}

func _Override_DeleteOverride_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(config.Override)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverrideServer).DeleteOverride(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.config.rpc.Override/DeleteOverride",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverrideServer).DeleteOverride(ctx, req.(*config.Override))
	}
	return interceptor(ctx, in, info, handler)
}

func _Override_ApplyOverride_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyOverrideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverrideServer).ApplyOverride(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.config.rpc.Override/ApplyOverride",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverrideServer).ApplyOverride(ctx, req.(*ApplyOverrideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Override_AddDefaultOverrides_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDefaultOverridesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverrideServer).AddDefaultOverrides(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.config.rpc.Override/AddDefaultOverrides",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverrideServer).AddDefaultOverrides(ctx, req.(*AddDefaultOverridesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Override_ServiceDesc is the grpc.ServiceDesc for Override service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Override_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rafay.dev.config.rpc.Override",
	HandlerType: (*OverrideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOverride",
			Handler:    _Override_CreateOverride_Handler,
		},
		{
			MethodName: "GetOverride",
			Handler:    _Override_GetOverride_Handler,
		},
		{
			MethodName: "GetOverrides",
			Handler:    _Override_GetOverrides_Handler,
		},
		{
			MethodName: "UpdateOverride",
			Handler:    _Override_UpdateOverride_Handler,
		},
		{
			MethodName: "DeleteOverride",
			Handler:    _Override_DeleteOverride_Handler,
		},
		{
			MethodName: "ApplyOverride",
			Handler:    _Override_ApplyOverride_Handler,
		},
		{
			MethodName: "AddDefaultOverrides",
			Handler:    _Override_AddDefaultOverrides_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rpc/config/override.proto",
}