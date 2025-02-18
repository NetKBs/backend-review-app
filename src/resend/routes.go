package resend

import (
	"github.com/NetKBs/backend-reviewapp/src/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	resend := router.Group("/code", middlewares.AuthMiddleware())
	{

		resend.POST("/generate", generateVerificationCodeController)
		resend.POST("/verify", verifyVerificationCodeController)

	}
}
