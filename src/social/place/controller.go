package place

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getPlaceController(c *gin.Context) {
	mapsID := c.Query("maps_id")
	if mapsID == "" {
		c.Status(http.StatusBadRequest)
	}

	placeDetails, err := GetPlaceDetailsByMapsId(c.Request.Context(), mapsID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get place details",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": placeDetails,
	})
}
