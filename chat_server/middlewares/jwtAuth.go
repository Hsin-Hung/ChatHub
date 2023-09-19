package middlewares

import (
	"chat-server/token"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Validate JWT token middleware
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := token.ExtractTokenID(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}
		username_sent := c.Query("username")
		fmt.Println(username_sent)
		fmt.Println(username)
		if username != username_sent {
			c.String(http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}

}
