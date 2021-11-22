package provider

import (
	"context"

	commonmodel "github.com/nooocode/pkg/model"
	apipb "github.com/nooocode/usercenter/api"
	"github.com/nooocode/usercenter/model"
)

type MenuProvider struct {
	apipb.UnimplementedMenuServer
}

func (u *MenuProvider) Add(ctx context.Context, in *apipb.MenuInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: commonmodel.Success,
	}
	err := model.AddMenu(model.PBToMenu(in))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *MenuProvider) Update(ctx context.Context, in *apipb.MenuInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: commonmodel.Success,
	}
	err := model.UpdateMenu(model.PBToMenu(in))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *MenuProvider) Delete(ctx context.Context, in *apipb.DelRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: commonmodel.Success,
	}
	err := model.DeleteMenu(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *MenuProvider) Query(ctx context.Context, in *apipb.QueryMenuRequest) (*apipb.QueryMenuResponse, error) {
	resp := &apipb.QueryMenuResponse{
		Code: commonmodel.Success,
	}
	model.QueryMenu(in, resp)
	return resp, nil
}

func (u *MenuProvider) GetAll(ctx context.Context, in *apipb.GetAllRequest) (*apipb.GetAllMenuResponse, error) {
	resp := &apipb.GetAllMenuResponse{
		Code: commonmodel.Success,
	}
	menus, err := model.GetAllMenus()
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MenusToPB(menus)
	}

	return resp, nil
}

func (u *MenuProvider) GetDetail(ctx context.Context, in *apipb.GetDetailRequest) (*apipb.GetMenuDetailResponse, error) {
	resp := &apipb.GetMenuDetailResponse{
		Code: commonmodel.Success,
	}
	menu, err := model.GetMenuByID(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	resp.Data = model.MenuToPB(menu)
	return resp, nil
}
