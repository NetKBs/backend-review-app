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

func GetPlacesByCoordsService(ctx context.Context, lat string, lon string) (places geoapify.Places, err error) {
	places, err = geoapify.GetPlacesAroundCoords(lat, lon)
	if err != nil {
		return places, err
	}
	// place, err := findPlaceByMapsIdRepo(placeDetails.MapsID)
	// if err != nil {
	// 	return placeDetailsDTO, err
	// }

	// placeDetailsDTO = PlaceDetailsResponseDTO{
	// 	ID:        place.ID,
	// 	Details:   placeDetails,
	// 	CreatedAt: place.CreatedAt.String(),
	// 	UpdatedAt: place.UpdatedAt.String(),
	// }
	return places, err
}
