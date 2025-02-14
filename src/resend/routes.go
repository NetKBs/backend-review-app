package resend

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	resend := router.Group("/code")
	{

		resend.POST("/generate", generateVerificationCodeController)
		resend.POST("/verify", verifyVerificationCodeController)

	}
}
