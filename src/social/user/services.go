package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/NetKBs/backend-reviewapp/src/schema"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetUserByIdService(id uint) (userDTO UserResponseDTO, err error) {

	user, err := GetUserByIdRepository(id)
	if err != nil {
		return userDTO, err
	}

	userDTO = UserResponseDTO{
		ID:          user.ID,
		Username:    user.Username,
		AvatarUrl:   getStringPointer(user.AvatarUrl),
		DisplayName: user.DisplayName,
		Email:       user.Email,
		Password:    user.Password,
		CreatedAt:   user.CreatedAt.String(),
		UpdatedAt:   user.UpdatedAt.String(),
		DeletedAt:   user.DeletedAt.Time.Format(time.RFC3339),
	}

	return userDTO, nil
}

func getStringPointer(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

func CreateUserService(db *gorm.DB, userDTO UserResponseDTO) (UserResponseDTO, error) {
	if userDTO.Username == "" || userDTO.Email == "" || userDTO.Password == "" {
		return UserResponseDTO{}, errors.New("username, email, and password are required")
	}

	bytePassword := []byte(userDTO.Password)
	hash, error := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost) //DefaultCost es 10
	if error != nil {
		fmt.Println(error)
	}

	hashedPassword := string(hash)

	// Create the user struct
	user := schema.User{
		Username:    userDTO.Username,
		AvatarUrl:   &userDTO.AvatarUrl,
		DisplayName: userDTO.DisplayName,
		Email:       userDTO.Email,
		Password:    hashedPassword,
	}

	// Create the user in the database
	err := db.Create(&user).Error
	if err != nil {
		return UserResponseDTO{}, fmt.Errorf("failed to create user: %w", err)
	}

	//Convert time.Time to string for the response
	return UserResponseDTO{
		ID:          user.ID,
		Username:    user.Username,
		AvatarUrl:   *user.AvatarUrl,
		DisplayName: user.DisplayName,
		Email:       user.Email,
		Password:    user.Password,
		CreatedAt:   user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   user.UpdatedAt.Format(time.RFC3339),
		DeletedAt:   user.DeletedAt.Time.Format(time.RFC3339),
	}, nil

}
