// SPDX-FileCopyrightText: 2023 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: rpc/daemon.proto

package rpc

import (
	context "context"
	proto "github.com/stv0g/cunicu/pkg/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Daemon_GetBuildInfo_FullMethodName = "/cunicu.rpc.Daemon/GetBuildInfo"
	Daemon_StreamEvents_FullMethodName = "/cunicu.rpc.Daemon/StreamEvents"
	Daemon_UnWait_FullMethodName       = "/cunicu.rpc.Daemon/UnWait"
	Daemon_Stop_FullMethodName         = "/cunicu.rpc.Daemon/Stop"
	Daemon_Restart_FullMethodName      = "/cunicu.rpc.Daemon/Restart"
	Daemon_Sync_FullMethodName         = "/cunicu.rpc.Daemon/Sync"
	Daemon_GetStatus_FullMethodName    = "/cunicu.rpc.Daemon/GetStatus"
	Daemon_SetConfig_FullMethodName    = "/cunicu.rpc.Daemon/SetConfig"
	Daemon_GetConfig_FullMethodName    = "/cunicu.rpc.Daemon/GetConfig"
	Daemon_ReloadConfig_FullMethodName = "/cunicu.rpc.Daemon/ReloadConfig"
	Daemon_AddPeer_FullMethodName      = "/cunicu.rpc.Daemon/AddPeer"
)

// DaemonClient is the client API for Daemon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DaemonClient interface {
	GetBuildInfo(ctx context.Context, in *proto.Empty, opts ...grpc.CallOption) (*proto.BuildInfo, error)
	StreamEvents(ctx context.Context, in *proto.Empty, opts ...grpc.CallOption) (Daemon_StreamEventsClient, error)
	UnWait(ctx context.Context, in *proto.Empty, opts ...grpc.CallOption) (*proto.Empty, error)
	Stop(ctx context.Context, in *proto.Empty, opts ...grpc.CallOption) (*proto.Empty, error)
	Restart(ctx context.Context, in *proto.Empty, opts ...grpc.CallOption) (*proto.Empty, error)
	Sync(ctx context.Context, in *proto.Empty, opts ...grpc.CallOption) (*proto.Empty, error)
	GetStatus(ctx context.Context, in *GetStatusParams, opts ...grpc.CallOption) (*GetStatusResp, error)
	SetConfig(ctx context.Context, in *SetConfigParams, opts ...grpc.CallOption) (*proto.Empty, error)
	GetConfig(ctx context.Context, in *GetConfigParams, opts ...grpc.CallOption) (*GetConfigResp, error)
	ReloadConfig(ctx context.Context, in *proto.Empty, opts ...grpc.CallOption) (*proto.Empty, error)
	AddPeer(ctx context.Context, in *AddPeerParams, opts ...grpc.CallOption) (*AddPeerResp, error)
}

type daemonClient struct {
	cc grpc.ClientConnInterface
}

func NewDaemonClient(cc grpc.ClientConnInterface) DaemonClient {
	return &daemonClient{cc}
}

func (c *daemonClient) GetBuildInfo(ctx context.Context, in *proto.Empty, opts ...grpc.CallOption) (*proto.BuildInfo, error) {
	out := new(proto.BuildInfo)
	err := c.cc.Invoke(ctx, Daemon_GetBuildInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) StreamEvents(ctx context.Context, in *proto.Empty, opts ...grpc.CallOption) (Daemon_StreamEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Daemon_ServiceDesc.Streams[0], Daemon_StreamEvents_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &daemonStreamEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Daemon_StreamEventsClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type daemonStreamEventsClient struct {
	grpc.ClientStream
}

func (x *daemonStreamEventsClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *daemonClient) UnWait(ctx context.Context, in *proto.Empty, opts ...grpc.CallOption) (*proto.Empty, error) {
	out := new(proto.Empty)
	err := c.cc.Invoke(ctx, Daemon_UnWait_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) Stop(ctx context.Context, in *proto.Empty, opts ...grpc.CallOption) (*proto.Empty, error) {
	out := new(proto.Empty)
	err := c.cc.Invoke(ctx, Daemon_Stop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) Restart(ctx context.Context, in *proto.Empty, opts ...grpc.CallOption) (*proto.Empty, error) {
	out := new(proto.Empty)
	err := c.cc.Invoke(ctx, Daemon_Restart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) Sync(ctx context.Context, in *proto.Empty, opts ...grpc.CallOption) (*proto.Empty, error) {
	out := new(proto.Empty)
	err := c.cc.Invoke(ctx, Daemon_Sync_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) GetStatus(ctx context.Context, in *GetStatusParams, opts ...grpc.CallOption) (*GetStatusResp, error) {
	out := new(GetStatusResp)
	err := c.cc.Invoke(ctx, Daemon_GetStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) SetConfig(ctx context.Context, in *SetConfigParams, opts ...grpc.CallOption) (*proto.Empty, error) {
	out := new(proto.Empty)
	err := c.cc.Invoke(ctx, Daemon_SetConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) GetConfig(ctx context.Context, in *GetConfigParams, opts ...grpc.CallOption) (*GetConfigResp, error) {
	out := new(GetConfigResp)
	err := c.cc.Invoke(ctx, Daemon_GetConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) ReloadConfig(ctx context.Context, in *proto.Empty, opts ...grpc.CallOption) (*proto.Empty, error) {
	out := new(proto.Empty)
	err := c.cc.Invoke(ctx, Daemon_ReloadConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) AddPeer(ctx context.Context, in *AddPeerParams, opts ...grpc.CallOption) (*AddPeerResp, error) {
	out := new(AddPeerResp)
	err := c.cc.Invoke(ctx, Daemon_AddPeer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DaemonServer is the server API for Daemon service.
// All implementations must embed UnimplementedDaemonServer
// for forward compatibility
type DaemonServer interface {
	GetBuildInfo(context.Context, *proto.Empty) (*proto.BuildInfo, error)
	StreamEvents(*proto.Empty, Daemon_StreamEventsServer) error
	UnWait(context.Context, *proto.Empty) (*proto.Empty, error)
	Stop(context.Context, *proto.Empty) (*proto.Empty, error)
	Restart(context.Context, *proto.Empty) (*proto.Empty, error)
	Sync(context.Context, *proto.Empty) (*proto.Empty, error)
	GetStatus(context.Context, *GetStatusParams) (*GetStatusResp, error)
	SetConfig(context.Context, *SetConfigParams) (*proto.Empty, error)
	GetConfig(context.Context, *GetConfigParams) (*GetConfigResp, error)
	ReloadConfig(context.Context, *proto.Empty) (*proto.Empty, error)
	AddPeer(context.Context, *AddPeerParams) (*AddPeerResp, error)
	mustEmbedUnimplementedDaemonServer()
}

// UnimplementedDaemonServer must be embedded to have forward compatible implementations.
type UnimplementedDaemonServer struct {
}

func (UnimplementedDaemonServer) GetBuildInfo(context.Context, *proto.Empty) (*proto.BuildInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBuildInfo not implemented")
}
func (UnimplementedDaemonServer) StreamEvents(*proto.Empty, Daemon_StreamEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamEvents not implemented")
}
func (UnimplementedDaemonServer) UnWait(context.Context, *proto.Empty) (*proto.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnWait not implemented")
}
func (UnimplementedDaemonServer) Stop(context.Context, *proto.Empty) (*proto.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedDaemonServer) Restart(context.Context, *proto.Empty) (*proto.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (UnimplementedDaemonServer) Sync(context.Context, *proto.Empty) (*proto.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedDaemonServer) GetStatus(context.Context, *GetStatusParams) (*GetStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedDaemonServer) SetConfig(context.Context, *SetConfigParams) (*proto.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConfig not implemented")
}
func (UnimplementedDaemonServer) GetConfig(context.Context, *GetConfigParams) (*GetConfigResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedDaemonServer) ReloadConfig(context.Context, *proto.Empty) (*proto.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReloadConfig not implemented")
}
func (UnimplementedDaemonServer) AddPeer(context.Context, *AddPeerParams) (*AddPeerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPeer not implemented")
}
func (UnimplementedDaemonServer) mustEmbedUnimplementedDaemonServer() {}

// UnsafeDaemonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DaemonServer will
// result in compilation errors.
type UnsafeDaemonServer interface {
	mustEmbedUnimplementedDaemonServer()
}

func RegisterDaemonServer(s grpc.ServiceRegistrar, srv DaemonServer) {
	s.RegisterService(&Daemon_ServiceDesc, srv)
}

func _Daemon_GetBuildInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).GetBuildInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Daemon_GetBuildInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).GetBuildInfo(ctx, req.(*proto.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_StreamEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(proto.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DaemonServer).StreamEvents(m, &daemonStreamEventsServer{stream})
}

type Daemon_StreamEventsServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type daemonStreamEventsServer struct {
	grpc.ServerStream
}

func (x *daemonStreamEventsServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _Daemon_UnWait_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).UnWait(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Daemon_UnWait_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).UnWait(ctx, req.(*proto.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Daemon_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).Stop(ctx, req.(*proto.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Daemon_Restart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).Restart(ctx, req.(*proto.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Daemon_Sync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).Sync(ctx, req.(*proto.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Daemon_GetStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).GetStatus(ctx, req.(*GetStatusParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_SetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetConfigParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).SetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Daemon_SetConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).SetConfig(ctx, req.(*SetConfigParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Daemon_GetConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).GetConfig(ctx, req.(*GetConfigParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_ReloadConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).ReloadConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Daemon_ReloadConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).ReloadConfig(ctx, req.(*proto.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_AddPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPeerParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).AddPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Daemon_AddPeer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).AddPeer(ctx, req.(*AddPeerParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Daemon_ServiceDesc is the grpc.ServiceDesc for Daemon service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Daemon_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cunicu.rpc.Daemon",
	HandlerType: (*DaemonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBuildInfo",
			Handler:    _Daemon_GetBuildInfo_Handler,
		},
		{
			MethodName: "UnWait",
			Handler:    _Daemon_UnWait_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Daemon_Stop_Handler,
		},
		{
			MethodName: "Restart",
			Handler:    _Daemon_Restart_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _Daemon_Sync_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _Daemon_GetStatus_Handler,
		},
		{
			MethodName: "SetConfig",
			Handler:    _Daemon_SetConfig_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _Daemon_GetConfig_Handler,
		},
		{
			MethodName: "ReloadConfig",
			Handler:    _Daemon_ReloadConfig_Handler,
		},
		{
			MethodName: "AddPeer",
			Handler:    _Daemon_AddPeer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEvents",
			Handler:       _Daemon_StreamEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpc/daemon.proto",
}
