package feed

import (
	"github.com/NetKBs/backend-reviewapp/src/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	feed := router.Group("/feed", middlewares.AuthMiddleware())
	{
		feed.GET("/:user_id", getFeedController)
	}
}
