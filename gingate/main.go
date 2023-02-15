package main

import (
	"gingate/api"
	"gingate/component"
	"gingate/core"
	"os"
)

func main() {
	router := api.InitRouter()
	srv := core.NewServer()
	defer srv.ShutDown()
	component.InitComponent()
	defer component.CloseComponent()
	srv.Router = router
	err := srv.Start()
	if err != nil {
		srv.Log.Error(err.Error())
		os.Exit(0)
	}
}
