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

func AddTenant(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &ucmodel.Tenant{}
	resp := &model.CommonResponse{
		Code: model.Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建租户请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = ucmodel.CreateTenant(req)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func UpdateTenant(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &ucmodel.Tenant{}
	resp := &model.CommonResponse{
		Code: model.Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建租户请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = ucmodel.UpdateTenant(req)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func DeleteTenant(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &ucmodel.Tenant{}
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

	if req.ID == "" {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = ucmodel.DeleteTenant(req.ID)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func EnableTenant(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &ucmodel.Tenant{}
	resp := &model.CommonResponse{
		Code: model.Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建租户请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = ucmodel.EnableTenant(req.ID, req.Enable)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func QueryTenant(c *gin.Context) {
	req := &apipb.QueryTenantRequest{}
	resp := &apipb.QueryTenantResponse{
		Code: model.Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	ucmodel.QueryTenant(req, resp)

	c.JSON(http.StatusOK, resp)
}

func GetAllTenant(c *gin.Context) {
	resp := &ucmodel.QueryTenantResponse{
		CommonResponse: model.CommonResponse{
			Code: model.Success,
		},
	}
	apis, err := ucmodel.GetAllTenant()
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = apis
	resp.Records = int64(len(apis))
	resp.Pages = 1
	c.JSON(http.StatusOK, resp)
}

func GetTenantDetail(c *gin.Context) {
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

	resp.Data, err = ucmodel.GetTenantByID(idStr)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func CopyTenant(c *gin.Context) {
	resp := model.CommonResponse{}
	idStr := c.Query("id")
	if idStr == "" {
		resp.Code = model.BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}

	err := ucmodel.CopyTenant(idStr)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterTenantRouter(r *gin.Engine) {
	g := r.Group("/api/core/auth/tenant")

	g.POST("add", AddTenant)
	g.PUT("update", UpdateTenant)
	g.GET("query", QueryTenant)
	g.DELETE("delete", DeleteTenant)
	g.GET("all", GetAllTenant)
	g.GET("detail", GetTenantDetail)
	g.POST("copy", CopyTenant)
	g.POST("enable", EnableTenant)
}
