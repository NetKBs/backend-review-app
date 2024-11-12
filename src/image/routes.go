package image

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	images := r.Group("/images")
	{
		images.GET("/:name", GetImageByNameController)
	}
}
