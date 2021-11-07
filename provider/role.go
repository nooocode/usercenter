package provider

import (
	"context"

	apipb "codeup.aliyun.com/atali/usercenter/api"
	"codeup.aliyun.com/atali/usercenter/model"
)

type RoleProvider struct {
	apipb.UnimplementedRoleServer
}

func (u *RoleProvider) Add(ctx context.Context, in *apipb.RoleInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	// err := model.CreateRole(userInfoToRole(in), false)
	// if err != nil {
	// 	resp.Code = apipb.Code_InternalServerError
	// 	resp.Message = err.Error()
	// }
	return resp, nil
}

func (u *RoleProvider) Update(ctx context.Context, in *apipb.RoleInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	// err := model.UpdateRole(userInfoToRole(in))
	// if err != nil {
	// 	resp.Code = apipb.Code_InternalServerError
	// 	resp.Message = err.Error()
	// }
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

func (u *RoleProvider) Enable(ctx context.Context, in *apipb.EnableRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	// err := model.EnableRole(in.Id, in.Enable)
	// if err != nil {
	// 	resp.Code = apipb.Code_InternalServerError
	// 	resp.Message = err.Error()
	// }
	return resp, nil
}

func (u *RoleProvider) GetAll(ctx context.Context, in *apipb.GetAllRequest) (*apipb.GetAllRoleResponse, error) {
	resp := &apipb.GetAllRoleResponse{}
	// users, err := model.GetAllRole()
	// if err != nil {
	// 	resp.Code = apipb.Code_InternalServerError
	// 	resp.Message = err.Error()
	// } else {
	// 	resp.Data = userToRoleInfos(users)
	// }

	return resp, nil
}

func (u *RoleProvider) GetDetail(ctx context.Context, in *apipb.GetDetailRequest) (*apipb.GetRoleDetailResponse, error) {
	resp := &apipb.GetRoleDetailResponse{}
	// user, err := model.GetRoleById(in.Id)
	// if err != nil {
	// 	resp.Code = apipb.Code_InternalServerError
	// 	resp.Message = err.Error()
	// }
	// resp.Data = userToRoleInfo(&user)
	return resp, nil
}
