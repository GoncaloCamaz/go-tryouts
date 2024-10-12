// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: class.proto

package class

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request to get class info by ID
type ClassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassId int64 `protobuf:"varint,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
}

func (x *ClassRequest) Reset() {
	*x = ClassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_class_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassRequest) ProtoMessage() {}

func (x *ClassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_class_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassRequest.ProtoReflect.Descriptor instead.
func (*ClassRequest) Descriptor() ([]byte, []int) {
	return file_class_proto_rawDescGZIP(), []int{0}
}

func (x *ClassRequest) GetClassId() int64 {
	if x != nil {
		return x.ClassId
	}
	return 0
}

// Response with class information
type ClassResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Year    string `protobuf:"bytes,2,opt,name=year,proto3" json:"year,omitempty"`
	Number  int32  `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	Updated string `protobuf:"bytes,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Created string `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *ClassResponse) Reset() {
	*x = ClassResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_class_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassResponse) ProtoMessage() {}

func (x *ClassResponse) ProtoReflect() protoreflect.Message {
	mi := &file_class_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassResponse.ProtoReflect.Descriptor instead.
func (*ClassResponse) Descriptor() ([]byte, []int) {
	return file_class_proto_rawDescGZIP(), []int{1}
}

func (x *ClassResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ClassResponse) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *ClassResponse) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *ClassResponse) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

func (x *ClassResponse) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

type ClassListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Classes []*ClassResponse `protobuf:"bytes,1,rep,name=classes,proto3" json:"classes,omitempty"`
}

func (x *ClassListResponse) Reset() {
	*x = ClassListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_class_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassListResponse) ProtoMessage() {}

func (x *ClassListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_class_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassListResponse.ProtoReflect.Descriptor instead.
func (*ClassListResponse) Descriptor() ([]byte, []int) {
	return file_class_proto_rawDescGZIP(), []int{2}
}

func (x *ClassListResponse) GetClasses() []*ClassResponse {
	if x != nil {
		return x.Classes
	}
	return nil
}

var File_class_proto protoreflect.FileDescriptor

var file_class_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x29, 0x0a, 0x0c, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x22, 0x7f, 0x0a, 0x0d,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x43, 0x0a,
	0x11, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x73, 0x32, 0x8b, 0x01, 0x0a, 0x0c, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x13, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_class_proto_rawDescOnce sync.Once
	file_class_proto_rawDescData = file_class_proto_rawDesc
)

func file_class_proto_rawDescGZIP() []byte {
	file_class_proto_rawDescOnce.Do(func() {
		file_class_proto_rawDescData = protoimpl.X.CompressGZIP(file_class_proto_rawDescData)
	})
	return file_class_proto_rawDescData
}

var file_class_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_class_proto_goTypes = []any{
	(*ClassRequest)(nil),      // 0: class.ClassRequest
	(*ClassResponse)(nil),     // 1: class.ClassResponse
	(*ClassListResponse)(nil), // 2: class.ClassListResponse
	(*emptypb.Empty)(nil),     // 3: google.protobuf.Empty
}
var file_class_proto_depIdxs = []int32{
	1, // 0: class.ClassListResponse.classes:type_name -> class.ClassResponse
	0, // 1: class.ClassService.GetClassInfo:input_type -> class.ClassRequest
	3, // 2: class.ClassService.GetClassList:input_type -> google.protobuf.Empty
	1, // 3: class.ClassService.GetClassInfo:output_type -> class.ClassResponse
	2, // 4: class.ClassService.GetClassList:output_type -> class.ClassListResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_class_proto_init() }
func file_class_proto_init() {
	if File_class_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_class_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ClassRequest); i {
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
		file_class_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ClassResponse); i {
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
		file_class_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ClassListResponse); i {
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
			RawDescriptor: file_class_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_class_proto_goTypes,
		DependencyIndexes: file_class_proto_depIdxs,
		MessageInfos:      file_class_proto_msgTypes,
	}.Build()
	File_class_proto = out.File
	file_class_proto_rawDesc = nil
	file_class_proto_goTypes = nil
	file_class_proto_depIdxs = nil
}
