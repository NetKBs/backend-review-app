package inference

import (
	"github.com/NetKBs/backend-reviewapp/src/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	inferences := router.Group("/inferences", middlewares.AuthMiddleware())
	{
		inferences.POST("/", InferenceController)
	}
}
