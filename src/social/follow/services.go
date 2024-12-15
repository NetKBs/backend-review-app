package follow

func GetFollowersCountService(id uint) (followersCount uint, err error) {
	followersCount, err = GetFollowersCount(id)
	if err != nil {
		return followersCount, err
	}
	return followersCount, nil
}

func GetFollowingCountService(id uint) (followingsCount uint, err error) {
	followingsCount, err = GetFollowingCount(id)
	if err != nil {
		return followingsCount, err
	}
	return followingsCount, nil
}
