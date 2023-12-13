// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: api/user_v1/service.proto

package __

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

// GetUser
type GetRequestUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdUser uint32 `protobuf:"varint,1,opt,name=id_user,json=idUser,proto3" json:"id_user,omitempty"`
}

func (x *GetRequestUser) Reset() {
	*x = GetRequestUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequestUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequestUser) ProtoMessage() {}

func (x *GetRequestUser) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequestUser.ProtoReflect.Descriptor instead.
func (*GetRequestUser) Descriptor() ([]byte, []int) {
	return file_api_user_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequestUser) GetIdUser() uint32 {
	if x != nil {
		return x.IdUser
	}
	return 0
}

type GetResponseUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdUser    uint32 `protobuf:"varint,1,opt,name=id_user,json=idUser,proto3" json:"id_user,omitempty"`
	Login     string `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	IsManager bool   `protobuf:"varint,3,opt,name=isManager,proto3" json:"isManager,omitempty"`
}

func (x *GetResponseUser) Reset() {
	*x = GetResponseUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponseUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponseUser) ProtoMessage() {}

func (x *GetResponseUser) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponseUser.ProtoReflect.Descriptor instead.
func (*GetResponseUser) Descriptor() ([]byte, []int) {
	return file_api_user_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponseUser) GetIdUser() uint32 {
	if x != nil {
		return x.IdUser
	}
	return 0
}

func (x *GetResponseUser) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *GetResponseUser) GetIsManager() bool {
	if x != nil {
		return x.IsManager
	}
	return false
}

// CheckPassword
type GetRequestAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *GetRequestAuth) Reset() {
	*x = GetRequestAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequestAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequestAuth) ProtoMessage() {}

func (x *GetRequestAuth) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequestAuth.ProtoReflect.Descriptor instead.
func (*GetRequestAuth) Descriptor() ([]byte, []int) {
	return file_api_user_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequestAuth) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *GetRequestAuth) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GetResponseAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsValidated bool `protobuf:"varint,1,opt,name=isValidated,proto3" json:"isValidated,omitempty"`
}

func (x *GetResponseAuth) Reset() {
	*x = GetResponseAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponseAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponseAuth) ProtoMessage() {}

func (x *GetResponseAuth) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponseAuth.ProtoReflect.Descriptor instead.
func (*GetResponseAuth) Descriptor() ([]byte, []int) {
	return file_api_user_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetResponseAuth) GetIsValidated() bool {
	if x != nil {
		return x.IsValidated
	}
	return false
}

// Update mail
type PutRequestMail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdUser uint32 `protobuf:"varint,1,opt,name=id_user,json=idUser,proto3" json:"id_user,omitempty"`
	Email  string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *PutRequestMail) Reset() {
	*x = PutRequestMail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutRequestMail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutRequestMail) ProtoMessage() {}

func (x *PutRequestMail) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutRequestMail.ProtoReflect.Descriptor instead.
func (*PutRequestMail) Descriptor() ([]byte, []int) {
	return file_api_user_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *PutRequestMail) GetIdUser() uint32 {
	if x != nil {
		return x.IdUser
	}
	return 0
}

func (x *PutRequestMail) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type PutResponseMail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *PutResponseMail) Reset() {
	*x = PutResponseMail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_v1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutResponseMail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutResponseMail) ProtoMessage() {}

func (x *PutResponseMail) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutResponseMail.ProtoReflect.Descriptor instead.
func (*PutResponseMail) Descriptor() ([]byte, []int) {
	return file_api_user_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *PutResponseMail) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// Create user
type PostRequestUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login     string `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Password  string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	IsManager bool   `protobuf:"varint,5,opt,name=isManager,proto3" json:"isManager,omitempty"`
}

func (x *PostRequestUser) Reset() {
	*x = PostRequestUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_v1_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostRequestUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostRequestUser) ProtoMessage() {}

func (x *PostRequestUser) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostRequestUser.ProtoReflect.Descriptor instead.
func (*PostRequestUser) Descriptor() ([]byte, []int) {
	return file_api_user_v1_service_proto_rawDescGZIP(), []int{6}
}

func (x *PostRequestUser) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *PostRequestUser) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *PostRequestUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PostRequestUser) GetIsManager() bool {
	if x != nil {
		return x.IsManager
	}
	return false
}

type PostResponseUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdBoard  uint32   `protobuf:"varint,1,opt,name=id_board,json=idBoard,proto3" json:"id_board,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Private  bool     `protobuf:"varint,3,opt,name=private,proto3" json:"private,omitempty"`
	IdColumn []uint32 `protobuf:"varint,4,rep,packed,name=id_column,json=idColumn,proto3" json:"id_column,omitempty"`
}

func (x *PostResponseUser) Reset() {
	*x = PostResponseUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_v1_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostResponseUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostResponseUser) ProtoMessage() {}

func (x *PostResponseUser) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostResponseUser.ProtoReflect.Descriptor instead.
func (*PostResponseUser) Descriptor() ([]byte, []int) {
	return file_api_user_v1_service_proto_rawDescGZIP(), []int{7}
}

func (x *PostResponseUser) GetIdUser() string {
	if x != nil {
		return x.IdUser
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdUser    uint32 `protobuf:"varint,1,opt,name=id_user,json=idUser,proto3" json:"id_user,omitempty"`
	Login     string `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Password  string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	IsManager bool   `protobuf:"varint,5,opt,name=isManager,proto3" json:"isManager,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_v1_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_v1_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_api_user_v1_service_proto_rawDescGZIP(), []int{8}
}

func (x *User) GetIdUser() uint32 {
	if x != nil {
		return x.IdUser
	}
	return 0
}

func (x *User) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetIsManager() bool {
	if x != nil {
		return x.IsManager
	}
	return false
}

var File_api_user_v1_service_proto protoreflect.FileDescriptor

var file_api_user_v1_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x76, 0x31, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x69, 0x64, 0x55, 0x73, 0x65, 0x72, 0x22,
	0x5e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x69, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x22,
	0x42, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x33, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x41, 0x75, 0x74, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x3f, 0x0a, 0x0e, 0x50, 0x75, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x69, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x27, 0x0a, 0x0f, 0x50, 0x75, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x77, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a,
	0x09, 0x69, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x69, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x22, 0x2b, 0x0a, 0x10, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x69, 0x64, 0x55, 0x73, 0x65, 0x72, 0x22, 0x85, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x69, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x32, 0x8d, 0x02, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x56, 0x31, 0x12, 0x3c, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x1a,
	0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x0c, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x1a, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x75, 0x74, 0x68, 0x12, 0x3f, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d,
	0x61, 0x69, 0x6c, 0x1a, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x50, 0x75,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x41, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_user_v1_service_proto_rawDescOnce sync.Once
	file_api_user_v1_service_proto_rawDescData = file_api_user_v1_service_proto_rawDesc
)

func file_api_user_v1_service_proto_rawDescGZIP() []byte {
	file_api_user_v1_service_proto_rawDescOnce.Do(func() {
		file_api_user_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_user_v1_service_proto_rawDescData)
	})
	return file_api_user_v1_service_proto_rawDescData
}

var file_api_user_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_user_v1_service_proto_goTypes = []interface{}{
	(*GetRequestUser)(nil),   // 0: user_v1.GetRequestUser
	(*GetResponseUser)(nil),  // 1: user_v1.GetResponseUser
	(*GetRequestAuth)(nil),   // 2: user_v1.GetRequestAuth
	(*GetResponseAuth)(nil),  // 3: user_v1.GetResponseAuth
	(*PutRequestMail)(nil),   // 4: user_v1.PutRequestMail
	(*PutResponseMail)(nil),  // 5: user_v1.PutResponseMail
	(*PostRequestUser)(nil),  // 6: user_v1.PostRequestUser
	(*PostResponseUser)(nil), // 7: user_v1.PostResponseUser
	(*User)(nil),             // 8: user_v1.User
}
var file_api_user_v1_service_proto_depIdxs = []int32{
	0, // 0: user_v1.UserV1.GetUser:input_type -> user_v1.GetRequestUser
	2, // 1: user_v1.UserV1.ValidateUser:input_type -> user_v1.GetRequestAuth
	4, // 2: user_v1.UserV1.UpdateMail:input_type -> user_v1.PutRequestMail
	6, // 3: user_v1.UserV1.CreateUser:input_type -> user_v1.PostRequestUser
	1, // 4: user_v1.UserV1.GetUser:output_type -> user_v1.GetResponseUser
	3, // 5: user_v1.UserV1.ValidateUser:output_type -> user_v1.GetResponseAuth
	5, // 6: user_v1.UserV1.UpdateMail:output_type -> user_v1.PutResponseMail
	7, // 7: user_v1.UserV1.CreateUser:output_type -> user_v1.PostResponseUser
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_user_v1_service_proto_init() }
func file_api_user_v1_service_proto_init() {
	if File_api_user_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_user_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequestUser); i {
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
		file_api_user_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponseUser); i {
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
		file_api_user_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequestAuth); i {
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
		file_api_user_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponseAuth); i {
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
		file_api_user_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutRequestMail); i {
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
		file_api_user_v1_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutResponseMail); i {
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
		file_api_user_v1_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostRequestUser); i {
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
		file_api_user_v1_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostResponseUser); i {
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
		file_api_user_v1_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_api_user_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_user_v1_service_proto_goTypes,
		DependencyIndexes: file_api_user_v1_service_proto_depIdxs,
		MessageInfos:      file_api_user_v1_service_proto_msgTypes,
	}.Build()
	File_api_user_v1_service_proto = out.File
	file_api_user_v1_service_proto_rawDesc = nil
	file_api_user_v1_service_proto_goTypes = nil
	file_api_user_v1_service_proto_depIdxs = nil
}
