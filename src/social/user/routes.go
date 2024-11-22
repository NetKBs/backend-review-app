package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	users := router.Group("/user")
	{
		users.GET("/:id", GetUserByIdController)       // Obtener un usuario por ID (Read)
		users.POST("/", CreateUserController)          // Crear un nuevo usuario (Create)
		users.PUT("/:id", UpdateUserController)        // Actualizar un usuario por ID (Update)
		users.DELETE("/:id", DeleteUserbyIdController) // Eliminar un usuario por ID (Delete)
	}
}
