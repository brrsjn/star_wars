// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: pb/InformerToBroker.proto

package protos

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

type CityNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planet    string `protobuf:"bytes,1,opt,name=Planet,proto3" json:"Planet,omitempty"`
	City      string `protobuf:"bytes,2,opt,name=City,proto3" json:"City,omitempty"`
	Survivors int32  `protobuf:"varint,3,opt,name=Survivors,proto3" json:"Survivors,omitempty"`
}

func (x *CityNumber) Reset() {
	*x = CityNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_InformerToBroker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityNumber) ProtoMessage() {}

func (x *CityNumber) ProtoReflect() protoreflect.Message {
	mi := &file_pb_InformerToBroker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityNumber.ProtoReflect.Descriptor instead.
func (*CityNumber) Descriptor() ([]byte, []int) {
	return file_pb_InformerToBroker_proto_rawDescGZIP(), []int{0}
}

func (x *CityNumber) GetPlanet() string {
	if x != nil {
		return x.Planet
	}
	return ""
}

func (x *CityNumber) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *CityNumber) GetSurvivors() int32 {
	if x != nil {
		return x.Survivors
	}
	return 0
}

type CityName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planet  string `protobuf:"bytes,1,opt,name=Planet,proto3" json:"Planet,omitempty"`
	City    string `protobuf:"bytes,2,opt,name=City,proto3" json:"City,omitempty"`
	NewName string `protobuf:"bytes,3,opt,name=NewName,proto3" json:"NewName,omitempty"`
}

func (x *CityName) Reset() {
	*x = CityName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_InformerToBroker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityName) ProtoMessage() {}

func (x *CityName) ProtoReflect() protoreflect.Message {
	mi := &file_pb_InformerToBroker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityName.ProtoReflect.Descriptor instead.
func (*CityName) Descriptor() ([]byte, []int) {
	return file_pb_InformerToBroker_proto_rawDescGZIP(), []int{1}
}

func (x *CityName) GetPlanet() string {
	if x != nil {
		return x.Planet
	}
	return ""
}

func (x *CityName) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *CityName) GetNewName() string {
	if x != nil {
		return x.NewName
	}
	return ""
}

type CityDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planet    string `protobuf:"bytes,1,opt,name=Planet,proto3" json:"Planet,omitempty"`
	City      string `protobuf:"bytes,2,opt,name=City,proto3" json:"City,omitempty"`
	Survivors int32  `protobuf:"varint,3,opt,name=Survivors,proto3" json:"Survivors,omitempty"`
}

func (x *CityDelete) Reset() {
	*x = CityDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_InformerToBroker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityDelete) ProtoMessage() {}

func (x *CityDelete) ProtoReflect() protoreflect.Message {
	mi := &file_pb_InformerToBroker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityDelete.ProtoReflect.Descriptor instead.
func (*CityDelete) Descriptor() ([]byte, []int) {
	return file_pb_InformerToBroker_proto_rawDescGZIP(), []int{2}
}

func (x *CityDelete) GetPlanet() string {
	if x != nil {
		return x.Planet
	}
	return ""
}

func (x *CityDelete) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *CityDelete) GetSurvivors() int32 {
	if x != nil {
		return x.Survivors
	}
	return 0
}

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
		mi := &file_pb_InformerToBroker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Instruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instruct) ProtoMessage() {}

func (x *Instruct) ProtoReflect() protoreflect.Message {
	mi := &file_pb_InformerToBroker_proto_msgTypes[3]
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
	return file_pb_InformerToBroker_proto_rawDescGZIP(), []int{3}
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
}

func (x *Servidor) Reset() {
	*x = Servidor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_InformerToBroker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Servidor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Servidor) ProtoMessage() {}

func (x *Servidor) ProtoReflect() protoreflect.Message {
	mi := &file_pb_InformerToBroker_proto_msgTypes[4]
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
	return file_pb_InformerToBroker_proto_rawDescGZIP(), []int{4}
}

func (x *Servidor) GetAddres() string {
	if x != nil {
		return x.Addres
	}
	return ""
}

var File_pb_InformerToBroker_proto protoreflect.FileDescriptor

var file_pb_InformerToBroker_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x62, 0x2f, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x54, 0x6f, 0x42,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x0a, 0x43,
	0x69, 0x74, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6c, 0x61,
	0x6e, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6c, 0x61, 0x6e, 0x65,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x75, 0x72, 0x76, 0x69, 0x76, 0x6f,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x53, 0x75, 0x72, 0x76, 0x69, 0x76,
	0x6f, 0x72, 0x73, 0x22, 0x50, 0x0a, 0x08, 0x43, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x4e,
	0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4e, 0x65,
	0x77, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x56, 0x0a, 0x0a, 0x43, 0x69, 0x74, 0x79, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x43,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12,
	0x1c, 0x0a, 0x09, 0x53, 0x75, 0x72, 0x76, 0x69, 0x76, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x53, 0x75, 0x72, 0x76, 0x69, 0x76, 0x6f, 0x72, 0x73, 0x22, 0x3e, 0x0a,
	0x08, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x61, 0x22, 0x22, 0x0a,
	0x08, 0x53, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x32, 0xea, 0x01, 0x0a, 0x10, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x54, 0x6f,
	0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x43, 0x69, 0x74,
	0x79, 0x54, 0x6f, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x12, 0x0b, 0x2e, 0x43, 0x69, 0x74, 0x79,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x09, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f,
	0x72, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x54, 0x6f, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x12, 0x09, 0x2e, 0x43, 0x69, 0x74, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x09, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x22,
	0x00, 0x12, 0x28, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x0b, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x09,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x69, 0x74, 0x79, 0x12, 0x0b, 0x2e, 0x43, 0x69, 0x74, 0x79,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x1a, 0x09, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f,
	0x72, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x09, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x1a, 0x09, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x22, 0x00, 0x42, 0x09,
	0x5a, 0x07, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pb_InformerToBroker_proto_rawDescOnce sync.Once
	file_pb_InformerToBroker_proto_rawDescData = file_pb_InformerToBroker_proto_rawDesc
)

func file_pb_InformerToBroker_proto_rawDescGZIP() []byte {
	file_pb_InformerToBroker_proto_rawDescOnce.Do(func() {
		file_pb_InformerToBroker_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_InformerToBroker_proto_rawDescData)
	})
	return file_pb_InformerToBroker_proto_rawDescData
}

var file_pb_InformerToBroker_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pb_InformerToBroker_proto_goTypes = []interface{}{
	(*CityNumber)(nil), // 0: CityNumber
	(*CityName)(nil),   // 1: CityName
	(*CityDelete)(nil), // 2: CityDelete
	(*Instruct)(nil),   // 3: instruct
	(*Servidor)(nil),   // 4: Servidor
}
var file_pb_InformerToBroker_proto_depIdxs = []int32{
	0, // 0: InformerToBroker.AddCityToBroker:input_type -> CityNumber
	1, // 1: InformerToBroker.UpdateNameToBroker:input_type -> CityName
	0, // 2: InformerToBroker.UpdateNumber:input_type -> CityNumber
	2, // 3: InformerToBroker.DeleteCity:input_type -> CityDelete
	3, // 4: InformerToBroker.ConnectToServer:input_type -> instruct
	4, // 5: InformerToBroker.AddCityToBroker:output_type -> Servidor
	4, // 6: InformerToBroker.UpdateNameToBroker:output_type -> Servidor
	4, // 7: InformerToBroker.UpdateNumber:output_type -> Servidor
	4, // 8: InformerToBroker.DeleteCity:output_type -> Servidor
	4, // 9: InformerToBroker.ConnectToServer:output_type -> Servidor
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_InformerToBroker_proto_init() }
func file_pb_InformerToBroker_proto_init() {
	if File_pb_InformerToBroker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_InformerToBroker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CityNumber); i {
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
		file_pb_InformerToBroker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CityName); i {
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
		file_pb_InformerToBroker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CityDelete); i {
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
		file_pb_InformerToBroker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_pb_InformerToBroker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_pb_InformerToBroker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_InformerToBroker_proto_goTypes,
		DependencyIndexes: file_pb_InformerToBroker_proto_depIdxs,
		MessageInfos:      file_pb_InformerToBroker_proto_msgTypes,
	}.Build()
	File_pb_InformerToBroker_proto = out.File
	file_pb_InformerToBroker_proto_rawDesc = nil
	file_pb_InformerToBroker_proto_goTypes = nil
	file_pb_InformerToBroker_proto_depIdxs = nil
}
