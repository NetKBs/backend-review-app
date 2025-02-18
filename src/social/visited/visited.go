package visited

type VisitedPlaceInput struct {
	UserId  uint `json:"user_id" binding:"required"`
	PlaceId uint `json:"place_id" binding:"required"`
}
