// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.2
// source: proto/v1/modelserving.proto

package modelserving

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Provider            string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	CsvName             string `protobuf:"bytes,3,opt,name=csvName,proto3" json:"csvName,omitempty"`
	MinSupportedVersion string `protobuf:"bytes,4,opt,name=minSupportedVersion,proto3" json:"minSupportedVersion,omitempty"`
	Installed           bool   `protobuf:"varint,5,opt,name=installed,proto3" json:"installed,omitempty"`
	Template            string `protobuf:"bytes,6,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_modelserving_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_modelserving_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_proto_v1_modelserving_proto_rawDescGZIP(), []int{0}
}

func (x *Application) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Application) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *Application) GetCsvName() string {
	if x != nil {
		return x.CsvName
	}
	return ""
}

func (x *Application) GetMinSupportedVersion() string {
	if x != nil {
		return x.MinSupportedVersion
	}
	return ""
}

func (x *Application) GetInstalled() bool {
	if x != nil {
		return x.Installed
	}
	return false
}

func (x *Application) GetTemplate() string {
	if x != nil {
		return x.Template
	}
	return ""
}

type Support struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Isv    bool `protobuf:"varint,1,opt,name=isv,proto3" json:"isv,omitempty"`
	Redhat bool `protobuf:"varint,2,opt,name=redhat,proto3" json:"redhat,omitempty"`
}

func (x *Support) Reset() {
	*x = Support{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_modelserving_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Support) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Support) ProtoMessage() {}

func (x *Support) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_modelserving_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Support.ProtoReflect.Descriptor instead.
func (*Support) Descriptor() ([]byte, []int) {
	return file_proto_v1_modelserving_proto_rawDescGZIP(), []int{1}
}

func (x *Support) GetIsv() bool {
	if x != nil {
		return x.Isv
	}
	return false
}

func (x *Support) GetRedhat() bool {
	if x != nil {
		return x.Redhat
	}
	return false
}

type Integration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstallDocUrl string `protobuf:"bytes,1,opt,name=install_doc_url,json=installDocUrl,proto3" json:"install_doc_url,omitempty"`
	Level         string `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *Integration) Reset() {
	*x = Integration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_modelserving_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Integration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Integration) ProtoMessage() {}

func (x *Integration) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_modelserving_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Integration.ProtoReflect.Descriptor instead.
func (*Integration) Descriptor() ([]byte, []int) {
	return file_proto_v1_modelserving_proto_rawDescGZIP(), []int{2}
}

func (x *Integration) GetInstallDocUrl() string {
	if x != nil {
		return x.InstallDocUrl
	}
	return ""
}

func (x *Integration) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

type GetAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Applications []*Application `protobuf:"bytes,1,rep,name=applications,proto3" json:"applications,omitempty"`
}

func (x *GetAppResponse) Reset() {
	*x = GetAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_modelserving_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppResponse) ProtoMessage() {}

func (x *GetAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_modelserving_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppResponse.ProtoReflect.Descriptor instead.
func (*GetAppResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_modelserving_proto_rawDescGZIP(), []int{3}
}

func (x *GetAppResponse) GetApplications() []*Application {
	if x != nil {
		return x.Applications
	}
	return nil
}

type GetAppParamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppName     string `protobuf:"bytes,1,opt,name=appName,proto3" json:"appName,omitempty"`
	StorageName string `protobuf:"bytes,2,opt,name=storageName,proto3" json:"storageName,omitempty"`
	Namespace   string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *GetAppParamsRequest) Reset() {
	*x = GetAppParamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_modelserving_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppParamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppParamsRequest) ProtoMessage() {}

func (x *GetAppParamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_modelserving_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppParamsRequest.ProtoReflect.Descriptor instead.
func (*GetAppParamsRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_modelserving_proto_rawDescGZIP(), []int{4}
}

func (x *GetAppParamsRequest) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *GetAppParamsRequest) GetStorageName() string {
	if x != nil {
		return x.StorageName
	}
	return ""
}

func (x *GetAppParamsRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type GetAppParamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppName     string `protobuf:"bytes,1,opt,name=appName,proto3" json:"appName,omitempty"`
	StorageName string `protobuf:"bytes,2,opt,name=storageName,proto3" json:"storageName,omitempty"`
	Namespace   string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Parameters  []byte `protobuf:"bytes,4,opt,name=parameters,proto3" json:"parameters,omitempty"`
}

func (x *GetAppParamsResponse) Reset() {
	*x = GetAppParamsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_modelserving_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppParamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppParamsResponse) ProtoMessage() {}

func (x *GetAppParamsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_modelserving_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppParamsResponse.ProtoReflect.Descriptor instead.
func (*GetAppParamsResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_modelserving_proto_rawDescGZIP(), []int{5}
}

func (x *GetAppParamsResponse) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *GetAppParamsResponse) GetStorageName() string {
	if x != nil {
		return x.StorageName
	}
	return ""
}

func (x *GetAppParamsResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *GetAppParamsResponse) GetParameters() []byte {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type GetRenderedCRRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppName     string            `protobuf:"bytes,1,opt,name=appName,proto3" json:"appName,omitempty"`
	StorageName string            `protobuf:"bytes,2,opt,name=storageName,proto3" json:"storageName,omitempty"`
	Namespace   string            `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Parameters  map[string]string `protobuf:"bytes,4,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetRenderedCRRequest) Reset() {
	*x = GetRenderedCRRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_modelserving_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRenderedCRRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRenderedCRRequest) ProtoMessage() {}

func (x *GetRenderedCRRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_modelserving_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRenderedCRRequest.ProtoReflect.Descriptor instead.
func (*GetRenderedCRRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_modelserving_proto_rawDescGZIP(), []int{6}
}

func (x *GetRenderedCRRequest) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *GetRenderedCRRequest) GetStorageName() string {
	if x != nil {
		return x.StorageName
	}
	return ""
}

func (x *GetRenderedCRRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *GetRenderedCRRequest) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type GetRenderedCRResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Manifest []byte `protobuf:"bytes,1,opt,name=manifest,proto3" json:"manifest,omitempty"`
}

func (x *GetRenderedCRResponse) Reset() {
	*x = GetRenderedCRResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_modelserving_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRenderedCRResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRenderedCRResponse) ProtoMessage() {}

func (x *GetRenderedCRResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_modelserving_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRenderedCRResponse.ProtoReflect.Descriptor instead.
func (*GetRenderedCRResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_modelserving_proto_rawDescGZIP(), []int{7}
}

func (x *GetRenderedCRResponse) GetManifest() []byte {
	if x != nil {
		return x.Manifest
	}
	return nil
}

var File_proto_v1_modelserving_proto protoreflect.FileDescriptor

var file_proto_v1_modelserving_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61,
	0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x01,
	0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x73, 0x76, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x73, 0x76, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x6d, 0x69, 0x6e, 0x53, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6d, 0x69, 0x6e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x22, 0x33, 0x0a, 0x07, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x73, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x69, 0x73, 0x76,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x64, 0x68, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x72, 0x65, 0x64, 0x68, 0x61, 0x74, 0x22, 0x4b, 0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x5f, 0x64, 0x6f, 0x63, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x44, 0x6f, 0x63, 0x55, 0x72, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x46, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x6f, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x90,
	0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x22, 0xfa, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65,
	0x64, 0x43, 0x52, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x64, 0x43, 0x52, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x1a,
	0x3d, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x33,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x64, 0x43, 0x52, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x6e, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x61, 0x6e, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x32, 0xd0, 0x02, 0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x12, 0x54, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x12, 0x6b, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x41, 0x70, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70,
	0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x7b, 0x61,
	0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x7d, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x70,
	0x70, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65,
	0x64, 0x43, 0x52, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x64, 0x43, 0x52, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x22, 0x23,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x7d, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0x3a, 0x01, 0x2a, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x3b, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v1_modelserving_proto_rawDescOnce sync.Once
	file_proto_v1_modelserving_proto_rawDescData = file_proto_v1_modelserving_proto_rawDesc
)

func file_proto_v1_modelserving_proto_rawDescGZIP() []byte {
	file_proto_v1_modelserving_proto_rawDescOnce.Do(func() {
		file_proto_v1_modelserving_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_modelserving_proto_rawDescData)
	})
	return file_proto_v1_modelserving_proto_rawDescData
}

var file_proto_v1_modelserving_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_v1_modelserving_proto_goTypes = []interface{}{
	(*Application)(nil),           // 0: api.Application
	(*Support)(nil),               // 1: api.Support
	(*Integration)(nil),           // 2: api.Integration
	(*GetAppResponse)(nil),        // 3: api.GetAppResponse
	(*GetAppParamsRequest)(nil),   // 4: api.GetAppParamsRequest
	(*GetAppParamsResponse)(nil),  // 5: api.GetAppParamsResponse
	(*GetRenderedCRRequest)(nil),  // 6: api.GetRenderedCRRequest
	(*GetRenderedCRResponse)(nil), // 7: api.GetRenderedCRResponse
	nil,                           // 8: api.GetRenderedCRRequest.ParametersEntry
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_proto_v1_modelserving_proto_depIdxs = []int32{
	0, // 0: api.GetAppResponse.applications:type_name -> api.Application
	8, // 1: api.GetRenderedCRRequest.parameters:type_name -> api.GetRenderedCRRequest.ParametersEntry
	9, // 2: api.ModelServing.ListApp:input_type -> google.protobuf.Empty
	4, // 3: api.ModelServing.GetAppParams:input_type -> api.GetAppParamsRequest
	6, // 4: api.ModelServing.GetAppCustomResource:input_type -> api.GetRenderedCRRequest
	3, // 5: api.ModelServing.ListApp:output_type -> api.GetAppResponse
	5, // 6: api.ModelServing.GetAppParams:output_type -> api.GetAppParamsResponse
	7, // 7: api.ModelServing.GetAppCustomResource:output_type -> api.GetRenderedCRResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_v1_modelserving_proto_init() }
func file_proto_v1_modelserving_proto_init() {
	if File_proto_v1_modelserving_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1_modelserving_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_proto_v1_modelserving_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Support); i {
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
		file_proto_v1_modelserving_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Integration); i {
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
		file_proto_v1_modelserving_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppResponse); i {
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
		file_proto_v1_modelserving_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppParamsRequest); i {
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
		file_proto_v1_modelserving_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppParamsResponse); i {
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
		file_proto_v1_modelserving_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRenderedCRRequest); i {
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
		file_proto_v1_modelserving_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRenderedCRResponse); i {
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
			RawDescriptor: file_proto_v1_modelserving_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1_modelserving_proto_goTypes,
		DependencyIndexes: file_proto_v1_modelserving_proto_depIdxs,
		MessageInfos:      file_proto_v1_modelserving_proto_msgTypes,
	}.Build()
	File_proto_v1_modelserving_proto = out.File
	file_proto_v1_modelserving_proto_rawDesc = nil
	file_proto_v1_modelserving_proto_goTypes = nil
	file_proto_v1_modelserving_proto_depIdxs = nil
}
