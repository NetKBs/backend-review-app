package geoapify

import (
	"encoding/json"
	"fmt"
	"strings"
)

const placeDetailsV2 = "/v2/place-details?"
const pDV2ByID = "id="

func GetPlaceDetailsById(params string) (PlaceDetails, error) {
	url := GEOAPIFY_SITE + placeDetailsV2 + pDV2ByID + params
	url += "&apiKey=" + apiKey

	status, body, err := getJSON(url)
	if err != nil || status != 200 {
		fmt.Println(err)
		return PlaceDetails{}, err
	}
	pd := PlaceDetails{}

	var parsed any
	err = json.Unmarshal(body, &parsed)
	parseFailed := false

	fc_parsed, ok := parsed.(map[string]any)
	parseFailed = parseFailed || !ok

	parsed_map, ok := fc_parsed["features"].([]any)[0].(map[string]any)["properties"].(map[string]any)
	parseFailed = parseFailed || !ok

	pd.Name, ok = parsed_map["name"].(string)
	parseFailed = parseFailed || !ok
	pd.Address, ok = parsed_map["formatted"].(string)
	parseFailed = parseFailed || !ok

	categories, ok := parsed_map["categories"]
	parseFailed = parseFailed || !ok
	categories_string := fmt.Sprintf("%v", categories)
	categories_string = strings.Trim(categories_string, "[]")
	category, categories_string, _ := strings.Cut(categories_string, " ")
	categoryAlt, _, _ := strings.Cut(categories_string, " ")
	if categoryAlt != "" {
		_, categoryAlt, _ = strings.Cut(categoryAlt, ".")
		category = categoryAlt
	}
	pd.Category = category

	pd.Longitude, ok = parsed_map["lon"].(float64)
	parseFailed = parseFailed || !ok
	pd.Latitude, ok = parsed_map["lat"].(float64)
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
	contacts.Twitter, ok = orNil(raw_datasource["contact:twitter"])
	contacts.Facebook, ok = orNil(raw_datasource["contact:facebook"])
	contacts.Instagram, ok = orNil(raw_datasource["contact:instagram"])
	pd.Contacts = contacts

	if parseFailed {
		err = fmt.Errorf("failed to parse place details")
		return PlaceDetails{}, err
	}
	return pd, nil
}
