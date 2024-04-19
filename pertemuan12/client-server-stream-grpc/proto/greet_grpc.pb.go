// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: greet.proto

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

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetServiceClient interface {
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponses, error)
	GreetStream(ctx context.Context, in *GreetStreamRequest, opts ...grpc.CallOption) (GreetService_GreetStreamClient, error)
	StoreStream(ctx context.Context, opts ...grpc.CallOption) (GreetService_StoreStreamClient, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponses, error) {
	out := new(GreetResponses)
	err := c.cc.Invoke(ctx, "/greet.GreetService/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) GreetStream(ctx context.Context, in *GreetStreamRequest, opts ...grpc.CallOption) (GreetService_GreetStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[0], "/greet.GreetService/GreetStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_GreetStreamClient interface {
	Recv() (*GreetStreamResponses, error)
	grpc.ClientStream
}

type greetServiceGreetStreamClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetStreamClient) Recv() (*GreetStreamResponses, error) {
	m := new(GreetStreamResponses)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) StoreStream(ctx context.Context, opts ...grpc.CallOption) (GreetService_StoreStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[1], "/greet.GreetService/StoreStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceStoreStreamClient{stream}
	return x, nil
}

type GreetService_StoreStreamClient interface {
	Send(*GreetStreamRequest) error
	Recv() (*StoreStreamResponses, error)
	grpc.ClientStream
}

type greetServiceStoreStreamClient struct {
	grpc.ClientStream
}

func (x *greetServiceStoreStreamClient) Send(m *GreetStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceStoreStreamClient) Recv() (*StoreStreamResponses, error) {
	m := new(StoreStreamResponses)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
// All implementations must embed UnimplementedGreetServiceServer
// for forward compatibility
type GreetServiceServer interface {
	Greet(context.Context, *GreetRequest) (*GreetResponses, error)
	GreetStream(*GreetStreamRequest, GreetService_GreetStreamServer) error
	StoreStream(GreetService_StoreStreamServer) error
	mustEmbedUnimplementedGreetServiceServer()
}

// UnimplementedGreetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (UnimplementedGreetServiceServer) Greet(context.Context, *GreetRequest) (*GreetResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (UnimplementedGreetServiceServer) GreetStream(*GreetStreamRequest, GreetService_GreetStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetStream not implemented")
}
func (UnimplementedGreetServiceServer) StoreStream(GreetService_StoreStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method StoreStream not implemented")
}
func (UnimplementedGreetServiceServer) mustEmbedUnimplementedGreetServiceServer() {}

// UnsafeGreetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServiceServer will
// result in compilation errors.
type UnsafeGreetServiceServer interface {
	mustEmbedUnimplementedGreetServiceServer()
}

func RegisterGreetServiceServer(s grpc.ServiceRegistrar, srv GreetServiceServer) {
	s.RegisterService(&GreetService_ServiceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_GreetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).GreetStream(m, &greetServiceGreetStreamServer{stream})
}

type GreetService_GreetStreamServer interface {
	Send(*GreetStreamResponses) error
	grpc.ServerStream
}

type greetServiceGreetStreamServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetStreamServer) Send(m *GreetStreamResponses) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetService_StoreStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).StoreStream(&greetServiceStoreStreamServer{stream})
}

type GreetService_StoreStreamServer interface {
	Send(*StoreStreamResponses) error
	Recv() (*GreetStreamRequest, error)
	grpc.ServerStream
}

type greetServiceStoreStreamServer struct {
	grpc.ServerStream
}

func (x *greetServiceStoreStreamServer) Send(m *StoreStreamResponses) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceStoreStreamServer) Recv() (*GreetStreamRequest, error) {
	m := new(GreetStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetService_ServiceDesc is the grpc.ServiceDesc for GreetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetService_Greet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetStream",
			Handler:       _GreetService_GreetStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StoreStream",
			Handler:       _GreetService_StoreStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "greet.proto",
}
