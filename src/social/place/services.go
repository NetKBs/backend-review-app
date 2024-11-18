package place

import (
	"context"

	"github.com/NetKBs/backend-reviewapp/geoapify"
)

func GetPlaceDetailsByMapsId(ctx context.Context, mapsID string) (geoapify.PlaceDetails, error) {
	return geoapify.GetPlaceDetailsById(mapsID)
}
