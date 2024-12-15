package user

import (
	"errors"

	"github.com/NetKBs/backend-reviewapp/src/schema"
	"github.com/NetKBs/backend-reviewapp/src/social/bookmark"
	"github.com/NetKBs/backend-reviewapp/src/social/follow"
	"github.com/NetKBs/backend-reviewapp/src/social/visited"
	"golang.org/x/crypto/bcrypt"
)

func VerifyUsernameService(username string) (exists bool, err error) {

	_, err = GetUsernameUserRepository(username)
	if err != nil {
		return false, err
	}
	return true, err
}

func CreateUserService(userDTO UserCreateDTO, avatarPath string) (uint, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	user := schema.User{
		Username:    userDTO.Username,
		AvatarUrl:   &avatarPath,
		DisplayName: userDTO.DisplayName,
		Email:       userDTO.Email,
		Password:    string(hashedPassword),
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

func UpdatePasswordUserService(id uint, password UpdatePasswordDTO) error {
	dbPassword, err := GetPasswordUserRepository(id)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password.OldPassword)); err != nil {
		return errors.New("invalid old password")
	}

	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(password.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	err = UpdatePasswordUserRepository(id, string(hashedNewPassword))
	return err
}

func DeleteUserByIdService(id uint) error {
	err := DeleteUserbyIDRepository(id)
	return err
}
