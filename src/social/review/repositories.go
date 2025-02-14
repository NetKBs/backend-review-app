package review

import (
	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func GetReviewByIdRepository(id uint) (review schema.Review, err error) {
	db := config.DB

	if err = db.Where("id = ?", id).First(&review).Error; err != nil {
		return review, err
	}
	return review, nil
}

func GetReviewsByUserIdRepository(userId uint, limit int, page int) ([]schema.Review, error) {
	var reviews []schema.Review
	offset := (page - 1) * limit

	db := config.DB

	err := db.Where("user_id = ?", userId).Limit(limit).Offset(offset).Find(&reviews).Error
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func CreateReviewRepository(review schema.Review) (id uint, err error) {
	db := config.DB

	if err = db.Create(&review).Error; err != nil {
		return id, err
	}
	return review.ID, nil
}

func UpdateReviewRepository(id uint, review schema.Review) (err error) {
	db := config.DB

	if err = db.Where("id = ?", id).First(&schema.Review{}).Error; err != nil {
		return err
	}

	if err = db.Model(&schema.Review{}).Where("id = ?", id).Update("Text", review.Text).Error; err != nil {
		return err
	}
	return nil
}

func DeleteReviewRepository(id uint) (err error) {
	db := config.DB

	if err = db.Where("id = ?", id).First(&schema.Review{}).Error; err != nil {
		return err
	}

	if err = db.Where("id = ?", id).Delete(&schema.Review{}).Error; err != nil {
		return err
	}
	return nil
}
