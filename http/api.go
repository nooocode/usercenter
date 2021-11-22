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

// AddAPI godoc
// @Summary 新增API
// @Description 新增API
// @Tags API管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.APIInfo true "Add API"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/core/auth/api/add [post]
func AddAPI(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &ucmodel.API{}
	resp := &model.CommonResponse{
		Code: model.Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建API请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = ucmodel.CreateAPI(req)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateAPI godoc
// @Summary 更新API
// @Description 更新API
// @Tags API管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.APIInfo true "Update API"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/core/auth/api/update [put]
func UpdateAPI(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &ucmodel.API{}
	resp := &model.CommonResponse{
		Code: model.Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建API请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = ucmodel.UpdateAPI(req)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteAPI godoc
// @Summary 删除API
// @Description 软删除API
// @Tags API管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.DelRequest true "Delete API"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/core/auth/api/delete [delete]
func DeleteAPI(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &ucmodel.API{}
	resp := &model.CommonResponse{
		Code: model.Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建API请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = ucmodel.DeleteApi(req.ID)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// EnableAPI godoc
// @Summary 禁用/启用API
// @Description 禁用/启用API
// @Tags API管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.EnableRequest true "Enable/Disable API"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/core/auth/api/enable [post]
func EnableAPI(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &ucmodel.API{}
	resp := &model.CommonResponse{
		Code: model.Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建API请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = ucmodel.EnableAPI(req.ID, req.Enable)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryAPI godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags API管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param path query string false "路径"
// @Param method query string false "方法"
// @Param group query string false "分组"
// @Param checkAuth query string false "是否检查权限"
// @Param checkLogin query string false "是否需要登录"
// @Success 200 {object} apipb.QueryAPIResponse
// @Router /api/core/auth/api/query [get]
func QueryAPI(c *gin.Context) {
	req := &apipb.QueryAPIRequest{}
	resp := &apipb.QueryAPIResponse{
		Code: model.Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	ucmodel.QueryAPI(req, resp)

	c.JSON(http.StatusOK, resp)
}

// GetAllAPI godoc
// @Summary 查询所有API
// @Description 查询所有API
// @Tags API管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetAllAPIResponse
// @Router /api/core/auth/api/all [get]
func GetAllAPI(c *gin.Context) {
	resp := &apipb.QueryAPIResponse{
		Code: model.Success,
	}
	apis, err := ucmodel.GetAllAPIs()
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = ucmodel.APIsToPB(apis)
	resp.Records = int64(len(apis))
	resp.Pages = 1
	c.JSON(http.StatusOK, resp)
}

// GetAPIDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags API管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetAPIDetailResponse
// @Router /api/core/auth/api/detail [get]
func GetAPIDetail(c *gin.Context) {
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

	resp.Data, err = ucmodel.GetAPIById(idStr)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}
