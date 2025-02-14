package geoapify

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

const GEOAPIFY_SITE = "https://api.geoapify.com"
const CategoriesString = `accommodation,activity,airport,commercial,catering,education,childcare,entertainment,healthcare,heritage,highway,leisure,man_made,natural,national_park,office,parking,pet,power,production,railway,rental,service,tourism,religion,camping,amenity,beach,adult,airport,building,ski,sport,public_transport,administrative,postal_code,political,low_emission_zone,populated_place,production`

var Categories = [...]string{
	"accommodation",
	"activity",
	"airport",
	"commercial",
	"catering",
	"education",
	"childcare",
	"entertainment",
	"healthcare",
	"heritage",
	"highway",
	"leisure",
	"man_made",
	"natural",
	"national_park",
	"office",
	"parking",
	"pet",
	"power",
	"production",
	"railway",
	"rental",
	"service",
	"tourism",
	"religion",
	"camping",
	"amenity",
	"beach",
	"adult",
	"airport",
	"building",
	"ski",
	"sport",
	"public_transport",
	"administrative",
	"postal_code",
	"political",
	"low_emission_zone",
	"populated_place",
	"production",
}

var apiKey string

func SetGeoapifyKey(extApiKey string) {
	if extApiKey == "" {
		log.Println("Geoapify API key is empty. Geoapify client will not work.")
		return
	}

	apiKey = extApiKey
}

func getJSON(apiURL string) (int, []byte, error) {
	res, err := http.Get(apiURL)
	if err != nil {
		return 0, []byte{}, err
	}

	if res.StatusCode == 200 {
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		return res.StatusCode, body, err
	}

	if res.StatusCode == 202 {
		return res.StatusCode, []byte{}, nil
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	e := fmt.Sprintf("GET: Expecting 200 or 202, but got status %d\n%s",
		res.StatusCode,
		string(body))
	return 0, []byte{}, errors.New(e)
}
