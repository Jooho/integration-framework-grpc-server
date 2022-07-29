// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: proto/v1/k8scall.proto

package k8scall

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

// K8SCallClient is the client API for K8SCall service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type K8SCallClient interface {
	CreateObjectByStringJson(ctx context.Context, in *K8SStringJson, opts ...grpc.CallOption) (*CreateObjectByFileResponse, error)
}

type k8SCallClient struct {
	cc grpc.ClientConnInterface
}

func NewK8SCallClient(cc grpc.ClientConnInterface) K8SCallClient {
	return &k8SCallClient{cc}
}

func (c *k8SCallClient) CreateObjectByStringJson(ctx context.Context, in *K8SStringJson, opts ...grpc.CallOption) (*CreateObjectByFileResponse, error) {
	out := new(CreateObjectByFileResponse)
	err := c.cc.Invoke(ctx, "/api.K8SCall/CreateObjectByStringJson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// K8SCallServer is the server API for K8SCall service.
// All implementations must embed UnimplementedK8SCallServer
// for forward compatibility
type K8SCallServer interface {
	CreateObjectByStringJson(context.Context, *K8SStringJson) (*CreateObjectByFileResponse, error)
	mustEmbedUnimplementedK8SCallServer()
}

// UnimplementedK8SCallServer must be embedded to have forward compatible implementations.
type UnimplementedK8SCallServer struct {
}

func (UnimplementedK8SCallServer) CreateObjectByStringJson(context.Context, *K8SStringJson) (*CreateObjectByFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateObjectByStringJson not implemented")
}
func (UnimplementedK8SCallServer) mustEmbedUnimplementedK8SCallServer() {}

// UnsafeK8SCallServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to K8SCallServer will
// result in compilation errors.
type UnsafeK8SCallServer interface {
	mustEmbedUnimplementedK8SCallServer()
}

func RegisterK8SCallServer(s grpc.ServiceRegistrar, srv K8SCallServer) {
	s.RegisterService(&K8SCall_ServiceDesc, srv)
}

func _K8SCall_CreateObjectByStringJson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(K8SStringJson)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SCallServer).CreateObjectByStringJson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.K8SCall/CreateObjectByStringJson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SCallServer).CreateObjectByStringJson(ctx, req.(*K8SStringJson))
	}
	return interceptor(ctx, in, info, handler)
}

// K8SCall_ServiceDesc is the grpc.ServiceDesc for K8SCall service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var K8SCall_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.K8SCall",
	HandlerType: (*K8SCallServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateObjectByStringJson",
			Handler:    _K8SCall_CreateObjectByStringJson_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/k8scall.proto",
}