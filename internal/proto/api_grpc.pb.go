// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: internal/proto/api.proto

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

// ApiServiceClient is the client API for ApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiServiceClient interface {
	CreateBackupPolicy(ctx context.Context, in *CreateBackupPolicyRequest, opts ...grpc.CallOption) (*CreateBackupPolicyResponse, error)
	StartBackup(ctx context.Context, in *StartBackupRequest, opts ...grpc.CallOption) (*StartBackupResponse, error)
}

type apiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiServiceClient(cc grpc.ClientConnInterface) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) CreateBackupPolicy(ctx context.Context, in *CreateBackupPolicyRequest, opts ...grpc.CallOption) (*CreateBackupPolicyResponse, error) {
	out := new(CreateBackupPolicyResponse)
	err := c.cc.Invoke(ctx, "/proto.ApiService/CreateBackupPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) StartBackup(ctx context.Context, in *StartBackupRequest, opts ...grpc.CallOption) (*StartBackupResponse, error) {
	out := new(StartBackupResponse)
	err := c.cc.Invoke(ctx, "/proto.ApiService/StartBackup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServiceServer is the server API for ApiService service.
// All implementations must embed UnimplementedApiServiceServer
// for forward compatibility
type ApiServiceServer interface {
	CreateBackupPolicy(context.Context, *CreateBackupPolicyRequest) (*CreateBackupPolicyResponse, error)
	StartBackup(context.Context, *StartBackupRequest) (*StartBackupResponse, error)
	mustEmbedUnimplementedApiServiceServer()
}

// UnimplementedApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApiServiceServer struct {
}

func (UnimplementedApiServiceServer) CreateBackupPolicy(context.Context, *CreateBackupPolicyRequest) (*CreateBackupPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBackupPolicy not implemented")
}
func (UnimplementedApiServiceServer) StartBackup(context.Context, *StartBackupRequest) (*StartBackupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartBackup not implemented")
}
func (UnimplementedApiServiceServer) mustEmbedUnimplementedApiServiceServer() {}

// UnsafeApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServiceServer will
// result in compilation errors.
type UnsafeApiServiceServer interface {
	mustEmbedUnimplementedApiServiceServer()
}

func RegisterApiServiceServer(s grpc.ServiceRegistrar, srv ApiServiceServer) {
	s.RegisterService(&ApiService_ServiceDesc, srv)
}

func _ApiService_CreateBackupPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBackupPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateBackupPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ApiService/CreateBackupPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateBackupPolicy(ctx, req.(*CreateBackupPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_StartBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).StartBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ApiService/StartBackup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).StartBackup(ctx, req.(*StartBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiService_ServiceDesc is the grpc.ServiceDesc for ApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ApiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBackupPolicy",
			Handler:    _ApiService_CreateBackupPolicy_Handler,
		},
		{
			MethodName: "StartBackup",
			Handler:    _ApiService_StartBackup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/api.proto",
}
