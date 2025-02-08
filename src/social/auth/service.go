package auth

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("SECRECT_KEY"))

func LoginService(username, password string) (string, error) {
	user, err := LoginRepository(username, password)
	if err != nil {
		return "", err
	}

	if jwtKey == nil {
		jwtKey = []byte("tu_clave_secreta")
	}

	expirationTime := time.Now().Add(24 * time.Hour) // 24h
	var claims = Claims{
		UserId:   user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
