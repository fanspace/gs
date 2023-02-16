package main

import (
	"gingate/api"
	"gingate/core"
	_ "gingate/docs"
	"os"
)

// @title Gate Api
// @version 1.0
// @description an api gateway.

// @contact.email lf6128@163.com

// @license.name GPL v3
// @license.url http://www.gnu.org/licenses/quick-guide-gplv3.html

// @host 127.0.0.1:8001
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	router := api.InitRouter()
	srv := core.NewServer()
	defer srv.ShutDown()
	core.InitComponent()
	srv.Router = router
	err := srv.Start()
	if err != nil {
		srv.Log.Error(err.Error())
		os.Exit(0)
	}
}
