package answer

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAwnsersByCommentIdController(c *gin.Context) {
	id := c.Param("id")
	ansCommentsId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	answer, err := GetAnswersByCommentIdService(uint(ansCommentsId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Answer": answer})
}

func GetAnswerByIdController(c *gin.Context) {
	id := c.Param("id")
	answerId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	answer, err := GetAnswerByIdService(uint(answerId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Answer": answer})
}

func CreateAnswerController(c *gin.Context) {
	AnswerCreateDTO := AnswerCreateDTO{}
	if err := c.ShouldBind(&AnswerCreateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := CreateAnswerService(AnswerCreateDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func UpdateAnswerController(c *gin.Context) {
	AnswerUpdateDTO := AnswerUpdateDTO{}
	if err := c.ShouldBindJSON(&AnswerUpdateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	answerId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = UpdateAnswerService(uint(answerId), AnswerUpdateDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Answer updated successfully"})
}

func DeleteAnswerController(c *gin.Context) {
	id := c.Param("id")
	answerId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = DeleteAnswerService(uint(answerId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Answer deleted successfully"})
}
