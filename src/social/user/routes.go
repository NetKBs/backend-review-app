package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.GET("/:id", GetUserByIdController)
		users.POST("/", CreateUserController)
		users.PUT("/:id", UpdateUserController)
		users.DELETE("/:id", DeleteUserbyIdController)

		users.GET("/verify/:username", VerifyUsernameController)
	}
}
