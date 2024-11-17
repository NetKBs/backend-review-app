package auth

import (
	"errors"
	"os"

	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
	// "golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("SECRECT_KEY"))

func LoginRepository(username, password string) (*schema.User, error) {
	db := config.DB
	var user schema.User

	if jwtKey == nil {
		jwtKey = []byte("tu_clave_secreta")
	}

	if err := db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}
