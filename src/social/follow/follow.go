package follow

type InputFollow struct {
	FollowedId uint `json:"followed_id" binding:"required"`
}

type FollowResponseDTO struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}
