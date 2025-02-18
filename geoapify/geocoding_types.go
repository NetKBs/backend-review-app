package geoapify

type AutocompleteRank struct {
	Importance float32 `json:"importance"`
	Confidence float64 `json:"confidence"`
	MatchType  string  `json:"match_type"`
}

type Geocoding struct {
	MapsID     string           `json:"maps_id"`
	Name       *string          `json:"name"`
	ResultType string           `json:"result_type"`
	Category   string           `json:"category"`
	Address    string           `json:"address"`
	Latitude   float64          `json:"latitude"`
	Longitude  float64          `json:"longitude"`
	Rank       AutocompleteRank `json:"rank"`
}

type Geocodings []Geocoding
