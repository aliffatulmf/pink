package routers

import (
	"github.com/aliffatulmf/pink/handlers"
	"github.com/gin-gonic/gin"
)

func NewRouter(router gin.IRouter) {
	v1 := router.Group("/api/v1")
	v1.GET("/check", handlers.VerifyDomain)
	v1.GET("/list", handlers.ListDomains)
}
