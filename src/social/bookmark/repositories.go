package bookmark

import (
	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func GetBookmarkCountRepository(userId uint) (bookmarkCount uint, err error) {
	db := config.DB
	var user schema.User

	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		return bookmarkCount, err
	}

	bookmarkCount = uint(db.Model(&user).Association("BookmarkedPlaces").Count())
	return bookmarkCount, nil
}

func GetBookmarksRepository(userId uint) ([]uint, error) {
	db := config.DB
	var user schema.User
	var bookmarkedPlaces []schema.Place

	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	if err := db.Model(&user).Association("BookmarkedPlaces").Find(&bookmarkedPlaces); err != nil {
		return nil, err
	}

	var bookmarkedPlacesIDs []uint
	for _, place := range bookmarkedPlaces {
		bookmarkedPlacesIDs = append(bookmarkedPlacesIDs, place.ID)
	}

	return bookmarkedPlacesIDs, nil
}

func CreateBookmarkRepository(userId uint, placeId uint) error {
	db := config.DB

	var place schema.Place
	if err := db.First(&place, placeId).Error; err != nil {
		return err
	}

	var user schema.User
	if err := db.First(&user, userId).Error; err != nil {
		return err
	}

	if err := db.Model(&user).Association("BookmarkedPlaces").Append(&place); err != nil {
		return err
	}

	return nil
}

func DeleteBookmarkRepository(userId uint, placeId uint) error {
	db := config.DB

	var place schema.Place
	if err := db.First(&place, placeId).Error; err != nil {
		return err
	}

	var user schema.User
	if err := db.First(&user, userId).Error; err != nil {
		return err
	}

	if err := db.Model(&user).Association("BookmarkedPlaces").Delete(&place); err != nil {
		return err
	}

	return nil
}
