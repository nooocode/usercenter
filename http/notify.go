package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

///api/wechat/notify?signature=779ee239e53c506537b56e530cd96bd5869c890a&echostr=7007211687744363958×tamp=1632722226&nonce=1266362590
func WechatNotify(c *gin.Context) {
	c.String(http.StatusOK, "%s", c.Query("echostr"))
}
