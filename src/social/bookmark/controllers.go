package bookmark

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBookmarksController(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID"})
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	cursor := c.DefaultQuery("cursor", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
		return
	}

	cursorUint, err := strconv.ParseUint(cursor, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cursor value"})
		return
	}

	bookmarkedPlaces, nextCursor, err := GetBookmarksService(userId.(uint), limit, uint(cursorUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookmarkedPlaces, "next_cursor": nextCursor})
}

func CreateBookmarkController(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID"})
		return
	}

	var placeInput BookmarkInputDTO
	if err := c.ShouldBindJSON(&placeInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := CreateBookmarkService(userId.(uint), placeInput.PlaceID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "bookmark created"})
}

func DeleteBookmarkController(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID"})
		return
	}

	var placeInput BookmarkInputDTO
	if err := c.ShouldBindJSON(&placeInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := DeleteBookmarkService(userId.(uint), placeInput.PlaceID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "bookmark deleted"})
}
