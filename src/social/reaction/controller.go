package reaction

import (
	"net/http"

	"github.com/NetKBs/backend-reviewapp/src/schema"
	"github.com/gin-gonic/gin"
)

func GetAllReactions(c *gin.Context) {
	reactions, err := GetAllReactionsServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch reactions"})
		return
	}
	c.JSON(http.StatusOK, reactions)
}

// GetReactionByUserIdAndContentId maneja la solicitud GET para obtener una reacción por UserId y ContentId
func GetReactionByUserIdAndContentId(c *gin.Context) {
	userId := c.DefaultQuery("user_id", "")
	contentId := c.DefaultQuery("content_id", "")

	if userId == "" || contentId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserId and ContentId are required"})
		return
	}

	reaction, err := GetReactionByUserIdAndContentIdService(userId, contentId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reaction not found"})
		return
	}

	c.JSON(http.StatusOK, reaction)
}

// CreateReaction maneja la solicitud POST para crear una nueva reacción
func CreateReaction(c *gin.Context) {
	var reaction schema.Reaction
	if err := c.ShouldBindJSON(&reaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	createdReaction, err := CreateReactionServices(reaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create reaction"})
		return
	}

	c.JSON(http.StatusCreated, createdReaction)
}
