package review

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	reviews := r.Group("/review")
	{
		reviews.GET("/:id", GetReviewByIdController)
	}
}
