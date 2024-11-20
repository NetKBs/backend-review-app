package place

import (
	"github.com/NetKBs/backend-reviewapp/geoapify"
)

type PlaceDetailsResponseDTO struct {
	PlaceID   uint                  `json:"place_id"`
	MapsId    string                `json:"maps_id"`
	Details   geoapify.PlaceDetails `json:"details"`
	CreatedAt string                `json:"created_at"`
	UpdatedAt string                `json:"updated_at"`
}
