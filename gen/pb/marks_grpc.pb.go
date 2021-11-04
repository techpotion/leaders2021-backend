// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// MicroserviceClient is the client API for Microservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MicroserviceClient interface {
	GetMark(ctx context.Context, in *Marks_GetRequest, opts ...grpc.CallOption) (*Marks_GetResponse, error)
}

type microserviceClient struct {
	cc grpc.ClientConnInterface
}

func NewMicroserviceClient(cc grpc.ClientConnInterface) MicroserviceClient {
	return &microserviceClient{cc}
}

func (c *microserviceClient) GetMark(ctx context.Context, in *Marks_GetRequest, opts ...grpc.CallOption) (*Marks_GetResponse, error) {
	out := new(Marks_GetResponse)
	err := c.cc.Invoke(ctx, "/pb.Microservice/GetMark", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MicroserviceServer is the server API for Microservice service.
// All implementations must embed UnimplementedMicroserviceServer
// for forward compatibility
type MicroserviceServer interface {
	GetMark(context.Context, *Marks_GetRequest) (*Marks_GetResponse, error)
	mustEmbedUnimplementedMicroserviceServer()
}

// UnimplementedMicroserviceServer must be embedded to have forward compatible implementations.
type UnimplementedMicroserviceServer struct {
}

func (UnimplementedMicroserviceServer) GetMark(context.Context, *Marks_GetRequest) (*Marks_GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMark not implemented")
}
func (UnimplementedMicroserviceServer) mustEmbedUnimplementedMicroserviceServer() {}

// UnsafeMicroserviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MicroserviceServer will
// result in compilation errors.
type UnsafeMicroserviceServer interface {
	mustEmbedUnimplementedMicroserviceServer()
}

func RegisterMicroserviceServer(s grpc.ServiceRegistrar, srv MicroserviceServer) {
	s.RegisterService(&Microservice_ServiceDesc, srv)
}

func _Microservice_GetMark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Marks_GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroserviceServer).GetMark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Microservice/GetMark",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroserviceServer).GetMark(ctx, req.(*Marks_GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Microservice_ServiceDesc is the grpc.ServiceDesc for Microservice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Microservice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Microservice",
	HandlerType: (*MicroserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMark",
			Handler:    _Microservice_GetMark_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "marks.proto",
}