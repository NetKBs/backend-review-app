package follow

import (
	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
	"gorm.io/gorm"
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

func GetFollowersByIdRespository(id uint) (followers []schema.User, err error) {
	db := config.DB

	if err := db.Model(&schema.User{Model: gorm.Model{ID: id}}).Association("Followers").Find(&followers); err != nil {
		return followers, err
	}

	return followers, nil
}

func GetFollowingsByIdRepository(id uint) (followings []schema.User, err error) {
	db := config.DB

	if err := db.Model(&schema.User{Model: gorm.Model{ID: id}}).Association("Following").Find(&followings); err != nil {
		return followings, err
	}
	return followings, nil
}

func CreateFollowRepository(follower_id, followed_id uint) (err error) {
	db := config.DB

	if err := db.Model(&schema.User{Model: gorm.Model{ID: follower_id}}).Association("Following").Append(&schema.User{Model: gorm.Model{ID: followed_id}}); err != nil {
		return err
	}

	if err := db.Model(&schema.User{Model: gorm.Model{ID: followed_id}}).Association("Followers").Append(&schema.User{Model: gorm.Model{ID: follower_id}}); err != nil {
		return err
	}

	return nil
}

func DeleteFollowRespository(follower_id, followed_id uint) (err error) {
	db := config.DB

	if err := db.Model(&schema.User{Model: gorm.Model{ID: follower_id}}).Association("Following").Delete(&schema.User{Model: gorm.Model{ID: followed_id}}); err != nil {
		return err
	}

	if err := db.Model(&schema.User{Model: gorm.Model{ID: followed_id}}).Association("Followers").Delete(&schema.User{Model: gorm.Model{ID: follower_id}}); err != nil {
		return err
	}

	return nil
}
