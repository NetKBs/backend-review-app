package follow

import (
	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func GetFollowersCount(id uint) (followersCount uint, err error) {
	db := config.DB
	var user schema.User

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return followersCount, err
	}

	followersCount = uint(db.Model(&user).Association("Followers").Count())
	return followersCount, nil
}

func GetFollowingCount(id uint) (followingCount uint, err error) {
	db := config.DB
	var user schema.User

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return followingCount, err
	}

	followingCount = uint(db.Model(&user).Association("Following").Count())
	return followingCount, nil
}
