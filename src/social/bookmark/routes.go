package bookmark

import (
	"github.com/NetKBs/backend-reviewapp/src/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	bookmark := router.Group("/bookmarks", middlewares.AuthMiddleware())
	{
		bookmark.GET("/", GetBookmarksController)
		bookmark.POST("/", CreateBookmarkController)
		bookmark.DELETE("/", DeleteBookmarkController)

	}
}
