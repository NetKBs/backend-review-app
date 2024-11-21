package place

import (
	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func findPlaceByMapsIdRepo(mapsID string) (place schema.Place, err error) {
	db := config.DB
	place = schema.Place{MapsId: mapsID}

	result := db.First(&place)
	// if result.Error.Error() != "record not found" {
	if result.RowsAffected == 0 {
		result = db.Create(&place)
	}
	return place, result.Error
}
