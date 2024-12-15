package user

import (
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func handleExceptions(err error) (int, string) {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return http.StatusNotFound, "User not found"
	default:
		return http.StatusInternalServerError, err.Error()
	}
}

func VerifyUsernameController(c *gin.Context) {
	username := c.Param("username")
	exists, _ := VerifyUsernameService(username)
	c.JSON(http.StatusOK, gin.H{
		"exists": exists,
	})
}

func CreateUserController(c *gin.Context) {
	var userDTO UserCreateDTO
	if err := c.ShouldBind(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// avatar
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	newFilename := fmt.Sprintf("%s.png", strings.TrimSpace(string(newUUID)))

	path := fmt.Sprintf("images/%s", newFilename)
	if err := c.SaveUploadedFile(userDTO.AvatarImage, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newUser, err := CreateUserService(userDTO, path)
	if err != nil {
		status, errorMessage := handleExceptions(err)
		c.JSON(status, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": newUser,
	})

}

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

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
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

	data := []UserResponseDTO{newUser}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func UpdatePasswordController(c *gin.Context) {
	id := c.Param("id")
	revId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var passwordStruct struct {
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(passwordStruct); err != nil || passwordStruct.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password is required"})
		return
	}

	err = UpdatePasswordUserService(uint(revId), passwordStruct.Password)

	if err != nil {
		status, errorMessage := handleExceptions(err)
		c.JSON(status, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Contraseña cambiada exitosamente",
	})
}

func DeleteUserbyIdController(c *gin.Context) {
	id := c.Param("id")
	revId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = DeleteUserByIdService(uint(revId))

	if err != nil {
		status, errorMessage := handleExceptions(err)
		c.JSON(status, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Usuario eliminado Exitosamente",
	})
}
