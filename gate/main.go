package main

import (
	_ "gate/controller"
	"github.com/go-spring/spring-core/gs"
	_ "github.com/go-spring/starter-gin"
	"github.com/labstack/gommon/log"
)

func main() {
	log.Fatal(gs.Run())
}
