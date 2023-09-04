package middlewares

import (
	"fmt"
	"net/http"

	"Society-Synergy/base/token"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		fmt.Print(err)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, Token Expired", "message": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
