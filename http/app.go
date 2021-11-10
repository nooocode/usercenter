package http

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nooocode/pkg/model"
	"github.com/nooocode/pkg/utils/log"
	"github.com/nooocode/pkg/utils/middleware"
	ucmodel "github.com/nooocode/usercenter/model"
)

func AddAPP(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &ucmodel.APP{}
	resp := &model.CommonResponse{
		Code: model.Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建APP请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = ucmodel.CreateAPP(req)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func UpdateAPP(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &ucmodel.APP{}
	resp := &model.CommonResponse{
		Code: model.Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建APP请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = ucmodel.UpdateAPP(req)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func DeleteAPP(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &ucmodel.APP{}
	resp := &model.CommonResponse{
		Code: model.Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建APP请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = ucmodel.DeleteAPP(req.ID)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func QueryAPP(c *gin.Context) {
	req := &ucmodel.QueryAPPRequest{}
	resp := &ucmodel.QueryAPPResponse{
		CommonResponse: model.CommonResponse{
			Code: model.Success,
		},
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	ucmodel.QueryAPP(req, resp)

	c.JSON(http.StatusOK, resp)
}

func GetAllAPP(c *gin.Context) {
	resp := &ucmodel.QueryAPPResponse{
		CommonResponse: model.CommonResponse{
			Code: model.Success,
		},
	}
	metadatas, err := ucmodel.GetAllAPPs()
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = metadatas
	resp.Records = int64(len(metadatas))
	resp.Pages = 1
	c.JSON(http.StatusOK, resp)
}

func GetAPPDetail(c *gin.Context) {
	resp := model.CommonDetailResponse{}
	idStr := c.Query("id")
	if idStr == "" {
		resp.Code = model.BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	resp.Data, err = ucmodel.GetAPPById(idStr)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}
