package http

import (
	"context"
	"fmt"
	"net/http"

	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/gin-gonic/gin"
	"github.com/nooocode/pkg/constants"
	"github.com/nooocode/pkg/model"
	"github.com/nooocode/pkg/utils/log"
	"github.com/nooocode/pkg/utils/middleware"
	apipb "github.com/nooocode/usercenter/api"
	ucmodel "github.com/nooocode/usercenter/model"
	userm "github.com/nooocode/usercenter/utils/middleware"
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
	JsCode          string `json:"jsCode"`
	PhoneNumberCode string `json:"phoneNumberCode"`
	EncryptedData   string `json:"encryptedData"`
	IV              string `json:"iv"`
	Register        bool   `json:"register"`
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
		user.Avatar = plainData.AvatarURL
		user.Nickname = plainData.NickName
		user.City = plainData.City
		user.Country = plainData.Country
		user.Province = plainData.Province
		user.Gender = plainData.Gender == 1
		user.Mobile = plainData.PhoneNumber
		user.Enable = true
		user.UserRoles = []*apipb.UserRole{
			{RoleID: constants.DefaultRoleID},
		}
		//获取手机号
		if req.PhoneNumberCode != "" {
			result2, err := miniProgram.GetAuth().GetPhoneNumber(req.PhoneNumberCode)
			if err != nil {
				log.Warnf(context.Background(), "TransID:%s,GetPhoneNumber:%v", transID, err)
			} else {
				user.Mobile = result2.PhoneInfo.PhoneNumber
			}
		}
		if user.Mobile != "" {
			user.UserName = user.Mobile
		} else if result.UnionID != "" {
			user.UserName = result.UnionID
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

func bindPhone(c *gin.Context) {
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
	//获取手机号
	phoneNumber := ""
	if req.PhoneNumberCode != "" {
		result2, err := miniProgram.GetAuth().GetPhoneNumber(req.PhoneNumberCode)
		if err != nil {
			log.Warnf(context.Background(), "TransID:%s,GetPhoneNumber:%v", transID, err)
			resp.Code = model.InternalServerError
			resp.Message = err.Error()
		} else {
			phoneNumber = result2.PhoneInfo.PhoneNumber
		}
	}
	err = ucmodel.BindPhone(userm.GetUserID(c), phoneNumber)
	if err != nil {
		resp.Code = model.InternalServerError
		resp.Message = err.Error()
		log.Warnf(context.Background(), "TransID:%s,BindPhone Error:%v", transID, err)
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterWechatRouter(r *gin.Engine) {
	g := r.Group("/api/wechat")
	g.GET("notify", WechatNotify)
	g.POST("mini/login", wechatMiniLogin)
	g.POST("mini/register/check", wechatMiniCheckRegister)
	g.POST("mini/phone/bind", bindPhone)
	InitMini()
}
