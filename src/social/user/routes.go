package user

import (
	"github.com/NetKBs/backend-reviewapp/src/middlewares"
	"github.com/NetKBs/backend-reviewapp/src/social/follow"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.GET("/verify/:username", UserExistsByUsernameController)
		users.GET("/id/:id", middlewares.AuthMiddleware(), GetUserByIdController)
		users.GET("/username/:username", middlewares.AuthMiddleware(), GetUserByUsernameController)

		users.POST("/", CreateUserController)
		users.DELETE("/:id", middlewares.AuthMiddleware(), DeleteUserbyIdController)

		users.PUT("/:id", middlewares.AuthMiddleware(), UpdateUserController)                  // General user update
		users.PUT("/password/:id", middlewares.AuthMiddleware(), UpdatePasswordUserController) // Password update

		users.GET("/:id/followers", middlewares.AuthMiddleware(), follow.GetFollowersByIdController)
		users.GET("/:id/followings", middlewares.AuthMiddleware(), follow.GetFollowingsByIdController)
	}

	follows := users.Group("/follow")
	{
		follows.POST("/:follower_id/:followed_id", middlewares.AuthMiddleware(), follow.CreateFollowController)
		follows.DELETE("/:follower_id/:followed_id", middlewares.AuthMiddleware(), follow.DeleteFollowController)
	}
}
