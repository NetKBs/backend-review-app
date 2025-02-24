package bookmark

import (
	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func GetBookmarksRepository(userId uint) ([]schema.Place, error) {
	db := config.DB
	var user schema.User
	var bookmarkedPlaces []schema.Place

	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	if err := db.Model(&user).Association("BookmarkedPlaces").Find(&bookmarkedPlaces); err != nil {
		return nil, err
	}

	return bookmarkedPlaces, nil
}

func bookmarkpostrepository(userid uint, placeID uint) error {
	db := config.DB
	var user schema.User
	var place schema.Place

	if err := db.Where("id = ?", userid).First(&user).Error; err != nil {
		return err
	}

	if err := db.Where("id = ?", placeID).First(&place).Error; err != nil {
		return err
	}

	if err := db.Model(&user).Association("BookmarkedPlaces").Append(&place); err != nil {
		return err
	}

	return nil
}

func bookmarkdeleterepository(userID uint, placeID uint) error {
	db := config.DB
	var user schema.User
	var place schema.Place

	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		return err
	}

	if err := db.Where("id = ?", placeID).First(&place).Error; err != nil {
		return err
	}

	if err := db.Model(&user).Association("BookmarkedPlaces").Delete(&place); err != nil {
		return err
	}

	return nil
}
