package visited

import (
	"context"
	"strconv"

	"github.com/NetKBs/backend-reviewapp/src/social/place"
)

func GetVisitedPlacesByUserIdService(userId uint, limit int, cursor uint) ([]place.PlaceDetailsResponseDTO, string, error) {
	placeIds, err := GetVisitedPlacesByUserIdRepository(userId, limit, cursor)
	if err != nil {
		return []place.PlaceDetailsResponseDTO{}, "", err
	}

	visitedPlaces := []place.PlaceDetailsResponseDTO{}
	for _, placeId := range placeIds {
		placeDetail, err := place.GetPlaceDetailsByPlaceIdService(context.TODO(), int(placeId))
		if err != nil {
			return nil, "", err
		}
		visitedPlaces = append(visitedPlaces, placeDetail)
	}

	nextCursor := ""
	if len(placeIds) > 0 {
		nextCursor = strconv.FormatUint(uint64(placeIds[len(placeIds)-1]), 10)
	}

	return visitedPlaces, nextCursor, nil
}

func GetVisitedCountService(userId uint) (visitedCount uint, err error) {
	visitedCount, err = GetVisitedCountRepository(userId)
	return visitedCount, err
}

func GetVisitorsCountService(placeID uint) (uint, error) {
	count, err := GetVisitorsCountRepository(placeID)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CreateVisitedPlaceService(userID, placeID uint) error {
	return CreateVisitedPlaceRepository(userID, placeID)
}

func DeleteVisitedPlaceService(userID, placeID uint) error {
	return DeleteVisitedPlaceRepository(userID, placeID)
}
