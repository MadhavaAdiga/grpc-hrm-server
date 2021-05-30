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

// OrganizatoinServiceClient is the client API for OrganizatoinService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrganizatoinServiceClient interface {
	// rpc to create a new organization
	CreateOrganization(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*CreateOrganizationResponse, error)
	// rpc to search organization
	FindOrganization(ctx context.Context, in *FindOrganizationRequest, opts ...grpc.CallOption) (*FindOrganizationResponse, error)
}

type organizatoinServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrganizatoinServiceClient(cc grpc.ClientConnInterface) OrganizatoinServiceClient {
	return &organizatoinServiceClient{cc}
}

func (c *organizatoinServiceClient) CreateOrganization(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*CreateOrganizationResponse, error) {
	out := new(CreateOrganizationResponse)
	err := c.cc.Invoke(ctx, "/OrganizatoinService/CreateOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizatoinServiceClient) FindOrganization(ctx context.Context, in *FindOrganizationRequest, opts ...grpc.CallOption) (*FindOrganizationResponse, error) {
	out := new(FindOrganizationResponse)
	err := c.cc.Invoke(ctx, "/OrganizatoinService/FindOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizatoinServiceServer is the server API for OrganizatoinService service.
// All implementations must embed UnimplementedOrganizatoinServiceServer
// for forward compatibility
type OrganizatoinServiceServer interface {
	// rpc to create a new organization
	CreateOrganization(context.Context, *CreateOrganizationRequest) (*CreateOrganizationResponse, error)
	// rpc to search organization
	FindOrganization(context.Context, *FindOrganizationRequest) (*FindOrganizationResponse, error)
	mustEmbedUnimplementedOrganizatoinServiceServer()
}

// UnimplementedOrganizatoinServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrganizatoinServiceServer struct {
}

func (UnimplementedOrganizatoinServiceServer) CreateOrganization(context.Context, *CreateOrganizationRequest) (*CreateOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrganization not implemented")
}
func (UnimplementedOrganizatoinServiceServer) FindOrganization(context.Context, *FindOrganizationRequest) (*FindOrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOrganization not implemented")
}
func (UnimplementedOrganizatoinServiceServer) mustEmbedUnimplementedOrganizatoinServiceServer() {}

// UnsafeOrganizatoinServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrganizatoinServiceServer will
// result in compilation errors.
type UnsafeOrganizatoinServiceServer interface {
	mustEmbedUnimplementedOrganizatoinServiceServer()
}

func RegisterOrganizatoinServiceServer(s grpc.ServiceRegistrar, srv OrganizatoinServiceServer) {
	s.RegisterService(&OrganizatoinService_ServiceDesc, srv)
}

func _OrganizatoinService_CreateOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizatoinServiceServer).CreateOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrganizatoinService/CreateOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizatoinServiceServer).CreateOrganization(ctx, req.(*CreateOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizatoinService_FindOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizatoinServiceServer).FindOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrganizatoinService/FindOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizatoinServiceServer).FindOrganization(ctx, req.(*FindOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrganizatoinService_ServiceDesc is the grpc.ServiceDesc for OrganizatoinService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrganizatoinService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OrganizatoinService",
	HandlerType: (*OrganizatoinServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrganization",
			Handler:    _OrganizatoinService_CreateOrganization_Handler,
		},
		{
			MethodName: "FindOrganization",
			Handler:    _OrganizatoinService_FindOrganization_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "organization_service.proto",
}
