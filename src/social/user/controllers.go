package user

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
		var status int
		var errorMessage string

		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
			errorMessage = "User not found"
		} else {
			status = http.StatusInternalServerError
			errorMessage = fmt.Sprintf("%s: %s", "Internal Server Error", err.Error())
		}

		c.JSON(status, gin.H{"error": errorMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{"User": user})
}