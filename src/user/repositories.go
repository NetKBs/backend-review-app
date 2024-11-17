package user

import (
	"errors"
	"time"

	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
	"github.com/dgrijalva/jwt-go"
	// "golang.org/x/crypto/bcrypt"
)

// GetAllUsersRepository obtiene todos los usuarios de la base de datos
func GetAllUsersRepository() ([]schema.User, error) {
	var users []schema.User
	db := config.DB

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

var jwtKey = []byte("tu_clave_secreta") // Cambia esto por una clave segura

// Claims define la estructura de los datos que se incluirán en el token
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// LoginRepository verifica las credenciales del usuario y genera un token
func LoginRepository(username, password string) (string, *schema.User, error) {
	var user schema.User
	db := config.DB

	// Buscar el usuario por nombre de usuario
	if err := db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return "", nil, errors.New("credenciales inválidas")
	}

	// Crear el token
	expirationTime := time.Now().Add(24 * time.Hour) // El token expira en 24 horas
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", nil, err
	}

	return tokenString, &user, nil
}
