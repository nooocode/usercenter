package provider

import (
	"context"

	commonmodel "github.com/nooocode/pkg/model"
	apipb "github.com/nooocode/usercenter/api"
	"github.com/nooocode/usercenter/model"
	"github.com/nooocode/usercenter/model/token"
)

type IdentityProvider struct {
	apipb.UnimplementedIdentityServer
}

func (u *IdentityProvider) Authenticate(ctx context.Context, in *apipb.AuthenticateRequest) (*apipb.AuthenticateResponse, error) {
	resp := &apipb.AuthenticateResponse{
		Code: commonmodel.Success,
	}
	currentUser, code, err := model.Authenticate(in.Token, in.Method, in.Url, in.CheckAuth)
	if code != int(apipb.Code_Success) {
		resp.Code = apipb.Code(code)
		if err != nil {
			resp.Message = err.Error()
		}
	} else {
		resp.CurrentUser = currentUser
		resp.Code = apipb.Code(code)
	}
	return resp, nil
}

func (u *IdentityProvider) DecodeToken(ctx context.Context, in *apipb.DecodeTokenRequest) (*apipb.AuthenticateResponse, error) {
	resp := &apipb.AuthenticateResponse{
		Code: commonmodel.Success,
	}
	currentUser, err := token.DecodeToken(in.Token)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.CurrentUser = currentUser
		resp.Code = commonmodel.Success
	}
	return resp, nil
}
