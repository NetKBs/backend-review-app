package place

import (
	"net/http"

	"github.com/NetKBs/backend-reviewapp/geoapify"
	"github.com/gin-gonic/gin"
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
	data := []geoapify.PlaceDetails{placeDetails}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
