package review

import (
	"github.com/NetKBs/backend-reviewapp/src/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	reviews := router.Group("/reviews", middlewares.AuthMiddleware())
	{
		reviews.GET("/:id", GetReviewByIdController)
		reviews.GET("/:id/likes", GetReviewLikesController)
		reviews.GET("/:id/dislikes", GetReviewDislikesController)
		reviews.GET("/place/:id", GetReviewsByPlaceIdController)
		reviews.GET("/user/:id", GetReviewsByUserIdController)
		reviews.POST("/", CreateReviewController)
		reviews.PUT("/:id", UpdateReviewController)
		reviews.DELETE("/:id", middlewares.AdminRequired(), DeleteReviewController)
	}
}
