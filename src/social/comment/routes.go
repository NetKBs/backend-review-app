package comment

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	comments := router.Group("/comments")
	{
		comments.GET("/:id", GetCommentByIdController)
		comments.GET("/review/:id", getCommentsByIdReviewController)
		comments.POST("/", CreateCommentController)
		comments.PUT("/:id", UpdateCommentController)
		comments.DELETE("/:id", DeleteCommentController)
	}
}
