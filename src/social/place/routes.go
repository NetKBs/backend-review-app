package place

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	group := router.Group("/place")

	group.GET("/", getPlaceController)
}
