package image

import (
	"os"

	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func RegisterReviewImagesService(id uint, imagePaths []string) (err error) {

	var reviewImages []schema.ReviewImage
	for _, imagePath := range imagePaths {
		reviewImages = append(reviewImages, schema.ReviewImage{ReviewId: id, ImageURL: imagePath})
	}

	if err = RegisterReviewImagesRepository(reviewImages); err != nil {
		return err
	}
	return nil
}

func GetReviewImagesService(id uint) (imagePaths []string, err error) {
	reviewImages, err := GetReviewImagesRepository(id)
	if err != nil {
		return imagePaths, err
	}

	for _, reviewImage := range reviewImages {
		imagePaths = append(imagePaths, reviewImage.ImageURL)
	}
	return imagePaths, nil
}

func DeleteReviewImagesService(id uint) (err error) {
	imagePaths, err := GetReviewImagesService(id)
	if err != nil {
		return err
	}

	for _, imagePath := range imagePaths {
		os.Remove(imagePath)
		/*if err = os.Remove(imagePath); err != nil {
			return err
		}*/
	}

	if err = DeleteReviewImagesRepository(id); err != nil {
		return err
	}

	return nil
}

func DeleteImageByPathService(imagePath string) (err error) {
	if err = os.Remove(imagePath); err != nil {
		return err
	}
	return nil
}
