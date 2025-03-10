package user

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func UserExistsByUsernameController(c *gin.Context) {
	username := c.Param("username")
	exists, err := UserExistsByFieldService("username", username, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
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

	id, err := CreateUserService(userDTO)
	if err != nil {
		status, errorMessage := handleExceptions(err)
		c.JSON(status, gin.H{"error": errorMessage})
		return
	}

	// avatar
	newUUID := uuid.New()
	newFilename := fmt.Sprintf("%s.png", strings.TrimSpace(newUUID.String()))

	path := fmt.Sprintf("images/%s", newFilename)
	if err := c.SaveUploadedFile(userDTO.AvatarImage, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := UpdateAvatarUserService(id, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})

}

func GetUserByIdController(c *gin.Context) {
	id := c.Param("id")

	user, err := GetUserByFieldService("id", id)
	if err != nil {
		status, errorMessage := handleExceptions(err)
		c.JSON(status, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func GetUserByUsernameController(c *gin.Context) {
	username := c.Param("username")

	user, err := GetUserByFieldService("username", username)
	if err != nil {
		status, errorMessage := handleExceptions(err)
		c.JSON(status, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user})
}

func SearchUserByUsernameController(c *gin.Context) {
	username := c.Query("text")

	user, err := SearchUserByUsernameService(username)
	if err != nil {
		status, errorMessage := handleExceptions(err)
		c.JSON(status, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user})
}

func UpdatePasswordUserController(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var passwordDTO UserUpdatePasswordDTO
	if err := c.ShouldBindJSON(&passwordDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := UpdatePasswordUserService(uint(userId), passwordDTO); err != nil {
		status, errorMessage := handleExceptions(err)
		c.JSON(status, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password updated successfully",
	})
}

func UpdateUserController(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var userDTO UserUpdateDTO
	if err := c.ShouldBind(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var avatarPath string = ""
	if userDTO.AvatarImage != nil {
		newUUID := uuid.New()
		newFilename := fmt.Sprintf("%s.png", strings.TrimSpace(newUUID.String()))
		avatarPath = fmt.Sprintf("images/%s", newFilename)
		if err := c.SaveUploadedFile(userDTO.AvatarImage, avatarPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if err := UpdateUserService(uint(userId), userDTO, avatarPath); err != nil {
		status, errorMessage := handleExceptions(err)
		c.JSON(status, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
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
