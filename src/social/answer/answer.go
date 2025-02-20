package answer

type AnswerResponseDTO struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	CommentID uint   `json:"comment_id"`
	Text      string `json:"text"`
	Likes     uint   `json:"likes"`
	Dislikes  uint   `json:"dislikes"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AnswerCreateDTO struct {
	CommentID uint   `json:"comment_id" binding:"required"`
	Text      string `json:"text" binding:"required"`
}

type AnswerUpdateDTO struct {
	Text string `json:"text" binding:"required"`
}
