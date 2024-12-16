package user

import (
	"errors"

	"github.com/NetKBs/backend-reviewapp/src/image"
	"github.com/NetKBs/backend-reviewapp/src/schema"
	"github.com/NetKBs/backend-reviewapp/src/social/bookmark"
	"github.com/NetKBs/backend-reviewapp/src/social/follow"
	"github.com/NetKBs/backend-reviewapp/src/social/visited"
	"golang.org/x/crypto/bcrypt"
)

func UserExistsByUsernameService(username string) (exists bool, err error) {
	return UserExistsByUsernameRepository(username)
}

func UserExistsByEmailService(email string) (exists bool, err error) {
	return UserExistsByEmailRepository(email)
}

func UserExistsByIdService(id uint) (exists bool, err error) {
	return UserExistsByIdRepository(id)
}

func CreateUserService(userDTO UserCreateDTO) (uint, error) {

	if exists, _ := UserExistsByUsernameService(userDTO.Username); exists {
		return 0, errors.New("username already exists")
	}
	if exists, _ := UserExistsByEmailService(userDTO.Email); exists {
		return 0, errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	user := schema.User{
		Username:    userDTO.Username,
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

func UpdatePasswordUserService(id uint, password UserUpdatePasswordDTO) error {
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

func UpdateAvatarUserService(id uint, newAvatarPath string) error {
	oldPath, err := UpdateAvatarUserRepository(id, newAvatarPath)
	if err != nil {
		return err
	}

	if oldPath != "" {
		if err := image.DeleteImageByPathService(oldPath); err != nil {
			return err
		}
	}

	return nil
}

func UpdateEmailUserService(id uint, email UserUpdateEmailDTO) error {
	if exists, _ := UserExistsByEmailService(email.Email); exists {
		return errors.New("email already exists")
	}
	return UpdateEmailUserRepository(id, email.Email)
}

func UpdateUserService(id uint, userDTO UserUpdateDTO) error {
	return UpdateUserRepository(id, userDTO)
}

func DeleteUserByIdService(id uint) error {
	avatarPath, err := DeleteUserbyIDRepository(id)
	if err != nil {
		return err
	}

	if err := image.DeleteImageByPathService(avatarPath); err != nil {
		return err
	}

	return nil
}
