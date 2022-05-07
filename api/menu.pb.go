// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: menu.proto

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

type MenuInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Level    uint32 `protobuf:"varint,2,opt,name=level,proto3" json:"level"`
	ParentID string `protobuf:"bytes,3,opt,name=parentID,proto3" json:"parentID"`
	Path     string `protobuf:"bytes,4,opt,name=path,proto3" json:"path"`
	// @inject_tag: validate:"required"
	Name        string           `protobuf:"bytes,5,opt,name=name,proto3" json:"name" validate:"required"`
	Hidden      bool             `protobuf:"varint,6,opt,name=hidden,proto3" json:"hidden"`
	Component   string           `protobuf:"bytes,7,opt,name=component,proto3" json:"component"`
	Sort        int32            `protobuf:"varint,8,opt,name=sort,proto3" json:"sort"`
	Cache       bool             `protobuf:"varint,9,opt,name=cache,proto3" json:"cache"`
	DefaultMenu bool             `protobuf:"varint,10,opt,name=defaultMenu,proto3" json:"defaultMenu"`
	Title       string           `protobuf:"bytes,11,opt,name=title,proto3" json:"title"`
	Icon        string           `protobuf:"bytes,12,opt,name=icon,proto3" json:"icon"`
	CloseTab    bool             `protobuf:"varint,13,opt,name=closeTab,proto3" json:"closeTab"`
	Children    []*MenuInfo      `protobuf:"bytes,14,rep,name=children,proto3" json:"children"`
	Parameters  []*MenuParameter `protobuf:"bytes,15,rep,name=parameters,proto3" json:"parameters"`
	MenuFuncs   []*MenuFunc      `protobuf:"bytes,16,rep,name=menuFuncs,proto3" json:"menuFuncs"`
}

func (x *MenuInfo) Reset() {
	*x = MenuInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuInfo) ProtoMessage() {}

func (x *MenuInfo) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuInfo.ProtoReflect.Descriptor instead.
func (*MenuInfo) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{0}
}

func (x *MenuInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MenuInfo) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *MenuInfo) GetParentID() string {
	if x != nil {
		return x.ParentID
	}
	return ""
}

func (x *MenuInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *MenuInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MenuInfo) GetHidden() bool {
	if x != nil {
		return x.Hidden
	}
	return false
}

func (x *MenuInfo) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *MenuInfo) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *MenuInfo) GetCache() bool {
	if x != nil {
		return x.Cache
	}
	return false
}

func (x *MenuInfo) GetDefaultMenu() bool {
	if x != nil {
		return x.DefaultMenu
	}
	return false
}

func (x *MenuInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MenuInfo) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *MenuInfo) GetCloseTab() bool {
	if x != nil {
		return x.CloseTab
	}
	return false
}

func (x *MenuInfo) GetChildren() []*MenuInfo {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *MenuInfo) GetParameters() []*MenuParameter {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *MenuInfo) GetMenuFuncs() []*MenuFunc {
	if x != nil {
		return x.MenuFuncs
	}
	return nil
}

type MenuParameter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	MenuID string `protobuf:"bytes,2,opt,name=menuID,proto3" json:"menuID"`
	Type   string `protobuf:"bytes,3,opt,name=type,proto3" json:"type"`
	Key    string `protobuf:"bytes,4,opt,name=key,proto3" json:"key"`
	Value  string `protobuf:"bytes,5,opt,name=value,proto3" json:"value"`
}

func (x *MenuParameter) Reset() {
	*x = MenuParameter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuParameter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuParameter) ProtoMessage() {}

func (x *MenuParameter) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuParameter.ProtoReflect.Descriptor instead.
func (*MenuParameter) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{1}
}

func (x *MenuParameter) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MenuParameter) GetMenuID() string {
	if x != nil {
		return x.MenuID
	}
	return ""
}

func (x *MenuParameter) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MenuParameter) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *MenuParameter) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type MenuFunc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	MenuID       string         `protobuf:"bytes,2,opt,name=menuID,proto3" json:"menuID"`
	Name         string         `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	Title        string         `protobuf:"bytes,4,opt,name=title,proto3" json:"title"`
	Hidden       bool           `protobuf:"varint,5,opt,name=hidden,proto3" json:"hidden"`
	MenuFuncApis []*MenuFuncApi `protobuf:"bytes,6,rep,name=menuFuncApis,proto3" json:"menuFuncApis"`
}

func (x *MenuFunc) Reset() {
	*x = MenuFunc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuFunc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuFunc) ProtoMessage() {}

func (x *MenuFunc) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuFunc.ProtoReflect.Descriptor instead.
func (*MenuFunc) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{2}
}

func (x *MenuFunc) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MenuFunc) GetMenuID() string {
	if x != nil {
		return x.MenuID
	}
	return ""
}

func (x *MenuFunc) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MenuFunc) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MenuFunc) GetHidden() bool {
	if x != nil {
		return x.Hidden
	}
	return false
}

func (x *MenuFunc) GetMenuFuncApis() []*MenuFuncApi {
	if x != nil {
		return x.MenuFuncApis
	}
	return nil
}

type MenuFuncApi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	MenuFuncID string   `protobuf:"bytes,2,opt,name=menuFuncID,proto3" json:"menuFuncID"`
	ApiID      string   `protobuf:"bytes,3,opt,name=apiID,proto3" json:"apiID"`
	ApiInfo    *APIInfo `protobuf:"bytes,4,opt,name=apiInfo,proto3" json:"apiInfo"`
}

func (x *MenuFuncApi) Reset() {
	*x = MenuFuncApi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuFuncApi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuFuncApi) ProtoMessage() {}

func (x *MenuFuncApi) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuFuncApi.ProtoReflect.Descriptor instead.
func (*MenuFuncApi) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{3}
}

func (x *MenuFuncApi) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MenuFuncApi) GetMenuFuncID() string {
	if x != nil {
		return x.MenuFuncID
	}
	return ""
}

func (x *MenuFuncApi) GetApiID() string {
	if x != nil {
		return x.ApiID
	}
	return ""
}

func (x *MenuFuncApi) GetApiInfo() *APIInfo {
	if x != nil {
		return x.ApiInfo
	}
	return nil
}

type QueryMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: uri:"pageIndex" form:"pageIndex"
	PageIndex int64 `protobuf:"varint,1,opt,name=pageIndex,proto3" json:"pageIndex" uri:"pageIndex" form:"pageIndex"`
	// @inject_tag: uri:"pageSize" form:"pageSize"
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize" uri:"pageSize" form:"pageSize"`
	// @inject_tag: uri:"orderField" form:"orderField"
	OrderField string `protobuf:"bytes,3,opt,name=orderField,proto3" json:"orderField" uri:"orderField" form:"orderField"`
	// @inject_tag: uri:"desc" form:"desc"
	Desc bool `protobuf:"varint,4,opt,name=desc,proto3" json:"desc" uri:"desc" form:"desc"`
	// @inject_tag: uri:"name" form:"name"
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name" uri:"name" form:"name"`
	// @inject_tag: uri:"path" form:"path"
	Path string `protobuf:"bytes,6,opt,name=path,proto3" json:"path" uri:"path" form:"path"`
	// @inject_tag: uri:"title" form:"title"
	Title string `protobuf:"bytes,7,opt,name=title,proto3" json:"title" uri:"title" form:"title"`
	// @inject_tag: uri:"parentID" form:"parentID"
	ParentID string `protobuf:"bytes,8,opt,name=parentID,proto3" json:"parentID" uri:"parentID" form:"parentID"`
	// @inject_tag: uri:"level" form:"level"
	Level int32 `protobuf:"varint,9,opt,name=level,proto3" json:"level" uri:"level" form:"level"`
	// @inject_tag: uri:"ids" form:"ids"
	Ids []string `protobuf:"bytes,10,rep,name=ids,proto3" json:"ids" uri:"ids" form:"ids"`
}

func (x *QueryMenuRequest) Reset() {
	*x = QueryMenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMenuRequest) ProtoMessage() {}

func (x *QueryMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMenuRequest.ProtoReflect.Descriptor instead.
func (*QueryMenuRequest) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{4}
}

func (x *QueryMenuRequest) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *QueryMenuRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *QueryMenuRequest) GetOrderField() string {
	if x != nil {
		return x.OrderField
	}
	return ""
}

func (x *QueryMenuRequest) GetDesc() bool {
	if x != nil {
		return x.Desc
	}
	return false
}

func (x *QueryMenuRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *QueryMenuRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *QueryMenuRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *QueryMenuRequest) GetParentID() string {
	if x != nil {
		return x.ParentID
	}
	return ""
}

func (x *QueryMenuRequest) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *QueryMenuRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type QueryMenuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code        `protobuf:"varint,1,opt,name=code,proto3,enum=usercenter.Code" json:"code"`
	Message string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*MenuInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
	Pages   int64       `protobuf:"varint,4,opt,name=pages,proto3" json:"pages"`
	Records int64       `protobuf:"varint,5,opt,name=records,proto3" json:"records"`
	Total   int64       `protobuf:"varint,6,opt,name=total,proto3" json:"total"`
}

func (x *QueryMenuResponse) Reset() {
	*x = QueryMenuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMenuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMenuResponse) ProtoMessage() {}

func (x *QueryMenuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMenuResponse.ProtoReflect.Descriptor instead.
func (*QueryMenuResponse) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{5}
}

func (x *QueryMenuResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *QueryMenuResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryMenuResponse) GetData() []*MenuInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryMenuResponse) GetPages() int64 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *QueryMenuResponse) GetRecords() int64 {
	if x != nil {
		return x.Records
	}
	return 0
}

func (x *QueryMenuResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAllMenuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code        `protobuf:"varint,1,opt,name=code,proto3,enum=usercenter.Code" json:"code"`
	Message string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    []*MenuInfo `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetAllMenuResponse) Reset() {
	*x = GetAllMenuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMenuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMenuResponse) ProtoMessage() {}

func (x *GetAllMenuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMenuResponse.ProtoReflect.Descriptor instead.
func (*GetAllMenuResponse) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllMenuResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetAllMenuResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllMenuResponse) GetData() []*MenuInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetMenuDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code      `protobuf:"varint,1,opt,name=code,proto3,enum=usercenter.Code" json:"code"`
	Message string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	Data    *MenuInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *GetMenuDetailResponse) Reset() {
	*x = GetMenuDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMenuDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMenuDetailResponse) ProtoMessage() {}

func (x *GetMenuDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_menu_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMenuDetailResponse.ProtoReflect.Descriptor instead.
func (*GetMenuDetailResponse) Descriptor() ([]byte, []int) {
	return file_menu_proto_rawDescGZIP(), []int{7}
}

func (x *GetMenuDetailResponse) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_None
}

func (x *GetMenuDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetMenuDetailResponse) GetData() *MenuInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_menu_proto protoreflect.FileDescriptor

var file_menu_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xdd, 0x03, 0x0a, 0x08, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x69, 0x64, 0x64,
	0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x63, 0x61, 0x63, 0x68, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x69, 0x63, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x61, 0x62,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x61, 0x62,
	0x12, 0x30, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x0e, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x4d, 0x65, 0x6e, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72,
	0x65, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x32, 0x0a,
	0x09, 0x6d, 0x65, 0x6e, 0x75, 0x46, 0x75, 0x6e, 0x63, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x65,
	0x6e, 0x75, 0x46, 0x75, 0x6e, 0x63, 0x52, 0x09, 0x6d, 0x65, 0x6e, 0x75, 0x46, 0x75, 0x6e, 0x63,
	0x73, 0x22, 0x73, 0x0a, 0x0d, 0x4d, 0x65, 0x6e, 0x75, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6e, 0x75, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6e, 0x75, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xb1, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x6e, 0x75, 0x46,
	0x75, 0x6e, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6e, 0x75, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6e, 0x75, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x3b, 0x0a,
	0x0c, 0x6d, 0x65, 0x6e, 0x75, 0x46, 0x75, 0x6e, 0x63, 0x41, 0x70, 0x69, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x46, 0x75, 0x6e, 0x63, 0x41, 0x70, 0x69, 0x52, 0x0c, 0x6d, 0x65,
	0x6e, 0x75, 0x46, 0x75, 0x6e, 0x63, 0x41, 0x70, 0x69, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x0b, 0x4d,
	0x65, 0x6e, 0x75, 0x46, 0x75, 0x6e, 0x63, 0x41, 0x70, 0x69, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65,
	0x6e, 0x75, 0x46, 0x75, 0x6e, 0x63, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6d, 0x65, 0x6e, 0x75, 0x46, 0x75, 0x6e, 0x63, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70,
	0x69, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x69, 0x49, 0x44,
	0x12, 0x2d, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x41,
	0x50, 0x49, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x61, 0x70, 0x69, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x82, 0x02, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x65,
	0x73, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x22, 0xc3, 0x01, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65,
	0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x7e, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x24, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6e, 0x75,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x81, 0x01, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x4d, 0x65, 0x6e, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x9e,
	0x03, 0x0a, 0x04, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x39, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x14,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6e, 0x75,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x46, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x6e, 0x75,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_menu_proto_rawDescOnce sync.Once
	file_menu_proto_rawDescData = file_menu_proto_rawDesc
)

func file_menu_proto_rawDescGZIP() []byte {
	file_menu_proto_rawDescOnce.Do(func() {
		file_menu_proto_rawDescData = protoimpl.X.CompressGZIP(file_menu_proto_rawDescData)
	})
	return file_menu_proto_rawDescData
}

var file_menu_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_menu_proto_goTypes = []interface{}{
	(*MenuInfo)(nil),              // 0: usercenter.MenuInfo
	(*MenuParameter)(nil),         // 1: usercenter.MenuParameter
	(*MenuFunc)(nil),              // 2: usercenter.MenuFunc
	(*MenuFuncApi)(nil),           // 3: usercenter.MenuFuncApi
	(*QueryMenuRequest)(nil),      // 4: usercenter.QueryMenuRequest
	(*QueryMenuResponse)(nil),     // 5: usercenter.QueryMenuResponse
	(*GetAllMenuResponse)(nil),    // 6: usercenter.GetAllMenuResponse
	(*GetMenuDetailResponse)(nil), // 7: usercenter.GetMenuDetailResponse
	(*APIInfo)(nil),               // 8: usercenter.APIInfo
	(Code)(0),                     // 9: usercenter.Code
	(*DelRequest)(nil),            // 10: usercenter.DelRequest
	(*GetAllRequest)(nil),         // 11: usercenter.GetAllRequest
	(*GetDetailRequest)(nil),      // 12: usercenter.GetDetailRequest
	(*CommonResponse)(nil),        // 13: usercenter.CommonResponse
}
var file_menu_proto_depIdxs = []int32{
	0,  // 0: usercenter.MenuInfo.children:type_name -> usercenter.MenuInfo
	1,  // 1: usercenter.MenuInfo.parameters:type_name -> usercenter.MenuParameter
	2,  // 2: usercenter.MenuInfo.menuFuncs:type_name -> usercenter.MenuFunc
	3,  // 3: usercenter.MenuFunc.menuFuncApis:type_name -> usercenter.MenuFuncApi
	8,  // 4: usercenter.MenuFuncApi.apiInfo:type_name -> usercenter.APIInfo
	9,  // 5: usercenter.QueryMenuResponse.code:type_name -> usercenter.Code
	0,  // 6: usercenter.QueryMenuResponse.data:type_name -> usercenter.MenuInfo
	9,  // 7: usercenter.GetAllMenuResponse.code:type_name -> usercenter.Code
	0,  // 8: usercenter.GetAllMenuResponse.data:type_name -> usercenter.MenuInfo
	9,  // 9: usercenter.GetMenuDetailResponse.code:type_name -> usercenter.Code
	0,  // 10: usercenter.GetMenuDetailResponse.data:type_name -> usercenter.MenuInfo
	0,  // 11: usercenter.Menu.Add:input_type -> usercenter.MenuInfo
	0,  // 12: usercenter.Menu.Update:input_type -> usercenter.MenuInfo
	10, // 13: usercenter.Menu.Delete:input_type -> usercenter.DelRequest
	4,  // 14: usercenter.Menu.Query:input_type -> usercenter.QueryMenuRequest
	11, // 15: usercenter.Menu.GetAll:input_type -> usercenter.GetAllRequest
	12, // 16: usercenter.Menu.GetDetail:input_type -> usercenter.GetDetailRequest
	13, // 17: usercenter.Menu.Add:output_type -> usercenter.CommonResponse
	13, // 18: usercenter.Menu.Update:output_type -> usercenter.CommonResponse
	13, // 19: usercenter.Menu.Delete:output_type -> usercenter.CommonResponse
	5,  // 20: usercenter.Menu.Query:output_type -> usercenter.QueryMenuResponse
	6,  // 21: usercenter.Menu.GetAll:output_type -> usercenter.GetAllMenuResponse
	7,  // 22: usercenter.Menu.GetDetail:output_type -> usercenter.GetMenuDetailResponse
	17, // [17:23] is the sub-list for method output_type
	11, // [11:17] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_menu_proto_init() }
func file_menu_proto_init() {
	if File_menu_proto != nil {
		return
	}
	file_common_proto_init()
	file_api_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_menu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuInfo); i {
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
		file_menu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuParameter); i {
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
		file_menu_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuFunc); i {
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
		file_menu_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuFuncApi); i {
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
		file_menu_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMenuRequest); i {
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
		file_menu_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMenuResponse); i {
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
		file_menu_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllMenuResponse); i {
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
		file_menu_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMenuDetailResponse); i {
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
			RawDescriptor: file_menu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_menu_proto_goTypes,
		DependencyIndexes: file_menu_proto_depIdxs,
		MessageInfos:      file_menu_proto_msgTypes,
	}.Build()
	File_menu_proto = out.File
	file_menu_proto_rawDesc = nil
	file_menu_proto_goTypes = nil
	file_menu_proto_depIdxs = nil
}
