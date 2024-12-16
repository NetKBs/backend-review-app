package user

import (
	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func UserExistsByUsernameRepository(username string) (bool, error) {
	db := config.DB
	var count int64
	if err := db.Model(&schema.User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func UserExistsByIdRepository(id uint) (bool, error) {
	db := config.DB
	var count int64
	if err := db.Model(&schema.User{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func GetPasswordUserRepository(id uint) (string, error) {
	db := config.DB
	var dbPassword string

	if err := db.Where("id = ?", id).First(&schema.User{}).Error; err != nil {
		return "", err
	}

	if err := db.Model(&schema.User{}).Where("id = ?", id).Pluck("password", &dbPassword).Error; err != nil {
		return "", err
	}
	return dbPassword, nil
}

func UpdatePasswordUserRepository(id uint, newPassword string) error {
	db := config.DB
	if err := db.Model(&schema.User{}).Where("id = ?", id).Update("password", newPassword).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAvatarUserRepository(id uint, avatarPath string) (string, error) {
	db := config.DB
	var oldAvatar string

	if err := db.Model(&schema.User{}).Where("id = ?", id).Pluck("avatar_url", &oldAvatar).Error; err != nil {
		return "", err
	}
	if err := db.Model(&schema.User{}).Where("id = ?", id).Update("avatar_url", avatarPath).Error; err != nil {
		return "", err
	}

	return oldAvatar, nil
}

func UpdateEmailUserRepository(id uint, email string) error {
	db := config.DB

	if err := db.Where("id = ?", id).First(&schema.User{}).Error; err != nil {
		return err
	}
	if err := db.Model(&schema.User{}).Where("id = ?", id).Update("email", email).Error; err != nil {
		return err
	}
	return nil
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
