package middleware

import (
	"dentistry-clinic/internal/infrastructure/security"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		autho := c.GetHeader("Authorization")
		if autho == "" {
			c.AbortWithStatus(401)
			return
		}

		tS := strings.TrimPrefix(autho, "Bearer ")
		claims, err := security.ParseJWT(tS)
		if err != nil {
			c.AbortWithStatus(401)
			return
		}
		c.Set("user_id", claims.UserId)
		c.Set("role", claims.Role)
		c.Next()
	}
}
