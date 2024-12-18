package visited

func GetVisitedCount(userId uint) (visitedCount uint, err error) {
	visitedCount, err = GetVisitedCountRepository(userId)
	if err != nil {
		return visitedCount, err
	}
	return visitedCount, nil
}
