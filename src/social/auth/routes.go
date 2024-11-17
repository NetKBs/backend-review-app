package auth

import (
	"net/http"

	"github.com/NetKBs/backend-reviewapp/src/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", LoginController)

		// Protected Route. Example of how to protect a route (Delete later)
		auth.GET("/test", middlewares.AuthMiddleware(), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Acceso concedido"})
		})
	}
}
