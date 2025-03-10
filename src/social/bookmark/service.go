package bookmark

import (
	"context"
	"strconv"

	"github.com/NetKBs/backend-reviewapp/src/social/place"
)

func GetBookmarkCount(userId uint) (bookmarkCount uint, err error) {
	bookmarkCount, err = GetBookmarkCountRepository(userId)
	if err != nil {
		return bookmarkCount, err
	}
	return bookmarkCount, nil
}

func GetBookmarksService(userId uint, limit int, cursor uint) ([]place.PlaceDetailsResponseDTO, string, error) {
	placeIds, err := GetBookmarksRepository(userId, limit, cursor)
	if err != nil {
		return []place.PlaceDetailsResponseDTO{}, "", err
	}

	bookmarkedPlaces := []place.PlaceDetailsResponseDTO{}
	for _, placeId := range placeIds {
		placeDetail, err := place.GetPlaceDetailsByPlaceIdService(context.TODO(), int(placeId))
		if err != nil {
			return nil, "", err
		}
		bookmarkedPlaces = append(bookmarkedPlaces, placeDetail)
	}

	nextCursor := ""
	if len(placeIds) > 0 {
		nextCursor = strconv.FormatUint(uint64(placeIds[len(placeIds)-1]), 10)
	}

	return bookmarkedPlaces, nextCursor, nil
}

func CreateBookmarkService(userId uint, placeId uint) error {
	return CreateBookmarkRepository(userId, placeId)
}

func DeleteBookmarkService(userId uint, placeId uint) error {
	return DeleteBookmarkRepository(userId, placeId)
}
