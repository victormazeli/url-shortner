package middlewares

import (
	"github.com/gin-gonic/gin"
	"backend/pkg/ApiResponse"
)

func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		ApiResponse.SendNotFound(c, "resource not found")
		c.Next()
	}
}
