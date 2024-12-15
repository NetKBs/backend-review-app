package bookmark

func GetBookmarkCount(userId uint) (bookmarkCount uint, err error) {
	bookmarkCount, err = GetBookmarkCountRepository(userId)
	if err != nil {
		return bookmarkCount, err
	}
	return bookmarkCount, nil
}
