package visited

func GetVisitedCountService(userId uint) (visitedCount uint, err error) {
	visitedCount, err = GetVisitedCountRepository(userId)
	return visitedCount, err
}

func GetVisitorsCountService(placeID uint) (uint, error) {
	count, err := GetVisitorsCountRepository(placeID)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CreateVisitedPlaceService(userID, placeID uint) error {
	return CreateVisitedPlaceRepository(userID, placeID)
}

func DeleteVisitedPlaceService(userID, placeID uint) error {
	return DeleteVisitedPlaceRepository(userID, placeID)
}
