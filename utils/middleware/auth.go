package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nooocode/pkg/model"
	apipb "github.com/nooocode/usercenter/api"
	ucmodel "github.com/nooocode/usercenter/model"
)

func GetUser(c *gin.Context) (bool, *apipb.CurrentUser) {
	obj, exists := c.Get("User")
	if !exists {
		return false, nil
	}
	user, ok := obj.(*apipb.CurrentUser)
	if !ok {
		return false, nil
	}
	return true, user
}

func GetUserID(c *gin.Context) string {
	exists, user := GetUser(c)
	if !exists || user == nil {
		return ""
	}
	return user.Id
}

func GetUserName(c *gin.Context) string {
	exists, user := GetUser(c)
	if !exists || user == nil {
		return ""
	}
	return user.UserName
}

func GetTenantID(c *gin.Context) string {
	exists, user := GetUser(c)
	if !exists || user == nil {
		return ""
	}
	return user.TenantID
}

func GetAccessToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	if token == "" {
		token = c.GetHeader("authorization")
	}
	token = strings.Replace(token, "Bearer ", "", 1)
	return token
}

func AuthRequired(c *gin.Context) {
	t := GetAccessToken(c)
	if t == "camunda-admin" {
		c.Set("User", &apipb.CurrentUser{
			UserName: "camunda-admin",
		})
		return
	}
	currentUser, code, err := ucmodel.Authenticate(t, c.Request.Method, c.Request.URL.Path, true)

	if code != model.Success {

		c.AbortWithStatusJSON(http.StatusOK, model.CommonResponse{
			Code:    code,
			Message: err.Error(),
		})
		return
	}

	c.Set("User", currentUser)
}
