package user

import (
	"github.com/NetKBs/backend-reviewapp/src/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.GET("/verify/:username", UserExistsByUsernameController) // checks
		users.GET("/:id", middlewares.AuthMiddleware(), GetUserByIdController)

		users.POST("/", CreateUserController) // register
		users.DELETE("/:id", middlewares.AuthMiddleware(), DeleteUserbyIdController)

		users.PUT("/displayname/:id", middlewares.AuthMiddleware(), UpdateUserDisplayNameController)
		users.PUT("/username/:id", middlewares.AuthMiddleware(), UpdateUsernameUserController)
		users.PUT("/password/:id", middlewares.AuthMiddleware(), UpdatePasswordUserController)
		users.PUT("/avatar/:id", middlewares.AuthMiddleware(), UpdateAvatarUserController)
		users.PUT("/email/:id", middlewares.AuthMiddleware(), UpdateEmailUserController)
	}
}
