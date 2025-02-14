package auth

import "github.com/dgrijalva/jwt-go"

type inputDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Claims struct {
	UserId   uint   `json:"user_id" binding:"required"`
	Username string `json:"username"`
	jwt.StandardClaims
}
