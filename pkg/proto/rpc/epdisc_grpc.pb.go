// SPDX-FileCopyrightText: 2023 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.4
// source: rpc/epdisc.proto

package rpc

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
	EndpointDiscoverySocket_RestartPeer_FullMethodName = "/cunicu.rpc.EndpointDiscoverySocket/RestartPeer"
)

// EndpointDiscoverySocketClient is the client API for EndpointDiscoverySocket service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EndpointDiscoverySocketClient interface {
	RestartPeer(ctx context.Context, in *RestartPeerParams, opts ...grpc.CallOption) (*proto.Empty, error)
}

type endpointDiscoverySocketClient struct {
	cc grpc.ClientConnInterface
}

func NewEndpointDiscoverySocketClient(cc grpc.ClientConnInterface) EndpointDiscoverySocketClient {
	return &endpointDiscoverySocketClient{cc}
}

func (c *endpointDiscoverySocketClient) RestartPeer(ctx context.Context, in *RestartPeerParams, opts ...grpc.CallOption) (*proto.Empty, error) {
	out := new(proto.Empty)
	err := c.cc.Invoke(ctx, EndpointDiscoverySocket_RestartPeer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndpointDiscoverySocketServer is the server API for EndpointDiscoverySocket service.
// All implementations must embed UnimplementedEndpointDiscoverySocketServer
// for forward compatibility
type EndpointDiscoverySocketServer interface {
	RestartPeer(context.Context, *RestartPeerParams) (*proto.Empty, error)
	mustEmbedUnimplementedEndpointDiscoverySocketServer()
}

// UnimplementedEndpointDiscoverySocketServer must be embedded to have forward compatible implementations.
type UnimplementedEndpointDiscoverySocketServer struct {
}

func (UnimplementedEndpointDiscoverySocketServer) RestartPeer(context.Context, *RestartPeerParams) (*proto.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestartPeer not implemented")
}
func (UnimplementedEndpointDiscoverySocketServer) mustEmbedUnimplementedEndpointDiscoverySocketServer() {
}

// UnsafeEndpointDiscoverySocketServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EndpointDiscoverySocketServer will
// result in compilation errors.
type UnsafeEndpointDiscoverySocketServer interface {
	mustEmbedUnimplementedEndpointDiscoverySocketServer()
}

func RegisterEndpointDiscoverySocketServer(s grpc.ServiceRegistrar, srv EndpointDiscoverySocketServer) {
	s.RegisterService(&EndpointDiscoverySocket_ServiceDesc, srv)
}

func _EndpointDiscoverySocket_RestartPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestartPeerParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointDiscoverySocketServer).RestartPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EndpointDiscoverySocket_RestartPeer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointDiscoverySocketServer).RestartPeer(ctx, req.(*RestartPeerParams))
	}
	return interceptor(ctx, in, info, handler)
}

// EndpointDiscoverySocket_ServiceDesc is the grpc.ServiceDesc for EndpointDiscoverySocket service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EndpointDiscoverySocket_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cunicu.rpc.EndpointDiscoverySocket",
	HandlerType: (*EndpointDiscoverySocketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RestartPeer",
			Handler:    _EndpointDiscoverySocket_RestartPeer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/epdisc.proto",
}
