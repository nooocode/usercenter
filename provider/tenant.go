package provider

import (
	"context"

	apipb "github.com/nooocode/usercenter/api"
	"github.com/nooocode/usercenter/model"
)

type TenantProvider struct {
	apipb.UnimplementedTenantServer
}

func (u *TenantProvider) Add(ctx context.Context, in *apipb.TenantInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	// err := model.CreateTenant(userInfoToTenant(in), false)
	// if err != nil {
	// 	resp.Code = apipb.Code_InternalServerError
	// 	resp.Message = err.Error()
	// }
	return resp, nil
}

func (u *TenantProvider) Update(ctx context.Context, in *apipb.TenantInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	// err := model.UpdateTenant(userInfoToTenant(in))
	// if err != nil {
	// 	resp.Code = apipb.Code_InternalServerError
	// 	resp.Message = err.Error()
	// }
	return resp, nil
}

func (u *TenantProvider) Delete(ctx context.Context, in *apipb.DelRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	err := model.DeleteTenant(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *TenantProvider) Query(ctx context.Context, in *apipb.QueryTenantRequest) (*apipb.QueryTenantResponse, error) {
	resp := &apipb.QueryTenantResponse{}
	model.QueryTenant(in, resp)
	return resp, nil
}

func (u *TenantProvider) Enable(ctx context.Context, in *apipb.EnableRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	err := model.EnableTenant(in.Id, in.Enable)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *TenantProvider) GetAll(ctx context.Context, in *apipb.GetAllRequest) (*apipb.GetAllTenantResponse, error) {
	resp := &apipb.GetAllTenantResponse{}
	// users, err := model.GetAllTenant()
	// if err != nil {
	// 	resp.Code = apipb.Code_InternalServerError
	// 	resp.Message = err.Error()
	// } else {
	// 	resp.Data = userToTenantInfos(users)
	// }

	return resp, nil
}

func (u *TenantProvider) GetDetail(ctx context.Context, in *apipb.GetDetailRequest) (*apipb.GetTenantDetailResponse, error) {
	resp := &apipb.GetTenantDetailResponse{}
	// user, err := model.GetTenantById(in.Id)
	// if err != nil {
	// 	resp.Code = apipb.Code_InternalServerError
	// 	resp.Message = err.Error()
	// }
	// resp.Data = userToTenantInfo(&user)
	return resp, nil
}
