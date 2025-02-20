package review

import (
	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func GetCountReviewsByUserIdRepository(id uint) (uint, error) {
	db := config.DB
	var count int64

	if err := db.Model(&schema.Review{}).Where("user_id = ?", id).Count(&count).Error; err != nil {
		return 0, err
	}
	return uint(count), nil
}

func GetReviewByIdRepository(id uint) (review schema.Review, err error) {
	db := config.DB

	if err = db.Where("id = ?", id).First(&review).Error; err != nil {
		return review, err
	}
	return review, nil
}

func GetReviewsByUserIdRepository(userId uint, limit int, page int) ([]schema.Review, int64, error) {
	var reviews []schema.Review
	offset := (page - 1) * limit
	db := config.DB

	var total int64
	if err := db.Model(&schema.Review{}).Where("user_id = ?", userId).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := db.Where("user_id = ?", userId).Limit(limit).Offset(offset).Order("created_at DESC").Find(&reviews).Error
	if err != nil {
		return nil, 0, err
	}

	return reviews, total, nil
}

func GetReviewsByUserIdRepositoryCursor(userId uint, limit int, lastID uint) ([]schema.Review, error) {
	var reviews []schema.Review
	db := config.DB

	query := db.Where("user_id = ?", userId)

	if lastID > 0 {
		query = query.Where("id < ?", lastID)
	}

	err := query.Order("id DESC").Limit(limit).Find(&reviews).Error

	return reviews, err
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
