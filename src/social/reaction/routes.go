package reaction

import (
	"github.com/NetKBs/backend-reviewapp/src/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	reactions := router.Group("/reactions", middlewares.AuthMiddleware())
	{
		reactions.GET("/", GetAllReactions)                       // Obtener todas las reacciones
		reactions.GET("/search", GetReactionByUserIdAndContentId) // Obtener una reacción por ID
		reactions.POST("/", CreateReaction)                       // Crear una nueva reacción
		reactions.DELETE("/delete", DeleteReactionController)     // borrar una nueva reacción
	}
}
