package Middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HasKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Request.URL.Query()["key"]
		if len(key) != 1 || key[0] != "a97d6edd-ff2d-4ace-9ddd-9e784ab5bf5c" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		fmt.Print(key)
		c.Next()
	}
}
