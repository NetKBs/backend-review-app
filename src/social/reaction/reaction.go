package reaction

type ReactionResponseDTO struct {
	ID           uint   `json:"id"`
	UserId       uint   `json:"user_id"`
	ContentId    uint   `json:"content_id"`
	ContentType  string `json:"content_type"`
	ReactionType bool   `json:"reaction_type"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type ReactionCreateDTO struct {
	ContentId    uint   `json:"content_id" binding:"required"`
	ContentType  string `json:"content_type" binding:"required"`
	ReactionType bool   `json:"reaction_type" binding:"required"`
}
