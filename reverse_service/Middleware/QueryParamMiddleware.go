package Middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MatrixQueryRequest struct {
	Origins      string `form:"origins,omitempty" json:"origins,omitempty"`
	Origin       string `form:"origin,omitempty" json:"origin,omitempty"`
	Destinations string `form:"destinations,omitempty" json:"destinations,omitempty"`
	Destination  string `form:"destination,omitempty" json:"destination,omitempty"`
}

func HasCorrectParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		var u MatrixQueryRequest
		err := c.BindQuery(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.Next()
	}
}
