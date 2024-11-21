package comment

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCommentByIdController(c *gin.Context) {
	id := c.Param("id")
	commentId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	comment, err := GetCommentByIdService(uint(commentId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comment": comment})
}

func getCommentsByIdReviewController(c *gin.Context) {
	id := c.Param("id")
	revcommentsId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	comment, err := getCommentsByIdReviewService(uint(revcommentsId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comment": comment})
}

func CreateCommentController(c *gin.Context) {
	CommentCreateDTO := CommentCreateDTO{}
	if err := c.ShouldBind(&CommentCreateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := CreateCommentService(CommentCreateDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func UpdateCommentController(c *gin.Context) {
	CommentUpdateDTO := CommentUpdateDTO{}
	if err := c.ShouldBindJSON(&CommentUpdateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	commentId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = UpdateCommentService(uint(commentId), CommentUpdateDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
}

func DeleteCommentController(c *gin.Context) {
	id := c.Param("id")
	commentId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = DeleteCommentService(uint(commentId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
