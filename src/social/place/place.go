package place

import (
	"github.com/NetKBs/backend-reviewapp/geoapify"
)

type PlaceDetailsResponseDTO struct {
	ID        uint                  `json:"id"`
	Details   geoapify.PlaceDetails `json:"details"`
	CreatedAt string                `json:"created_at"`
	UpdatedAt string                `json:"updated_at"`
}
