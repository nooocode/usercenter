// Code generated by protoc-gen-go-triple. DO NOT EDIT.
// versions:
// - protoc-gen-go-triple v1.0.8
// - protoc             v3.19.1
// source: tenant.proto

package usercenter

import (
	context "context"
	protocol "dubbo.apache.org/dubbo-go/v3/protocol"
	dubbo3 "dubbo.apache.org/dubbo-go/v3/protocol/dubbo3"
	invocation "dubbo.apache.org/dubbo-go/v3/protocol/invocation"
	grpc_go "github.com/dubbogo/grpc-go"
	codes "github.com/dubbogo/grpc-go/codes"
	metadata "github.com/dubbogo/grpc-go/metadata"
	status "github.com/dubbogo/grpc-go/status"
	common "github.com/dubbogo/triple/pkg/common"
	constant "github.com/dubbogo/triple/pkg/common/constant"
	triple "github.com/dubbogo/triple/pkg/triple"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc_go.SupportPackageIsVersion7

// TenantClient is the client API for Tenant service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TenantClient interface {
	Add(ctx context.Context, in *TenantInfo, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment)
	Update(ctx context.Context, in *TenantInfo, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment)
	Delete(ctx context.Context, in *DelRequest, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment)
	Query(ctx context.Context, in *QueryTenantRequest, opts ...grpc_go.CallOption) (*QueryTenantResponse, common.ErrorWithAttachment)
	Enable(ctx context.Context, in *EnableRequest, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc_go.CallOption) (*GetAllTenantResponse, common.ErrorWithAttachment)
	GetDetail(ctx context.Context, in *GetDetailRequest, opts ...grpc_go.CallOption) (*GetTenantDetailResponse, common.ErrorWithAttachment)
	StatisticCount(ctx context.Context, in *StatisticTenantCountRequest, opts ...grpc_go.CallOption) (*StatisticCountResponse, common.ErrorWithAttachment)
}

type tenantClient struct {
	cc *triple.TripleConn
}

type TenantClientImpl struct {
	Add            func(ctx context.Context, in *TenantInfo) (*CommonResponse, error)
	Update         func(ctx context.Context, in *TenantInfo) (*CommonResponse, error)
	Delete         func(ctx context.Context, in *DelRequest) (*CommonResponse, error)
	Query          func(ctx context.Context, in *QueryTenantRequest) (*QueryTenantResponse, error)
	Enable         func(ctx context.Context, in *EnableRequest) (*CommonResponse, error)
	GetAll         func(ctx context.Context, in *GetAllRequest) (*GetAllTenantResponse, error)
	GetDetail      func(ctx context.Context, in *GetDetailRequest) (*GetTenantDetailResponse, error)
	StatisticCount func(ctx context.Context, in *StatisticTenantCountRequest) (*StatisticCountResponse, error)
}

func (c *TenantClientImpl) GetDubboStub(cc *triple.TripleConn) TenantClient {
	return NewTenantClient(cc)
}

func (c *TenantClientImpl) XXX_InterfaceName() string {
	return "usercenter.Tenant"
}

func NewTenantClient(cc *triple.TripleConn) TenantClient {
	return &tenantClient{cc}
}

func (c *tenantClient) Add(ctx context.Context, in *TenantInfo, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment) {
	out := new(CommonResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Add", in, out)
}

func (c *tenantClient) Update(ctx context.Context, in *TenantInfo, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment) {
	out := new(CommonResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Update", in, out)
}

func (c *tenantClient) Delete(ctx context.Context, in *DelRequest, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment) {
	out := new(CommonResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Delete", in, out)
}

func (c *tenantClient) Query(ctx context.Context, in *QueryTenantRequest, opts ...grpc_go.CallOption) (*QueryTenantResponse, common.ErrorWithAttachment) {
	out := new(QueryTenantResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Query", in, out)
}

func (c *tenantClient) Enable(ctx context.Context, in *EnableRequest, opts ...grpc_go.CallOption) (*CommonResponse, common.ErrorWithAttachment) {
	out := new(CommonResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/Enable", in, out)
}

func (c *tenantClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc_go.CallOption) (*GetAllTenantResponse, common.ErrorWithAttachment) {
	out := new(GetAllTenantResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/GetAll", in, out)
}

func (c *tenantClient) GetDetail(ctx context.Context, in *GetDetailRequest, opts ...grpc_go.CallOption) (*GetTenantDetailResponse, common.ErrorWithAttachment) {
	out := new(GetTenantDetailResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/GetDetail", in, out)
}

func (c *tenantClient) StatisticCount(ctx context.Context, in *StatisticTenantCountRequest, opts ...grpc_go.CallOption) (*StatisticCountResponse, common.ErrorWithAttachment) {
	out := new(StatisticCountResponse)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/StatisticCount", in, out)
}

// TenantServer is the server API for Tenant service.
// All implementations must embed UnimplementedTenantServer
// for forward compatibility
type TenantServer interface {
	Add(context.Context, *TenantInfo) (*CommonResponse, error)
	Update(context.Context, *TenantInfo) (*CommonResponse, error)
	Delete(context.Context, *DelRequest) (*CommonResponse, error)
	Query(context.Context, *QueryTenantRequest) (*QueryTenantResponse, error)
	Enable(context.Context, *EnableRequest) (*CommonResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllTenantResponse, error)
	GetDetail(context.Context, *GetDetailRequest) (*GetTenantDetailResponse, error)
	StatisticCount(context.Context, *StatisticTenantCountRequest) (*StatisticCountResponse, error)
	mustEmbedUnimplementedTenantServer()
}

// UnimplementedTenantServer must be embedded to have forward compatible implementations.
type UnimplementedTenantServer struct {
	proxyImpl protocol.Invoker
}

func (UnimplementedTenantServer) Add(context.Context, *TenantInfo) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedTenantServer) Update(context.Context, *TenantInfo) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTenantServer) Delete(context.Context, *DelRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTenantServer) Query(context.Context, *QueryTenantRequest) (*QueryTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedTenantServer) Enable(context.Context, *EnableRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (UnimplementedTenantServer) GetAll(context.Context, *GetAllRequest) (*GetAllTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedTenantServer) GetDetail(context.Context, *GetDetailRequest) (*GetTenantDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetail not implemented")
}
func (UnimplementedTenantServer) StatisticCount(context.Context, *StatisticTenantCountRequest) (*StatisticCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatisticCount not implemented")
}
func (s *UnimplementedTenantServer) XXX_SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *UnimplementedTenantServer) XXX_GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func (s *UnimplementedTenantServer) XXX_ServiceDesc() *grpc_go.ServiceDesc {
	return &Tenant_ServiceDesc
}
func (s *UnimplementedTenantServer) XXX_InterfaceName() string {
	return "usercenter.Tenant"
}

func (UnimplementedTenantServer) mustEmbedUnimplementedTenantServer() {}

// UnsafeTenantServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TenantServer will
// result in compilation errors.
type UnsafeTenantServer interface {
	mustEmbedUnimplementedTenantServer()
}

func RegisterTenantServer(s grpc_go.ServiceRegistrar, srv TenantServer) {
	s.RegisterService(&Tenant_ServiceDesc, srv)
}

func _Tenant_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("Add", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(TenantInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("Update", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("Delete", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("Query", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("Enable", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("GetAll", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_GetDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("GetDetail", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _Tenant_StatisticCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatisticTenantCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("StatisticCount", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

// Tenant_ServiceDesc is the grpc_go.ServiceDesc for Tenant service.
// It's only intended for direct use with grpc_go.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tenant_ServiceDesc = grpc_go.ServiceDesc{
	ServiceName: "usercenter.Tenant",
	HandlerType: (*TenantServer)(nil),
	Methods: []grpc_go.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Tenant_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Tenant_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Tenant_Delete_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _Tenant_Query_Handler,
		},
		{
			MethodName: "Enable",
			Handler:    _Tenant_Enable_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _Tenant_GetAll_Handler,
		},
		{
			MethodName: "GetDetail",
			Handler:    _Tenant_GetDetail_Handler,
		},
		{
			MethodName: "StatisticCount",
			Handler:    _Tenant_StatisticCount_Handler,
		},
	},
	Streams:  []grpc_go.StreamDesc{},
	Metadata: "tenant.proto",
}
