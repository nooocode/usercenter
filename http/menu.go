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

// AddMenu
// @Summary 新增菜单
// @Description 新增菜单
// @Tags 菜单管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.MenuInfo true "请求参数"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/core/auth/menu/add [post]
func AddMenu(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.MenuInfo{}
	resp := &model.CommonResponse{
		Code: model.Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建Menu请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = ucmodel.AddMenu(ucmodel.PBToMenu(req))
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateMenu
// @Summary 更新菜单
// @Description 更新菜单
// @Tags 菜单管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.MenuInfo true "请求参数"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/core/auth/menu/update [put]
func UpdateMenu(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.MenuInfo{}
	resp := &apipb.CommonResponse{
		Code: model.Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建Menu请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = ucmodel.UpdateMenu(ucmodel.PBToMenu(req))
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteMenu
// @Summary 删除菜单
// @Description 软删除菜单
// @Tags 菜单管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.DelRequest true "请求参数"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/core/auth/menu/delete [delete]
func DeleteMenu(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,新建Menu请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = ucmodel.DeleteMenu(req.Id)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryMenu
// @Summary 分页查询
// @Description 分页查询
// @Tags 菜单管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param path query string false "路径"
// @Param name query string false "名称"
// @Param title query string false "显示名称"
// @Param parentID query string false "父ID"
// @Param level query int false "层级"
// @Success 200 {object} apipb.QueryMenuResponse
// @Router /api/core/auth/menu/query [get]
func QueryMenu(c *gin.Context) {
	req := &apipb.QueryMenuRequest{}
	resp := &apipb.QueryMenuResponse{
		Code: model.Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	ucmodel.QueryMenu(req, resp)

	c.JSON(http.StatusOK, resp)
}

// GetMenuDetail
// @Summary 查询明细
// @Description 查询明细
// @Tags 菜单管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetMenuDetailResponse
// @Router /api/core/auth/menu/detail [get]
func GetMenuDetail(c *gin.Context) {
	resp := &apipb.GetMenuDetailResponse{
		Code: model.Success,
	}
	idStr := c.Query("id")
	if idStr == "" {
		resp.Code = model.BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := ucmodel.GetMenuByID(idStr)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = ucmodel.MenuToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// GetMenuTree
// @Summary 查询所有菜单（Tree）
// @Description 查询所有菜单（Tree）
// @Tags 菜单管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.QueryMenuResponse
// @Router /api/core/auth/menu/all [get]
func GetMenuTree(c *gin.Context) {
	resp := &apipb.QueryMenuResponse{
		Code: model.Success,
	}
	var err error
	data, records, err := ucmodel.GetInfoList()
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = ucmodel.MenusToPB(data)
		resp.Records = records
	}
	c.JSON(http.StatusOK, resp)
}
