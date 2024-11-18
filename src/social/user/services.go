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
		CreatedAt:   user.CreatedAt.String(), //Preguntar
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

func UpdateUserService(newUser UserResponseDTO) error {

	user, err := GetUserByIdRepository(newUser.ID)
	if err != nil {
		return err
	}

	handleChanges(&user, &newUser)
	err = UpdateUserRepository(&user)

	return err
}

func handleChanges(user *schema.User, newUser *UserResponseDTO) error {
	var hashedPassword string
	var err error
	if newUser.Username != "" {
		user.Username = newUser.Username
	}
	if newUser.AvatarUrl != "" {
		user.AvatarUrl = &newUser.AvatarUrl
	}
	if newUser.DisplayName != "" {
		user.DisplayName = newUser.DisplayName
	}
	if newUser.Email != "" {
		user.Email = newUser.Email
	}
	if newUser.Password != "" {
		hashedPassword, err = bcryptHash(newUser.Password)
		user.Password = hashedPassword
	}
	return err
}

func bcryptHash(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost) //DefaultCost es 10
	return string(hash), err
}
