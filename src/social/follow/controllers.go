package follow

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFollowersByIdController(c *gin.Context) {
	id := c.Param("id")
	parsedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	followers, err := GetFollowersByIdService(uint(parsedId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": followers})
}

func GetFollowingsByIdController(c *gin.Context) {
	id := c.Param("id")
	parsedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	follwings, err := GetFollowingsByIdService(uint(parsedId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": follwings})
}

func CreateFollowController(c *gin.Context) {
	follower_id, err := strconv.ParseUint(c.Param("follower_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid follower_id"})
		return
	}

	followed_id, err := strconv.ParseUint(c.Param("followed_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid followed_id"})
		return
	}

	if err := CreateFollowService(uint(follower_id), uint(followed_id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Follow created successfully"})
}

func DeleteFollowController(c *gin.Context) {
	follower_id, err := strconv.ParseUint(c.Param("follower_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid follower_id"})
		return
	}

	followed_id, err := strconv.ParseUint(c.Param("followed_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid followed_id"})
		return
	}

	if err := DeleteFollowService(uint(follower_id), uint(followed_id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Follow deleted successfully"})
}
