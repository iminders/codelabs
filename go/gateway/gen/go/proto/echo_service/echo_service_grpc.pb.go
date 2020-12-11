// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package echo_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// EchoServiceClient is the client API for EchoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EchoServiceClient interface {
	Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
	Downfile(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (EchoService_DownfileClient, error)
}

type echoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoServiceClient(cc grpc.ClientConnInterface) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	out := new(StringMessage)
	err := c.cc.Invoke(ctx, "/echo.service.v1.EchoService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) Downfile(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (EchoService_DownfileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EchoService_serviceDesc.Streams[0], "/echo.service.v1.EchoService/Downfile", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServiceDownfileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EchoService_DownfileClient interface {
	Recv() (*FileBinary, error)
	grpc.ClientStream
}

type echoServiceDownfileClient struct {
	grpc.ClientStream
}

func (x *echoServiceDownfileClient) Recv() (*FileBinary, error) {
	m := new(FileBinary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EchoServiceServer is the server API for EchoService service.
// All implementations must embed UnimplementedEchoServiceServer
// for forward compatibility
type EchoServiceServer interface {
	Echo(context.Context, *StringMessage) (*StringMessage, error)
	Downfile(*StringMessage, EchoService_DownfileServer) error
	mustEmbedUnimplementedEchoServiceServer()
}

// UnimplementedEchoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEchoServiceServer struct {
}

func (UnimplementedEchoServiceServer) Echo(context.Context, *StringMessage) (*StringMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedEchoServiceServer) Downfile(*StringMessage, EchoService_DownfileServer) error {
	return status.Errorf(codes.Unimplemented, "method Downfile not implemented")
}
func (UnimplementedEchoServiceServer) mustEmbedUnimplementedEchoServiceServer() {}

// UnsafeEchoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EchoServiceServer will
// result in compilation errors.
type UnsafeEchoServiceServer interface {
	mustEmbedUnimplementedEchoServiceServer()
}

func RegisterEchoServiceServer(s grpc.ServiceRegistrar, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.service.v1.EchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).Echo(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_Downfile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StringMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EchoServiceServer).Downfile(m, &echoServiceDownfileServer{stream})
}

type EchoService_DownfileServer interface {
	Send(*FileBinary) error
	grpc.ServerStream
}

type echoServiceDownfileServer struct {
	grpc.ServerStream
}

func (x *echoServiceDownfileServer) Send(m *FileBinary) error {
	return x.ServerStream.SendMsg(m)
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echo.service.v1.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoService_Echo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Downfile",
			Handler:       _EchoService_Downfile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/echo_service/echo_service.proto",
}
