package reaction

import (
	"net/http"
	"strconv"

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

	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID"})
		return
	}

	contentId := c.DefaultQuery("content_id", "")
	contentType := c.DefaultQuery("content_type", "")

	if userId == "" || contentId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserId and ContentId are required"})
		return
	}

	reaction, err := GetReactionByUserIdAndContentIdService(userId.(uint), contentId, contentType)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reaction not found"})
		return
	}

	c.JSON(http.StatusOK, reaction)
}

// CreateReaction maneja la solicitud POST para crear una nueva reacción
func CreateReaction(c *gin.Context) {
	var reaction ReactionCreateDTO
	if err := c.ShouldBindJSON(&reaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID"})
		return
	}

	createdReaction, err := CreateReactionServices(reaction, userId.(uint))
	if err != nil {
		// Aquí se devuelve el mensaje de error específico en lugar de un mensaje genérico
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdReaction)
}

// DeleteReactionController maneja la solicitud HTTP para borrar una reacción
func DeleteReactionController(c *gin.Context) {

	contentId := c.Query("contentId")

	contentType := c.Query("contentType") // Obtener ContentType del parámetro de consulta

	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID invalido"})
		return
	}

	// Convertir contentId a uint
	contentIdUint, err := strconv.ParseUint(contentId, 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ContentId invalido"})
		return
	}

	// Verificar que el ContentType sea válido
	if contentType != "review" && contentType != "comment" && contentType != "answer" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ContentType invalido"})
		return
	}

	// Llama al servicio para eliminar la reacción
	err = DeleteReactionServices(userId.(uint), uint(contentIdUint), contentType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete reaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reacción borrada exitosamente"})
}
