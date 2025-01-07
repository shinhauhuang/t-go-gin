package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Getl7HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
