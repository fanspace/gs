package api

import (
	"gingate/core"
	. "gingate/internal/controller"
	"gingate/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	const (
		prefix = core.APISITE_PREFIX
	)
	if core.Cfg.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	config := cors.DefaultConfig()
	config.AllowOrigins = core.Cfg.CorsSettings.AllowOrigins
	config.AllowMethods = core.Cfg.CorsSettings.AllowMethods
	//config.AddAllowHeaders("Authorization", "Logintype")
	router.Use(cors.New(config), gin.Recovery(), middleware.FilterBan())
	if !core.Cfg.ReleaseMode {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	router.GET(prefix+"/ping", Ping)
	userGroup := router.Group(prefix + "/user")
	userGroup.Use()
	{
		userGroup.GET("/:id", GetUser)
		userGroup.POST("/list", QueryUsers)
	}
	articleGroup := router.Group(prefix + "/article")
	//articleGroup.Use(middleware.MustLogin())
	articleGroup.Use()
	{
		articleGroup.GET("/:id", GetArticle)
		articleGroup.POST("/list", QueryArticles)
	}
	return router
}
