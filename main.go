package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"dubbo.apache.org/dubbo-go/v3/common/logger"
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

var (
	survivalTimeout = int(3 * time.Second)
)

func initSignal() {
	signals := make(chan os.Signal, 1)
	// It is not possible to block SIGKILL or syscall.SIGSTOP
	signal.Notify(signals, os.Interrupt, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-signals
		logger.Infof("get signal %s", sig.String())
		switch sig {
		case syscall.SIGHUP:
			// reload()
		default:
			time.Sleep(time.Second * 5)
			time.AfterFunc(time.Duration(survivalTimeout), func() {
				logger.Warnf("app exit now by force...")
				os.Exit(1)
			})

			// The program exits normally or timeout forcibly exits.
			fmt.Println("provider app exit now...")
			return
		}
	}
}
