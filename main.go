package main

import (
	"fmt"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/gin-gonic/gin"
	"github.com/nooocode/pkg/utils"
	"github.com/nooocode/usercenter/docs"
	"github.com/nooocode/usercenter/http"
	"github.com/nooocode/usercenter/model"
	"github.com/nooocode/usercenter/model/token"
	"github.com/nooocode/usercenter/provider"
	"github.com/nooocode/usercenter/utils/middleware"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// gin-swagger middleware
// swagger embed files
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
	// programatically set swagger info
	docs.SwaggerInfo.Title = "User Center API"
	docs.SwaggerInfo.Description = "This is a User Center server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "noocode.cn"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()
	r.Use(utils.Cors())
	http.RegisterAuthRouter(r)
	r.GET("/usercenter/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(middleware.AuthRequired)
	r.Run(fmt.Sprintf(":%d", port))
}
