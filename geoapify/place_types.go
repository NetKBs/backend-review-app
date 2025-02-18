package geoapify

type Place struct {
	MapsID     string        `json:"maps_id"`
	Name       string        `json:"name"`
	Categories []string      `json:"categories"`
	Address    string        `json:"address"`
	Latitude   float64       `json:"latitude"`
	Longitude  float64       `json:"longitude"`
	Contacts   PlaceContacts `json:"contacts"`
}

type Places []Place
