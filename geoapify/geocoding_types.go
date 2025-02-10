package geoapify

type AutocompleteRank struct {
	Importance float32
	Confidence float64
	MatchType  string
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
