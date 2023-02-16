package api

import (
	"gingate/commons"
	"gingate/core"
	. "gingate/internal/controller"
	"gingate/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	const (
		prefix = commons.APISITE_PREFIX
	)
	if core.Cfg.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	config := cors.DefaultConfig()
	config.AllowOrigins = core.Cfg.CorsSettings.AllowOrigins
	config.AllowMethods = core.Cfg.CorsSettings.AllowMethods
	//config.AddAllowHeaders("Authorization", "Logintype")
	router.Use(cors.New(config))
	router.GET(prefix+"/ping", Ping)
	userGroup := router.Group(prefix + "/user")
	userGroup.Use()
	{
		userGroup.GET("/:id", GetUser)
		userGroup.POST("/list", QueryUsers)
	}
	articleGroup := router.Group(prefix + "/article")
	articleGroup.Use(middleware.MustLogin())
	{
		articleGroup.GET("/:id", GetArticle)
		articleGroup.POST("/list", QueryArticles)
	}
	return router
}
