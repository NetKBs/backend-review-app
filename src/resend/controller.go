package resend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func generateVerificationCodeController(c *gin.Context) {
	var input generateInputDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := sendVerificationEmailService(input.UserId, input.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Code sent",
		"email":   input.Email,
	})
}

func verifyVerificationCodeController(c *gin.Context) {
	var verifyInput verifyinputDTO
	if err := c.ShouldBindJSON(&verifyInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := verifyVerificationCodeService(verifyInput.UserId, verifyInput.Code); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Code verified",
	})
}
