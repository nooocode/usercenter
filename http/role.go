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

// AddRole
// @Summary 新增角色
// @Description 新增角色
// @Tags 角色管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.RoleInfo true "请求参数"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/core/auth/role/add [post]
func AddRole(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.RoleInfo{}
	resp := &apipb.CommonResponse{
		Code: model.Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建Role请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	//只有平台租户才能为其他租户创建角色
	tenantID := middleware.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	err = ucmodel.CreateRole(ucmodel.PBToRole(req))
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateRole
// @Summary 更新角色
// @Description 更新角色
// @Tags 角色管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.RoleInfo true "请求参数"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/core/auth/role/update [put]
func UpdateRole(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.RoleInfo{}
	resp := &apipb.CommonResponse{
		Code: model.Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建Role请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	//只有平台租户才能更改角色的租户
	tenantID := middleware.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	err = ucmodel.UpdateRole(ucmodel.PBToRole(req))
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteRole
// @Summary 删除角色
// @Description 软删除角色
// @Tags 角色管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.DelRequest true "请求参数"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/core/auth/role/delete [delete]
func DeleteRole(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,新建Role请求参数无效:%v", transID, err)
		return
	}
	err = ucmodel.DeleteRole(req.Id)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryRole
// @Summary 分页查询
// @Description 分页查询
// @Tags 角色管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param tenantID query string false "租户ID"
// @Param name query string false "名称"
// @Success 200 {object} apipb.QueryRoleResponse
// @Router /api/core/auth/role/query [get]
func QueryRole(c *gin.Context) {
	req := &apipb.QueryRoleRequest{}
	resp := &apipb.QueryRoleResponse{
		Code: model.Success,
	}
	//只有平台租户才能查询其他租户的角色
	tenantID := middleware.GetTenantID(c)
	if tenantID != constants.PlatformTenantID {
		req.TenantID = tenantID
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	ucmodel.QueryRole(req, resp)

	c.JSON(http.StatusOK, resp)
}

// GetRoleDetail
// @Summary 查询明细
// @Description 查询明细
// @Tags 角色管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetRoleDetailResponse
// @Router /api/core/auth/role/detail [get]
func GetRoleDetail(c *gin.Context) {
	resp := &apipb.GetRoleDetailResponse{
		Code: model.Success,
	}
	idStr := c.Query("id")
	if idStr == "" {
		resp.Code = model.BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := ucmodel.GetRoleByID(idStr)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = ucmodel.RoleToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// GetAllRole
// @Summary 查询所有角色
// @Description 查询所有角色
// @Tags 角色管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param tenantID query string false "租户ID"
// @Param containerComm query bool false "是否包含公共角色"
// @Success 200 {object} apipb.QueryRoleResponse
// @Router /api/core/auth/role/all [get]
func GetAllRole(c *gin.Context) {
	resp := &apipb.QueryRoleResponse{
		Code: model.Success,
	}
	req := &apipb.GetAllRoleRequest{}
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
	roles, err := ucmodel.GetAllRole(req.TenantID, req.ContainerComm)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = ucmodel.RolesToPB(roles)
	resp.Records = int64(len(roles))
	resp.Pages = 1
	c.JSON(http.StatusOK, resp)
}
