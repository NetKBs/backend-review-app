package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginController(c *gin.Context) {
	inputDTO := inputDTO{}
	if err := c.ShouldBindJSON(&inputDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := LoginService(inputDTO.Username, inputDTO.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
