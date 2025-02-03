package place

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPlaceDetailsController(c *gin.Context) {
	placeDetailsDTO := PlaceDetailsResponseDTO{}
	var err error

	lat := c.Query("lat")
	lon := c.Query("lon")
	if lat != "" && lon != "" {
		placeDetailsDTO, err = GetPlaceDetailsByCoordsService(c.Request.Context(), lat, lon)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to get place details",
			})
			return
		}
	}

	if placeDetailsDTO == (PlaceDetailsResponseDTO{}) {
		c.Status(http.StatusBadRequest)
	}
	data := []PlaceDetailsResponseDTO{placeDetailsDTO}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func getPlacesController(c *gin.Context) {
	placeDetailsDTO := PlaceDetailsResponseDTO{}
	// var err error
	//
	// lat := c.Query("lat")
	// lon := c.Query("lon")
	// if lat != "" && lon != "" {
	// 	placeDetailsDTO, err = GetPlaceDetailsByCoordsService(c.Request.Context(), lat, lon)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{
	// 			"message": "failed to get place details",
	// 		})
	// 		return
	// 	}
	// }
	//
	// if placeDetailsDTO == (PlaceDetailsResponseDTO{}) {
	// 	c.Status(http.StatusBadRequest)
	// }
	data := []PlaceDetailsResponseDTO{placeDetailsDTO}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
