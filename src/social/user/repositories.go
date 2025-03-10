package user

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func UserExistsByFieldRepository(field string, value interface{}, excludeId uint) (bool, error) {
	db := config.DB
	var count int64
	query := fmt.Sprintf("%s = ? AND id != ? AND deleted_at IS NULL", field)
	if err := db.Model(&schema.User{}).Where(query, value, excludeId).Count(&count).Error; err != nil {
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

func SearchUserByUsernameRepository(username string) ([]schema.User, error) {
	db := config.DB
	user := []schema.User{}
	var err error

	if err = db.Where("username LIKE ?", "%"+username+"%").Limit(5).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByUsernameRepository(username string) (schema.User, error) {
	db := config.DB
	user := schema.User{}
	var err error

	if err = db.Where("username = ?", username).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
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

func UpdateUserRepository(id uint, userDTO UserUpdateDTO, avatarPath string) (string, error) {
	db := config.DB
	var oldAvatar sql.NullString

	user := schema.User{}

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return "", err
	}

	if avatarPath != "" {
		if err := db.Model(&schema.User{}).Where("id = ?", id).Pluck("avatar_url", &oldAvatar).Error; err != nil {
			return "", err
		}
		user.AvatarUrl = &avatarPath
	}

	if userDTO.DisplayName != "" {
		user.DisplayName = userDTO.DisplayName
	}
	if userDTO.Username != "" {
		user.Username = userDTO.Username
	}
	if userDTO.Email != "" {
		user.Email = userDTO.Email
	}
	if userDTO.Description != "" {
		user.Description = userDTO.Description
	}

	if err := db.Save(&user).Error; err != nil {
		return "", err
	}

	// avatar path updated
	if avatarPath != "" && oldAvatar.Valid {
		return oldAvatar.String, nil
	}

	return "", nil
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

func VerifyUserRepository(id uint) error {
	db := config.DB

	if err := db.Where("id = ?", id).First(&schema.User{}).Error; err != nil {
		return errors.New("user not found")
	}

	if err := db.Model(&schema.User{}).Where("id = ?", id).Update("verified", true).Error; err != nil {
		return err
	}

	return nil
}
