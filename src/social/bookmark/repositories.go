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
