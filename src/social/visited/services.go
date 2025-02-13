package visited

func GetVisitedCountService(userId uint) (visitedCount uint, err error) {
	visitedCount, err = GetVisitedCountRepository(userId)
	if err != nil {
		return visitedCount, err
	}
	return visitedCount, nil
}

func GetVisitorsCountService(placeID uint) (uint, error) {
	count, err := GetVisitorsCountRepository(placeID)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CreateVisitedPlaceService(userID, placeID uint) error {
	err := CreateVisitedPlaceRepository(userID, placeID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteVisitedPlaceService(userID, placeID uint) error {
	err := DeleteVisitedPlaceRepository(userID, placeID)
	if err != nil {
		return err
	}
	return nil
}
