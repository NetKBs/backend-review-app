package feed

import (
	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func GetFeedRepository(userId uint, limit int, lastID uint) ([]schema.Review, error) {
	db := config.DB
	var reviews = []schema.Review{}

	query := db.Table("follow").
		Select("reviews.*, follow.followed_id").
		Joins("inner join reviews on reviews.user_id = follow.followed_id").
		Where("follow.follower_id = ?", userId).
		Order("reviews.id DESC").
		Limit(limit)

	if lastID != 0 {
		query = query.Where("reviews.id < ?", lastID)
	}

	if err := query.Scan(&reviews).Error; err != nil {
		return reviews, err
	}

	return reviews, nil
}
