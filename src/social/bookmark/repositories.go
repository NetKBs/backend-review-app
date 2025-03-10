package bookmark

import (
	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
	"gorm.io/gorm"
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

func GetBookmarksRepository(userId uint, limit int, cursor uint) ([]uint, error) {
	db := config.DB
	var user schema.User
	var bookmarkedPlaces []schema.Place

	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	query := db.Model(&user).Preload("BookmarkedPlaces", func(db *gorm.DB) *gorm.DB {
		if cursor != 0 {
			return db.Order("id DESC").Limit(limit).Where("id < ?", cursor)
		}
		return db.Order("id DESC").Limit(limit)
	})

	if err := query.Find(&user).Error; err != nil {
		return nil, err
	}

	bookmarkedPlaces = user.BookmarkedPlaces
	var bookmarkedPlaceIDs []uint
	for _, place := range bookmarkedPlaces {
		bookmarkedPlaceIDs = append(bookmarkedPlaceIDs, place.ID)
	}

	return bookmarkedPlaceIDs, nil
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
