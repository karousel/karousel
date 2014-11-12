package middleware

import (
	r "github.com/dancannon/gorethink"
	"github.com/gin-gonic/gin"
)

func Database(session *r.Session) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("session", session)
		c.Next()
	}
}
