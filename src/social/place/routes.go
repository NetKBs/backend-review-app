package place

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	group := router.Group("/places")

	group.GET("/", getPlaceController)
}
