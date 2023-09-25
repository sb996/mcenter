// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: mcenter/apps/service/pb/rpc.proto

package service

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

const (
	RPC_ValidateCredential_FullMethodName = "/sb996.mcenter.service.RPC/ValidateCredential"
	RPC_QueryService_FullMethodName       = "/sb996.mcenter.service.RPC/QueryService"
	RPC_DescribeService_FullMethodName    = "/sb996.mcenter.service.RPC/DescribeService"
	RPC_QueryGitlabProject_FullMethodName = "/sb996.mcenter.service.RPC/QueryGitlabProject"
)

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCClient interface {
	ValidateCredential(ctx context.Context, in *ValidateCredentialRequest, opts ...grpc.CallOption) (*Service, error)
	QueryService(ctx context.Context, in *QueryServiceRequest, opts ...grpc.CallOption) (*ServiceSet, error)
	DescribeService(ctx context.Context, in *DescribeServiceRequest, opts ...grpc.CallOption) (*Service, error)
	QueryGitlabProject(ctx context.Context, in *QueryGitlabProjectRequest, opts ...grpc.CallOption) (*ServiceSet, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) ValidateCredential(ctx context.Context, in *ValidateCredentialRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, RPC_ValidateCredential_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) QueryService(ctx context.Context, in *QueryServiceRequest, opts ...grpc.CallOption) (*ServiceSet, error) {
	out := new(ServiceSet)
	err := c.cc.Invoke(ctx, RPC_QueryService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DescribeService(ctx context.Context, in *DescribeServiceRequest, opts ...grpc.CallOption) (*Service, error) {
	out := new(Service)
	err := c.cc.Invoke(ctx, RPC_DescribeService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) QueryGitlabProject(ctx context.Context, in *QueryGitlabProjectRequest, opts ...grpc.CallOption) (*ServiceSet, error) {
	out := new(ServiceSet)
	err := c.cc.Invoke(ctx, RPC_QueryGitlabProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility
type RPCServer interface {
	ValidateCredential(context.Context, *ValidateCredentialRequest) (*Service, error)
	QueryService(context.Context, *QueryServiceRequest) (*ServiceSet, error)
	DescribeService(context.Context, *DescribeServiceRequest) (*Service, error)
	QueryGitlabProject(context.Context, *QueryGitlabProjectRequest) (*ServiceSet, error)
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have forward compatible implementations.
type UnimplementedRPCServer struct {
}

func (UnimplementedRPCServer) ValidateCredential(context.Context, *ValidateCredentialRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateCredential not implemented")
}
func (UnimplementedRPCServer) QueryService(context.Context, *QueryServiceRequest) (*ServiceSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryService not implemented")
}
func (UnimplementedRPCServer) DescribeService(context.Context, *DescribeServiceRequest) (*Service, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeService not implemented")
}
func (UnimplementedRPCServer) QueryGitlabProject(context.Context, *QueryGitlabProjectRequest) (*ServiceSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryGitlabProject not implemented")
}
func (UnimplementedRPCServer) mustEmbedUnimplementedRPCServer() {}

// UnsafeRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCServer will
// result in compilation errors.
type UnsafeRPCServer interface {
	mustEmbedUnimplementedRPCServer()
}

func RegisterRPCServer(s grpc.ServiceRegistrar, srv RPCServer) {
	s.RegisterService(&RPC_ServiceDesc, srv)
}

func _RPC_ValidateCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).ValidateCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_ValidateCredential_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).ValidateCredential(ctx, req.(*ValidateCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_QueryService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).QueryService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_QueryService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).QueryService(ctx, req.(*QueryServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DescribeService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DescribeService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_DescribeService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DescribeService(ctx, req.(*DescribeServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_QueryGitlabProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGitlabProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).QueryGitlabProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_QueryGitlabProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).QueryGitlabProject(ctx, req.(*QueryGitlabProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPC_ServiceDesc is the grpc.ServiceDesc for RPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sb996.mcenter.service.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateCredential",
			Handler:    _RPC_ValidateCredential_Handler,
		},
		{
			MethodName: "QueryService",
			Handler:    _RPC_QueryService_Handler,
		},
		{
			MethodName: "DescribeService",
			Handler:    _RPC_DescribeService_Handler,
		},
		{
			MethodName: "QueryGitlabProject",
			Handler:    _RPC_QueryGitlabProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mcenter/apps/service/pb/rpc.proto",
}
