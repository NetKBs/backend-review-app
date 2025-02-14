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

func GetFollowersByIdService(id uint) ([]FollowResponseDTO, error) {
	followersResponse := []FollowResponseDTO{}

	followers, err := GetFollowersByIdRespository(id)
	if err != nil {
		return []FollowResponseDTO{}, nil
	}

	for _, follower := range followers {
		f := FollowResponseDTO{
			ID:        follower.ID,
			Username:  follower.Username,
			AvatarURL: *follower.AvatarUrl,
		}
		followersResponse = append(followersResponse, f)
	}

	return followersResponse, nil
}

func GetFollowingsByIdService(id uint) ([]FollowResponseDTO, error) {
	followingsResponse := []FollowResponseDTO{}
	followings, err := GetFollowingsByIdRepository(id)
	if err != nil {
		return []FollowResponseDTO{}, err
	}

	for _, following := range followings {
		f := FollowResponseDTO{
			ID:        following.ID,
			Username:  following.Username,
			AvatarURL: *following.AvatarUrl,
		}
		followingsResponse = append(followingsResponse, f)
	}
	return followingsResponse, nil
}

func CreateFollowService(follower_id, followed_id uint) (err error) {
	return CreateFollowRepository(follower_id, followed_id)
}

func DeleteFollowService(follower_id, followed_id uint) (err error) {
	return DeleteFollowRespository(follower_id, followed_id)
}
