package user

import "mime/multipart"

type UserResponseDTO struct {
	ID            uint   `json:"id"`
	Username      string `json:"username"`
	AvatarUrl     string `json:"avatar_url"`
	DisplayName   string `json:"display_name"`
	Email         string `json:"email"`
	Followers     uint   `json:"followers"`
	Following     uint   `json:"following"`
	Bookmarks     uint   `json:"bookmarks"`
	VisitedPlaces uint   `json:"visited_places"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type UserCreateDTO struct {
	Username    string                `form:"username" binding:"required"`
	AvatarImage *multipart.FileHeader `form:"avatar_image"`
	DisplayName string                `form:"display_name" binding:"required"`
	Email       string                `form:"email" binding:"required"`
	Password    string                `form:"password" binding:"required"`
}

type UpdatePasswordDTO struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}
