package user

import (
	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func GetUserByIdRepository(id uint) (user schema.User, err error) {
	db := config.DB

	if err = db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func UpdateUserRepository(user *schema.User) error {
	db := config.DB
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Save(user).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	return err
}
