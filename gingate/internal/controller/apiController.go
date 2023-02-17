package controller

import (
	"fmt"
	"gingate/core"
	"github.com/gin-gonic/gin"
)

// @Summary 存活测试
// @Description 存活测试
// @Tags Public
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.String(200, "pong")
	//service.TestXorm()
	fmt.Println("*************  Now Smark is : " + core.Cfg.Smark)
	return
}
