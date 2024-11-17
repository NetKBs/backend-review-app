package image

import (
	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func RegisterReviewImagesRepository(reviewImages []schema.ReviewImage) (err error) {
	db := config.DB

	if err = db.Create(&reviewImages).Error; err != nil {
		return err
	}
	return nil
}

func GetReviewImagesRepository(id uint) (reviewImages []schema.ReviewImage, err error) {
	db := config.DB

	if err = db.Where("review_id = ?", id).Find(&reviewImages).Error; err != nil {
		return reviewImages, err
	}
	return reviewImages, nil
}

func DeleteReviewImagesRepository(id uint) (err error) {
	db := config.DB

	if err = db.Where("review_id = ?", id).First(&schema.ReviewImage{}).Error; err != nil {
		return err
	}

	if err = db.Where("review_id = ?", id).Delete(&schema.ReviewImage{}).Error; err != nil {
		return err
	}

	return nil
}
