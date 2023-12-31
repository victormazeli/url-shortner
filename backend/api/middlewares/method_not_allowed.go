package middlewares

import (
	"github.com/gin-gonic/gin"
	"backend/pkg/ApiResponse"
)

func MethodNotAllowed() gin.HandlerFunc {
	return func(c *gin.Context) {
		ApiResponse.SendMethodNotAllowedError(c, "method not allowed")
		c.Next()
	}
}
