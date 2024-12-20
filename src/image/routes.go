package image

import (
	"github.com/NetKBs/backend-reviewapp/src/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	images := r.Group("/images", middlewares.AuthMiddleware())
	{
		images.GET("/:name", GetImageByNameController)
	}
}
