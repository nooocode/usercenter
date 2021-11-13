package http

import (
	"github.com/gin-gonic/gin"
	"github.com/nooocode/usercenter/utils/middleware"
)

func RegisterAuthRouter(r *gin.Engine) {
	r.Use(middleware.AuthRequired)
	userGroup := r.Group("/api/core/auth/user")
	userGroup.POST("login", Login)
	userGroup.POST("logout", Logout)
	userGroup.GET("profile", Profile)
	userGroup.POST("add", AddUser)
	userGroup.PUT("update", UpdateUser)
	userGroup.GET("query", QueryUser)
	userGroup.DELETE("delete", DeleteUser)
	userGroup.POST("enable", EnableUser)
	userGroup.GET("all", GetAllUser)
	userGroup.GET("detail", GetUserDetail)
	userGroup.POST("resetpwd", ResetPwd)
	userGroup.POST("changepwd", ChangePwd)

	roleGroup := r.Group("/api/core/auth/role")
	roleGroup.POST("add", AddRole)
	roleGroup.PUT("update", UpdateRole)
	roleGroup.GET("query", QueryRole)
	roleGroup.DELETE("delete", DeleteRole)
	roleGroup.GET("all", GetAllRole)
	roleGroup.GET("detail", GetRoleDetail)

	titleGroup := r.Group("/api/core/auth/title")
	titleGroup.POST("add", AddTitle)
	titleGroup.PUT("update", UpdateTitle)
	titleGroup.GET("query", QueryTitle)
	titleGroup.DELETE("delete", DeleteTitle)
	titleGroup.GET("all", GetAllTitle)
	titleGroup.GET("detail", GetTitleDetail)

	appGroup := r.Group("/api/core/auth/app")
	appGroup.POST("add", AddAPP)
	appGroup.PUT("update", UpdateAPP)
	appGroup.GET("query", QueryAPP)
	appGroup.DELETE("delete", DeleteAPP)
	appGroup.GET("all", GetAllAPP)
	appGroup.GET("detail", GetAPPDetail)

	menuGroup := r.Group("/api/core/auth/menu")
	menuGroup.POST("add", AddMenu)
	menuGroup.PUT("update", UpdateMenu)
	menuGroup.GET("query", QueryMenu)
	menuGroup.DELETE("delete", DeleteMenu)
	menuGroup.GET("detail", GetMenuDetail)
	menuGroup.GET("tree", GetMenuTree)

	apiGroup := r.Group("/api/core/auth/api")
	apiGroup.POST("add", AddAPI)
	apiGroup.PUT("update", UpdateAPI)
	apiGroup.GET("query", QueryAPI)
	apiGroup.DELETE("delete", DeleteAPI)
	apiGroup.POST("enable", EnableAPI)
	apiGroup.GET("all", GetAllAPI)
	apiGroup.GET("detail", GetAPIDetail)

	RegisterTenantRouter(r)
}
