package provider

import (
	"context"

	apipb "github.com/nooocode/usercenter/api"
	"github.com/nooocode/usercenter/model"
)

type RoleProvider struct {
	apipb.UnimplementedRoleServer
}

func (u *RoleProvider) Add(ctx context.Context, in *apipb.RoleInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	err := model.CreateRole(model.PBToRole(in))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *RoleProvider) Update(ctx context.Context, in *apipb.RoleInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	err := model.UpdateRole(model.PBToRole(in))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *RoleProvider) Delete(ctx context.Context, in *apipb.DelRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	err := model.DeleteRole(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *RoleProvider) Query(ctx context.Context, in *apipb.QueryRoleRequest) (*apipb.QueryRoleResponse, error) {
	resp := &apipb.QueryRoleResponse{}
	model.QueryRole(in, resp)
	return resp, nil
}

func (u *RoleProvider) GetAll(ctx context.Context, in *apipb.GetAllRequest) (*apipb.GetAllRoleResponse, error) {
	resp := &apipb.GetAllRoleResponse{}
	roles, err := model.GetAllRole()
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.RolesToPB(roles)
	}

	return resp, nil
}

func (u *RoleProvider) GetDetail(ctx context.Context, in *apipb.GetDetailRequest) (*apipb.GetRoleDetailResponse, error) {
	resp := &apipb.GetRoleDetailResponse{}
	role, err := model.GetRoleByID(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	resp.Data = model.RoleToPB(role)
	return resp, nil
}
