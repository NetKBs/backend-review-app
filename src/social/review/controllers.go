package review

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/NetKBs/backend-reviewapp/src/image"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetReviewByIdController(c *gin.Context) {
	id := c.Param("id")
	revId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	review, err := GetReviewByIdService(uint(revId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": review})
}

func GetReviewLikesController(c *gin.Context) {
	id := c.Param("id")
	revId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit := 10
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
			return
		}
	}

	cursor := c.DefaultQuery("cursor", "")

	likes, nextCursor, err := GetReviewLikesByIdService(revId, limit, cursor)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": likes, "next_cursor": nextCursor})
}

func GetReviewDislikesController(c *gin.Context) {
	id := c.Param("id")
	revId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit := 10
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
			return
		}
	}

	cursor := c.DefaultQuery("cursor", "")

	dislikes, nextCursor, err := GetReviewDislikesByIdService(revId, limit, cursor)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dislikes, "next_cursor": nextCursor})
}

func GetReviewsByPlaceIdController(c *gin.Context) {
	placeIdStr := c.Param("id")
	placeId, err := strconv.ParseUint(placeIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	pageStr := c.DefaultQuery("page", "1")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page value"})
		return
	}

	reviews, pagination, err := GetReviewsByPlaceIdService(uint(placeId), limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reviews, "pagination": pagination})
}

func GetReviewsByUserIdController(c *gin.Context) {
	userIdStr := c.Param("id")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	pageStr := c.DefaultQuery("page", "1")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page value"})
		return
	}

	reviews, pagination, err := GetReviewsByUserIdService(uint(userId), limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reviews, "pagination": pagination})
}

func CreateReviewController(c *gin.Context) {
	ReviewCreateDTO := ReviewCreateDTO{}
	if err := c.ShouldBind(&ReviewCreateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID"})
		return
	}

	id, err := CreateReviewService(ReviewCreateDTO, userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var Imagepaths []string
	for _, file := range ReviewCreateDTO.Images {

		newUUID := uuid.New()
		newFilename := fmt.Sprintf("%s.png", strings.TrimSpace(newUUID.String()))

		path := fmt.Sprintf("images/%s", newFilename)
		if err := c.SaveUploadedFile(file, path); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		Imagepaths = append(Imagepaths, path)
	}

	if err := image.RegisterReviewImagesService(id, Imagepaths); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})

}

func UpdateReviewController(c *gin.Context) {
	ReviewUpdateDTO := ReviewUpdateDTO{}
	if err := c.ShouldBindJSON(&ReviewUpdateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	revId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = UpdateReviewService(uint(revId), ReviewUpdateDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review updated successfully"})
}

func DeleteReviewController(c *gin.Context) {
	id := c.Param("id")
	revId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = DeleteReviewService(uint(revId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = image.DeleteReviewImagesService(uint(revId)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}
