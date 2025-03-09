package visited

import (
	"github.com/NetKBs/backend-reviewapp/src/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	visited := router.Group("/visited", middlewares.AuthMiddleware())
	{
		visited.GET("/count/:user_id", GetVisitedCountController)
		visited.GET("/visitors/:place_id", GetVisitorsCount)

		visited.GET("/", GetVisitedPlacesByUserIdController)

		visited.POST("/", CreateVisitedPlace)
		visited.DELETE("/", DeleteVisitedPlace)
	}
}
