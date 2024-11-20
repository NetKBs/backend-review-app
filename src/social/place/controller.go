package place

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPlaceController(c *gin.Context) {
	mapsID := c.Query("maps_id")
	if mapsID == "" {
		c.Status(http.StatusBadRequest)
	}

	placeDetailsDTO, err := GetPlaceDetailsByMapsId(c.Request.Context(), mapsID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get place details",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": placeDetailsDTO,
	})
}
