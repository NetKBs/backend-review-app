package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	users := router.Group("/user")
	{
		users.GET("/:id", GetUserByIdController)
	}
}
