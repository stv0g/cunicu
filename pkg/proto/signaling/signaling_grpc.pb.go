// SPDX-FileCopyrightText: 2023 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: signaling/signaling.proto

package signaling

import (
	context "context"
	proto "cunicu.li/cunicu/pkg/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Signaling_GetBuildInfo_FullMethodName = "/cunicu.signaling.Signaling/GetBuildInfo"
	Signaling_Subscribe_FullMethodName    = "/cunicu.signaling.Signaling/Subscribe"
	Signaling_Publish_FullMethodName      = "/cunicu.signaling.Signaling/Publish"
)

// SignalingClient is the client API for Signaling service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignalingClient interface {
	GetBuildInfo(ctx context.Context, in *proto.Empty, opts ...grpc.CallOption) (*proto.BuildInfo, error)
	Subscribe(ctx context.Context, in *SubscribeParams, opts ...grpc.CallOption) (Signaling_SubscribeClient, error)
	Publish(ctx context.Context, in *Envelope, opts ...grpc.CallOption) (*proto.Empty, error)
}

type signalingClient struct {
	cc grpc.ClientConnInterface
}

func NewSignalingClient(cc grpc.ClientConnInterface) SignalingClient {
	return &signalingClient{cc}
}

func (c *signalingClient) GetBuildInfo(ctx context.Context, in *proto.Empty, opts ...grpc.CallOption) (*proto.BuildInfo, error) {
	out := new(proto.BuildInfo)
	err := c.cc.Invoke(ctx, Signaling_GetBuildInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalingClient) Subscribe(ctx context.Context, in *SubscribeParams, opts ...grpc.CallOption) (Signaling_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Signaling_ServiceDesc.Streams[0], Signaling_Subscribe_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &signalingSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Signaling_SubscribeClient interface {
	Recv() (*Envelope, error)
	grpc.ClientStream
}

type signalingSubscribeClient struct {
	grpc.ClientStream
}

func (x *signalingSubscribeClient) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *signalingClient) Publish(ctx context.Context, in *Envelope, opts ...grpc.CallOption) (*proto.Empty, error) {
	out := new(proto.Empty)
	err := c.cc.Invoke(ctx, Signaling_Publish_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignalingServer is the server API for Signaling service.
// All implementations must embed UnimplementedSignalingServer
// for forward compatibility
type SignalingServer interface {
	GetBuildInfo(context.Context, *proto.Empty) (*proto.BuildInfo, error)
	Subscribe(*SubscribeParams, Signaling_SubscribeServer) error
	Publish(context.Context, *Envelope) (*proto.Empty, error)
	mustEmbedUnimplementedSignalingServer()
}

// UnimplementedSignalingServer must be embedded to have forward compatible implementations.
type UnimplementedSignalingServer struct {
}

func (UnimplementedSignalingServer) GetBuildInfo(context.Context, *proto.Empty) (*proto.BuildInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBuildInfo not implemented")
}
func (UnimplementedSignalingServer) Subscribe(*SubscribeParams, Signaling_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedSignalingServer) Publish(context.Context, *Envelope) (*proto.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedSignalingServer) mustEmbedUnimplementedSignalingServer() {}

// UnsafeSignalingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignalingServer will
// result in compilation errors.
type UnsafeSignalingServer interface {
	mustEmbedUnimplementedSignalingServer()
}

func RegisterSignalingServer(s grpc.ServiceRegistrar, srv SignalingServer) {
	s.RegisterService(&Signaling_ServiceDesc, srv)
}

func _Signaling_GetBuildInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalingServer).GetBuildInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Signaling_GetBuildInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalingServer).GetBuildInfo(ctx, req.(*proto.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signaling_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SignalingServer).Subscribe(m, &signalingSubscribeServer{stream})
}

type Signaling_SubscribeServer interface {
	Send(*Envelope) error
	grpc.ServerStream
}

type signalingSubscribeServer struct {
	grpc.ServerStream
}

func (x *signalingSubscribeServer) Send(m *Envelope) error {
	return x.ServerStream.SendMsg(m)
}

func _Signaling_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Envelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalingServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Signaling_Publish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalingServer).Publish(ctx, req.(*Envelope))
	}
	return interceptor(ctx, in, info, handler)
}

// Signaling_ServiceDesc is the grpc.ServiceDesc for Signaling service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Signaling_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cunicu.signaling.Signaling",
	HandlerType: (*SignalingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBuildInfo",
			Handler:    _Signaling_GetBuildInfo_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _Signaling_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Signaling_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "signaling/signaling.proto",
}
