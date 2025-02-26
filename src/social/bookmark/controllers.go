package bookmark

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func bookmarkgetController(c *gin.Context) {

	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El user_id es requerido"})
		return
	}

	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id debe ser un número válido"})
		return
	}

	bookmarkedPlaces, err := GetBookmarks(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los lugares marcados"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"bookmarked_places": bookmarkedPlaces})
}

func bookmarkpostController(c *gin.Context) {
	var input BookmarkInputDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID"})
		return
	}

	err := bookmarkpostservice(userId.(uint), input.PlaceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Lugar marcado como favorito correctamente",
	})
}

func bookmarkdeletecontroller(c *gin.Context) {
	userIDStr := c.Query("user_id")
	placeIDStr := c.Query("place_id")

	if userIDStr == "" || placeIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id y place_id son requeridos"})
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id inválido"})
		return
	}

	placeID, err := strconv.ParseUint(placeIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "place_id inválido"})
		return
	}

	err = bookmarkdeleteservice(uint(userID), uint(placeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Lugar eliminado de favoritos correctamente", //quitar userid ya que recibe por token
	})
}
