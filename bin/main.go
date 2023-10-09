package main

import (
	"github.com/aliffatulmf/pink/routers"
	"github.com/aliffatulmf/pink/server"
)

func main() {
	serve := server.NewServer()
	defer serve.Run()

	routers.NewRouter(serve.Engine())
}
