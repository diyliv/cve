// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/cve/cve.proto

package cvepb

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

type Creds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Creds) Reset() {
	*x = Creds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cve_cve_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Creds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Creds) ProtoMessage() {}

func (x *Creds) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cve_cve_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Creds.ProtoReflect.Descriptor instead.
func (*Creds) Descriptor() ([]byte, []int) {
	return file_proto_cve_cve_proto_rawDescGZIP(), []int{0}
}

func (x *Creds) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *Creds) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Cve struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Link        string `protobuf:"bytes,2,opt,name=link,proto3" json:"link,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Cve) Reset() {
	*x = Cve{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cve_cve_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cve) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cve) ProtoMessage() {}

func (x *Cve) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cve_cve_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cve.ProtoReflect.Descriptor instead.
func (*Cve) Descriptor() ([]byte, []int) {
	return file_proto_cve_cve_proto_rawDescGZIP(), []int{1}
}

func (x *Cve) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cve) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *Cve) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserCreds *Creds `protobuf:"bytes,1,opt,name=userCreds,proto3" json:"userCreds,omitempty"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cve_cve_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cve_cve_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_proto_cve_cve_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterReq) GetUserCreds() *Creds {
	if x != nil {
		return x.UserCreds
	}
	return nil
}

type RegisterResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *RegisterResp) Reset() {
	*x = RegisterResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cve_cve_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResp) ProtoMessage() {}

func (x *RegisterResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cve_cve_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResp.ProtoReflect.Descriptor instead.
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return file_proto_cve_cve_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterResp) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserCreds *Creds `protobuf:"bytes,1,opt,name=userCreds,proto3" json:"userCreds,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cve_cve_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cve_cve_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_proto_cve_cve_proto_rawDescGZIP(), []int{4}
}

func (x *LoginReq) GetUserCreds() *Creds {
	if x != nil {
		return x.UserCreds
	}
	return nil
}

type LoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cve_cve_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cve_cve_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_proto_cve_cve_proto_rawDescGZIP(), []int{5}
}

func (x *LoginResp) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type FindVulnerabilitiesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search string `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *FindVulnerabilitiesReq) Reset() {
	*x = FindVulnerabilitiesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cve_cve_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindVulnerabilitiesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindVulnerabilitiesReq) ProtoMessage() {}

func (x *FindVulnerabilitiesReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cve_cve_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindVulnerabilitiesReq.ProtoReflect.Descriptor instead.
func (*FindVulnerabilitiesReq) Descriptor() ([]byte, []int) {
	return file_proto_cve_cve_proto_rawDescGZIP(), []int{6}
}

func (x *FindVulnerabilitiesReq) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type FindVulnerabilitiesResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Cve `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *FindVulnerabilitiesResp) Reset() {
	*x = FindVulnerabilitiesResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cve_cve_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindVulnerabilitiesResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindVulnerabilitiesResp) ProtoMessage() {}

func (x *FindVulnerabilitiesResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cve_cve_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindVulnerabilitiesResp.ProtoReflect.Descriptor instead.
func (*FindVulnerabilitiesResp) Descriptor() ([]byte, []int) {
	return file_proto_cve_cve_proto_rawDescGZIP(), []int{7}
}

func (x *FindVulnerabilitiesResp) GetResult() *Cve {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_proto_cve_cve_proto protoreflect.FileDescriptor

var file_proto_cve_cve_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x76, 0x65, 0x2f, 0x63, 0x76, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63, 0x76, 0x65, 0x22, 0x39, 0x0a, 0x05, 0x43, 0x72,
	0x65, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x4f, 0x0a, 0x03, 0x43, 0x76, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x28, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x76, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x64, 0x73, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x73, 0x22,
	0x26, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x34, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x28, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x76, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x64, 0x73, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x73, 0x22, 0x23, 0x0a,
	0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x30, 0x0a, 0x16, 0x46, 0x69, 0x6e, 0x64, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x22, 0x3b, 0x0a, 0x17, 0x46, 0x69, 0x6e, 0x64, 0x56, 0x75, 0x6c, 0x6e,
	0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x20, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x63, 0x76, 0x65, 0x2e, 0x43, 0x76, 0x65, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x32, 0xbd, 0x01, 0x0a, 0x0a, 0x43, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x31, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x63,
	0x76, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x11,
	0x2e, 0x63, 0x76, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0d, 0x2e, 0x63,
	0x76, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x63, 0x76,
	0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x52, 0x0a,
	0x13, 0x46, 0x69, 0x6e, 0x64, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x12, 0x1b, 0x2e, 0x63, 0x76, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x56,
	0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x1a, 0x1c, 0x2e, 0x63, 0x76, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x56, 0x75, 0x6c, 0x6e,
	0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x63, 0x76, 0x65, 0x70, 0x62, 0x3b, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_cve_cve_proto_rawDescOnce sync.Once
	file_proto_cve_cve_proto_rawDescData = file_proto_cve_cve_proto_rawDesc
)

func file_proto_cve_cve_proto_rawDescGZIP() []byte {
	file_proto_cve_cve_proto_rawDescOnce.Do(func() {
		file_proto_cve_cve_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cve_cve_proto_rawDescData)
	})
	return file_proto_cve_cve_proto_rawDescData
}

var file_proto_cve_cve_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_cve_cve_proto_goTypes = []interface{}{
	(*Creds)(nil),                   // 0: cve.Creds
	(*Cve)(nil),                     // 1: cve.Cve
	(*RegisterReq)(nil),             // 2: cve.RegisterReq
	(*RegisterResp)(nil),            // 3: cve.RegisterResp
	(*LoginReq)(nil),                // 4: cve.LoginReq
	(*LoginResp)(nil),               // 5: cve.LoginResp
	(*FindVulnerabilitiesReq)(nil),  // 6: cve.FindVulnerabilitiesReq
	(*FindVulnerabilitiesResp)(nil), // 7: cve.FindVulnerabilitiesResp
}
var file_proto_cve_cve_proto_depIdxs = []int32{
	0, // 0: cve.RegisterReq.userCreds:type_name -> cve.Creds
	0, // 1: cve.LoginReq.userCreds:type_name -> cve.Creds
	1, // 2: cve.FindVulnerabilitiesResp.result:type_name -> cve.Cve
	2, // 3: cve.CveService.Register:input_type -> cve.RegisterReq
	4, // 4: cve.CveService.Login:input_type -> cve.LoginReq
	6, // 5: cve.CveService.FindVulnerabilities:input_type -> cve.FindVulnerabilitiesReq
	3, // 6: cve.CveService.Register:output_type -> cve.RegisterResp
	5, // 7: cve.CveService.Login:output_type -> cve.LoginResp
	7, // 8: cve.CveService.FindVulnerabilities:output_type -> cve.FindVulnerabilitiesResp
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_cve_cve_proto_init() }
func file_proto_cve_cve_proto_init() {
	if File_proto_cve_cve_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_cve_cve_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Creds); i {
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
		file_proto_cve_cve_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cve); i {
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
		file_proto_cve_cve_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReq); i {
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
		file_proto_cve_cve_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResp); i {
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
		file_proto_cve_cve_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_proto_cve_cve_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResp); i {
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
		file_proto_cve_cve_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindVulnerabilitiesReq); i {
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
		file_proto_cve_cve_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindVulnerabilitiesResp); i {
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
			RawDescriptor: file_proto_cve_cve_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cve_cve_proto_goTypes,
		DependencyIndexes: file_proto_cve_cve_proto_depIdxs,
		MessageInfos:      file_proto_cve_cve_proto_msgTypes,
	}.Build()
	File_proto_cve_cve_proto = out.File
	file_proto_cve_cve_proto_rawDesc = nil
	file_proto_cve_cve_proto_goTypes = nil
	file_proto_cve_cve_proto_depIdxs = nil
}
