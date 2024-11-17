package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("/", GetAllUsersHandler)
		users.POST("/login", LoginHandler)
		// Ruta protegida. EJEMPLO DE COMO DAR PROTECCION DE RUTA (BORRAR DESPUES)
		users.GET("/protected", AuthMiddleware(), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Acceso concedido"})
		})
	}
}
