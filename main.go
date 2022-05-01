package main

import (
	"fmt"
	"strconv"
	"strings"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/gin-gonic/gin"
	"github.com/nooocode/pkg/constants"
	"github.com/nooocode/pkg/utils"
	ucconfig "github.com/nooocode/usercenter/config"
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
	nacosAddr := config.GetRootConfig().ConfigCenter.Address
	list := strings.Split(nacosAddr, ":")
	port, err := strconv.ParseUint(list[1], 10, 64)
	if err != nil {
		panic(err)
	}
	ucconfig.Init(list[0], port, config.GetRootConfig().ConfigCenter.Username, config.GetRootConfig().ConfigCenter.Password)
	model.Init(ucconfig.DefaultConfig.Mysql, ucconfig.DefaultConfig.Debug)
	token.InitTokenCache(ucconfig.DefaultConfig.Token.Key, ucconfig.DefaultConfig.Token.RedisAddr, ucconfig.DefaultConfig.Token.RedisName, ucconfig.DefaultConfig.Token.RedisPwd, ucconfig.DefaultConfig.Token.Expired)
	constants.SetPlatformTenantID(ucconfig.DefaultConfig.PlatformTenantID)
	constants.SetSuperAdminRoleID(ucconfig.DefaultConfig.SuperAdminRoleID)
	constants.SetDefaultRoleID(ucconfig.DefaultConfig.DefaultRoleID)
	constants.SetEnabelTenant(ucconfig.DefaultConfig.EnableTenant)
	model.SetDefaultPwd(ucconfig.DefaultConfig.DefaultPwd)
	fmt.Println("started server")
	Start(48080)
}

func Start(port int) {
	// programatically set swagger info
	docs.SwaggerInfo.Title = "User Center API"
	docs.SwaggerInfo.Description = "This is a User Center server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()
	r.Use(middleware.AuthRequired)
	r.Use(utils.Cors())
	http.RegisterAuthRouter(r)
	r.GET("/swagger/usercenter/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(fmt.Sprintf(":%d", port))
}
