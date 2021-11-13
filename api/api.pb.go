// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: api.proto

package usercenter

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

type APIInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TenantID    string `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	Path        string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Group       string `protobuf:"bytes,4,opt,name=group,proto3" json:"group,omitempty"`
	Method      string `protobuf:"bytes,5,opt,name=method,proto3" json:"method,omitempty"`
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Enable      bool   `protobuf:"varint,7,opt,name=enable,proto3" json:"enable,omitempty"`
	CheckAuth   bool   `protobuf:"varint,8,opt,name=checkAuth,proto3" json:"checkAuth,omitempty"`
	CheckLogin  bool   `protobuf:"varint,9,opt,name=checkLogin,proto3" json:"checkLogin,omitempty"`
}

func (x *APIInfo) Reset() {
	*x = APIInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIInfo) ProtoMessage() {}

func (x *APIInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIInfo.ProtoReflect.Descriptor instead.
func (*APIInfo) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *APIInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *APIInfo) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *APIInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *APIInfo) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *APIInfo) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *APIInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *APIInfo) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *APIInfo) GetCheckAuth() bool {
	if x != nil {
		return x.CheckAuth
	}
	return false
}

func (x *APIInfo) GetCheckLogin() bool {
	if x != nil {
		return x.CheckLogin
	}
	return false
}

type QueryAPIRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageIndex  int64  `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	PageSize   int64  `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderField string `protobuf:"bytes,3,opt,name=orderField,proto3" json:"orderField,omitempty"`
	Desc       bool   `protobuf:"varint,4,opt,name=desc,proto3" json:"desc,omitempty"`
	Path       string `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	Method     string `protobuf:"bytes,6,opt,name=method,proto3" json:"method,omitempty"`
	Group      string `protobuf:"bytes,7,opt,name=group,proto3" json:"group,omitempty"`
	CheckAuth  int32  `protobuf:"varint,8,opt,name=checkAuth,proto3" json:"checkAuth,omitempty"`
	CheckLogin int32  `protobuf:"varint,9,opt,name=checkLogin,proto3" json:"checkLogin,omitempty"`
}

func (x *QueryAPIRequest) Reset() {
	*x = QueryAPIRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAPIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAPIRequest) ProtoMessage() {}

func (x *QueryAPIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAPIRequest.ProtoReflect.Descriptor instead.
func (*QueryAPIRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *QueryAPIRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryAPIRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryAPIRequest) GetOrderField() string {
	if x != nil {
		return x.OrderField
	}
	return ""
}

func (x *QueryAPIRequest) GetDesc() bool {
	if x != nil {
		return x.Desc
	}
	return false
}

func (x *QueryAPIRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *QueryAPIRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *QueryAPIRequest) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *QueryAPIRequest) GetCheckAuth() int32 {
	if x != nil {
		return x.CheckAuth
	}
	return 0
}

func (x *QueryAPIRequest) GetCheckLogin() int32 {
	if x != nil {
		return x.CheckLogin
	}
	return 0
}

type QueryAPIResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code       `protobuf:"varint,1,opt,name=code,proto3,enum=usercenter.Code" json:"code,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*APIInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	Pages   int64      `protobuf:"varint,4,opt,name=pages,proto3" json:"pages,omitempty"`
	Records int64      `protobuf:"varint,5,opt,name=records,proto3" json:"records,omitempty"`
	Total   int64      `protobuf:"varint,6,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *QueryAPIResponse) Reset() {
	*x = QueryAPIResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAPIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAPIResponse) ProtoMessage() {}

func (x *QueryAPIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAPIResponse.ProtoReflect.Descriptor instead.
func (*QueryAPIResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *QueryAPIResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_Success
}

func (x *QueryAPIResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryAPIResponse) GetData() []*APIInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryAPIResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryAPIResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryAPIResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllAPIResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code       `protobuf:"varint,1,opt,name=code,proto3,enum=usercenter.Code" json:"code,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*APIInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAllAPIResponse) Reset() {
	*x = GetAllAPIResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAPIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAPIResponse) ProtoMessage() {}

func (x *GetAllAPIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAPIResponse.ProtoReflect.Descriptor instead.
func (*GetAllAPIResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllAPIResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_Success
}

func (x *GetAllAPIResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllAPIResponse) GetData() []*APIInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetAPIDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code     `protobuf:"varint,1,opt,name=code,proto3,enum=usercenter.Code" json:"code,omitempty"`
	Message string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *APIInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAPIDetailResponse) Reset() {
	*x = GetAPIDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAPIDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAPIDetailResponse) ProtoMessage() {}

func (x *GetAPIDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAPIDetailResponse.ProtoReflect.Descriptor instead.
func (*GetAPIDetailResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetAPIDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_Success
}

func (x *GetAPIDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAPIDetailResponse) GetData() *APIInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x75, 0x73, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x01, 0x0a, 0x07, 0x41, 0x50, 0x49, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0xff, 0x01, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x41, 0x50, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0xc1, 0x01, 0x0a, 0x10, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x41, 0x50, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x27,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x50, 0x49, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x7c, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x50, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x50,
	0x49, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x7f, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x41, 0x50, 0x49, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x41,
	0x50, 0x49, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xda, 0x03, 0x0a,
	0x03, 0x41, 0x50, 0x49, 0x12, 0x38, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x13, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x50, 0x49, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x50, 0x49, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1a, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x05, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x50, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x41, 0x50, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x41, 0x0a, 0x06, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x19, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x19,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x50, 0x49,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x50, 0x49, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x3b,
	0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_proto_goTypes = []interface{}{
	(*APIInfo)(nil),              // 0: usercenter.APIInfo
	(*QueryAPIRequest)(nil),      // 1: usercenter.QueryAPIRequest
	(*QueryAPIResponse)(nil),     // 2: usercenter.QueryAPIResponse
	(*GetAllAPIResponse)(nil),    // 3: usercenter.GetAllAPIResponse
	(*GetAPIDetailResponse)(nil), // 4: usercenter.GetAPIDetailResponse
	(Code)(0),                    // 5: usercenter.Code
	(*DelRequest)(nil),           // 6: usercenter.DelRequest
	(*EnableRequest)(nil),        // 7: usercenter.EnableRequest
	(*GetAllRequest)(nil),        // 8: usercenter.GetAllRequest
	(*GetDetailRequest)(nil),     // 9: usercenter.GetDetailRequest
	(*CommonResponse)(nil),       // 10: usercenter.CommonResponse
}
var file_api_proto_depIdxs = []int32{
	5,  // 0: usercenter.QueryAPIResponse.code:type_name -> usercenter.Code
	0,  // 1: usercenter.QueryAPIResponse.data:type_name -> usercenter.APIInfo
	5,  // 2: usercenter.GetAllAPIResponse.code:type_name -> usercenter.Code
	0,  // 3: usercenter.GetAllAPIResponse.data:type_name -> usercenter.APIInfo
	5,  // 4: usercenter.GetAPIDetailResponse.code:type_name -> usercenter.Code
	0,  // 5: usercenter.GetAPIDetailResponse.data:type_name -> usercenter.APIInfo
	0,  // 6: usercenter.API.Add:input_type -> usercenter.APIInfo
	0,  // 7: usercenter.API.Update:input_type -> usercenter.APIInfo
	6,  // 8: usercenter.API.Delete:input_type -> usercenter.DelRequest
	1,  // 9: usercenter.API.Query:input_type -> usercenter.QueryAPIRequest
	7,  // 10: usercenter.API.Enable:input_type -> usercenter.EnableRequest
	8,  // 11: usercenter.API.GetAll:input_type -> usercenter.GetAllRequest
	9,  // 12: usercenter.API.GetDetail:input_type -> usercenter.GetDetailRequest
	10, // 13: usercenter.API.Add:output_type -> usercenter.CommonResponse
	10, // 14: usercenter.API.Update:output_type -> usercenter.CommonResponse
	10, // 15: usercenter.API.Delete:output_type -> usercenter.CommonResponse
	2,  // 16: usercenter.API.Query:output_type -> usercenter.QueryAPIResponse
	10, // 17: usercenter.API.Enable:output_type -> usercenter.CommonResponse
	3,  // 18: usercenter.API.GetAll:output_type -> usercenter.GetAllAPIResponse
	4,  // 19: usercenter.API.GetDetail:output_type -> usercenter.GetAPIDetailResponse
	13, // [13:20] is the sub-list for method output_type
	6,  // [6:13] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIInfo); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAPIRequest); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAPIResponse); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAPIResponse); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAPIDetailResponse); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
