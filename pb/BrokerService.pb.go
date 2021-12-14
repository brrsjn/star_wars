// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: pb/BrokerService.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Instruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Lectura bool   `protobuf:"varint,2,opt,name=lectura,proto3" json:"lectura,omitempty"`
}

func (x *Instruct) Reset() {
	*x = Instruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_BrokerService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Instruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instruct) ProtoMessage() {}

func (x *Instruct) ProtoReflect() protoreflect.Message {
	mi := &file_pb_BrokerService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instruct.ProtoReflect.Descriptor instead.
func (*Instruct) Descriptor() ([]byte, []int) {
	return file_pb_BrokerService_proto_rawDescGZIP(), []int{0}
}

func (x *Instruct) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Instruct) GetLectura() bool {
	if x != nil {
		return x.Lectura
	}
	return false
}

type Servidor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addres string `protobuf:"bytes,1,opt,name=addres,proto3" json:"addres,omitempty"`
	Error  bool   `protobuf:"varint,2,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *Servidor) Reset() {
	*x = Servidor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_BrokerService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Servidor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Servidor) ProtoMessage() {}

func (x *Servidor) ProtoReflect() protoreflect.Message {
	mi := &file_pb_BrokerService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Servidor.ProtoReflect.Descriptor instead.
func (*Servidor) Descriptor() ([]byte, []int) {
	return file_pb_BrokerService_proto_rawDescGZIP(), []int{1}
}

func (x *Servidor) GetAddres() string {
	if x != nil {
		return x.Addres
	}
	return ""
}

func (x *Servidor) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

var File_pb_BrokerService_proto protoreflect.FileDescriptor

var file_pb_BrokerService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x62, 0x2f, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x61, 0x22, 0x38, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x64, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x32, 0x33, 0x0a, 0x06, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x0f,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x09, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x09, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x64, 0x6f, 0x72, 0x22, 0x00, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_BrokerService_proto_rawDescOnce sync.Once
	file_pb_BrokerService_proto_rawDescData = file_pb_BrokerService_proto_rawDesc
)

func file_pb_BrokerService_proto_rawDescGZIP() []byte {
	file_pb_BrokerService_proto_rawDescOnce.Do(func() {
		file_pb_BrokerService_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_BrokerService_proto_rawDescData)
	})
	return file_pb_BrokerService_proto_rawDescData
}

var file_pb_BrokerService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_BrokerService_proto_goTypes = []interface{}{
	(*Instruct)(nil), // 0: instruct
	(*Servidor)(nil), // 1: Servidor
}
var file_pb_BrokerService_proto_depIdxs = []int32{
	0, // 0: Broker.ConnectToServer:input_type -> instruct
	1, // 1: Broker.ConnectToServer:output_type -> Servidor
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_BrokerService_proto_init() }
func file_pb_BrokerService_proto_init() {
	if File_pb_BrokerService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_BrokerService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Instruct); i {
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
		file_pb_BrokerService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Servidor); i {
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
			RawDescriptor: file_pb_BrokerService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_BrokerService_proto_goTypes,
		DependencyIndexes: file_pb_BrokerService_proto_depIdxs,
		MessageInfos:      file_pb_BrokerService_proto_msgTypes,
	}.Build()
	File_pb_BrokerService_proto = out.File
	file_pb_BrokerService_proto_rawDesc = nil
	file_pb_BrokerService_proto_goTypes = nil
	file_pb_BrokerService_proto_depIdxs = nil
}
