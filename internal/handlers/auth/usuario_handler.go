package auth

import (
	"dentistry-clinic/internal/infrastructure/security"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func InitHanlderAuth(r *gin.Engine) {

	r.POST("/login", func(c *gin.Context) {
		var user struct {
			Email string `json:"email"`
			Senha string `json:"senha"`
		}
		println(&user)
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if user.Email == "admin@consultorio.com" && user.Senha == "123" {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"email": user.Email,
				"exp":   time.Now().Add(time.Hour * 24).Unix(),
			})
			tokenS, err := token.SignedString(security.SecretKey)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"token": tokenS})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha inválidos"})
		}
	})
}
