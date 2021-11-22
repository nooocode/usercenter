package provider

import (
	"context"

	commonmodel "github.com/nooocode/pkg/model"
	apipb "github.com/nooocode/usercenter/api"
	"github.com/nooocode/usercenter/model"
)

type APIProvider struct {
	apipb.UnimplementedAPIServer
}

func (u *APIProvider) Add(ctx context.Context, in *apipb.APIInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: commonmodel.Success,
	}
	err := model.CreateAPI(model.PBToAPI(in))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *APIProvider) Update(ctx context.Context, in *apipb.APIInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: commonmodel.Success,
	}
	err := model.UpdateAPI(model.PBToAPI(in))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *APIProvider) Delete(ctx context.Context, in *apipb.DelRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: commonmodel.Success,
	}
	err := model.DeleteApi(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *APIProvider) Query(ctx context.Context, in *apipb.QueryAPIRequest) (*apipb.QueryAPIResponse, error) {
	resp := &apipb.QueryAPIResponse{
		Code: commonmodel.Success,
	}
	model.QueryAPI(in, resp)
	return resp, nil
}

func (u *APIProvider) Enable(ctx context.Context, in *apipb.EnableRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: commonmodel.Success,
	}
	err := model.EnableAPI(in.Id, in.Enable)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *APIProvider) GetAll(ctx context.Context, in *apipb.GetAllRequest) (*apipb.GetAllAPIResponse, error) {
	resp := &apipb.GetAllAPIResponse{
		Code: commonmodel.Success,
	}
	apis, err := model.GetAllAPIs()
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.APIsToPB(apis)
	}

	return resp, nil
}

func (u *APIProvider) GetDetail(ctx context.Context, in *apipb.GetDetailRequest) (*apipb.GetAPIDetailResponse, error) {
	resp := &apipb.GetAPIDetailResponse{
		Code: commonmodel.Success,
	}
	api, err := model.GetAPIById(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	resp.Data = model.APIToPB(&api)
	return resp, nil
}
