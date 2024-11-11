package review

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetReviewByIdController(c *gin.Context) {
	id := c.Param("id")
	revId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	review, err := GetReviewByIdService(revId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"review": review})
}
