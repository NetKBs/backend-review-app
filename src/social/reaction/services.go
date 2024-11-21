package reaction

func GetReactionsCountService(id uint, contentType string) (reactionsCount map[string]uint, err error) {
	reactionsCount, err = GetReactionsCountRepository(id, contentType)
	if err != nil {
		return reactionsCount, err
	}

	return reactionsCount, nil
}
