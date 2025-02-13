package visited

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	visited := router.Group("/visited", middlewares.AuthMiddleware())
	{
		visited.GET("/count/:user_id", GetVisitedCountController)
		visited.GET("/visitors/:place_id", GetVisitorsCount)
		visited.POST("/:user_id/:place_id", CreateVisitedPlace)
		visited.DELETE("/:user_id/:place_id", DeleteVisitedPlace)
	}
}
