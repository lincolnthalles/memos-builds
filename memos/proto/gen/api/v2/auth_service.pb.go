// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/v2/auth_service.proto

package apiv2

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetAuthStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAuthStatusRequest) Reset() {
	*x = GetAuthStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_auth_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthStatusRequest) ProtoMessage() {}

func (x *GetAuthStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_auth_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthStatusRequest.ProtoReflect.Descriptor instead.
func (*GetAuthStatusRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_auth_service_proto_rawDescGZIP(), []int{0}
}

type GetAuthStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetAuthStatusResponse) Reset() {
	*x = GetAuthStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_auth_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthStatusResponse) ProtoMessage() {}

func (x *GetAuthStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_auth_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthStatusResponse.ProtoReflect.Descriptor instead.
func (*GetAuthStatusResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_auth_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetAuthStatusResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type SignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username    string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password    string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	NeverExpire bool   `protobuf:"varint,3,opt,name=never_expire,json=neverExpire,proto3" json:"never_expire,omitempty"`
}

func (x *SignInRequest) Reset() {
	*x = SignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_auth_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInRequest) ProtoMessage() {}

func (x *SignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_auth_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInRequest.ProtoReflect.Descriptor instead.
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_auth_service_proto_rawDescGZIP(), []int{2}
}

func (x *SignInRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SignInRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SignInRequest) GetNeverExpire() bool {
	if x != nil {
		return x.NeverExpire
	}
	return false
}

type SignInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *SignInResponse) Reset() {
	*x = SignInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_auth_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInResponse) ProtoMessage() {}

func (x *SignInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_auth_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInResponse.ProtoReflect.Descriptor instead.
func (*SignInResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_auth_service_proto_rawDescGZIP(), []int{3}
}

func (x *SignInResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type SignInWithSSORequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdpId       int32  `protobuf:"varint,1,opt,name=idp_id,json=idpId,proto3" json:"idp_id,omitempty"`
	Code        string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	RedirectUri string `protobuf:"bytes,3,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
}

func (x *SignInWithSSORequest) Reset() {
	*x = SignInWithSSORequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_auth_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInWithSSORequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInWithSSORequest) ProtoMessage() {}

func (x *SignInWithSSORequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_auth_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInWithSSORequest.ProtoReflect.Descriptor instead.
func (*SignInWithSSORequest) Descriptor() ([]byte, []int) {
	return file_api_v2_auth_service_proto_rawDescGZIP(), []int{4}
}

func (x *SignInWithSSORequest) GetIdpId() int32 {
	if x != nil {
		return x.IdpId
	}
	return 0
}

func (x *SignInWithSSORequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SignInWithSSORequest) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

type SignInWithSSOResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *SignInWithSSOResponse) Reset() {
	*x = SignInWithSSOResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_auth_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInWithSSOResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInWithSSOResponse) ProtoMessage() {}

func (x *SignInWithSSOResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_auth_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInWithSSOResponse.ProtoReflect.Descriptor instead.
func (*SignInWithSSOResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_auth_service_proto_rawDescGZIP(), []int{5}
}

func (x *SignInWithSSOResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type SignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SignUpRequest) Reset() {
	*x = SignUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_auth_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpRequest) ProtoMessage() {}

func (x *SignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_auth_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpRequest.ProtoReflect.Descriptor instead.
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_auth_service_proto_rawDescGZIP(), []int{6}
}

func (x *SignUpRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SignUpRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *SignUpResponse) Reset() {
	*x = SignUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_auth_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpResponse) ProtoMessage() {}

func (x *SignUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_auth_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpResponse.ProtoReflect.Descriptor instead.
func (*SignUpResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_auth_service_proto_rawDescGZIP(), []int{7}
}

func (x *SignUpResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type SignOutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignOutRequest) Reset() {
	*x = SignOutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_auth_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignOutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignOutRequest) ProtoMessage() {}

func (x *SignOutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_auth_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignOutRequest.ProtoReflect.Descriptor instead.
func (*SignOutRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_auth_service_proto_rawDescGZIP(), []int{8}
}

type SignOutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignOutResponse) Reset() {
	*x = SignOutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_auth_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignOutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignOutResponse) ProtoMessage() {}

func (x *SignOutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_auth_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignOutResponse.ProtoReflect.Descriptor instead.
func (*SignOutResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_auth_service_proto_rawDescGZIP(), []int{9}
}

var File_api_v2_auth_service_proto protoreflect.FileDescriptor

var file_api_v2_auth_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x65, 0x6d,
	0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x1a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x32, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x6a, 0x0a, 0x0d, 0x53,
	0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x76, 0x65, 0x72, 0x5f, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x6e, 0x65, 0x76, 0x65,
	0x72, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x22, 0x38, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x49,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x22, 0x64, 0x0a, 0x14, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x53,
	0x53, 0x4f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x64, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x64, 0x70, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x5f, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x69, 0x22, 0x3f, 0x0a, 0x15, 0x53, 0x69, 0x67, 0x6e, 0x49,
	0x6e, 0x57, 0x69, 0x74, 0x68, 0x53, 0x53, 0x4f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x47, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e,
	0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x38, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x10, 0x0a, 0x0e, 0x53,
	0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a,
	0x0f, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xa9, 0x04, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x75, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x22, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x22, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x60, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x49,
	0x6e, 0x12, 0x1b, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x22, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x12, 0x79, 0x0a, 0x0d, 0x53, 0x69, 0x67,
	0x6e, 0x49, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x53, 0x53, 0x4f, 0x12, 0x22, 0x2e, 0x6d, 0x65, 0x6d,
	0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e,
	0x57, 0x69, 0x74, 0x68, 0x53, 0x53, 0x4f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x49, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x53, 0x53, 0x4f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x17, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e,
	0x2f, 0x73, 0x73, 0x6f, 0x12, 0x60, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x1b,
	0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x65,
	0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x15, 0x22, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x64, 0x0a, 0x07, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75,
	0x74, 0x12, 0x1c, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x6f, 0x75, 0x74, 0x42, 0xa8, 0x01, 0x0a,
	0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x32, 0x42, 0x10, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x75, 0x73, 0x65, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x32, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x4d, 0x41, 0x58, 0xaa, 0x02, 0x0c,
	0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x0c, 0x4d,
	0x65, 0x6d, 0x6f, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x32, 0xe2, 0x02, 0x18, 0x4d, 0x65,
	0x6d, 0x6f, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x3a, 0x3a,
	0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v2_auth_service_proto_rawDescOnce sync.Once
	file_api_v2_auth_service_proto_rawDescData = file_api_v2_auth_service_proto_rawDesc
)

func file_api_v2_auth_service_proto_rawDescGZIP() []byte {
	file_api_v2_auth_service_proto_rawDescOnce.Do(func() {
		file_api_v2_auth_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v2_auth_service_proto_rawDescData)
	})
	return file_api_v2_auth_service_proto_rawDescData
}

var file_api_v2_auth_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_v2_auth_service_proto_goTypes = []interface{}{
	(*GetAuthStatusRequest)(nil),  // 0: memos.api.v2.GetAuthStatusRequest
	(*GetAuthStatusResponse)(nil), // 1: memos.api.v2.GetAuthStatusResponse
	(*SignInRequest)(nil),         // 2: memos.api.v2.SignInRequest
	(*SignInResponse)(nil),        // 3: memos.api.v2.SignInResponse
	(*SignInWithSSORequest)(nil),  // 4: memos.api.v2.SignInWithSSORequest
	(*SignInWithSSOResponse)(nil), // 5: memos.api.v2.SignInWithSSOResponse
	(*SignUpRequest)(nil),         // 6: memos.api.v2.SignUpRequest
	(*SignUpResponse)(nil),        // 7: memos.api.v2.SignUpResponse
	(*SignOutRequest)(nil),        // 8: memos.api.v2.SignOutRequest
	(*SignOutResponse)(nil),       // 9: memos.api.v2.SignOutResponse
	(*User)(nil),                  // 10: memos.api.v2.User
}
var file_api_v2_auth_service_proto_depIdxs = []int32{
	10, // 0: memos.api.v2.GetAuthStatusResponse.user:type_name -> memos.api.v2.User
	10, // 1: memos.api.v2.SignInResponse.user:type_name -> memos.api.v2.User
	10, // 2: memos.api.v2.SignInWithSSOResponse.user:type_name -> memos.api.v2.User
	10, // 3: memos.api.v2.SignUpResponse.user:type_name -> memos.api.v2.User
	0,  // 4: memos.api.v2.AuthService.GetAuthStatus:input_type -> memos.api.v2.GetAuthStatusRequest
	2,  // 5: memos.api.v2.AuthService.SignIn:input_type -> memos.api.v2.SignInRequest
	4,  // 6: memos.api.v2.AuthService.SignInWithSSO:input_type -> memos.api.v2.SignInWithSSORequest
	6,  // 7: memos.api.v2.AuthService.SignUp:input_type -> memos.api.v2.SignUpRequest
	8,  // 8: memos.api.v2.AuthService.SignOut:input_type -> memos.api.v2.SignOutRequest
	1,  // 9: memos.api.v2.AuthService.GetAuthStatus:output_type -> memos.api.v2.GetAuthStatusResponse
	3,  // 10: memos.api.v2.AuthService.SignIn:output_type -> memos.api.v2.SignInResponse
	5,  // 11: memos.api.v2.AuthService.SignInWithSSO:output_type -> memos.api.v2.SignInWithSSOResponse
	7,  // 12: memos.api.v2.AuthService.SignUp:output_type -> memos.api.v2.SignUpResponse
	9,  // 13: memos.api.v2.AuthService.SignOut:output_type -> memos.api.v2.SignOutResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_v2_auth_service_proto_init() }
func file_api_v2_auth_service_proto_init() {
	if File_api_v2_auth_service_proto != nil {
		return
	}
	file_api_v2_user_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v2_auth_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthStatusRequest); i {
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
		file_api_v2_auth_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthStatusResponse); i {
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
		file_api_v2_auth_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInRequest); i {
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
		file_api_v2_auth_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInResponse); i {
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
		file_api_v2_auth_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInWithSSORequest); i {
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
		file_api_v2_auth_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInWithSSOResponse); i {
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
		file_api_v2_auth_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpRequest); i {
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
		file_api_v2_auth_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpResponse); i {
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
		file_api_v2_auth_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignOutRequest); i {
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
		file_api_v2_auth_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignOutResponse); i {
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
			RawDescriptor: file_api_v2_auth_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v2_auth_service_proto_goTypes,
		DependencyIndexes: file_api_v2_auth_service_proto_depIdxs,
		MessageInfos:      file_api_v2_auth_service_proto_msgTypes,
	}.Build()
	File_api_v2_auth_service_proto = out.File
	file_api_v2_auth_service_proto_rawDesc = nil
	file_api_v2_auth_service_proto_goTypes = nil
	file_api_v2_auth_service_proto_depIdxs = nil
}