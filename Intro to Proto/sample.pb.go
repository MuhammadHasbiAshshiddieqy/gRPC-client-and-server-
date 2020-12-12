// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: sample.proto

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

type DataAnggota struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nama         string                    `protobuf:"bytes,1,opt,name=nama,proto3" json:"nama,omitempty"`
	Usia         int64                     `protobuf:"varint,2,opt,name=usia,proto3" json:"usia,omitempty"`
	DataTambahan *DataAnggota_DATATAMBAHAN `protobuf:"bytes,3,opt,name=data_tambahan,proto3" json:"data_tambahan,omitempty"`
}

func (x *DataAnggota) Reset() {
	*x = DataAnggota{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataAnggota) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataAnggota) ProtoMessage() {}

func (x *DataAnggota) ProtoReflect() protoreflect.Message {
	mi := &file_sample_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataAnggota.ProtoReflect.Descriptor instead.
func (*DataAnggota) Descriptor() ([]byte, []int) {
	return file_sample_proto_rawDescGZIP(), []int{0}
}

func (x *DataAnggota) GetNama() string {
	if x != nil {
		return x.Nama
	}
	return ""
}

func (x *DataAnggota) GetUsia() int64 {
	if x != nil {
		return x.Usia
	}
	return 0
}

func (x *DataAnggota) GetDataTambahan() *DataAnggota_DATATAMBAHAN {
	if x != nil {
		return x.DataTambahan
	}
	return nil
}

type Kelompok struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total   int64        `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Anggota *DataAnggota `protobuf:"bytes,2,opt,name=anggota,json=DataAnggota,proto3" json:"anggota,omitempty"`
}

func (x *Kelompok) Reset() {
	*x = Kelompok{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kelompok) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kelompok) ProtoMessage() {}

func (x *Kelompok) ProtoReflect() protoreflect.Message {
	mi := &file_sample_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kelompok.ProtoReflect.Descriptor instead.
func (*Kelompok) Descriptor() ([]byte, []int) {
	return file_sample_proto_rawDescGZIP(), []int{1}
}

func (x *Kelompok) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Kelompok) GetAnggota() *DataAnggota {
	if x != nil {
		return x.Anggota
	}
	return nil
}

type DataAnggota_DATATAMBAHAN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hobi  string `protobuf:"bytes,1,opt,name=hobi,proto3" json:"hobi,omitempty"`
	Hidup bool   `protobuf:"varint,2,opt,name=hidup,proto3" json:"hidup,omitempty"`
}

func (x *DataAnggota_DATATAMBAHAN) Reset() {
	*x = DataAnggota_DATATAMBAHAN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataAnggota_DATATAMBAHAN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataAnggota_DATATAMBAHAN) ProtoMessage() {}

func (x *DataAnggota_DATATAMBAHAN) ProtoReflect() protoreflect.Message {
	mi := &file_sample_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataAnggota_DATATAMBAHAN.ProtoReflect.Descriptor instead.
func (*DataAnggota_DATATAMBAHAN) Descriptor() ([]byte, []int) {
	return file_sample_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DataAnggota_DATATAMBAHAN) GetHobi() string {
	if x != nil {
		return x.Hobi
	}
	return ""
}

func (x *DataAnggota_DATATAMBAHAN) GetHidup() bool {
	if x != nil {
		return x.Hidup
	}
	return false
}

var File_sample_proto protoreflect.FileDescriptor

var file_sample_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0xb3, 0x01, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x41, 0x6e, 0x67, 0x67, 0x6f,
	0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x69, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x73, 0x69, 0x61, 0x12, 0x42, 0x0a, 0x0d, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x74, 0x61, 0x6d, 0x62, 0x61, 0x68, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x41, 0x6e, 0x67, 0x67, 0x6f,
	0x74, 0x61, 0x2e, 0x44, 0x41, 0x54, 0x41, 0x54, 0x41, 0x4d, 0x42, 0x41, 0x48, 0x41, 0x4e, 0x52,
	0x0d, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x61, 0x6d, 0x62, 0x61, 0x68, 0x61, 0x6e, 0x1a, 0x38,
	0x0a, 0x0c, 0x44, 0x41, 0x54, 0x41, 0x54, 0x41, 0x4d, 0x42, 0x41, 0x48, 0x41, 0x4e, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x6f, 0x62, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f,
	0x62, 0x69, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x69, 0x64, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x68, 0x69, 0x64, 0x75, 0x70, 0x22, 0x4f, 0x0a, 0x08, 0x4b, 0x65, 0x6c, 0x6f,
	0x6d, 0x70, 0x6f, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2d, 0x0a, 0x07, 0x61, 0x6e,
	0x67, 0x67, 0x6f, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x41, 0x6e, 0x67, 0x67, 0x6f, 0x74, 0x61, 0x52, 0x0b, 0x44, 0x61,
	0x74, 0x61, 0x41, 0x6e, 0x67, 0x67, 0x6f, 0x74, 0x61, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sample_proto_rawDescOnce sync.Once
	file_sample_proto_rawDescData = file_sample_proto_rawDesc
)

func file_sample_proto_rawDescGZIP() []byte {
	file_sample_proto_rawDescOnce.Do(func() {
		file_sample_proto_rawDescData = protoimpl.X.CompressGZIP(file_sample_proto_rawDescData)
	})
	return file_sample_proto_rawDescData
}

var file_sample_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sample_proto_goTypes = []interface{}{
	(*DataAnggota)(nil),              // 0: pb.DataAnggota
	(*Kelompok)(nil),                 // 1: pb.Kelompok
	(*DataAnggota_DATATAMBAHAN)(nil), // 2: pb.DataAnggota.DATATAMBAHAN
}
var file_sample_proto_depIdxs = []int32{
	2, // 0: pb.DataAnggota.data_tambahan:type_name -> pb.DataAnggota.DATATAMBAHAN
	0, // 1: pb.Kelompok.anggota:type_name -> pb.DataAnggota
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sample_proto_init() }
func file_sample_proto_init() {
	if File_sample_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sample_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataAnggota); i {
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
		file_sample_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kelompok); i {
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
		file_sample_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataAnggota_DATATAMBAHAN); i {
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
			RawDescriptor: file_sample_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sample_proto_goTypes,
		DependencyIndexes: file_sample_proto_depIdxs,
		MessageInfos:      file_sample_proto_msgTypes,
	}.Build()
	File_sample_proto = out.File
	file_sample_proto_rawDesc = nil
	file_sample_proto_goTypes = nil
	file_sample_proto_depIdxs = nil
}
