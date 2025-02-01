package reaction

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	reactions := router.Group("/reactions")
	{
		reactions.GET("/", GetAllReactions)                       // Obtener todas las reacciones
		reactions.GET("/search", GetReactionByUserIdAndContentId) // Obtener una reacción por ID
		reactions.POST("/", CreateReaction)                       // Crear una nueva reacción
	}
}
