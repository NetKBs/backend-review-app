package review

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	reviews := router.Group("/review")
	{
		reviews.GET("/:id", GetReviewByIdController)
		reviews.POST("/", CreateReviewController)
		reviews.PUT("/:id", UpdateReviewController)
		reviews.DELETE("/:id", DeleteReviewController)
	}
}
