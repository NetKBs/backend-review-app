package user

import (
	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func VerifyUsernameRepository(username string) (bool, error) {
	db := config.DB
	var user schema.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return false, err
	}
	return true, nil
}

func CreateUserRepository(user schema.User) (uint, error) {
	db := config.DB
	if err := db.Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

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

func DeleteUserbyIDRepository(id uint) error {
	db := config.DB

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var user schema.User
	if err := tx.Where("id = ?", id).First(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Confirmar transacción
	return tx.Commit().Error
}

func UpdatePasswordUserRepository(id uint, password string) error {
	db := config.DB

	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var user schema.User
	if err := tx.First(&user, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	user.Password = password
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
