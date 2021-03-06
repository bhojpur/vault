// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// VaultUIClient is the client API for VaultUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VaultUIClient interface {
	// ListEngineSpecs returns a list of Vault Engine(s) that can be started through the UI.
	ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (VaultUI_ListEngineSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type vaultUIClient struct {
	cc grpc.ClientConnInterface
}

func NewVaultUIClient(cc grpc.ClientConnInterface) VaultUIClient {
	return &vaultUIClient{cc}
}

func (c *vaultUIClient) ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (VaultUI_ListEngineSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &VaultUI_ServiceDesc.Streams[0], "/v1.VaultUI/ListEngineSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &vaultUIListEngineSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VaultUI_ListEngineSpecsClient interface {
	Recv() (*ListEngineSpecsResponse, error)
	grpc.ClientStream
}

type vaultUIListEngineSpecsClient struct {
	grpc.ClientStream
}

func (x *vaultUIListEngineSpecsClient) Recv() (*ListEngineSpecsResponse, error) {
	m := new(ListEngineSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *vaultUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.VaultUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VaultUIServer is the server API for VaultUI service.
// All implementations must embed UnimplementedVaultUIServer
// for forward compatibility
type VaultUIServer interface {
	// ListEngineSpecs returns a list of Vault Engine(s) that can be started through the UI.
	ListEngineSpecs(*ListEngineSpecsRequest, VaultUI_ListEngineSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedVaultUIServer()
}

// UnimplementedVaultUIServer must be embedded to have forward compatible implementations.
type UnimplementedVaultUIServer struct {
}

func (UnimplementedVaultUIServer) ListEngineSpecs(*ListEngineSpecsRequest, VaultUI_ListEngineSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListEngineSpecs not implemented")
}
func (UnimplementedVaultUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedVaultUIServer) mustEmbedUnimplementedVaultUIServer() {}

// UnsafeVaultUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VaultUIServer will
// result in compilation errors.
type UnsafeVaultUIServer interface {
	mustEmbedUnimplementedVaultUIServer()
}

func RegisterVaultUIServer(s grpc.ServiceRegistrar, srv VaultUIServer) {
	s.RegisterService(&VaultUI_ServiceDesc, srv)
}

func _VaultUI_ListEngineSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListEngineSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VaultUIServer).ListEngineSpecs(m, &vaultUIListEngineSpecsServer{stream})
}

type VaultUI_ListEngineSpecsServer interface {
	Send(*ListEngineSpecsResponse) error
	grpc.ServerStream
}

type vaultUIListEngineSpecsServer struct {
	grpc.ServerStream
}

func (x *vaultUIListEngineSpecsServer) Send(m *ListEngineSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _VaultUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VaultUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.VaultUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VaultUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VaultUI_ServiceDesc is the grpc.ServiceDesc for VaultUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VaultUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.VaultUI",
	HandlerType: (*VaultUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _VaultUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListEngineSpecs",
			Handler:       _VaultUI_ListEngineSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "vault-ui.proto",
}
