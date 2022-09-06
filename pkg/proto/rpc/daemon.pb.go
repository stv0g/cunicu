// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: rpc/daemon.proto

package rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	proto "riasc.eu/wice/pkg/proto"
	core "riasc.eu/wice/pkg/proto/core"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StatusResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interfaces []*core.Interface `protobuf:"bytes,1,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
}

func (x *StatusResp) Reset() {
	*x = StatusResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_daemon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResp) ProtoMessage() {}

func (x *StatusResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_daemon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResp.ProtoReflect.Descriptor instead.
func (*StatusResp) Descriptor() ([]byte, []int) {
	return file_rpc_daemon_proto_rawDescGZIP(), []int{0}
}

func (x *StatusResp) GetInterfaces() []*core.Interface {
	if x != nil {
		return x.Interfaces
	}
	return nil
}

type StatusParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Intf string `protobuf:"bytes,1,opt,name=intf,proto3" json:"intf,omitempty"`
	Peer []byte `protobuf:"bytes,2,opt,name=peer,proto3" json:"peer,omitempty"`
}

func (x *StatusParams) Reset() {
	*x = StatusParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_daemon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusParams) ProtoMessage() {}

func (x *StatusParams) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_daemon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusParams.ProtoReflect.Descriptor instead.
func (*StatusParams) Descriptor() ([]byte, []int) {
	return file_rpc_daemon_proto_rawDescGZIP(), []int{1}
}

func (x *StatusParams) GetIntf() string {
	if x != nil {
		return x.Intf
	}
	return ""
}

func (x *StatusParams) GetPeer() []byte {
	if x != nil {
		return x.Peer
	}
	return nil
}

var File_rpc_daemon_proto protoreflect.FileDescriptor

var file_rpc_daemon_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x1a, 0x14, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0f, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x42, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x34, 0x0a, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x73, 0x22, 0x36, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x74, 0x66, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x74, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x65, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x32, 0xbc, 0x02,
	0x0a, 0x06, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x24, 0x0a, 0x06, 0x55, 0x6e,
	0x57, 0x61, 0x69, 0x74, 0x12, 0x0b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x0b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x22, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x0b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x0b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0b, 0x2e, 0x77,
	0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x04, 0x53,
	0x79, 0x6e, 0x63, 0x12, 0x0b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x3b, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x77,
	0x69, 0x63, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x1a, 0x14, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b,
	0x72, 0x69, 0x61, 0x73, 0x63, 0x2e, 0x65, 0x75, 0x2f, 0x77, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_rpc_daemon_proto_rawDescOnce sync.Once
	file_rpc_daemon_proto_rawDescData = file_rpc_daemon_proto_rawDesc
)

func file_rpc_daemon_proto_rawDescGZIP() []byte {
	file_rpc_daemon_proto_rawDescOnce.Do(func() {
		file_rpc_daemon_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_daemon_proto_rawDescData)
	})
	return file_rpc_daemon_proto_rawDescData
}

var file_rpc_daemon_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_daemon_proto_goTypes = []interface{}{
	(*StatusResp)(nil),      // 0: wice.rpc.StatusResp
	(*StatusParams)(nil),    // 1: wice.rpc.StatusParams
	(*core.Interface)(nil),  // 2: wice.core.Interface
	(*proto.Empty)(nil),     // 3: wice.Empty
	(*proto.BuildInfo)(nil), // 4: wice.BuildInfo
	(*Event)(nil),           // 5: wice.rpc.Event
}
var file_rpc_daemon_proto_depIdxs = []int32{
	2, // 0: wice.rpc.StatusResp.interfaces:type_name -> wice.core.Interface
	3, // 1: wice.rpc.Daemon.GetBuildInfo:input_type -> wice.Empty
	3, // 2: wice.rpc.Daemon.StreamEvents:input_type -> wice.Empty
	3, // 3: wice.rpc.Daemon.UnWait:input_type -> wice.Empty
	3, // 4: wice.rpc.Daemon.Stop:input_type -> wice.Empty
	3, // 5: wice.rpc.Daemon.Restart:input_type -> wice.Empty
	3, // 6: wice.rpc.Daemon.Sync:input_type -> wice.Empty
	1, // 7: wice.rpc.Daemon.GetStatus:input_type -> wice.rpc.StatusParams
	4, // 8: wice.rpc.Daemon.GetBuildInfo:output_type -> wice.BuildInfo
	5, // 9: wice.rpc.Daemon.StreamEvents:output_type -> wice.rpc.Event
	3, // 10: wice.rpc.Daemon.UnWait:output_type -> wice.Empty
	3, // 11: wice.rpc.Daemon.Stop:output_type -> wice.Empty
	3, // 12: wice.rpc.Daemon.Restart:output_type -> wice.Empty
	3, // 13: wice.rpc.Daemon.Sync:output_type -> wice.Empty
	0, // 14: wice.rpc.Daemon.GetStatus:output_type -> wice.rpc.StatusResp
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_daemon_proto_init() }
func file_rpc_daemon_proto_init() {
	if File_rpc_daemon_proto != nil {
		return
	}
	file_rpc_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_daemon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_daemon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_daemon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_daemon_proto_goTypes,
		DependencyIndexes: file_rpc_daemon_proto_depIdxs,
		MessageInfos:      file_rpc_daemon_proto_msgTypes,
	}.Build()
	File_rpc_daemon_proto = out.File
	file_rpc_daemon_proto_rawDesc = nil
	file_rpc_daemon_proto_goTypes = nil
	file_rpc_daemon_proto_depIdxs = nil
}
