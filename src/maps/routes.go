package maps

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	maps := r.Group("/map")
	{
		maps.GET("/styles", GetMapStyles)
	}
}
