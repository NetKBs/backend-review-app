package user

import (
	"time"

	"github.com/NetKBs/backend-reviewapp/src/schema"
	"golang.org/x/crypto/bcrypt"
)

func DeleteUserByIdService(id uint) error {
	err := DeleteUserbyIDRepository(id)
	return err
}

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

func CreateUserService(userDTO UserResponseDTO) (UserResponseDTO, error) {

	hashedPassword, err := bcryptHash(userDTO.Password)
	if err != nil {
		return UserResponseDTO{}, err
	}
	// Create the user struct
	user := schema.User{
		Username:    userDTO.Username,
		AvatarUrl:   &userDTO.AvatarUrl,
		DisplayName: userDTO.DisplayName,
		Email:       userDTO.Email,
		Password:    hashedPassword,
	}

	err = CreateUserRepository(&user)

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
	}, err

}

func UpdateUserService(newUser UserResponseDTO) error {

	user, err := GetUserByIdRepository(uint(newUser.ID))
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
