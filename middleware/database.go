package middleware

import (
	"github.com/coopernurse/gorp"
	"github.com/gin-gonic/gin"
)

func Database(db *gorp.DbMap) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
