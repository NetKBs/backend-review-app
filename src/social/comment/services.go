package comment

func GetCommentsCountService(id uint) (commentsCount uint, err error) {

	commentsCount, err = GetCommentsCountRepository(id)
	if err != nil {
		return commentsCount, err
	}

	return commentsCount, nil
}
