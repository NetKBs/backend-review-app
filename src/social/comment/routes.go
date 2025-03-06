package comment

import (
	"github.com/NetKBs/backend-reviewapp/src/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	comments := router.Group("/comments", middlewares.AuthMiddleware())
	{
		comments.GET("/:id", GetCommentByIdController)
		comments.GET("/:id/likes", GetCommentLikesController)
		comments.GET("/:id/dislikes", GetCommentDislikesController)
		comments.GET("/review/:id", GetCommentsByIdReviewController)
		comments.POST("/", CreateCommentController)
		comments.PUT("/:id", UpdateCommentController)
		comments.DELETE("/:id", middlewares.AdminRequired(), DeleteCommentController)
	}
}
