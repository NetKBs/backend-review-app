package user

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin" //Preguntar
	"gorm.io/gorm"
)

func GetUserByIdController(c *gin.Context) {
	id := c.Param("id")
	revId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	user, err := GetUserByIdService(uint(revId))
	if err != nil {
		status, errorMessage := handleExceptions(err)
		c.JSON(status, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{"User": user})
}

func handleExceptions(err error) (int, string) {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return http.StatusNotFound, "User not found"
	default:
		return http.StatusInternalServerError, err.Error()
	}
}

func UpdateUserController(c *gin.Context) {

	var newUser UserResponseDTO

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := UpdateUserService(newUser)

	if err != nil {
		status, errorMessage := handleExceptions(err)
		c.JSON(status, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{"User": newUser})
}

func CreateUserController(c *gin.Context) {
	var userDTO UserResponseDTO

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userDTO.Username == "" || userDTO.Email == "" || userDTO.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username, email, and password are required"})
		return
	}

	newUser, err := CreateUserService(userDTO)

	if err != nil {
		status, errorMessage := handleExceptions(err)
		c.JSON(status, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"User": newUser})
}
