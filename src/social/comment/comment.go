package comment

type CommentResponseDTO struct {
	ID        uint   `json:"id"`
	UserId    uint   `json:"user_id"`
	ReviewId  uint   `json:"review_id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CommentCreateDTO struct {
	UserId   uint   `json:"user_id" binding:"required"`
	ReviewId uint   `json:"review_id" binding:"required"`
	Text     string `json:"text" binding:"required"`
}

type CommentUpdateDTO struct {
	Text string `json:"text" binding:"required"`
}
