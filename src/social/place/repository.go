package place

import (
	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
	"gorm.io/gorm"
)

func findPlaceByMapsIdRepo(mapsID string) (place schema.Place, err error) {
	db := config.DB
	place = schema.Place{MapsId: mapsID}

	if err := db.Where("maps_id = ?", mapsID).First(&place).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&place).Error; err != nil {
				return schema.Place{}, err
			}
		} else {
			return schema.Place{}, err
		}

	}
	return place, nil
}
