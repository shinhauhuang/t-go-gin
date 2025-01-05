package routers

import (
	"t-go-gin/routers/handler"

	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/l7healthcheck", handler.Getl7HealthCheck)

	return r
}