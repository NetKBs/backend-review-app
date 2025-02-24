package bookmark

import (
	"errors"
)

func GetBookmarks(userId uint) ([]uint, error) {
	bookmarkedPlaces, err := GetBookmarksRepository(userId)
	if err != nil {
		return nil, err
	}

	var placeIDs []uint
	for _, place := range bookmarkedPlaces {
		placeIDs = append(placeIDs, place.ID)
	}

	return placeIDs, nil
}

func bookmarkpostservice(userId uint, placeID uint) error {
	err := bookmarkpostrepository(userId, placeID)
	if err != nil {
		return errors.New("error al marcar el lugar como favorito")
	}
	return nil
}

func bookmarkdeleteservice(userID uint, placeID uint) error {
	err := bookmarkdeleterepository(userID, placeID)
	if err != nil {
		return err
	}
	return nil
}
