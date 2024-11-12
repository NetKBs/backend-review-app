package reaction

func GetReactionsCountService(id uint) (reactionsCount map[string]uint, err error) {
	reactionsCount, err = GetReactionsCountRepository(id)
	if err != nil {
		return reactionsCount, err
	}

	return reactionsCount, nil
}
