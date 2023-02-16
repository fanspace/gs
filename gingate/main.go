package main

import (
	"gingate/api"
	"gingate/core"
	"os"
)

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
