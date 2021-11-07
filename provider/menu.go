package provider

import (
	"context"

	apipb "codeup.aliyun.com/atali/usercenter/api"
	"codeup.aliyun.com/atali/usercenter/model"
)

type MenuProvider struct {
	apipb.UnimplementedMenuServer
}

func (u *MenuProvider) Add(ctx context.Context, in *apipb.MenuInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	// err := model.CreateMenu(userInfoToMenu(in), false)
	// if err != nil {
	// 	resp.Code = apipb.Code_InternalServerError
	// 	resp.Message = err.Error()
	// }
	return resp, nil
}

func (u *MenuProvider) Update(ctx context.Context, in *apipb.MenuInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	// err := model.UpdateMenu(userInfoToMenu(in))
	// if err != nil {
	// 	resp.Code = apipb.Code_InternalServerError
	// 	resp.Message = err.Error()
	// }
	return resp, nil
}

func (u *MenuProvider) Delete(ctx context.Context, in *apipb.DelRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	err := model.DeleteMenu(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *MenuProvider) Query(ctx context.Context, in *apipb.QueryMenuRequest) (*apipb.QueryMenuResponse, error) {
	resp := &apipb.QueryMenuResponse{}
	model.QueryMenu(in, resp)
	return resp, nil
}

func (u *MenuProvider) Enable(ctx context.Context, in *apipb.EnableRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	// err := model.EnableMenu(in.Id, in.Enable)
	// if err != nil {
	// 	resp.Code = apipb.Code_InternalServerError
	// 	resp.Message = err.Error()
	// }
	return resp, nil
}

func (u *MenuProvider) GetAll(ctx context.Context, in *apipb.GetAllRequest) (*apipb.GetAllMenuResponse, error) {
	resp := &apipb.GetAllMenuResponse{}
	// users, err := model.GetAllMenu()
	// if err != nil {
	// 	resp.Code = apipb.Code_InternalServerError
	// 	resp.Message = err.Error()
	// } else {
	// 	resp.Data = userToMenuInfos(users)
	// }

	return resp, nil
}

func (u *MenuProvider) GetDetail(ctx context.Context, in *apipb.GetDetailRequest) (*apipb.GetMenuDetailResponse, error) {
	resp := &apipb.GetMenuDetailResponse{}
	// user, err := model.GetMenuById(in.Id)
	// if err != nil {
	// 	resp.Code = apipb.Code_InternalServerError
	// 	resp.Message = err.Error()
	// }
	// resp.Data = userToMenuInfo(&user)
	return resp, nil
}
