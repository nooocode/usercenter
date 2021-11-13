package http

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nooocode/pkg/model"
	"github.com/nooocode/pkg/utils/log"
	apipb "github.com/nooocode/usercenter/api"
	ucmodel "github.com/nooocode/usercenter/model"
	"github.com/nooocode/usercenter/utils/middleware"
)

func Login(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.LoginRequest{}
	resp := &apipb.LoginResponse{
		Code: model.Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建User请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	ucmodel.Login(req, resp)

	c.JSON(http.StatusOK, resp)
}

func Profile(c *gin.Context) {
	userID := middleware.GetUserID(c)
	userProfile, err := ucmodel.GetUserProfile(userID)
	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code":    model.InternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": model.Success,
		"data": userProfile,
	})
}

func AddUser(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &ucmodel.User{}
	resp := &model.CommonResponse{
		Code: model.Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建User请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	req.TenantID = middleware.GetTenantID(c)
	err = ucmodel.CreateUser(req, false)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func UpdateUser(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &ucmodel.User{}
	resp := &model.CommonResponse{
		Code: model.Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建User请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = ucmodel.UpdateUser(req)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func DeleteUser(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &ucmodel.User{}
	resp := &model.CommonResponse{
		Code: model.Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建User请求参数无效:%v", transID, err)
		return
	}
	err = ucmodel.DeleteUser(req.ID)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func EnableUser(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &ucmodel.User{}
	resp := &model.CommonResponse{
		Code: model.Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建User请求参数无效:%v", transID, err)
		return
	}

	err = ucmodel.EnableUser(req.ID, req.Enable)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func QueryUser(c *gin.Context) {
	req := &apipb.QueryUserRequest{}
	resp := &apipb.QueryUserResponse{
		Code: model.Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	ucmodel.QueryUser(req, resp)

	c.JSON(http.StatusOK, resp)
}

func GetAllUser(c *gin.Context) {
	resp := &ucmodel.QueryUserResponse{
		CommonResponse: model.CommonResponse{
			Code: model.Success,
		},
	}
	users, err := ucmodel.GetAllUsers()
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = users
	resp.Records = int64(len(users))
	resp.Pages = 1
	c.JSON(http.StatusOK, resp)
}

func GetUserDetail(c *gin.Context) {
	resp := model.CommonDetailResponse{
		CommonResponse: model.CommonResponse{
			Code: model.Success,
		},
	}
	idStr := c.Query("id")
	if idStr == "" {
		resp.Code = model.BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	resp.Data, err = ucmodel.GetUserById(idStr)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

type ResetPwdRequest struct {
	model.CommonRequest
	ID string `json:"id" form:"id" uri:"id"`
}

func ResetPwd(c *gin.Context) {
	resp := model.CommonDetailResponse{}
	req := &ResetPwdRequest{}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = ucmodel.ResetPwd(req.ID, ucmodel.DefaultPwd)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

type ChangePwdRequest struct {
	model.CommonRequest
	ID     string `json:"id" form:"id" uri:"id"`
	OldPwd string `json:"oldPwd" form:"oldPwd" uri:"oldPwd"`
	NewPwd string `json:"newPwd" form:"newPwd" uri:"newPwd"`
}

func ChangePwd(c *gin.Context) {
	resp := model.CommonDetailResponse{}
	req := &ChangePwdRequest{}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = ucmodel.UpdatePwd(req.ID, req.OldPwd, req.NewPwd)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func Logout(c *gin.Context) {
	resp := model.CommonResponse{
		Code: model.Success,
	}
	t := middleware.GetAccessToken(c)
	err := ucmodel.Logout(t)
	if err != nil {
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}
