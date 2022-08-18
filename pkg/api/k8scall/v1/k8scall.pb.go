// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.2
// source: proto/v1/k8scall.proto

package k8scall

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

type K8SStringJson struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Manifest  string `protobuf:"bytes,1,opt,name=manifest,proto3" json:"manifest,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *K8SStringJson) Reset() {
	*x = K8SStringJson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_k8scall_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *K8SStringJson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*K8SStringJson) ProtoMessage() {}

func (x *K8SStringJson) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_k8scall_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use K8SStringJson.ProtoReflect.Descriptor instead.
func (*K8SStringJson) Descriptor() ([]byte, []int) {
	return file_proto_v1_k8scall_proto_rawDescGZIP(), []int{0}
}

func (x *K8SStringJson) GetManifest() string {
	if x != nil {
		return x.Manifest
	}
	return ""
}

func (x *K8SStringJson) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type CreateObjectByFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok          bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateObjectByFileResponse) Reset() {
	*x = CreateObjectByFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_k8scall_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateObjectByFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateObjectByFileResponse) ProtoMessage() {}

func (x *CreateObjectByFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_k8scall_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateObjectByFileResponse.ProtoReflect.Descriptor instead.
func (*CreateObjectByFileResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_k8scall_proto_rawDescGZIP(), []int{1}
}

func (x *CreateObjectByFileResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *CreateObjectByFileResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_proto_v1_k8scall_proto protoreflect.FileDescriptor

var file_proto_v1_k8scall_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x38, 0x73, 0x63, 0x61,
	0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x49, 0x0a,
	0x0d, 0x4b, 0x38, 0x53, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x4e, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x5a, 0x0a, 0x07, 0x4b, 0x38, 0x53, 0x43,
	0x61, 0x6c, 0x6c, 0x12, 0x4f, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x42, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4a, 0x73, 0x6f, 0x6e, 0x12,
	0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4b, 0x38, 0x53, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4a,
	0x73, 0x6f, 0x6e, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x6b, 0x38, 0x73, 0x63, 0x61, 0x6c,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v1_k8scall_proto_rawDescOnce sync.Once
	file_proto_v1_k8scall_proto_rawDescData = file_proto_v1_k8scall_proto_rawDesc
)

func file_proto_v1_k8scall_proto_rawDescGZIP() []byte {
	file_proto_v1_k8scall_proto_rawDescOnce.Do(func() {
		file_proto_v1_k8scall_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_k8scall_proto_rawDescData)
	})
	return file_proto_v1_k8scall_proto_rawDescData
}

var file_proto_v1_k8scall_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_v1_k8scall_proto_goTypes = []interface{}{
	(*K8SStringJson)(nil),              // 0: api.K8SStringJson
	(*CreateObjectByFileResponse)(nil), // 1: api.CreateObjectByFileResponse
}
var file_proto_v1_k8scall_proto_depIdxs = []int32{
	0, // 0: api.K8SCall.CreateObjectByStringJson:input_type -> api.K8SStringJson
	1, // 1: api.K8SCall.CreateObjectByStringJson:output_type -> api.CreateObjectByFileResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_v1_k8scall_proto_init() }
func file_proto_v1_k8scall_proto_init() {
	if File_proto_v1_k8scall_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1_k8scall_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*K8SStringJson); i {
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
		file_proto_v1_k8scall_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateObjectByFileResponse); i {
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
			RawDescriptor: file_proto_v1_k8scall_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1_k8scall_proto_goTypes,
		DependencyIndexes: file_proto_v1_k8scall_proto_depIdxs,
		MessageInfos:      file_proto_v1_k8scall_proto_msgTypes,
	}.Build()
	File_proto_v1_k8scall_proto = out.File
	file_proto_v1_k8scall_proto_rawDesc = nil
	file_proto_v1_k8scall_proto_goTypes = nil
	file_proto_v1_k8scall_proto_depIdxs = nil
}
