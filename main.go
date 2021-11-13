package main

import (
	"fmt"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/gin-gonic/gin"
	"github.com/nooocode/pkg/utils"
	"github.com/nooocode/usercenter/http"
	"github.com/nooocode/usercenter/model"
	"github.com/nooocode/usercenter/model/token"
	"github.com/nooocode/usercenter/provider"
)

func main() {
	config.SetProviderService(&provider.UserProvider{})
	config.SetProviderService(&provider.TenantProvider{})
	config.SetProviderService(&provider.RoleProvider{})
	config.SetProviderService(&provider.MenuProvider{})
	config.SetProviderService(&provider.APIProvider{})
	config.SetProviderService(&provider.IdentityProvider{})
	if err := config.Load(); err != nil {
		panic(err)
	}
	params := config.GetRootConfig().ConfigCenter.Params
	fmt.Println(params)
	model.Init(params["mysql"], params["debug"] == "true")
	token.InitTokenCache(params["token-key"], params["redis-addr"], params["redis-user-name"], params["redis-pwd"], 120)
	fmt.Println("started server")
	Start(48080)
}

func Start(port int) {
	r := gin.Default()
	r.Use(utils.Cors())
	http.RegisterAuthRouter(r)

	r.Run(fmt.Sprintf(":%d", port))
}
