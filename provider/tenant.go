package provider

import (
	"context"

	"github.com/nooocode/pkg/constants"
	commonmodel "github.com/nooocode/pkg/model"
	apipb "github.com/nooocode/usercenter/api"
	"github.com/nooocode/usercenter/model"
)

type TenantProvider struct {
	apipb.UnimplementedTenantServer
}

func (u *TenantProvider) Add(ctx context.Context, in *apipb.TenantInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: commonmodel.Success,
	}
	err := model.CreateTenant(model.PBToTenant(in))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *TenantProvider) Update(ctx context.Context, in *apipb.TenantInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: commonmodel.Success,
	}
	if in.Id == constants.PlatformTenantID {
		resp.Code = commonmodel.BadRequest
		resp.Message = "平台租户不允许更新"
		return resp, nil
	}
	err := model.UpdateTenant(model.PBToTenant(in))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *TenantProvider) Delete(ctx context.Context, in *apipb.DelRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: commonmodel.Success,
	}
	if in.Id == constants.PlatformTenantID {
		resp.Code = commonmodel.BadRequest
		resp.Message = "平台租户不允许删除"
		return resp, nil
	}
	err := model.DeleteTenant(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *TenantProvider) Query(ctx context.Context, in *apipb.QueryTenantRequest) (*apipb.QueryTenantResponse, error) {
	resp := &apipb.QueryTenantResponse{
		Code: commonmodel.Success,
	}
	model.QueryTenant(in, resp)
	return resp, nil
}

func (u *TenantProvider) Enable(ctx context.Context, in *apipb.EnableRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: commonmodel.Success,
	}
	if in.Id == constants.PlatformTenantID {
		resp.Code = commonmodel.BadRequest
		resp.Message = "平台租户不允许更新"
		return resp, nil
	}
	err := model.EnableTenant(in.Id, in.Enable)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *TenantProvider) GetAll(ctx context.Context, in *apipb.GetAllRequest) (*apipb.GetAllTenantResponse, error) {
	resp := &apipb.GetAllTenantResponse{
		Code: commonmodel.Success,
	}
	users, err := model.GetAllTenant()
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.TenantsToPB(users)
	}

	return resp, nil
}

func (u *TenantProvider) GetDetail(ctx context.Context, in *apipb.GetDetailRequest) (*apipb.GetTenantDetailResponse, error) {
	resp := &apipb.GetTenantDetailResponse{
		Code: commonmodel.Success,
	}
	tenant, err := model.GetTenantByID(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	resp.Data = model.TenantToPB(tenant)
	return resp, nil
}
