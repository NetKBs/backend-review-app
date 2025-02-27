package answer

import (
	"github.com/NetKBs/backend-reviewapp/src/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	answer := router.Group("/answers", middlewares.AuthMiddleware())
	{
		answer.GET("/comment/:id", GetAwnsersByCommentIdController)
		answer.GET("/:id", GetAnswerByIdController)
		answer.GET("/:id/likes", GetAnswerLikesController)
		answer.GET("/:id/dislikes", GetAnswerDislikesController)
		answer.POST("/", CreateAnswerController)
		answer.PUT("/:id", UpdateAnswerController)
		answer.DELETE("/:id", DeleteAnswerController)
	}
}
