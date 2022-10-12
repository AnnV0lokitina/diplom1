// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: proto/handler.proto

package proto

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type UserRegRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserRegRequest) Reset() {
	*x = UserRegRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_handler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegRequest) ProtoMessage() {}

func (x *UserRegRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_handler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegRequest.ProtoReflect.Descriptor instead.
func (*UserRegRequest) Descriptor() ([]byte, []int) {
	return file_proto_handler_proto_rawDescGZIP(), []int{0}
}

func (x *UserRegRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *UserRegRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserRegResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session string `protobuf:"bytes,3,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *UserRegResponse) Reset() {
	*x = UserRegResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_handler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegResponse) ProtoMessage() {}

func (x *UserRegResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_handler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegResponse.ProtoReflect.Descriptor instead.
func (*UserRegResponse) Descriptor() ([]byte, []int) {
	return file_proto_handler_proto_rawDescGZIP(), []int{1}
}

func (x *UserRegResponse) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string               `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Type string               `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Time *timestamp.Timestamp `protobuf:"bytes,14,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_handler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_handler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_proto_handler_proto_rawDescGZIP(), []int{2}
}

func (x *FileInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileInfo) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *FileInfo) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type StoreFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info    *FileInfo `protobuf:"bytes,7,opt,name=info,proto3" json:"info,omitempty"`
	Content []byte    `protobuf:"bytes,8,opt,name=content,proto3" json:"content,omitempty"`
	Session string    `protobuf:"bytes,4,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *StoreFileRequest) Reset() {
	*x = StoreFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_handler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreFileRequest) ProtoMessage() {}

func (x *StoreFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_handler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreFileRequest.ProtoReflect.Descriptor instead.
func (*StoreFileRequest) Descriptor() ([]byte, []int) {
	return file_proto_handler_proto_rawDescGZIP(), []int{3}
}

func (x *StoreFileRequest) GetInfo() *FileInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *StoreFileRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *StoreFileRequest) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

type StoreFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size uint32 `protobuf:"varint,10,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *StoreFileResponse) Reset() {
	*x = StoreFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_handler_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreFileResponse) ProtoMessage() {}

func (x *StoreFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_handler_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreFileResponse.ProtoReflect.Descriptor instead.
func (*StoreFileResponse) Descriptor() ([]byte, []int) {
	return file_proto_handler_proto_rawDescGZIP(), []int{4}
}

func (x *StoreFileResponse) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type RestoreFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session string `protobuf:"bytes,11,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *RestoreFileRequest) Reset() {
	*x = RestoreFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_handler_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestoreFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestoreFileRequest) ProtoMessage() {}

func (x *RestoreFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_handler_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestoreFileRequest.ProtoReflect.Descriptor instead.
func (*RestoreFileRequest) Descriptor() ([]byte, []int) {
	return file_proto_handler_proto_rawDescGZIP(), []int{5}
}

func (x *RestoreFileRequest) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

type RestoreFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info    *FileInfo `protobuf:"bytes,16,opt,name=info,proto3" json:"info,omitempty"`
	Content []byte    `protobuf:"bytes,15,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *RestoreFileResponse) Reset() {
	*x = RestoreFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_handler_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestoreFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestoreFileResponse) ProtoMessage() {}

func (x *RestoreFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_handler_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestoreFileResponse.ProtoReflect.Descriptor instead.
func (*RestoreFileResponse) Descriptor() ([]byte, []int) {
	return file_proto_handler_proto_rawDescGZIP(), []int{6}
}

func (x *RestoreFileResponse) GetInfo() *FileInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *RestoreFileResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_proto_handler_proto protoreflect.FileDescriptor

var file_proto_handler_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a,
	0x0e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x2b, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x62,
	0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x22, 0x6b, 0x0a, 0x10, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x27, 0x0a, 0x11, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x2e, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x54, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x23, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0x90,
	0x02, 0x0a, 0x0d, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x12, 0x39, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x42, 0x0a,
	0x09, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x41, 0x6e, 0x6e, 0x56, 0x30, 0x6c, 0x6f, 0x6b, 0x69, 0x74, 0x69, 0x6e, 0x61, 0x2f, 0x64, 0x69,
	0x70, 0x6c, 0x6f, 0x6d, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_handler_proto_rawDescOnce sync.Once
	file_proto_handler_proto_rawDescData = file_proto_handler_proto_rawDesc
)

func file_proto_handler_proto_rawDescGZIP() []byte {
	file_proto_handler_proto_rawDescOnce.Do(func() {
		file_proto_handler_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_handler_proto_rawDescData)
	})
	return file_proto_handler_proto_rawDescData
}

var file_proto_handler_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_handler_proto_goTypes = []interface{}{
	(*UserRegRequest)(nil),      // 0: proto.UserRegRequest
	(*UserRegResponse)(nil),     // 1: proto.UserRegResponse
	(*FileInfo)(nil),            // 2: proto.FileInfo
	(*StoreFileRequest)(nil),    // 3: proto.StoreFileRequest
	(*StoreFileResponse)(nil),   // 4: proto.StoreFileResponse
	(*RestoreFileRequest)(nil),  // 5: proto.RestoreFileRequest
	(*RestoreFileResponse)(nil), // 6: proto.RestoreFileResponse
	(*timestamp.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_proto_handler_proto_depIdxs = []int32{
	7, // 0: proto.FileInfo.time:type_name -> google.protobuf.Timestamp
	2, // 1: proto.StoreFileRequest.info:type_name -> proto.FileInfo
	2, // 2: proto.RestoreFileResponse.info:type_name -> proto.FileInfo
	0, // 3: proto.SecureStorage.Register:input_type -> proto.UserRegRequest
	0, // 4: proto.SecureStorage.Login:input_type -> proto.UserRegRequest
	5, // 5: proto.SecureStorage.RestoreFile:input_type -> proto.RestoreFileRequest
	3, // 6: proto.SecureStorage.StoreFile:input_type -> proto.StoreFileRequest
	1, // 7: proto.SecureStorage.Register:output_type -> proto.UserRegResponse
	1, // 8: proto.SecureStorage.Login:output_type -> proto.UserRegResponse
	6, // 9: proto.SecureStorage.RestoreFile:output_type -> proto.RestoreFileResponse
	4, // 10: proto.SecureStorage.StoreFile:output_type -> proto.StoreFileResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_handler_proto_init() }
func file_proto_handler_proto_init() {
	if File_proto_handler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_handler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegRequest); i {
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
		file_proto_handler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegResponse); i {
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
		file_proto_handler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
		file_proto_handler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreFileRequest); i {
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
		file_proto_handler_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreFileResponse); i {
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
		file_proto_handler_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestoreFileRequest); i {
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
		file_proto_handler_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestoreFileResponse); i {
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
			RawDescriptor: file_proto_handler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_handler_proto_goTypes,
		DependencyIndexes: file_proto_handler_proto_depIdxs,
		MessageInfos:      file_proto_handler_proto_msgTypes,
	}.Build()
	File_proto_handler_proto = out.File
	file_proto_handler_proto_rawDesc = nil
	file_proto_handler_proto_goTypes = nil
	file_proto_handler_proto_depIdxs = nil
}
