package maps

import (
	"github.com/NetKBs/backend-reviewapp/src/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	maps := r.Group("/map", middlewares.AuthMiddleware())
	{
		maps.GET("/styles", GetMapStyles)
	}
}
