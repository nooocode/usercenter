package model

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/nooocode/pkg/model"
	apipb "github.com/nooocode/usercenter/api"
	"github.com/nooocode/usercenter/model/token"
)

func Authenticate(t, method, url string, checkAuth bool) (*apipb.CurrentUser, int, error) {
	// 判断是否不需要登录
	ok, err := enforcer.Enforce("-1", url, method)
	if err != nil {
		return nil, model.InternalServerError, err
	}
	if ok {
		return nil, model.Success, nil
	}

	if ok, err := token.DefaultTokenCache.Exists("", t); err != nil {
		fmt.Println(err.Error())
		return nil, model.InternalServerError, err
	} else if !ok {
		return nil, model.TokenInvalid, errors.New("token invalid")
	}

	currentUser, err := token.DecodeToken(t)
	if err != nil {
		var expErr *jwt.TokenExpiredError
		if errors.As(err, &expErr) {
			return nil, model.TokenExpired, errors.New("token is expired")
		} else {
			return nil, model.TokenInvalid, err
		}
	}

	if !checkAuth {
		return currentUser, model.Success, nil
	}

	// 判断是否不需要校验权限
	ok, err = enforcer.Enforce("0", url, method)
	if err != nil {
		return nil, model.InternalServerError, err
	}
	if ok {
		return currentUser, model.Success, nil
	}

	for _, roleID := range currentUser.RoleIDs {
		ok, err := enforcer.Enforce(roleID, url, method)
		if err != nil {
			return nil, model.InternalServerError, err
		}
		fmt.Println("====>check auth:", roleID, url, method, ok)
		if ok {
			return currentUser, model.Success, nil
		}
	}

	return currentUser, model.Unauthorized, errors.New("Unauthorized")
}
