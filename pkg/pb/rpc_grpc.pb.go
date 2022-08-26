// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: rpc.proto

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

// DaemonClient is the client API for Daemon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DaemonClient interface {
	StreamEvents(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Daemon_StreamEventsClient, error)
	UnWait(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Restart(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type daemonClient struct {
	cc grpc.ClientConnInterface
}

func NewDaemonClient(cc grpc.ClientConnInterface) DaemonClient {
	return &daemonClient{cc}
}

func (c *daemonClient) StreamEvents(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Daemon_StreamEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Daemon_ServiceDesc.Streams[0], "/wice.Daemon/StreamEvents", opts...)
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

func (c *daemonClient) UnWait(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/wice.Daemon/UnWait", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/wice.Daemon/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) Restart(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/wice.Daemon/Restart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DaemonServer is the server API for Daemon service.
// All implementations must embed UnimplementedDaemonServer
// for forward compatibility
type DaemonServer interface {
	StreamEvents(*Empty, Daemon_StreamEventsServer) error
	UnWait(context.Context, *Empty) (*Empty, error)
	Stop(context.Context, *Empty) (*Empty, error)
	Restart(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedDaemonServer()
}

// UnimplementedDaemonServer must be embedded to have forward compatible implementations.
type UnimplementedDaemonServer struct {
}

func (UnimplementedDaemonServer) StreamEvents(*Empty, Daemon_StreamEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamEvents not implemented")
}
func (UnimplementedDaemonServer) UnWait(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnWait not implemented")
}
func (UnimplementedDaemonServer) Stop(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedDaemonServer) Restart(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
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

func _Daemon_StreamEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
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
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).UnWait(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wice.Daemon/UnWait",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).UnWait(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wice.Daemon/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).Stop(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wice.Daemon/Restart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).Restart(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Daemon_ServiceDesc is the grpc.ServiceDesc for Daemon service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Daemon_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wice.Daemon",
	HandlerType: (*DaemonServer)(nil),
	Methods: []grpc.MethodDesc{
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
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEvents",
			Handler:       _Daemon_StreamEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpc.proto",
}

// WatcherClient is the client API for Watcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WatcherClient interface {
	Sync(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	GetStatus(ctx context.Context, in *StatusParams, opts ...grpc.CallOption) (*StatusResp, error)
	RemoveInterface(ctx context.Context, in *RemoveInterfaceParams, opts ...grpc.CallOption) (*Empty, error)
	SyncInterfaceConfig(ctx context.Context, in *InterfaceConfigParams, opts ...grpc.CallOption) (*Empty, error)
	AddInterfaceConfig(ctx context.Context, in *InterfaceConfigParams, opts ...grpc.CallOption) (*Empty, error)
	SetInterfaceConfig(ctx context.Context, in *InterfaceConfigParams, opts ...grpc.CallOption) (*Empty, error)
	// For manual signaling backend
	GetSignalingMessage(ctx context.Context, in *GetSignalingMessageParams, opts ...grpc.CallOption) (*GetSignalingMessageResp, error)
	PutSignalingMessage(ctx context.Context, in *PutSignalingMessageParams, opts ...grpc.CallOption) (*Empty, error)
}

type watcherClient struct {
	cc grpc.ClientConnInterface
}

func NewWatcherClient(cc grpc.ClientConnInterface) WatcherClient {
	return &watcherClient{cc}
}

func (c *watcherClient) Sync(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/wice.Watcher/Sync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watcherClient) GetStatus(ctx context.Context, in *StatusParams, opts ...grpc.CallOption) (*StatusResp, error) {
	out := new(StatusResp)
	err := c.cc.Invoke(ctx, "/wice.Watcher/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watcherClient) RemoveInterface(ctx context.Context, in *RemoveInterfaceParams, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/wice.Watcher/RemoveInterface", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watcherClient) SyncInterfaceConfig(ctx context.Context, in *InterfaceConfigParams, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/wice.Watcher/SyncInterfaceConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watcherClient) AddInterfaceConfig(ctx context.Context, in *InterfaceConfigParams, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/wice.Watcher/AddInterfaceConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watcherClient) SetInterfaceConfig(ctx context.Context, in *InterfaceConfigParams, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/wice.Watcher/SetInterfaceConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watcherClient) GetSignalingMessage(ctx context.Context, in *GetSignalingMessageParams, opts ...grpc.CallOption) (*GetSignalingMessageResp, error) {
	out := new(GetSignalingMessageResp)
	err := c.cc.Invoke(ctx, "/wice.Watcher/GetSignalingMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watcherClient) PutSignalingMessage(ctx context.Context, in *PutSignalingMessageParams, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/wice.Watcher/PutSignalingMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WatcherServer is the server API for Watcher service.
// All implementations must embed UnimplementedWatcherServer
// for forward compatibility
type WatcherServer interface {
	Sync(context.Context, *Empty) (*Empty, error)
	GetStatus(context.Context, *StatusParams) (*StatusResp, error)
	RemoveInterface(context.Context, *RemoveInterfaceParams) (*Empty, error)
	SyncInterfaceConfig(context.Context, *InterfaceConfigParams) (*Empty, error)
	AddInterfaceConfig(context.Context, *InterfaceConfigParams) (*Empty, error)
	SetInterfaceConfig(context.Context, *InterfaceConfigParams) (*Empty, error)
	// For manual signaling backend
	GetSignalingMessage(context.Context, *GetSignalingMessageParams) (*GetSignalingMessageResp, error)
	PutSignalingMessage(context.Context, *PutSignalingMessageParams) (*Empty, error)
	mustEmbedUnimplementedWatcherServer()
}

// UnimplementedWatcherServer must be embedded to have forward compatible implementations.
type UnimplementedWatcherServer struct {
}

func (UnimplementedWatcherServer) Sync(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedWatcherServer) GetStatus(context.Context, *StatusParams) (*StatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedWatcherServer) RemoveInterface(context.Context, *RemoveInterfaceParams) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveInterface not implemented")
}
func (UnimplementedWatcherServer) SyncInterfaceConfig(context.Context, *InterfaceConfigParams) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncInterfaceConfig not implemented")
}
func (UnimplementedWatcherServer) AddInterfaceConfig(context.Context, *InterfaceConfigParams) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddInterfaceConfig not implemented")
}
func (UnimplementedWatcherServer) SetInterfaceConfig(context.Context, *InterfaceConfigParams) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetInterfaceConfig not implemented")
}
func (UnimplementedWatcherServer) GetSignalingMessage(context.Context, *GetSignalingMessageParams) (*GetSignalingMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSignalingMessage not implemented")
}
func (UnimplementedWatcherServer) PutSignalingMessage(context.Context, *PutSignalingMessageParams) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutSignalingMessage not implemented")
}
func (UnimplementedWatcherServer) mustEmbedUnimplementedWatcherServer() {}

// UnsafeWatcherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WatcherServer will
// result in compilation errors.
type UnsafeWatcherServer interface {
	mustEmbedUnimplementedWatcherServer()
}

func RegisterWatcherServer(s grpc.ServiceRegistrar, srv WatcherServer) {
	s.RegisterService(&Watcher_ServiceDesc, srv)
}

func _Watcher_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatcherServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wice.Watcher/Sync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatcherServer).Sync(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watcher_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatcherServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wice.Watcher/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatcherServer).GetStatus(ctx, req.(*StatusParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watcher_RemoveInterface_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveInterfaceParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatcherServer).RemoveInterface(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wice.Watcher/RemoveInterface",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatcherServer).RemoveInterface(ctx, req.(*RemoveInterfaceParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watcher_SyncInterfaceConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InterfaceConfigParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatcherServer).SyncInterfaceConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wice.Watcher/SyncInterfaceConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatcherServer).SyncInterfaceConfig(ctx, req.(*InterfaceConfigParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watcher_AddInterfaceConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InterfaceConfigParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatcherServer).AddInterfaceConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wice.Watcher/AddInterfaceConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatcherServer).AddInterfaceConfig(ctx, req.(*InterfaceConfigParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watcher_SetInterfaceConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InterfaceConfigParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatcherServer).SetInterfaceConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wice.Watcher/SetInterfaceConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatcherServer).SetInterfaceConfig(ctx, req.(*InterfaceConfigParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watcher_GetSignalingMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSignalingMessageParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatcherServer).GetSignalingMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wice.Watcher/GetSignalingMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatcherServer).GetSignalingMessage(ctx, req.(*GetSignalingMessageParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watcher_PutSignalingMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutSignalingMessageParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatcherServer).PutSignalingMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wice.Watcher/PutSignalingMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatcherServer).PutSignalingMessage(ctx, req.(*PutSignalingMessageParams))
	}
	return interceptor(ctx, in, info, handler)
}

// Watcher_ServiceDesc is the grpc.ServiceDesc for Watcher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Watcher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wice.Watcher",
	HandlerType: (*WatcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sync",
			Handler:    _Watcher_Sync_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _Watcher_GetStatus_Handler,
		},
		{
			MethodName: "RemoveInterface",
			Handler:    _Watcher_RemoveInterface_Handler,
		},
		{
			MethodName: "SyncInterfaceConfig",
			Handler:    _Watcher_SyncInterfaceConfig_Handler,
		},
		{
			MethodName: "AddInterfaceConfig",
			Handler:    _Watcher_AddInterfaceConfig_Handler,
		},
		{
			MethodName: "SetInterfaceConfig",
			Handler:    _Watcher_SetInterfaceConfig_Handler,
		},
		{
			MethodName: "GetSignalingMessage",
			Handler:    _Watcher_GetSignalingMessage_Handler,
		},
		{
			MethodName: "PutSignalingMessage",
			Handler:    _Watcher_PutSignalingMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

// EndpointDiscoverySocketClient is the client API for EndpointDiscoverySocket service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EndpointDiscoverySocketClient interface {
	RestartPeer(ctx context.Context, in *RestartPeerParams, opts ...grpc.CallOption) (*Empty, error)
}

type endpointDiscoverySocketClient struct {
	cc grpc.ClientConnInterface
}

func NewEndpointDiscoverySocketClient(cc grpc.ClientConnInterface) EndpointDiscoverySocketClient {
	return &endpointDiscoverySocketClient{cc}
}

func (c *endpointDiscoverySocketClient) RestartPeer(ctx context.Context, in *RestartPeerParams, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/wice.EndpointDiscoverySocket/RestartPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndpointDiscoverySocketServer is the server API for EndpointDiscoverySocket service.
// All implementations must embed UnimplementedEndpointDiscoverySocketServer
// for forward compatibility
type EndpointDiscoverySocketServer interface {
	RestartPeer(context.Context, *RestartPeerParams) (*Empty, error)
	mustEmbedUnimplementedEndpointDiscoverySocketServer()
}

// UnimplementedEndpointDiscoverySocketServer must be embedded to have forward compatible implementations.
type UnimplementedEndpointDiscoverySocketServer struct {
}

func (UnimplementedEndpointDiscoverySocketServer) RestartPeer(context.Context, *RestartPeerParams) (*Empty, error) {
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
		FullMethod: "/wice.EndpointDiscoverySocket/RestartPeer",
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
	ServiceName: "wice.EndpointDiscoverySocket",
	HandlerType: (*EndpointDiscoverySocketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RestartPeer",
			Handler:    _EndpointDiscoverySocket_RestartPeer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

// SignalingSocketClient is the client API for SignalingSocket service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignalingSocketClient interface {
	GetSignalingMessage(ctx context.Context, in *GetSignalingMessageParams, opts ...grpc.CallOption) (*GetSignalingMessageResp, error)
	PutSignalingMessage(ctx context.Context, in *PutSignalingMessageParams, opts ...grpc.CallOption) (*Empty, error)
}

type signalingSocketClient struct {
	cc grpc.ClientConnInterface
}

func NewSignalingSocketClient(cc grpc.ClientConnInterface) SignalingSocketClient {
	return &signalingSocketClient{cc}
}

func (c *signalingSocketClient) GetSignalingMessage(ctx context.Context, in *GetSignalingMessageParams, opts ...grpc.CallOption) (*GetSignalingMessageResp, error) {
	out := new(GetSignalingMessageResp)
	err := c.cc.Invoke(ctx, "/wice.SignalingSocket/GetSignalingMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalingSocketClient) PutSignalingMessage(ctx context.Context, in *PutSignalingMessageParams, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/wice.SignalingSocket/PutSignalingMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignalingSocketServer is the server API for SignalingSocket service.
// All implementations must embed UnimplementedSignalingSocketServer
// for forward compatibility
type SignalingSocketServer interface {
	GetSignalingMessage(context.Context, *GetSignalingMessageParams) (*GetSignalingMessageResp, error)
	PutSignalingMessage(context.Context, *PutSignalingMessageParams) (*Empty, error)
	mustEmbedUnimplementedSignalingSocketServer()
}

// UnimplementedSignalingSocketServer must be embedded to have forward compatible implementations.
type UnimplementedSignalingSocketServer struct {
}

func (UnimplementedSignalingSocketServer) GetSignalingMessage(context.Context, *GetSignalingMessageParams) (*GetSignalingMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSignalingMessage not implemented")
}
func (UnimplementedSignalingSocketServer) PutSignalingMessage(context.Context, *PutSignalingMessageParams) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutSignalingMessage not implemented")
}
func (UnimplementedSignalingSocketServer) mustEmbedUnimplementedSignalingSocketServer() {}

// UnsafeSignalingSocketServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignalingSocketServer will
// result in compilation errors.
type UnsafeSignalingSocketServer interface {
	mustEmbedUnimplementedSignalingSocketServer()
}

func RegisterSignalingSocketServer(s grpc.ServiceRegistrar, srv SignalingSocketServer) {
	s.RegisterService(&SignalingSocket_ServiceDesc, srv)
}

func _SignalingSocket_GetSignalingMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSignalingMessageParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalingSocketServer).GetSignalingMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wice.SignalingSocket/GetSignalingMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalingSocketServer).GetSignalingMessage(ctx, req.(*GetSignalingMessageParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalingSocket_PutSignalingMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutSignalingMessageParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalingSocketServer).PutSignalingMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wice.SignalingSocket/PutSignalingMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalingSocketServer).PutSignalingMessage(ctx, req.(*PutSignalingMessageParams))
	}
	return interceptor(ctx, in, info, handler)
}

// SignalingSocket_ServiceDesc is the grpc.ServiceDesc for SignalingSocket service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SignalingSocket_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wice.SignalingSocket",
	HandlerType: (*SignalingSocketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSignalingMessage",
			Handler:    _SignalingSocket_GetSignalingMessage_Handler,
		},
		{
			MethodName: "PutSignalingMessage",
			Handler:    _SignalingSocket_PutSignalingMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}
