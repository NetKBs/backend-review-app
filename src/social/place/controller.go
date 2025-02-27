package place

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getPlaceDetailsByPlaceIdController(c *gin.Context) {
	var err error
	placeIdStr := c.Param("place_id")
	placeId, err := strconv.Atoi(placeIdStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid place_id",
		})
		return
	}

	placeDetailsDTO, err := GetPlaceDetailsByPlaceIdService(c.Request.Context(), placeId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if placeDetailsDTO == (PlaceDetailsResponseDTO{}) {
		c.Status(http.StatusBadRequest)
	}
	data := []PlaceDetailsResponseDTO{placeDetailsDTO}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}

func getPlaceDetailsController(c *gin.Context) {
	placeDetailsDTO := PlaceDetailsResponseDTO{}
	var err error

	mapsId := c.Query("mapsId")
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

	} else if mapsId != "" {
		placeDetailsDTO, err = GetPlaceDetailsByMapsIdService(c.Request.Context(), mapsId)
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
