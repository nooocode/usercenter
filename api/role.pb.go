// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: role.proto

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

type StatisticRoleCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantID string `protobuf:"bytes,1,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
}

func (x *StatisticRoleCountRequest) Reset() {
	*x = StatisticRoleCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatisticRoleCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatisticRoleCountRequest) ProtoMessage() {}

func (x *StatisticRoleCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatisticRoleCountRequest.ProtoReflect.Descriptor instead.
func (*StatisticRoleCountRequest) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{0}
}

func (x *StatisticRoleCountRequest) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

type RoleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TenantID string `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	// @inject_tag: validate:"required"
	Name          string      `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" validate:"required"`
	ParentID      string      `protobuf:"bytes,4,opt,name=parentID,proto3" json:"parentID,omitempty"`
	Children      []*RoleInfo `protobuf:"bytes,5,rep,name=children,proto3" json:"children,omitempty"`
	RoleMenus     []*RoleMenu `protobuf:"bytes,6,rep,name=roleMenus,proto3" json:"roleMenus,omitempty"`
	DefaultRouter string      `protobuf:"bytes,7,opt,name=defaultRouter,proto3" json:"defaultRouter,omitempty"`
	Description   string      `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	CanDel        bool        `protobuf:"varint,9,opt,name=canDel,proto3" json:"canDel,omitempty"`
	TenantName    string      `protobuf:"bytes,10,opt,name=tenantName,proto3" json:"tenantName,omitempty"`
}

func (x *RoleInfo) Reset() {
	*x = RoleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleInfo) ProtoMessage() {}

func (x *RoleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleInfo.ProtoReflect.Descriptor instead.
func (*RoleInfo) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{1}
}

func (x *RoleInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoleInfo) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *RoleInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleInfo) GetParentID() string {
	if x != nil {
		return x.ParentID
	}
	return ""
}

func (x *RoleInfo) GetChildren() []*RoleInfo {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *RoleInfo) GetRoleMenus() []*RoleMenu {
	if x != nil {
		return x.RoleMenus
	}
	return nil
}

func (x *RoleInfo) GetDefaultRouter() string {
	if x != nil {
		return x.DefaultRouter
	}
	return ""
}

func (x *RoleInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RoleInfo) GetCanDel() bool {
	if x != nil {
		return x.CanDel
	}
	return false
}

func (x *RoleInfo) GetTenantName() string {
	if x != nil {
		return x.TenantName
	}
	return ""
}

type RoleMenu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RoleID string    `protobuf:"bytes,2,opt,name=roleID,proto3" json:"roleID,omitempty"`
	MenuID string    `protobuf:"bytes,3,opt,name=menuID,proto3" json:"menuID,omitempty"`
	Funcs  string    `protobuf:"bytes,4,opt,name=funcs,proto3" json:"funcs,omitempty"`
	Menu   *MenuInfo `protobuf:"bytes,5,opt,name=menu,proto3" json:"menu,omitempty"`
}

func (x *RoleMenu) Reset() {
	*x = RoleMenu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleMenu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleMenu) ProtoMessage() {}

func (x *RoleMenu) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleMenu.ProtoReflect.Descriptor instead.
func (*RoleMenu) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{2}
}

func (x *RoleMenu) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoleMenu) GetRoleID() string {
	if x != nil {
		return x.RoleID
	}
	return ""
}

func (x *RoleMenu) GetMenuID() string {
	if x != nil {
		return x.MenuID
	}
	return ""
}

func (x *RoleMenu) GetFuncs() string {
	if x != nil {
		return x.Funcs
	}
	return ""
}

func (x *RoleMenu) GetMenu() *MenuInfo {
	if x != nil {
		return x.Menu
	}
	return nil
}

type QueryRoleRequest struct {
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
	//租户ID
	// @inject_tag: uri:"tenantID" form:"tenantID"
	TenantID string `protobuf:"bytes,6,opt,name=tenantID,proto3" json:"tenantID,omitempty" uri:"tenantID" form:"tenantID"`
	// true-包含公共角色 false-不包含公共角色
	//公共角色定义：不设置租户的角色
	// @inject_tag: uri:"containerComm" form:"containerComm"
	ContainerComm bool `protobuf:"varint,7,opt,name=containerComm,proto3" json:"containerComm,omitempty" uri:"containerComm" form:"containerComm"`
}

func (x *QueryRoleRequest) Reset() {
	*x = QueryRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRoleRequest) ProtoMessage() {}

func (x *QueryRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRoleRequest.ProtoReflect.Descriptor instead.
func (*QueryRoleRequest) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{3}
}

func (x *QueryRoleRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryRoleRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryRoleRequest) GetOrderField() string {
	if x != nil {
		return x.OrderField
	}
	return ""
}

func (x *QueryRoleRequest) GetDesc() bool {
	if x != nil {
		return x.Desc
	}
	return false
}

func (x *QueryRoleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *QueryRoleRequest) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *QueryRoleRequest) GetContainerComm() bool {
	if x != nil {
		return x.ContainerComm
	}
	return false
}

type QueryRoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code        `protobuf:"varint,1,opt,name=code,proto3,enum=usercenter.Code" json:"code,omitempty"`
	Message string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*RoleInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	Pages   int64       `protobuf:"varint,4,opt,name=pages,proto3" json:"pages,omitempty"`
	Records int64       `protobuf:"varint,5,opt,name=records,proto3" json:"records,omitempty"`
	Total   int64       `protobuf:"varint,6,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *QueryRoleResponse) Reset() {
	*x = QueryRoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRoleResponse) ProtoMessage() {}

func (x *QueryRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRoleResponse.ProtoReflect.Descriptor instead.
func (*QueryRoleResponse) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{4}
}

func (x *QueryRoleResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryRoleResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryRoleResponse) GetData() []*RoleInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryRoleResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryRoleResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryRoleResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllRoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code        `protobuf:"varint,1,opt,name=code,proto3,enum=usercenter.Code" json:"code,omitempty"`
	Message string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*RoleInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAllRoleResponse) Reset() {
	*x = GetAllRoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRoleResponse) ProtoMessage() {}

func (x *GetAllRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRoleResponse.ProtoReflect.Descriptor instead.
func (*GetAllRoleResponse) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllRoleResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllRoleResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllRoleResponse) GetData() []*RoleInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetAllRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"tenantID" form:"tenantID"
	TenantID string `protobuf:"bytes,1,opt,name=tenantID,proto3" json:"tenantID,omitempty" uri:"tenantID" form:"tenantID"`
	// true-包含公共角色 false-不包含公共角色
	//公共角色定义：不设置租户的角色
	// @inject_tag: uri:"containerComm" form:"containerComm"
	ContainerComm bool `protobuf:"varint,2,opt,name=containerComm,proto3" json:"containerComm,omitempty" uri:"containerComm" form:"containerComm"`
}

func (x *GetAllRoleRequest) Reset() {
	*x = GetAllRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRoleRequest) ProtoMessage() {}

func (x *GetAllRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRoleRequest.ProtoReflect.Descriptor instead.
func (*GetAllRoleRequest) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllRoleRequest) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *GetAllRoleRequest) GetContainerComm() bool {
	if x != nil {
		return x.ContainerComm
	}
	return false
}

type GetRoleDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code      `protobuf:"varint,1,opt,name=code,proto3,enum=usercenter.Code" json:"code,omitempty"`
	Message string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *RoleInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetRoleDetailResponse) Reset() {
	*x = GetRoleDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoleDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleDetailResponse) ProtoMessage() {}

func (x *GetRoleDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleDetailResponse.ProtoReflect.Descriptor instead.
func (*GetRoleDetailResponse) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{7}
}

func (x *GetRoleDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetRoleDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetRoleDetailResponse) GetData() *RoleInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_role_proto protoreflect.FileDescriptor

var file_role_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x37, 0x0a, 0x19, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x52,
	0x6f, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x22, 0xcc, 0x02, 0x0a, 0x08,
	0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x32, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x4d, 0x65,
	0x6e, 0x75, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52,
	0x09, 0x72, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x6e, 0x44, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x63, 0x61, 0x6e, 0x44, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x8a, 0x01, 0x0a, 0x08, 0x52,
	0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6e, 0x75, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x65, 0x6e, 0x75, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x75, 0x6e, 0x63, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x75, 0x6e, 0x63, 0x73, 0x12, 0x28, 0x0a,
	0x04, 0x6d, 0x65, 0x6e, 0x75, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x22, 0xd6, 0x01, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d,
	0x22, 0xc3, 0x01, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x7e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x55, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x22, 0x81, 0x01,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x32, 0x81, 0x04, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x39, 0x0a, 0x03, 0x41, 0x64,
	0x64, 0x12, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x52,
	0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1c, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x06, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x52,
	0x6f, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x3b, 0x75, 0x73, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_role_proto_rawDescOnce sync.Once
	file_role_proto_rawDescData = file_role_proto_rawDesc
)

func file_role_proto_rawDescGZIP() []byte {
	file_role_proto_rawDescOnce.Do(func() {
		file_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_role_proto_rawDescData)
	})
	return file_role_proto_rawDescData
}

var file_role_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_role_proto_goTypes = []interface{}{
	(*StatisticRoleCountRequest)(nil), // 0: usercenter.StatisticRoleCountRequest
	(*RoleInfo)(nil),                  // 1: usercenter.RoleInfo
	(*RoleMenu)(nil),                  // 2: usercenter.RoleMenu
	(*QueryRoleRequest)(nil),          // 3: usercenter.QueryRoleRequest
	(*QueryRoleResponse)(nil),         // 4: usercenter.QueryRoleResponse
	(*GetAllRoleResponse)(nil),        // 5: usercenter.GetAllRoleResponse
	(*GetAllRoleRequest)(nil),         // 6: usercenter.GetAllRoleRequest
	(*GetRoleDetailResponse)(nil),     // 7: usercenter.GetRoleDetailResponse
	(*MenuInfo)(nil),                  // 8: usercenter.MenuInfo
	(Code)(0),                         // 9: usercenter.Code
	(*DelRequest)(nil),                // 10: usercenter.DelRequest
	(*GetDetailRequest)(nil),          // 11: usercenter.GetDetailRequest
	(*CommonResponse)(nil),            // 12: usercenter.CommonResponse
	(*StatisticCountResponse)(nil),    // 13: usercenter.StatisticCountResponse
}
var file_role_proto_depIdxs = []int32{
	1,  // 0: usercenter.RoleInfo.children:type_name -> usercenter.RoleInfo
	2,  // 1: usercenter.RoleInfo.roleMenus:type_name -> usercenter.RoleMenu
	8,  // 2: usercenter.RoleMenu.menu:type_name -> usercenter.MenuInfo
	9,  // 3: usercenter.QueryRoleResponse.code:type_name -> usercenter.Code
	1,  // 4: usercenter.QueryRoleResponse.data:type_name -> usercenter.RoleInfo
	9,  // 5: usercenter.GetAllRoleResponse.code:type_name -> usercenter.Code
	1,  // 6: usercenter.GetAllRoleResponse.data:type_name -> usercenter.RoleInfo
	9,  // 7: usercenter.GetRoleDetailResponse.code:type_name -> usercenter.Code
	1,  // 8: usercenter.GetRoleDetailResponse.data:type_name -> usercenter.RoleInfo
	1,  // 9: usercenter.Role.Add:input_type -> usercenter.RoleInfo
	1,  // 10: usercenter.Role.Update:input_type -> usercenter.RoleInfo
	10, // 11: usercenter.Role.Delete:input_type -> usercenter.DelRequest
	3,  // 12: usercenter.Role.Query:input_type -> usercenter.QueryRoleRequest
	6,  // 13: usercenter.Role.GetAll:input_type -> usercenter.GetAllRoleRequest
	11, // 14: usercenter.Role.GetDetail:input_type -> usercenter.GetDetailRequest
	0,  // 15: usercenter.Role.StatisticCount:input_type -> usercenter.StatisticRoleCountRequest
	12, // 16: usercenter.Role.Add:output_type -> usercenter.CommonResponse
	12, // 17: usercenter.Role.Update:output_type -> usercenter.CommonResponse
	12, // 18: usercenter.Role.Delete:output_type -> usercenter.CommonResponse
	4,  // 19: usercenter.Role.Query:output_type -> usercenter.QueryRoleResponse
	5,  // 20: usercenter.Role.GetAll:output_type -> usercenter.GetAllRoleResponse
	7,  // 21: usercenter.Role.GetDetail:output_type -> usercenter.GetRoleDetailResponse
	13, // 22: usercenter.Role.StatisticCount:output_type -> usercenter.StatisticCountResponse
	16, // [16:23] is the sub-list for method output_type
	9,  // [9:16] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_role_proto_init() }
func file_role_proto_init() {
	if File_role_proto != nil {
		return
	}
	file_common_proto_init()
	file_menu_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatisticRoleCountRequest); i {
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
		file_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleInfo); i {
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
		file_role_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleMenu); i {
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
		file_role_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRoleRequest); i {
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
		file_role_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRoleResponse); i {
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
		file_role_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRoleResponse); i {
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
		file_role_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRoleRequest); i {
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
		file_role_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoleDetailResponse); i {
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
			RawDescriptor: file_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_role_proto_goTypes,
		DependencyIndexes: file_role_proto_depIdxs,
		MessageInfos:      file_role_proto_msgTypes,
	}.Build()
	File_role_proto = out.File
	file_role_proto_rawDesc = nil
	file_role_proto_goTypes = nil
	file_role_proto_depIdxs = nil
}
