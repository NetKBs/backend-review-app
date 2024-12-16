package user

import (
	"database/sql"

	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func UserExistsByUsernameRepository(username string) (bool, error) {
	db := config.DB
	var count int64
	if err := db.Model(&schema.User{}).Where("username = ? AND deleted_at IS NULL", username).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func UserExistsByEmailRepository(email string) (bool, error) {
	db := config.DB
	var count int64
	if err := db.Model(&schema.User{}).Where("email = ? AND deleted_at IS NULL", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func UserExistsByIdRepository(id uint) (bool, error) {
	db := config.DB
	var count int64
	if err := db.Model(&schema.User{}).Where("id = ? AND deleted_at IS NULL", id).Count(&count).Error; err != nil {
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
	var oldAvatar sql.NullString

	if err := db.Model(&schema.User{}).Where("id = ?", id).Pluck("avatar_url", &oldAvatar).Error; err != nil {
		return "", err
	}
	if err := db.Model(&schema.User{}).Where("id = ?", id).Update("avatar_url", avatarPath).Error; err != nil {
		return "", err
	}

	if oldAvatar.Valid {
		return oldAvatar.String, nil
	}
	return "", nil
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

func UpdateUserRepository(id uint, userDTO UserUpdateDTO) error {
	db := config.DB

	if err := db.Where("id = ?", id).First(&schema.User{}).Error; err != nil {
		return err
	}
	if err := db.Model(&schema.User{}).Where("id = ?", id).Update("display_name", userDTO.DisplayName).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUserbyIDRepository(id uint) (string, error) {
	db := config.DB
	var user schema.User

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return "", err
	}

	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return "", err
	}

	return *user.AvatarUrl, nil
}
