package bookmark

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	bookmark := router.Group("/bookmark")
	{

		bookmark.GET("/", bookmarkgetController)
		bookmark.POST("/", bookmarkpostController)
		bookmark.DELETE("/", bookmarkdeletecontroller)

	}
}
