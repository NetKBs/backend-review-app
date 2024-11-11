package review

import (
	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func GetReviewByIdRepository(id int) (review schema.Review, err error) {
	db := config.DB

	if err = db.Where("id = ?", id).First(&review).Error; err != nil {
		return review, err
	}

	return review, nil
}
