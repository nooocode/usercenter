package provider

import (
	"context"

	commonmodel "github.com/nooocode/pkg/model"
	apipb "github.com/nooocode/usercenter/api"
	"github.com/nooocode/usercenter/model"
)

type UserProvider struct {
	apipb.UnimplementedUserServer
}

func (u *UserProvider) Login(ctx context.Context, in *apipb.LoginRequest) (*apipb.LoginResponse, error) {
	resp := &apipb.LoginResponse{
		Code: commonmodel.Success,
	}
	model.Login(in, resp)
	return resp, nil
}

func (u *UserProvider) Add(ctx context.Context, in *apipb.UserInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: commonmodel.Success,
	}
	user := model.PBToUser(in)
	user.Password = in.Password
	err := model.CreateUser(user, false)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *UserProvider) Update(ctx context.Context, in *apipb.UserInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: commonmodel.Success,
	}
	err := model.UpdateUser(model.PBToUser(in))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *UserProvider) Delete(ctx context.Context, in *apipb.DelRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: commonmodel.Success,
	}
	err := model.DeleteUser(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *UserProvider) Query(ctx context.Context, in *apipb.QueryUserRequest) (*apipb.QueryUserResponse, error) {
	resp := &apipb.QueryUserResponse{
		Code: commonmodel.Success,
	}
	model.QueryUser(in, resp)
	return resp, nil
}

func (u *UserProvider) GetProfile(ctx context.Context, in *apipb.GetDetailRequest) (*apipb.GetProfileResponse, error) {
	resp := &apipb.GetProfileResponse{
		Code: commonmodel.Success,
	}
	//TODO
	user, err := model.GetUserById(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	resp.Data = model.UserToPB(&user)
	return resp, nil
}

func (u *UserProvider) Enable(ctx context.Context, in *apipb.EnableRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: commonmodel.Success,
	}
	err := model.EnableUser(in.Id, in.Enable)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *UserProvider) GetAll(ctx context.Context, in *apipb.GetAllUsersRequest) (*apipb.GetAllUsersResponse, error) {
	resp := &apipb.GetAllUsersResponse{
		Code: commonmodel.Success,
	}
	users, err := model.GetAllUsers(in)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.UsersToPB(users)
	}

	return resp, nil
}

func (u *UserProvider) GetDetail(ctx context.Context, in *apipb.GetDetailRequest) (*apipb.GetUserDetailResponse, error) {
	resp := &apipb.GetUserDetailResponse{
		Code: commonmodel.Success,
	}
	user, err := model.GetUserById(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	resp.Data = model.UserToPB(&user)
	return resp, nil
}

func (u *UserProvider) ResetPwd(ctx context.Context, in *apipb.GetDetailRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: commonmodel.Success,
	}
	err := model.ResetPwd(in.Id, model.DefaultPwd)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *UserProvider) ChangePwd(ctx context.Context, in *apipb.ChangePwdRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: commonmodel.Success,
	}
	err := model.UpdatePwd(in.Id, in.OldPwd, in.NewPwd)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *UserProvider) Logout(ctx context.Context, in *apipb.LogoutRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{
		Code: commonmodel.Success,
	}
	//TODO
	err := model.Logout(in.Token)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}
