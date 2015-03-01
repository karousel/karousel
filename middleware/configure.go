package middleware

import (
	"github.com/karousel/karousel/models"

	"github.com/gin-gonic/gin"
)

func Configure(config models.Configuration) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", config)
		c.Next()
	}
}
