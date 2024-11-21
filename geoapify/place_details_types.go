package geoapify

type PlaceDetails struct {
	MapsID    string        `json:"maps_id"`
	Name      string        `json:"name"`
	Category  string        `json:"category"`
	Address   string        `json:"address"`
	Latitude  float64       `json:"latitude"`
	Longitude float64       `json:"longitude"`
	Contacts  PlaceContacts `json:"contacts"`
}

type PlaceContacts struct {
	Website   *string `json:"website"`
	Email     *string `json:"email"`
	Mobile    *string `json:"mobile"`
	Facebook  *string `json:"facebook"`
	Twitter   *string `json:"twitter"`
	Instagram *string `json:"instagram"`
}
