// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package hrm

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

// RoleServiceClient is the client API for RoleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoleServiceClient interface {
	// create a role for the orgaization
	CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error)
}

type roleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleServiceClient(cc grpc.ClientConnInterface) RoleServiceClient {
	return &roleServiceClient{cc}
}

func (c *roleServiceClient) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error) {
	out := new(CreateRoleResponse)
	err := c.cc.Invoke(ctx, "/RoleService/CreateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServiceServer is the server API for RoleService service.
// All implementations must embed UnimplementedRoleServiceServer
// for forward compatibility
type RoleServiceServer interface {
	// create a role for the orgaization
	CreateRole(context.Context, *CreateRoleRequest) (*CreateRoleResponse, error)
	mustEmbedUnimplementedRoleServiceServer()
}

// UnimplementedRoleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRoleServiceServer struct {
}

func (UnimplementedRoleServiceServer) CreateRole(context.Context, *CreateRoleRequest) (*CreateRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedRoleServiceServer) mustEmbedUnimplementedRoleServiceServer() {}

// UnsafeRoleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoleServiceServer will
// result in compilation errors.
type UnsafeRoleServiceServer interface {
	mustEmbedUnimplementedRoleServiceServer()
}

func RegisterRoleServiceServer(s grpc.ServiceRegistrar, srv RoleServiceServer) {
	s.RegisterService(&RoleService_ServiceDesc, srv)
}

func _RoleService_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RoleService/CreateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).CreateRole(ctx, req.(*CreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoleService_ServiceDesc is the grpc.ServiceDesc for RoleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RoleService",
	HandlerType: (*RoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRole",
			Handler:    _RoleService_CreateRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "role_service.proto",
}
