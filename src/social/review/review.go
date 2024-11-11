package review

type ReviewResponseDTO struct {
	ID        uint   `json:"id"`
	UserId    uint   `json:"user_id"`
	PlaceId   uint   `json:"place_id"`
	Text      string `json:"text"`
	Rate      uint   `json:"rate"`
	Likes     uint   `json:"likes"`
	Dislikes  uint   `json:"dislikes"`
	Comments  uint   `json:"comments"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
