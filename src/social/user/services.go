package user

import (
	"github.com/NetKBs/backend-reviewapp/src/schema"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserService(userDTO UserRequestDTO) (UserResponseDTO, error) {

	hashedPassword, err := bcryptHash(userDTO.Password)
	if err != nil {
		return UserResponseDTO{}, err
	}

	user := schema.User{
		Username:    userDTO.Username,
		AvatarUrl:   &userDTO.AvatarUrl,
		DisplayName: userDTO.DisplayName,
		Email:       userDTO.Email,
		Password:    hashedPassword,
	}

	err = CreateUserRepository(&user)

	return UserResponseDTO{
		ID:          user.ID,
		Username:    user.Username,
		AvatarUrl:   *user.AvatarUrl,
		DisplayName: user.DisplayName,
		Email:       user.Email,
	}, err

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
	}

	return userDTO, nil
}

func getStringPointer(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
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
	return err
}

func UpdatePasswordUserService(id uint, password string) error {
	hashedPassword, err := bcryptHash(password)

	if err != nil {
		return err
	}
	err = UpdatePasswordUserRepository(id, hashedPassword)

	return err
}

func bcryptHash(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost) //DefaultCost es 10
	return string(hash), err
}

func DeleteUserByIdService(id uint) error {
	err := DeleteUserbyIDRepository(id)
	return err
}
