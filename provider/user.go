package provider

import (
	"context"

	commonmodel "codeup.aliyun.com/atali/pkg/model"
	apipb "codeup.aliyun.com/atali/usercenter/api"
	"codeup.aliyun.com/atali/usercenter/model"
)

type UserProvider struct {
	apipb.UnimplementedUserServer
}

func (u *UserProvider) Login(ctx context.Context, in *apipb.LoginRequest) (*apipb.LoginResponse, error) {
	resp := &apipb.LoginResponse{}
	model.Login(in, resp)
	return resp, nil
}

func (u *UserProvider) Add(ctx context.Context, in *apipb.UserInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	err := model.CreateUser(userInfoToUser(in), false)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *UserProvider) Update(ctx context.Context, in *apipb.UserInfo) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	err := model.UpdateUser(userInfoToUser(in))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *UserProvider) Delete(ctx context.Context, in *apipb.DelRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	err := model.DeleteUser(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *UserProvider) Query(ctx context.Context, in *apipb.QueryUserRequest) (*apipb.QueryUserResponse, error) {
	resp := &apipb.QueryUserResponse{}
	model.QueryUser(in, resp)
	return resp, nil
}

func (u *UserProvider) GetProfile(ctx context.Context, in *apipb.GetDetailRequest) (*apipb.GetProfileResponse, error) {
	resp := &apipb.GetProfileResponse{}
	//TODO
	user, err := model.GetUserById(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	resp.Data = userToUserInfo(&user)
	return resp, nil
}

func (u *UserProvider) Enable(ctx context.Context, in *apipb.EnableRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	err := model.EnableUser(in.Id, in.Enable)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *UserProvider) GetAll(ctx context.Context, in *apipb.GetAllRequest) (*apipb.GetAllUserResponse, error) {
	resp := &apipb.GetAllUserResponse{}
	users, err := model.GetAllUsers()
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = userToUserInfos(users)
	}

	return resp, nil
}

func (u *UserProvider) GetDetail(ctx context.Context, in *apipb.GetDetailRequest) (*apipb.GetUserDetailResponse, error) {
	resp := &apipb.GetUserDetailResponse{}
	user, err := model.GetUserById(in.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	resp.Data = userToUserInfo(&user)
	return resp, nil
}

func (u *UserProvider) ResetPwd(ctx context.Context, in *apipb.GetDetailRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	err := model.ResetPwd(in.Id, model.DefaultPwd)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *UserProvider) ChangePwd(ctx context.Context, in *apipb.ChangePwdRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	err := model.UpdatePwd(in.Id, in.OldPwd, in.NewPwd)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func (u *UserProvider) Logout(ctx context.Context, in *apipb.LogoutRequest) (*apipb.CommonResponse, error) {
	resp := &apipb.CommonResponse{}
	//TODO
	err := model.Logout("")
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	return resp, nil
}

func userInfoToUser(in *apipb.UserInfo) *model.User {
	return &model.User{
		TenantModel: commonmodel.TenantModel{
			Model: commonmodel.Model{
				ID: in.Id,
			},
			TenantID: in.Id,
		},
		UserName:    in.UserName,
		Nickname:    in.Nickname,
		UserRoles:   userRolesToUserRoles(in.UserRoles),
		RoleIDs:     in.RoleIDs,
		Enable:      in.Enable,
		Email:       in.Email,
		Mobile:      in.Mobile,
		IDCard:      in.IdCard,
		Avatar:      in.Avatar,
		EID:         in.Eid,
		Title:       in.Title,
		Description: in.Description,
		RealName:    in.RealName,
		Gender:      in.Gender,
	}
}

func userToUserInfo(in *model.User) *apipb.UserInfo {
	return &apipb.UserInfo{
		Id:          in.ID,
		TenantID:    in.ID,
		UserName:    in.UserName,
		Nickname:    in.Nickname,
		UserRoles:   userRolesToPBUserRoles(in.UserRoles),
		RoleIDs:     in.RoleIDs,
		Enable:      in.Enable,
		Email:       in.Email,
		Mobile:      in.Mobile,
		IdCard:      in.IDCard,
		Avatar:      in.Avatar,
		Eid:         in.EID,
		Title:       in.Title,
		Description: in.Description,
		RealName:    in.RealName,
		Gender:      in.Gender,
	}
}

func userToUserInfos(in []*model.User) []*apipb.UserInfo {
	var list []*apipb.UserInfo
	for _, user := range in {
		list = append(list, userToUserInfo(user))
	}
	return list
}

func userRolesToUserRoles(userRoles []*apipb.UserRole) []*model.UserRole {
	var list []*model.UserRole
	for _, userRole := range userRoles {
		list = append(list, &model.UserRole{
			Model: commonmodel.Model{
				ID: userRole.Id,
			},
			UserID: userRole.UserID,
			RoleID: userRole.RoleID,
		})
	}
	return list
}

func userRolesToPBUserRoles(userRoles []*model.UserRole) []*apipb.UserRole {
	var list []*apipb.UserRole
	for _, userRole := range userRoles {
		list = append(list, &apipb.UserRole{
			Id:     userRole.ID,
			UserID: userRole.UserID,
			RoleID: userRole.RoleID,
		})
	}
	return list
}
