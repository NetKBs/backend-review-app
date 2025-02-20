package comment

type CommentResponseDTO struct {
	ID        uint   `json:"id"`
	UserId    uint   `json:"user_id"`
	ReviewId  uint   `json:"review_id"`
	Text      string `json:"text"`
	Likes     uint   `json:"likes"`
	Dislikes  uint   `json:"dislikes"`
	Answers   uint   `json:"answers"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CommentCreateDTO struct {
	ReviewId uint   `json:"review_id" binding:"required"`
	Text     string `json:"text" binding:"required"`
}

type CommentUpdateDTO struct {
	Text string `json:"text" binding:"required"`
}
