package user

import (
	"time"
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
