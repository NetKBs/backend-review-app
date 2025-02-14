package place

import (
	"context"

	"github.com/NetKBs/backend-reviewapp/geoapify"
)

func GetPlaceDetailsByMapsIdService(ctx context.Context, mapsID string) (placeDetailsDTO PlaceDetailsResponseDTO, err error) {
	placeDetails, err := geoapify.GetPlaceDetailsById(mapsID)
	if err != nil {
		return placeDetailsDTO, err
	}
	place, err := findPlaceByMapsIdRepo(mapsID)
	if err != nil {
		return placeDetailsDTO, err
	}

	placeDetailsDTO = PlaceDetailsResponseDTO{
		ID:        place.ID,
		Details:   placeDetails,
		CreatedAt: place.CreatedAt.String(),
		UpdatedAt: place.UpdatedAt.String(),
	}
	return placeDetailsDTO, err
}

func GetPlaceDetailsByCoordsService(ctx context.Context, lat string, lon string) (placeDetailsDTO PlaceDetailsResponseDTO, err error) {
	placeDetails, err := geoapify.GetPlaceDetailsByCoord(lat, lon)
	if err != nil {
		return placeDetailsDTO, err
	}
	place, err := findPlaceByMapsIdRepo(placeDetails.MapsID)
	if err != nil {
		return placeDetailsDTO, err
	}

	placeDetailsDTO = PlaceDetailsResponseDTO{
		ID:        place.ID,
		Details:   placeDetails,
		CreatedAt: place.CreatedAt.String(),
		UpdatedAt: place.UpdatedAt.String(),
	}
	return placeDetailsDTO, err
}

// categories: slice of strings
func GetPlacesByCoordsService(ctx context.Context, categories []string, lat, lon string) (places geoapify.Places, err error) {
	var catString string
	for i, v := range categories {
		if i == 0 {
			catString = v
		} else {
			catString = catString + "," + v
		}
	}
	places, err = geoapify.GetPlacesAroundCoords(catString, lat, lon)
	if err != nil {
		return places, err
	}
	return places, err
}

func GetAutocompleteResultService(ctx context.Context, text string) (autocompleteResult AutocompleteResponseDTO, err error) {
	geocodings, err := geoapify.GetAutocompleteResponse(text)
	autocompleteResult.Query = text
	autocompleteResult.Result = geocodings

	return autocompleteResult, err
}
