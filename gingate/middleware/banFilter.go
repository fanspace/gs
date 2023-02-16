package middleware

import (
	"gingate/core"
	"gingate/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FilterBan() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := core.RemoteIp(c.Request)
		ipband := core.IsIpBand(ip)
		if ipband {
			c.Set("mc", new(core.MyClaim))
			c.JSON(http.StatusUnauthorized, gin.H{"code": 403, "result": model.SimpleResponse{Success: false, Msg: "检测到异常操作，该ip已被封禁60分钟"}})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}
