package inference

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	inferences := router.Group("/inferences") // need to be authenticated
	{
		inferences.POST("/", InferenceController)
	}
}
