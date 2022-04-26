// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: tenant.proto

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

type StatisticTenantCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatisticTenantCountRequest) Reset() {
	*x = StatisticTenantCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatisticTenantCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatisticTenantCountRequest) ProtoMessage() {}

func (x *StatisticTenantCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatisticTenantCountRequest.ProtoReflect.Descriptor instead.
func (*StatisticTenantCountRequest) Descriptor() ([]byte, []int) {
	return file_tenant_proto_rawDescGZIP(), []int{0}
}

type TenantInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: validate:"required"
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" validate:"required"`
	Contact       string `protobuf:"bytes,3,opt,name=contact,proto3" json:"contact,omitempty"`
	CellPhone     string `protobuf:"bytes,4,opt,name=cellPhone,proto3" json:"cellPhone,omitempty"`
	Address       string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	BusinessScope string `protobuf:"bytes,6,opt,name=businessScope,proto3" json:"businessScope,omitempty"`
	AreaCovered   string `protobuf:"bytes,7,opt,name=areaCovered,proto3" json:"areaCovered,omitempty"`
	StaffSize     int32  `protobuf:"varint,8,opt,name=staffSize,proto3" json:"staffSize,omitempty"`
	Enable        bool   `protobuf:"varint,9,opt,name=enable,proto3" json:"enable,omitempty"`
	//省份
	Province string `protobuf:"bytes,10,opt,name=province,proto3" json:"province,omitempty"`
	//城市
	City string `protobuf:"bytes,11,opt,name=city,proto3" json:"city,omitempty"`
	//区/县
	Area string `protobuf:"bytes,12,opt,name=area,proto3" json:"area,omitempty"`
	//街道/镇
	Town string `protobuf:"bytes,13,opt,name=town,proto3" json:"town,omitempty"`
}

func (x *TenantInfo) Reset() {
	*x = TenantInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantInfo) ProtoMessage() {}

func (x *TenantInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantInfo.ProtoReflect.Descriptor instead.
func (*TenantInfo) Descriptor() ([]byte, []int) {
	return file_tenant_proto_rawDescGZIP(), []int{1}
}

func (x *TenantInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TenantInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TenantInfo) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

func (x *TenantInfo) GetCellPhone() string {
	if x != nil {
		return x.CellPhone
	}
	return ""
}

func (x *TenantInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *TenantInfo) GetBusinessScope() string {
	if x != nil {
		return x.BusinessScope
	}
	return ""
}

func (x *TenantInfo) GetAreaCovered() string {
	if x != nil {
		return x.AreaCovered
	}
	return ""
}

func (x *TenantInfo) GetStaffSize() int32 {
	if x != nil {
		return x.StaffSize
	}
	return 0
}

func (x *TenantInfo) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *TenantInfo) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *TenantInfo) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *TenantInfo) GetArea() string {
	if x != nil {
		return x.Area
	}
	return ""
}

func (x *TenantInfo) GetTown() string {
	if x != nil {
		return x.Town
	}
	return ""
}

type QueryTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex,omitempty" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"orderField" form:"orderField"
	OrderField string `protobuf:"bytes,3,opt,name=orderField,proto3" json:"orderField,omitempty" uri:"orderField" form:"orderField"`
	// @inject_tag: uri:"desc" form:"desc"
	Desc bool `protobuf:"varint,4,opt,name=desc,proto3" json:"desc,omitempty" uri:"desc" form:"desc"`
	// @inject_tag: uri:"name" form:"name"
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty" uri:"name" form:"name"`
	//省份
	// @inject_tag: uri:"province" form:"province"
	Province string `protobuf:"bytes,6,opt,name=province,proto3" json:"province,omitempty" uri:"province" form:"province"`
	//城市
	// @inject_tag: uri:"city" form:"city"
	City string `protobuf:"bytes,7,opt,name=city,proto3" json:"city,omitempty" uri:"city" form:"city"`
	//区/县
	// @inject_tag: uri:"area" form:"area"
	Area string `protobuf:"bytes,8,opt,name=area,proto3" json:"area,omitempty" uri:"area" form:"area"`
	//街道/镇
	// @inject_tag: uri:"town" form:"town"
	Town string `protobuf:"bytes,9,opt,name=town,proto3" json:"town,omitempty" uri:"town" form:"town"`
}

func (x *QueryTenantRequest) Reset() {
	*x = QueryTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTenantRequest) ProtoMessage() {}

func (x *QueryTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTenantRequest.ProtoReflect.Descriptor instead.
func (*QueryTenantRequest) Descriptor() ([]byte, []int) {
	return file_tenant_proto_rawDescGZIP(), []int{2}
}

func (x *QueryTenantRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryTenantRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryTenantRequest) GetOrderField() string {
	if x != nil {
		return x.OrderField
	}
	return ""
}

func (x *QueryTenantRequest) GetDesc() bool {
	if x != nil {
		return x.Desc
	}
	return false
}

func (x *QueryTenantRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *QueryTenantRequest) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *QueryTenantRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *QueryTenantRequest) GetArea() string {
	if x != nil {
		return x.Area
	}
	return ""
}

func (x *QueryTenantRequest) GetTown() string {
	if x != nil {
		return x.Town
	}
	return ""
}

type QueryTenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code          `protobuf:"varint,1,opt,name=code,proto3,enum=usercenter.Code" json:"code,omitempty"`
	Message string        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*TenantInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	Pages   int64         `protobuf:"varint,4,opt,name=pages,proto3" json:"pages,omitempty"`
	Records int64         `protobuf:"varint,5,opt,name=records,proto3" json:"records,omitempty"`
	Total   int64         `protobuf:"varint,6,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *QueryTenantResponse) Reset() {
	*x = QueryTenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryTenantResponse) ProtoMessage() {}

func (x *QueryTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryTenantResponse.ProtoReflect.Descriptor instead.
func (*QueryTenantResponse) Descriptor() ([]byte, []int) {
	return file_tenant_proto_rawDescGZIP(), []int{3}
}

func (x *QueryTenantResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryTenantResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryTenantResponse) GetData() []*TenantInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryTenantResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryTenantResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryTenantResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllTenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code          `protobuf:"varint,1,opt,name=code,proto3,enum=usercenter.Code" json:"code,omitempty"`
	Message string        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*TenantInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAllTenantResponse) Reset() {
	*x = GetAllTenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTenantResponse) ProtoMessage() {}

func (x *GetAllTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTenantResponse.ProtoReflect.Descriptor instead.
func (*GetAllTenantResponse) Descriptor() ([]byte, []int) {
	return file_tenant_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllTenantResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllTenantResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllTenantResponse) GetData() []*TenantInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetTenantDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code        `protobuf:"varint,1,opt,name=code,proto3,enum=usercenter.Code" json:"code,omitempty"`
	Message string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *TenantInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetTenantDetailResponse) Reset() {
	*x = GetTenantDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTenantDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTenantDetailResponse) ProtoMessage() {}

func (x *GetTenantDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTenantDetailResponse.ProtoReflect.Descriptor instead.
func (*GetTenantDetailResponse) Descriptor() ([]byte, []int) {
	return file_tenant_proto_rawDescGZIP(), []int{5}
}

func (x *GetTenantDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetTenantDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetTenantDetailResponse) GetData() *TenantInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_tenant_proto protoreflect.FileDescriptor

var file_tenant_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x1b, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xd8, 0x02, 0x0a, 0x0a, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x65, 0x6c, 0x6c, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x65, 0x6c, 0x6c, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x0d,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x65,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x76,
	0x65, 0x72, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x66, 0x66, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x61, 0x66, 0x66, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72,
	0x65, 0x61, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x6f, 0x77, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x6f,
	0x77, 0x6e, 0x22, 0xee, 0x01, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x72, 0x65, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x6f, 0x77, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x6f, 0x77, 0x6e, 0x22, 0xc7, 0x01, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x82, 0x01,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x85, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xd0, 0x04, 0x0a, 0x06, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x16, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4a, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1e, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41,
	0x0a, 0x06, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x47, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x19, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x0e,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a,
	0x0d, 0x2e, 0x2f, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tenant_proto_rawDescOnce sync.Once
	file_tenant_proto_rawDescData = file_tenant_proto_rawDesc
)

func file_tenant_proto_rawDescGZIP() []byte {
	file_tenant_proto_rawDescOnce.Do(func() {
		file_tenant_proto_rawDescData = protoimpl.X.CompressGZIP(file_tenant_proto_rawDescData)
	})
	return file_tenant_proto_rawDescData
}

var file_tenant_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tenant_proto_goTypes = []interface{}{
	(*StatisticTenantCountRequest)(nil), // 0: usercenter.StatisticTenantCountRequest
	(*TenantInfo)(nil),                  // 1: usercenter.TenantInfo
	(*QueryTenantRequest)(nil),          // 2: usercenter.QueryTenantRequest
	(*QueryTenantResponse)(nil),         // 3: usercenter.QueryTenantResponse
	(*GetAllTenantResponse)(nil),        // 4: usercenter.GetAllTenantResponse
	(*GetTenantDetailResponse)(nil),     // 5: usercenter.GetTenantDetailResponse
	(Code)(0),                           // 6: usercenter.Code
	(*DelRequest)(nil),                  // 7: usercenter.DelRequest
	(*EnableRequest)(nil),               // 8: usercenter.EnableRequest
	(*GetAllRequest)(nil),               // 9: usercenter.GetAllRequest
	(*GetDetailRequest)(nil),            // 10: usercenter.GetDetailRequest
	(*CommonResponse)(nil),              // 11: usercenter.CommonResponse
	(*StatisticCountResponse)(nil),      // 12: usercenter.StatisticCountResponse
}
var file_tenant_proto_depIdxs = []int32{
	6,  // 0: usercenter.QueryTenantResponse.code:type_name -> usercenter.Code
	1,  // 1: usercenter.QueryTenantResponse.data:type_name -> usercenter.TenantInfo
	6,  // 2: usercenter.GetAllTenantResponse.code:type_name -> usercenter.Code
	1,  // 3: usercenter.GetAllTenantResponse.data:type_name -> usercenter.TenantInfo
	6,  // 4: usercenter.GetTenantDetailResponse.code:type_name -> usercenter.Code
	1,  // 5: usercenter.GetTenantDetailResponse.data:type_name -> usercenter.TenantInfo
	1,  // 6: usercenter.Tenant.Add:input_type -> usercenter.TenantInfo
	1,  // 7: usercenter.Tenant.Update:input_type -> usercenter.TenantInfo
	7,  // 8: usercenter.Tenant.Delete:input_type -> usercenter.DelRequest
	2,  // 9: usercenter.Tenant.Query:input_type -> usercenter.QueryTenantRequest
	8,  // 10: usercenter.Tenant.Enable:input_type -> usercenter.EnableRequest
	9,  // 11: usercenter.Tenant.GetAll:input_type -> usercenter.GetAllRequest
	10, // 12: usercenter.Tenant.GetDetail:input_type -> usercenter.GetDetailRequest
	0,  // 13: usercenter.Tenant.StatisticCount:input_type -> usercenter.StatisticTenantCountRequest
	11, // 14: usercenter.Tenant.Add:output_type -> usercenter.CommonResponse
	11, // 15: usercenter.Tenant.Update:output_type -> usercenter.CommonResponse
	11, // 16: usercenter.Tenant.Delete:output_type -> usercenter.CommonResponse
	3,  // 17: usercenter.Tenant.Query:output_type -> usercenter.QueryTenantResponse
	11, // 18: usercenter.Tenant.Enable:output_type -> usercenter.CommonResponse
	4,  // 19: usercenter.Tenant.GetAll:output_type -> usercenter.GetAllTenantResponse
	5,  // 20: usercenter.Tenant.GetDetail:output_type -> usercenter.GetTenantDetailResponse
	12, // 21: usercenter.Tenant.StatisticCount:output_type -> usercenter.StatisticCountResponse
	14, // [14:22] is the sub-list for method output_type
	6,  // [6:14] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_tenant_proto_init() }
func file_tenant_proto_init() {
	if File_tenant_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tenant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatisticTenantCountRequest); i {
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
		file_tenant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantInfo); i {
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
		file_tenant_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryTenantRequest); i {
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
		file_tenant_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryTenantResponse); i {
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
		file_tenant_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTenantResponse); i {
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
		file_tenant_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTenantDetailResponse); i {
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
			RawDescriptor: file_tenant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tenant_proto_goTypes,
		DependencyIndexes: file_tenant_proto_depIdxs,
		MessageInfos:      file_tenant_proto_msgTypes,
	}.Build()
	File_tenant_proto = out.File
	file_tenant_proto_rawDesc = nil
	file_tenant_proto_goTypes = nil
	file_tenant_proto_depIdxs = nil
}
