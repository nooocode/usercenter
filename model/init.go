package model

import (
	"github.com/nooocode/pkg/db/mysql"
	"go.uber.org/zap"
)

var dbClient *mysql.Mysql

//Init Init
func Init(connStr string, debug bool) {
	dbClient = mysql.NewMysql(connStr, debug)
	if debug {
		AutoMigrate()
	}
	InitCasbin()
	roles, err := GetAllRole("", true)
	if err != nil {
		panic(err)
	}
	for _, role := range roles {
		updateRoleAuth(role.ID)
	}
	updateNotCheckAuthRule()
	updateNotCheckLoginRule()

}

//AutoMigrate 自动生成表
func AutoMigrate() {
	dbClient.DB().AutoMigrate(&CasbinRule{}, &API{}, &Menu{}, &MenuParameter{}, &MenuFunc{}, &MenuFuncApi{}, &Role{}, &RoleMenu{}, &User{}, &UserRole{}, &Title{}, &APP{}, &APPProp{}, &Tenant{})
}

var Logger *zap.Logger

var DefaultPwd = "ABC123def"

func SetDefaultPwd(defaultPwd string) {
	if defaultPwd != "" {
		DefaultPwd = defaultPwd
	}
}
