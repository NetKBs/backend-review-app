package geoapify

type PlaceDetails struct {
	PlaceID   string        `json:"place_id"`
	Name      string        `json:"name"`
	Category  string        `json:"category"`
	Address   string        `json:"address"`
	Latitude  float64       `json:"latitude"`
	Longitude float64       `json:"longitude"`
	Contacts  PlaceContacts `json:"contacts"`
}

type PlaceContacts struct {
	Mobile    *string `json:"mobile",omitempty`
	Facebook  *string `json:"facebook",omitempty`
	Twitter   *string `json:"twitter",omitempty`
	Instagram *string `json:"instagram",omitempty`
}
