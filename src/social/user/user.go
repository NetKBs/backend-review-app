package user

type UserResponseDTO struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	AvatarUrl   string `json:"avatar_url"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	//Password    string `json:"password"`
}

type UserRequestDTO struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	AvatarUrl   string `json:"avatar_url"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}
