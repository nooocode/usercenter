package config

import (
	"dubbo.apache.org/dubbo-go/v3/common/yaml"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var DefaultConfig = &Config{}

func Init(nacosAddr string, port uint64, nacosUserName, nacosPwd string) {
	sc := []constant.ServerConfig{
		{
			IpAddr: nacosAddr,
			Port:   port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         "", //namespace id
		NotLoadCacheAtStart: true,
		LogDir:              "./log",
		CacheDir:            "./cache",
		LogLevel:            "debug",
		Username:            nacosUserName,
		Password:            nacosPwd,
	}

	// a more graceful way to create config client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	//get config
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: "usercenter-config",
		Group:  "platform",
	})
	err = yaml.UnmarshalYML([]byte(content), DefaultConfig)
	if err != nil {
		panic(err)
	}
}

type Config struct {
	Mysql            string          `yaml:"mysql"`
	Debug            bool            `yaml:"debug"`
	Token            TokenConfig     `yaml:"token"`
	SuperAdminRoleID string          `yaml:"superAdminRoleID"`
	PlatformTenantID string          `yaml:"platformTenantID"`
	DefaultRoleID    string          `yaml:"defaultRoleID"`
	DefaultPwd       string          `yaml:"defaultPwd"`
	MiniApp          []MiniAppConfig `yaml:"miniApp"`
	EnableTenant     bool            `yaml:"enableTenant"`
}

type MiniAppConfig struct {
	ID     string `yaml:"id"`
	Name   string `yaml:"name"`
	Secret string `yaml:"secret"`
}

type TokenConfig struct {
	Key       string `yaml:"key"`
	RedisAddr string `yaml:"redisAddr"`
	RedisName string `yaml:"redisName"`
	RedisPwd  string `yaml:"redisPwd"`
	Expired   int    `yaml:"expired"`
}
