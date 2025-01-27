package user

import (
	"github.com/NetKBs/backend-reviewapp/src/social/follow"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.GET("/verify/:username", UserExistsByUsernameController)
		users.GET("/:id", GetUserByIdController)

		users.POST("/", CreateUserController)
		users.DELETE("/:id", DeleteUserbyIdController)

		users.PUT("/displayname/:id", UpdateUserDisplayNameController)
		users.PUT("/username/:id", UpdateUsernameUserController)
		users.PUT("/password/:id", UpdatePasswordUserController)
		users.PUT("/avatar/:id", UpdateAvatarUserController)
		users.PUT("/email/:id", UpdateEmailUserController)

		users.GET("/:id/followers", follow.GetFollowersByIdController)
		users.GET("/:id/followings", follow.GetFollowingsByIdController)

	}

	follows := users.Group("/follow")
	{
		follows.POST("/:follower_id/:followed_id", follow.CreateFollowController)
		follows.DELETE("/:follower_id/:followed_id", follow.DeleteFollowController)
	}
}
