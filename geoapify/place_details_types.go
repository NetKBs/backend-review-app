package geoapify

type PlaceDetails struct {
	Name      string  `json:"name"`
	Category  string  `json:"category"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Contact   PlaceContacts
}

type PlaceContacts struct {
	Mobile    string `json:"mobile"`
	Facebook  string `json:"facebook"`
	Twitter   string `json:"twitter"`
	Instagram string `json:"instagram"`
}
