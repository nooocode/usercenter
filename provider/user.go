package provider

import (
	"context"

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
