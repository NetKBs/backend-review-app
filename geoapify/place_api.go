package geoapify

import (
	"encoding/json"
	"fmt"
	"log"
	// "strings"
)

const placesV2 = "/v2/places?"
const radius = "50"
const placesLimit = "10"

func GetPlacesAroundCoords(cat, lat, lon string) (places Places, err error) {
	catString := cat
	if cat == "" {
		catString = CategoriesString
	}
	categories := "categories=" + catString
	filter := "filter=circle:" + lat + "," + lon + "," + radius
	bias := "bias=proximity:" + lat + "," + lon
	limit := "limit=" + placesLimit
	url := GEOAPIFY_SITE + placesV2 + categories + "&" + filter + "&" + bias + "&" + limit + "&lang=es"
	log.Println(url)
	url += "&apiKey=" + apiKey

	status, body, err := getJSON(url)
	if err != nil || status != 200 {
		log.Println(err)
		return places, err
	}
	places, err = parsePlaces(body)
	return places, err
}

func parsePlaces(body []byte) (places Places, err error) {
	var parsed any
	err = json.Unmarshal(body, &parsed)
	parseFailed := false

	fc_parsed, ok := parsed.(map[string]any)
	parseFailed = parseFailed || !ok

	parsed_features, ok := fc_parsed["features"].([]any)
	parseFailed = parseFailed || !ok
	for _, raw_feature := range parsed_features {
		parsed_map, ok := raw_feature.(map[string]any)["properties"].(map[string]any)
		parseFailed = parseFailed || !ok
		if parseFailed {
			continue
		}
		var place Place

		place.MapsID, ok = parsed_map["place_id"].(string)
		parseFailed = parseFailed || !ok
		place.Name, ok = parsed_map["name"].(string)
		parseFailed = parseFailed || !ok
		place.Address, ok = parsed_map["formatted"].(string)
		parseFailed = parseFailed || !ok

		categories, ok := parsed_map["categories"].([]any)
		parseFailed = parseFailed || !ok
		for _, cat := range categories {
			parsed, ok := cat.(string)
			parseFailed = parseFailed || !ok

			place.Categories = append(place.Categories, parsed)
		}

		place.Longitude, ok = parsed_map["lon"].(float64)
		parseFailed = parseFailed || !ok
		place.Latitude, ok = parsed_map["lat"].(float64)
		parseFailed = parseFailed || !ok

		orNil := func(v any) (*string, bool) {
			result, ok := v.(string)
			if result == "" {
				return nil, ok
			}
			return &result, ok
		}
		raw_datasource, ok := parsed_map["datasource"].(map[string]any)["raw"].(map[string]any)
		contacts := PlaceContacts{}
		contacts.Mobile, ok = orNil(raw_datasource["phone"])
		contacts.Website, ok = orNil(raw_datasource["website"])
		contacts.Email, ok = orNil(raw_datasource["email"])
		contacts.Twitter, ok = orNil(raw_datasource["contact:twitter"])
		contacts.Facebook, ok = orNil(raw_datasource["contact:facebook"])
		contacts.Instagram, ok = orNil(raw_datasource["contact:instagram"])
		place.Contacts = contacts

		places = append(places, place)
	}

	if parseFailed {
		err = fmt.Errorf("failed to parse place details")
		return Places{}, err
	}
	return places, err
}
