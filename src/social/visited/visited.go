package visited

type VisitedPlaceInput struct {
	PlaceId uint `json:"place_id" binding:"required"`
}
