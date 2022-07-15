// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: api/proto/v1/modelserving/modelserving.proto

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

// ModelServingClient is the client API for ModelServing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModelServingClient interface {
	GetSubscribe(ctx context.Context, in *GetSubscribeRequest, opts ...grpc.CallOption) (*GetSubscribeResponse, error)
}

type modelServingClient struct {
	cc grpc.ClientConnInterface
}

func NewModelServingClient(cc grpc.ClientConnInterface) ModelServingClient {
	return &modelServingClient{cc}
}

func (c *modelServingClient) GetSubscribe(ctx context.Context, in *GetSubscribeRequest, opts ...grpc.CallOption) (*GetSubscribeResponse, error) {
	out := new(GetSubscribeResponse)
	err := c.cc.Invoke(ctx, "/v1.ModelServing/GetSubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelServingServer is the server API for ModelServing service.
// All implementations must embed UnimplementedModelServingServer
// for forward compatibility
type ModelServingServer interface {
	GetSubscribe(context.Context, *GetSubscribeRequest) (*GetSubscribeResponse, error)
	mustEmbedUnimplementedModelServingServer()
}

// UnimplementedModelServingServer must be embedded to have forward compatible implementations.
type UnimplementedModelServingServer struct {
}

func (UnimplementedModelServingServer) GetSubscribe(context.Context, *GetSubscribeRequest) (*GetSubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscribe not implemented")
}
func (UnimplementedModelServingServer) mustEmbedUnimplementedModelServingServer() {}

// UnsafeModelServingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModelServingServer will
// result in compilation errors.
type UnsafeModelServingServer interface {
	mustEmbedUnimplementedModelServingServer()
}

func RegisterModelServingServer(s grpc.ServiceRegistrar, srv ModelServingServer) {
	s.RegisterService(&ModelServing_ServiceDesc, srv)
}

func _ModelServing_GetSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServingServer).GetSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ModelServing/GetSubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServingServer).GetSubscribe(ctx, req.(*GetSubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ModelServing_ServiceDesc is the grpc.ServiceDesc for ModelServing service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModelServing_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ModelServing",
	HandlerType: (*ModelServingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSubscribe",
			Handler:    _ModelServing_GetSubscribe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/v1/modelserving/modelserving.proto",
}
