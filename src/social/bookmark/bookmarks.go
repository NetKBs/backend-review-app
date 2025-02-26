package bookmark

type BookmarkInputDTO struct {
	PlaceID uint `json:"place_id" binding:"required"`
}
