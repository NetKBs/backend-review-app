package resend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func verifyController(c *gin.Context) {
	var verifyInput verifyinputDTO
	if err := c.ShouldBindJSON(&verifyInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	if verifyInput.Code == "false" {
		if err := sendVerificationEmailService(verifyInput.Email); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo enviar el correo"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Código de verificación enviado",
			"email":   verifyInput.Email,
		})
		return
	}

	if verifyInput.Code == "true" {
		if verifyInput.Code == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El código de verificación es obligatorio"})
			return
		}

		if validateverificationCode(verifyInput.Email, verifyInput.Code) {
			c.JSON(http.StatusOK, gin.H{"message": "Código de verificación válido"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Código de verificación inválido o expirado"})
		}
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "Valor de codigo inválido"})
}
