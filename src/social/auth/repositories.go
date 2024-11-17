package auth

import (
	"errors"

	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
	// "golang.org/x/crypto/bcrypt"
)

func LoginRepository(username, password string) (*schema.User, error) {
	db := config.DB
	var user schema.User

	if err := db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}
