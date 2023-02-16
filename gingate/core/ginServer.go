package core

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

type GracefulServer struct {
	Server *http.Server
	Router *gin.Engine
	//SqlSupplier *SqlSupplier
	Log *Logger
}

var Srv *GracefulServer

func LoggerConfigFromLoggerConfig() *LoggerConfiguration {
	return &LoggerConfiguration{
		EnableConsole: Cfg.LogSettings.EnableConsole,
		ConsoleJson:   *Cfg.LogSettings.ConsoleJson,
		ConsoleLevel:  strings.ToLower(Cfg.LogSettings.ConsoleLevel),
		EnableFile:    Cfg.LogSettings.EnableFile,
		FileJson:      *Cfg.LogSettings.FileJson,
		FileLevel:     strings.ToLower(Cfg.LogSettings.FileLevel),
		FileLocation:  GetLogFileLocation(Cfg.LogSettings.FileLocation),
	}
}

func NewServer() *GracefulServer {
	log.Println("project professional And tecnical personnel api gate server")
	Srv = &GracefulServer{}
	Srv.Log = NewLogger(LoggerConfigFromLoggerConfig())
	RedirectStdLog(Srv.Log)
	// 使用server logger 作为全局的logger
	InitGlobalLogger(Srv.Log)
	return Srv
}

func (gracefulServer *GracefulServer) Start() error {
	log.Println("server start")
	gracefulServer.Server = &http.Server{
		Addr:    Cfg.HttpSettings.ListenAddress,
		Handler: gracefulServer.Router,

		//ReadTimeout:  time.Duration(Cfg.HttpSettings.ReadTimeout) * time.Second,
		//	WriteTimeout: time.Duration(Cfg.HttpSettings.WriteTimeout) * time.Second,
	}
	log.Println("Start Listening and serving HTTP on " + Cfg.HttpSettings.ListenAddress)

	return nil
}

func (gracefulServer *GracefulServer) ShutDown() {

	go func() {
		// service connections
		if err := Srv.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("listen: %s\n")
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("********************************************    Shutdown Server    ********************************************")
	defer CloseComponent()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := Srv.Server.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown err:" + err.Error())
	}
	Info("Server exiting")
}
