package middleware

import (
	"gingate/core"
	"gingate/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func MustLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwttokens := c.Request.Header.Get("Authorization")
		if strings.Contains(jwttokens, "Bearer ") {
			jwttokens = strings.Replace(jwttokens, "Bearer ", "", 1)
			isauth, mc := core.ParseJwt(jwttokens)
			if mc == nil || !isauth {
				c.Set("mc", new(core.MyClaim))
				c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "result": model.SimpleResponse{Success: false, Msg: "未登录或登录已过期"}})
				c.Abort()
				return
			} else {

				isbaned := core.IsUserBaned(mc.Username)
				if isbaned {
					c.Set("mc", new(core.MyClaim))
					c.JSON(http.StatusUnauthorized, gin.H{"code": 403, "result": model.SimpleResponse{Success: false, Msg: "账号已被锁定"}})
					c.Abort()
					return
				} else {
					c.Set("mc", mc)
					c.Next()
				}

			}
		} else {
			c.Set("mc", new(core.MyClaim))
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "result": model.SimpleResponse{Success: false, Msg: "未登录或登录已过期"}})
			c.Abort()
			return
		}
	}
}
