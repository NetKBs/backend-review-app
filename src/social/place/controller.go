package place

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPlaceDetailsController(c *gin.Context) {
	placeDetailsDTO := PlaceDetailsResponseDTO{}
	var err error

	lon := c.Query("lon")
	lat := c.Query("lat")
	if lat != "" && lon != "" {
		placeDetailsDTO, err = GetPlaceDetailsByCoordsService(c.Request.Context(), lon, lat)
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
	placesDTO := PlacesResponseDTO{}
	var err error

	categories := c.Query("categories")
	lon := c.Query("lon")
	lat := c.Query("lat")
	if lat != "" && lon != "" && categories != "" {
		placesDTO, err = GetPlacesByCoordsService(c.Request.Context(), []string{categories}, lon, lat)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to get place details",
			})
			return
		}
	}
	c.JSON(http.StatusOK, placesDTO)
}

func getAutocompleteResultController(c *gin.Context) {
	var autocompleteDTO AutocompleteResponseDTO

	text := c.Query("text")
	if text == "" {
		c.Status(http.StatusBadRequest)
	}

	autocompleteDTO, err := GetAutocompleteResultService(c.Request.Context(), text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get place details",
		})
		return
	}

	c.JSON(http.StatusOK, autocompleteDTO)
}
