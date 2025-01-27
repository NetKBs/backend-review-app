package resend

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	verify := router.Group("/Register")
	{

		verify.POST("/Register", verifyController)

	}
}
