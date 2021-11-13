package provider

import (
	"context"

	apipb "github.com/nooocode/usercenter/api"
	"github.com/nooocode/usercenter/model"
)

type IdentityProvider struct {
	apipb.UnimplementedAPIServer
}

func (u *IdentityProvider) Authenticate(ctx context.Context, in *apipb.AuthenticateRequest) (*apipb.AuthenticateResponse, error) {
	resp := &apipb.AuthenticateResponse{}
	currentUser, code, err := model.Authenticate(in.Token, in.Method, in.Url, in.CheckAuth)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.CurrentUser = currentUser
		resp.Code = apipb.Code(code)
	}
	return resp, nil
}
