package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte(os.Getenv("SECRECT_KEY"))

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado"})
			c.Abort()
			return
		}

		if jwtKey == nil {
			jwtKey = []byte("tu_clave_secreta")
		}

		parts := strings.Split(tokenString, "Bearer ")
		if len(parts) > 1 {
			tokenString = parts[1]
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := uint(claims["user_id"].(float64))
			username := claims["username"].(string)
			c.Set("user_id", userID)
			c.Set("username", username)

			var verified bool
			if err := config.DB.Model(&schema.User{}).Where("id = ?", userID).Pluck("verified", &verified).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			if !verified {
				c.JSON(http.StatusForbidden, gin.H{"error": "user not verified"})
				c.Abort()
				return
			}

		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid token claims"})
			c.Abort()
			return
		}

		c.Next()
	}
}
