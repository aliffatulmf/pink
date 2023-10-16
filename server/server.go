package main

import (
	"fmt"
	"log"

	"github.com/aliffatulmf/pink/cmd"
	"github.com/aliffatulmf/pink/handlers"
	"github.com/gin-gonic/gin"
)

const PinkLogo = `
  _____________       ______  
  ___  __ \__(_)_________  /__
  __  /_/ /_  /__  __ \_  //_/
  _  ____/_  / _  / / /  ,<   
  /_/     /_/  /_/ /_//_/|_|
`

func setupRouter(router *gin.Engine) *gin.Engine {
	router.GET("/api/check", handlers.VerifyDomain)
	router.GET("/api/list", handlers.ListDomains)
	return router
}

func setupGinEngine(mode string) *gin.Engine {
	switch mode {
	case cmd.ProductionMode:
		gin.SetMode(gin.ReleaseMode)
		engine := gin.New()
		engine.Use(gin.Recovery())

		fmt.Printf("%s\n", PinkLogo)
		fmt.Printf("Running in %s mode", mode)
		return engine
	case cmd.DevelopmentMode:
		return gin.Default()
	default:
		log.Fatal("Invalid environment mode")
		return nil
	}
}

func main() {
	args := cmd.ParseServerFlag()
	engine := setupGinEngine(args.Environment)
	router := setupRouter(engine)

	if err := router.Run(args.Addr()); err != nil {
		log.Fatal(err)
	}
}
