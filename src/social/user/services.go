package user

import (
	"errors"
	"strconv"

	"github.com/NetKBs/backend-reviewapp/src/image"
	"github.com/NetKBs/backend-reviewapp/src/schema"
	"github.com/NetKBs/backend-reviewapp/src/social/bookmark"
	"github.com/NetKBs/backend-reviewapp/src/social/follow"
	"github.com/NetKBs/backend-reviewapp/src/social/review"
	"github.com/NetKBs/backend-reviewapp/src/social/visited"
	"golang.org/x/crypto/bcrypt"
)

func HandleUniquenessError(type_ string) error {
	switch type_ {
	case "username":
		return errors.New("username already exists")
	case "email":
		return errors.New("email already exists")
	default:
		return nil
	}
}

func UserExistsByFieldService(field string, value interface{}, excludeId uint) (bool, error) {
	return UserExistsByFieldRepository(field, value, excludeId)
}

func CreateUserService(userDTO UserCreateDTO) (uint, error) {

	if exists, err := UserExistsByFieldService("username", userDTO.Username, 0); err != nil {
		return 0, err
	} else if exists {
		return 0, HandleUniquenessError("username")
	}

	if exists, err := UserExistsByFieldService("email", userDTO.Email, 0); err != nil {
		return 0, err
	} else if exists {
		return 0, HandleUniquenessError("email")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	user := schema.User{
		Username:    userDTO.Username,
		DisplayName: userDTO.DisplayName,
		Email:       userDTO.Email,
		Description: userDTO.Description,
		Password:    string(hashedPassword),
	}

	id, err := CreateUserRepository(user)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetUserByFieldService(field string, value interface{}) (userDTO UserResponseDTO, err error) {
	var user schema.User

	switch field {
	case "id":
		var idParsed int
		if idParsed, err = strconv.Atoi(value.(string)); err != nil {
			return userDTO, errors.New("invalid id type")
		}
		user, err = GetUserByIdRepository(uint(idParsed))

	case "username":
		username, ok := value.(string)
		if !ok {
			return userDTO, errors.New("invalid username type")
		}
		user, err = GetUserByUsernameRepository(username)

	default:
		return userDTO, errors.New("invalid field")

	}

	if err != nil {
		return userDTO, err
	}

	followersCount, err := follow.GetFollowersCountService(user.ID)
	if err != nil {
		return userDTO, err
	}
	followingCount, err := follow.GetFollowingCountService(user.ID)
	if err != nil {
		return userDTO, err
	}
	bookmarkCount, err := bookmark.GetBookmarkCount(user.ID)
	if err != nil {
		return userDTO, err
	}
	visitedCount, err := visited.GetVisitedCountService(user.ID)
	if err != nil {
		return userDTO, err
	}
	reviewsCount, err := review.GetCountReviewsByUserIdService(user.ID)
	if err != nil {
		return userDTO, err
	}

	userDTO = UserResponseDTO{
		ID:            user.ID,
		Username:      user.Username,
		AvatarUrl:     getStringPointer(user.AvatarUrl),
		DisplayName:   user.DisplayName,
		Email:         user.Email,
		Description:   user.Description,
		Verified:      user.Verified,
		Role:          user.Role,
		Followers:     followersCount,
		Following:     followingCount,
		Bookmarks:     bookmarkCount,
		VisitedPlaces: visitedCount,
		Reviews:       reviewsCount,
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

func UpdateUserService(id uint, userDTO UserUpdateDTO, avatarPath string) error {

	if userDTO.Username != "" {
		if exists, err := UserExistsByFieldService("username", userDTO.Username, id); err != nil {
			return err
		} else if exists {
			return HandleUniquenessError("username")
		}
	}

	if userDTO.Email != "" {
		if exists, err := UserExistsByFieldService("email", userDTO.Email, id); err != nil {
			return err
		} else if exists {
			return HandleUniquenessError("email")
		}
	}

	oldAvatar, err := UpdateUserRepository(id, userDTO, avatarPath)
	if err != nil {
		return err
	}

	if avatarPath != "" && oldAvatar != "" {
		image.DeleteImageByPathService(oldAvatar)
	}
	return nil

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

func VerifyUserService(userId uint) error {
	return VerifyUserRepository(userId)
}
