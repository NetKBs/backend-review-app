package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func getUser(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"user":     "carlos",
// 		"password": "cabello",
// 	})
// }

// GetAllUsersHandler maneja la solicitud para obtener todos los usuarios
func GetAllUsersHandler(c *gin.Context) {
	users, err := GetAllUsersRepository()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// LoginHandler maneja la solicitud de login
func LoginHandler(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Validar la entrada
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Intentar hacer login y obtener el token
	token, user, err := LoginRepository(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}
