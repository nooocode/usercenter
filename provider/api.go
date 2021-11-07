package provider

import (
	"context"

	apipb "codeup.aliyun.com/atali/usercenter/api"
	"codeup.aliyun.com/atali/usercenter/model"
)

type APIProvider struct {
	apipb.UnimplementedAPIServer
}

func (u *APIProvider) Add(ctx context.Context, in *apipb.APIInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	// err := model.CreateAPI(userInfoToAPI(in), false)
	// if err != nil {
	// 	resp.Code = apipb.Code_InternalServerError
	// 	resp.Message = err.Error()
	// }
	return resp, nil
}

func (u *APIProvider) Update(ctx context.Context, in *apipb.APIInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	// err := model.UpdateAPI(userInfoToAPI(in))
	// if err != nil {
	// 	resp.Code = apipb.Code_InternalServerError
	// 	resp.Message = err.Error()
	// }
	return resp, nil
}

func (u *APIProvider) Delete(ctx context.Context, in *apipb.DelRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	err := model.DeleteApi(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *APIProvider) Query(ctx context.Context, in *apipb.QueryAPIRequest) (*apipb.QueryAPIResponse, error) {
	resp := &apipb.QueryAPIResponse{}
	model.QueryAPI(in, resp)
	return resp, nil
}

func (u *APIProvider) Enable(ctx context.Context, in *apipb.EnableRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	// err := model.EnableAPI(in.Id, in.Enable)
	// if err != nil {
	// 	resp.Code = apipb.Code_InternalServerError
	// 	resp.Message = err.Error()
	// }
	return resp, nil
}

func (u *APIProvider) GetAll(ctx context.Context, in *apipb.GetAllRequest) (*apipb.GetAllAPIResponse, error) {
	resp := &apipb.GetAllAPIResponse{}
	// users, err := model.GetAllAPI()
	// if err != nil {
	// 	resp.Code = apipb.Code_InternalServerError
	// 	resp.Message = err.Error()
	// } else {
	// 	resp.Data = userToAPIInfos(users)
	// }

	return resp, nil
}

func (u *APIProvider) GetDetail(ctx context.Context, in *apipb.GetDetailRequest) (*apipb.GetAPIDetailResponse, error) {
	resp := &apipb.GetAPIDetailResponse{}
	// user, err := model.GetAPIById(in.Id)
	// if err != nil {
	// 	resp.Code = apipb.Code_InternalServerError
	// 	resp.Message = err.Error()
	// }
	// resp.Data = userToAPIInfo(&user)
	return resp, nil
}
