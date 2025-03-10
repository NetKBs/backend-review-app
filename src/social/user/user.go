package user

import "mime/multipart"

type UserResponseDTO struct {
	ID            uint   `json:"id"`
	Username      string `json:"username"`
	AvatarUrl     string `json:"avatar_url"`
	DisplayName   string `json:"display_name"`
	Email         string `json:"email"`
	Description   string `json:"description"`
	Verified      bool   `json:"verified"`
	Role          string `json:"role"`
	Followers     uint   `json:"followers"`
	Following     uint   `json:"following"`
	Bookmarks     uint   `json:"bookmarks"`
	VisitedPlaces uint   `json:"visited_places"`
	Reviews       uint   `json:"reviews"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type UserSearchResultDTO struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	AvatarUrl   string `json:"avatar_url"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Verified    bool   `json:"verified"`
	Role        string `json:"role"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UserCreateDTO struct {
	Username    string                `form:"username" binding:"required"`
	AvatarImage *multipart.FileHeader `form:"avatar_image" binding:"required"`
	DisplayName string                `form:"display_name" binding:"required"`
	Email       string                `form:"email" binding:"required"`
	Description string                `form:"description" binding:"max=200"`
	Password    string                `form:"password" binding:"required"`
}

type UserUpdateDTO struct {
	DisplayName string                `form:"display_name"`
	Username    string                `form:"username"`
	Email       string                `form:"email"`
	Description string                `form:"description"`
	AvatarImage *multipart.FileHeader `form:"avatar_image"`
}

type UserUpdatePasswordDTO struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}
