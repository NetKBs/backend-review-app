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

type AutocompleteResponseDTO struct {
	Query  string              `json:"query"`
	Result geoapify.Geocodings `json:"result"`
}

type PlacesResponseDTO struct {
	CenterLon float64         `json:"center_lon"`
	CenterLan float64         `json:"center_lat"`
	Data      geoapify.Places `json:"data"`
}
