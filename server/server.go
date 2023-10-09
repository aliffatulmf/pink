package server

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aliffatulmf/pink/env"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	Host   string
	Port   int
}

func (s *Server) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

func (s *Server) Run() {
	if err := s.engine.Run(s.Addr()); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}

func (s *Server) Engine() gin.IRouter {
	return s.engine
}

func NewServer() *Server {
	server := &Server{
		Host: "localhost",
		Port: 8080,
	}

	if env.EqualEnv("PINK_ENV", "production") {
		gin.SetMode(gin.ReleaseMode)
		server.engine = gin.Default()
	} else {
		gin.SetMode(gin.DebugMode)
		server.engine = gin.New()
	}

	if env.HasEnv("PINK_HOST") {
		server.Host = os.Getenv("PINK_HOST")
	}

	if env.HasEnv("PINK_PORT") {
		port, err := strconv.Atoi(os.Getenv("PINK_PORT"))
		if err != nil {
			log.Fatal("PINK_PORT must be a number")
		}
		server.Port = port
	}

	return server
}
