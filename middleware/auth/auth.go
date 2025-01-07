package auth

import (
	"github.com/gin-gonic/gin"
)

func Validation() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implement Validation

		c.Next()
	}
}
