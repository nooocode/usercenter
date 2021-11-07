package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"codeup.aliyun.com/atali/usercenter/provider"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

func main() {
	config.SetProviderService(&provider.UserProvider{})
	if err := config.Load(config.WithPath("./dubbogo.yaml")); err != nil {
		panic(err)
	}
	fmt.Println("started server")
	initSignal()
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
