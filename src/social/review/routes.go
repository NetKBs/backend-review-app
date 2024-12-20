package review

import (
	"github.com/NetKBs/backend-reviewapp/src/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	reviews := router.Group("/reviews", middlewares.AuthMiddleware())
	{
		reviews.GET("/:id", GetReviewByIdController)
		reviews.POST("/", CreateReviewController)
		reviews.PUT("/:id", UpdateReviewController)
		reviews.DELETE("/:id", DeleteReviewController)
	}
}
