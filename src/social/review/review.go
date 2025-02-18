package review

import (
	"mime/multipart"
)

type ReviewResponseDTO struct {
	ID        uint     `json:"id"`
	UserId    uint     `json:"user_id"`
	PlaceId   uint     `json:"place_id"`
	Text      string   `json:"text"`
	Rate      uint     `json:"rate"`
	Likes     uint     `json:"likes"`
	Dislikes  uint     `json:"dislikes"`
	Comments  uint     `json:"comments"`
	Images    []string `json:"images"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type ReviewCreateDTO struct {
	UserId  uint                    `form:"user_id" binding:"required"`
	PlaceId uint                    `form:"place_id" binding:"required"`
	Text    string                  `form:"text" binding:"required"`
	Rate    uint                    `form:"rate" binding:"required,oneof=1 2 3 4 5"`
	Images  []*multipart.FileHeader `form:"images" binding:"required"`
}

type ReviewUpdateDTO struct {
	Text string `json:"text" binding:"required"`
}
