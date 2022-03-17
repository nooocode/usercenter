package http

import (
	"context"
	"fmt"
	"net/http"

	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/gin-gonic/gin"
	"github.com/nooocode/pkg/model"
	"github.com/nooocode/pkg/utils/log"
	"github.com/nooocode/pkg/utils/middleware"
	apipb "github.com/nooocode/usercenter/api"
	ucmodel "github.com/nooocode/usercenter/model"
	wchat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
)

var (
	miniProgram   *miniprogram.MiniProgram
	MiniAppID     string
	MiniAppSecret string
)

func InitMini() {
	params := config.GetRootConfig().ConfigCenter.Params
	wc := wchat.NewWechat()
	memory := cache.NewMemory()
	cfg := &miniConfig.Config{
		AppID:     params["miniAppID"],
		AppSecret: params["miniAppSecret"],
		Cache:     memory,
	}
	miniProgram = wc.GetMiniProgram(cfg)

}

type MiniLoginRequest struct {
	JsCode        string `json:"jsCode"`
	EncryptedData string `json:"encryptedData"`
	IV            string `json:"iv"`
	Register      bool   `json:"register"`
}

func wechatMiniLogin(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &MiniLoginRequest{}
	resp := &apipb.LoginResponse{
		Code: model.Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,请求参数无效:%v", transID, err)
		return
	}
	fmt.Printf("req:%#v\n", req)
	result, err := miniProgram.GetAuth().Code2Session(req.JsCode)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,Code2Session出错:%v", transID, err)
		return
	}

	fmt.Printf("Code2Session:%#v\n", result)
	user := &apipb.UserInfo{
		WechatUnionID: result.UnionID,
		WechatOpenID:  result.OpenID,
	}
	if req.Register {
		plainData, err := miniProgram.GetEncryptor().Decrypt(result.SessionKey, req.EncryptedData, req.IV)
		if err != nil {
			resp.Code = model.BadRequest
			resp.Message = err.Error()
			c.JSON(http.StatusOK, resp)
			log.Warnf(context.Background(), "TransID:%s,解密出错:%v", transID, err)
			return
		}

		fmt.Printf("Decrypt:%#v\n", plainData)
		user.WechatUnionID = plainData.UnionID
		user.Avatar = plainData.AvatarURL
		user.Nickname = plainData.NickName
		user.City = plainData.City
		user.Country = plainData.Country
		user.Province = plainData.Province
		user.Gender = plainData.Gender == 1
		user.Mobile = plainData.PhoneNumber
		user.Enable = true
		user.UserRoles = []*apipb.UserRole{
			{RoleID: "9afecaae-2f6a-418e-892d-ce83ae013a42"},
		}
		if plainData.PhoneNumber != "" {
			user.UserName = plainData.PhoneNumber
		} else if plainData.UnionID != "" {
			user.UserName = plainData.UnionID
		} else {
			user.UserName = result.OpenID
		}
	}

	ucmodel.LoginByWechat(req.Register, ucmodel.PBToUser(user), resp)
	c.JSON(http.StatusOK, resp)
}

func wechatMiniCheckRegister(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &MiniLoginRequest{}
	resp := &ucmodel.CheckRegisterWithWechatResp{
		CommonResponse: model.CommonResponse{
			Code: model.Success,
		},
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,请求参数无效:%v", transID, err)
		return
	}
	fmt.Printf("req:%#v\n", req)
	result, err := miniProgram.GetAuth().Code2Session(req.JsCode)
	if err != nil {
		resp.Code = model.BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,Code2Session出错:%v", transID, err)
		return
	}
	ucmodel.CheckRegisterWithWechat(result.UnionID, result.OpenID, resp)
	c.JSON(http.StatusOK, resp)
}
func RegisterWechatRouter(r *gin.Engine) {
	g := r.Group("/api/wechat")
	g.GET("notify", WechatNotify)
	g.POST("mini/login", wechatMiniLogin)
	g.POST("mini/register/check", wechatMiniCheckRegister)
	InitMini()
}
