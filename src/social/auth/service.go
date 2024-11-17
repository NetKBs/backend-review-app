package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func LoginService(username, password string) (string, error) {
	_, err := LoginRepository(username, password)
	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(24 * time.Hour) // 24h
	var claims = Claims{
		Username: username,
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
