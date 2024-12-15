package user

import (
	"github.com/NetKBs/backend-reviewapp/src/schema"
	"github.com/NetKBs/backend-reviewapp/src/social/bookmark"
	"github.com/NetKBs/backend-reviewapp/src/social/follow"
	"github.com/NetKBs/backend-reviewapp/src/social/visited"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserService(userDTO UserCreateDTO, avatarPath string) (uint, error) {

	hashedPassword, err := bcryptHash(userDTO.Password)
	if err != nil {
		return 0, err
	}

	user := schema.User{
		Username:    userDTO.Username,
		AvatarUrl:   &avatarPath,
		DisplayName: userDTO.DisplayName,
		Email:       userDTO.Email,
		Password:    hashedPassword,
	}

	id, err := CreateUserRepository(user)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetUserByIdService(id uint) (userDTO UserResponseDTO, err error) {

	user, err := GetUserByIdRepository(id)
	if err != nil {
		return userDTO, err
	}

	followersCount, err := follow.GetFollowersCountService(id)
	if err != nil {
		return userDTO, err
	}
	followingCount, err := follow.GetFollowingCountService(id)
	if err != nil {
		return userDTO, err
	}
	bookmarkCount, err := bookmark.GetBookmarkCount(id)
	if err != nil {
		return userDTO, err
	}
	visitedCount, err := visited.GetVisitedCount(id)
	if err != nil {
		return userDTO, err
	}

	userDTO = UserResponseDTO{
		ID:            user.ID,
		Username:      user.Username,
		AvatarUrl:     getStringPointer(user.AvatarUrl),
		DisplayName:   user.DisplayName,
		Email:         user.Email,
		Followers:     followersCount,
		Following:     followingCount,
		Bookmarks:     bookmarkCount,
		VisitedPlaces: visitedCount,
		CreatedAt:     user.CreatedAt.String(),
		UpdatedAt:     user.UpdatedAt.String(),
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
