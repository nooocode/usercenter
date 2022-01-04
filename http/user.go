package http

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nooocode/pkg/constants"
	"github.com/nooocode/pkg/model"
	"github.com/nooocode/pkg/utils/log"
	apipb "github.com/nooocode/usercenter/api"
	ucmodel "github.com/nooocode/usercenter/model"
	"github.com/nooocode/usercenter/utils/middleware"
)

// Login
// @Summary 登录
// @Description 登录
// @Tags 用户管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "Bearer+空格+Token"
// @Param product body apipb.LoginRequest true "个人信息"
// @Success 200 {object} apipb.LoginResponse
// @Router /api/core/auth/user/login [post]
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

// UpdateProfile
// @Summary 获取个人信息
// @Description 获取个人信息
// @Tags 用户管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "Bearer+空格+Token"
// @Success 200 {object} apipb.UserProfile
// @Router /api/core/auth/user/profile [get]
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

// UpdateProfile
// @Summary 更新个人信息
// @Description 更新个人信息
// @Tags 用户管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "Bearer+空格+Token"
// @Param product body apipb.UserProfile true "个人信息"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/core/auth/user/profile [put]
func UpdateProfile(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.UserProfile{}
	resp := &apipb.CommonResponse{
		Code: model.Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新个人信息请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	err = ucmodel.UpdateProfile(&ucmodel.User{
		TenantModel: model.TenantModel{
			Model: model.Model{
				ID: req.Id,
			},
		},
		Nickname: req.Nickname,
		Email:    req.Email,
		Mobile:   req.Mobile,
		IDCard:   req.IdCard,
		Avatar:   req.Avatar,
		RealName: req.RealName,
		Gender:   req.Gender,
	})
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// AddUser godoc
// @Summary 新增用户
// @Description 新增用户
// @Tags 用户管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.UserInfo true "用户信息"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/core/auth/user/add [post]
func AddUser(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.UserInfo{}
	resp := &apipb.CommonResponse{
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
	//只有平台租户才能为其他租户创建用户
	tenantID := middleware.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	err = ucmodel.CreateUser(ucmodel.PBToUser(req), false)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateUser
// @Summary 更新用户
// @Description 更新用户
// @Tags 用户管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.UserInfo true "用户信息"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/core/auth/user/update [put]
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
	//只有平台租户才能更改用户的租户
	tenantID := middleware.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	err = ucmodel.UpdateUser(req)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteUser
// @Summary 删除用户
// @Description 删除用户
// @Tags 用户管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.DelRequest true "请求参数"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/core/auth/user/delete [delete]
func DeleteUser(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.DelRequest{}
	resp := &apipb.CommonResponse{
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
	err = ucmodel.DeleteUser(req.Id)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// EnableUser
// @Summary 禁用/启用用户
// @Description 禁用/启用用户
// @Tags 用户管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.EnableRequest true "请求参数"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/core/auth/user/enable [post]
func EnableUser(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.EnableRequest{}
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

	err = ucmodel.EnableUser(req.Id, req.Enable)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryUser
// @Summary 分页查询
// @Description 分页查询
// @Tags 用户管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param userName query string false "用户名"
// @Param nickname query string false "昵称"
// @Param idCard query string false "身份证号"
// @Param mobile query string false "手机号"
// @Param title query string false "职位"
// @Param type query int false "用户类型,从1开始,为0时查询全部"
// @Param tenantID query string false "租户ID"
// @Param group query string false "分组ID，例如属于某个组织的，或者某个个人"
// @Success 200 {object} apipb.QueryUserResponse
// @Router /api/core/auth/user/query [get]
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
	//只有平台租户才能查询其他租户的角色
	tenantID := middleware.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	ucmodel.QueryUser(req, resp)

	c.JSON(http.StatusOK, resp)
}

// GetAllUsers
// @Summary 查询所有用户
// @Description 查询所有用户
// @Tags 用户管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param type query int false "用户类型,从1开始,为0时查询全部"
// @Param tenantID query string false "租户ID"
// @Param group query string false "分组ID，例如属于某个组织的，或者某个个人"
// @Success 200 {object} apipb.GetAllUsersResponse
// @Router /api/core/auth/user/all [get]
func GetAllUsers(c *gin.Context) {
	resp := &apipb.GetAllUsersResponse{
		Code: model.Success,
	}
	req := &apipb.GetAllUsersRequest{}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	//只有平台租户才能查询其他租户的角色
	tenantID := middleware.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	users, err := ucmodel.GetAllUsers(req)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = ucmodel.UsersToPB(users)
	c.JSON(http.StatusOK, resp)
}

// GetUserDetail
// @Summary 查询明细
// @Description 查询明细
// @Tags 用户管理
// @Accept  json
// @Produce  json
// @Param id query string true "用户ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetUserDetailResponse
// @Router /api/core/auth/user/detail [get]
func GetUserDetail(c *gin.Context) {
	resp := &apipb.GetUserDetailResponse{
		Code: model.Success,
	}
	idStr := c.Query("id")
	if idStr == "" {
		resp.Code = model.BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := ucmodel.GetUserById(idStr)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = ucmodel.UserToPB(&data)
	}
	c.JSON(http.StatusOK, resp)
}

// ResetPwd
// @Summary 重置密码
// @Description 重置密码
// @Tags 用户管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.GetDetailRequest true "请求参数"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/core/auth/user/resetpwd [post]
func ResetPwd(c *gin.Context) {
	resp := &apipb.CommonResponse{
		Code: model.Success,
	}
	req := &apipb.GetDetailRequest{}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = ucmodel.ResetPwd(req.Id, ucmodel.DefaultPwd)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// ChangePwd
// @Summary 修改密码
// @Description 修改密码
// @Tags 用户管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.ChangePwdRequest true "请求参数"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/core/auth/user/changepwd [post]
func ChangePwd(c *gin.Context) {
	resp := &apipb.CommonResponse{
		Code: model.Success,
	}
	req := &apipb.ChangePwdRequest{}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	if req.NewPwd != req.NewConfirmPwd {
		resp.Code = model.BadRequest
		resp.Message = "新密码和确认密码不一样"
		c.JSON(http.StatusOK, resp)
		return
	}
	if req.Id == "" {
		req.Id = middleware.GetUserID(c)
	}
	err = ucmodel.UpdatePwd(req.Id, req.OldPwd, req.NewPwd)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// Logout
// @Summary 退出登录
// @Description 退出登录
// @Tags 用户管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/core/auth/user/logout [post]
func Logout(c *gin.Context) {
	resp := &apipb.CommonResponse{
		Code: model.Success,
	}
	t := middleware.GetAccessToken(c)
	err := ucmodel.Logout(t)
	if err != nil {
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}
