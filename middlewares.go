package main

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func RequestValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		u, _ := uuid.NewV4()
		c.Writer.Header().Set("X-Request-Id", u.String())
		c.Next()
	}
}

