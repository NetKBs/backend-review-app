package bookmark

import (
	"context"

	"github.com/NetKBs/backend-reviewapp/src/social/place"
)

func GetBookmarkCount(userId uint) (bookmarkCount uint, err error) {
	bookmarkCount, err = GetBookmarkCountRepository(userId)
	if err != nil {
		return bookmarkCount, err
	}
	return bookmarkCount, nil
}

func GetBookmarksService(userId uint) ([]place.PlaceDetailsResponseDTO, error) {
	placeIds, err := GetBookmarksRepository(userId)
	if err != nil {
		return []place.PlaceDetailsResponseDTO{}, err
	}

	bookmarkedPlaces := []place.PlaceDetailsResponseDTO{}
	for _, placeId := range placeIds {
		placeDetail, err := place.GetPlaceDetailsByPlaceIdService(context.TODO(), int(placeId))
		if err != nil {
			return nil, err
		}
		bookmarkedPlaces = append(bookmarkedPlaces, placeDetail)
	}

	return bookmarkedPlaces, nil
}

func CreateBookmarkService(userId uint, placeId uint) error {
	return CreateBookmarkRepository(userId, placeId)
}

func DeleteBookmarkService(userId uint, placeId uint) error {
	return DeleteBookmarkRepository(userId, placeId)
}
