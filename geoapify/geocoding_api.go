package geoapify

import (
	"encoding/json"
	"fmt"
	"log"
)

const GEOCODE_AUTOCOMPLETE = "/v1/geocode/autocomplete?"

func GetAutocompleteResponse(text string) (geocodings Geocodings, err error) {
	// filter := "filter=circle:" + lat + "," + lon + "," + radius
	text = "text=" + text
	format := "format=json"
	bias := "bias=countrycode:ve"
	limit := "limit=" + placesLimit
	url := GEOAPIFY_SITE + GEOCODE_AUTOCOMPLETE + text + "&" + format + "&" + bias + "&" + limit + "&lang=es"
	log.Println(url)
	url += "&apiKey=" + apiKey

	status, body, err := getJSON(url)
	if err != nil || status != 200 {
		log.Println(err)
		return geocodings, err
	}
	geocodings, err = parseGeocoding(body)
	return geocodings, err
}

func parseGeocoding(body []byte) (geocodings Geocodings, err error) {
	var parsed any
	err = json.Unmarshal(body, &parsed)
	parseFailed := false
	// log.Println(parsed)

	fc_parsed, ok := parsed.(map[string]any)
	parseFailed = parseFailed || !ok

	parsed_features, ok := fc_parsed["results"].([]any)
	parseFailed = parseFailed || !ok
	for _, raw_feature := range parsed_features {
		parsed_map, ok := raw_feature.(map[string]any)
		parseFailed = parseFailed || !ok
		if parseFailed {
			continue
		}
		var geocoding Geocoding

		orNil := func(v any) (*string, bool) {
			result, ok := v.(string)
			if result == "" {
				return nil, true
			}
			return &result, ok
		}
		geocoding.MapsID, ok = parsed_map["place_id"].(string)
		parseFailed = parseFailed || !ok
		geocoding.Name, ok = orNil(parsed_map["name"])
		parseFailed = parseFailed || !ok
		geocoding.Address, ok = parsed_map["formatted"].(string)
		parseFailed = parseFailed || !ok

		// log.Println(parsed_map["category"], reflect.TypeOf(parsed_map["category"]))
		geocoding.Category, ok = parsed_map["category"].(string)
		// parseFailed = parseFailed || !ok

		geocoding.Longitude, ok = parsed_map["lon"].(float64)
		parseFailed = parseFailed || !ok
		geocoding.Latitude, ok = parsed_map["lat"].(float64)
		parseFailed = parseFailed || !ok

		var rank AutocompleteRank
		{
			rankObj, ok := parsed_map["rank"].(map[string]any)
			importance, ok := rankObj["importance"].(float32)
			if ok {
				rank.Importance = importance
			}
			rank.Confidence, ok = rankObj["confidence"].(float64)
			parseFailed = parseFailed || !ok
			rank.MatchType, ok = rankObj["match_type"].(string)
			parseFailed = parseFailed || !ok
		}
		geocoding.Rank = rank

		geocodings = append(geocodings, geocoding)
	}

	if parseFailed {
		err = fmt.Errorf("failed to parse geocoding details")
		return Geocodings{}, err
	}
	return geocodings, err
}
