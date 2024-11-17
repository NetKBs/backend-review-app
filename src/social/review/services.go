package review

import (
	"github.com/NetKBs/backend-reviewapp/src/image"
	"github.com/NetKBs/backend-reviewapp/src/schema"
	"github.com/NetKBs/backend-reviewapp/src/social/comment"
	"github.com/NetKBs/backend-reviewapp/src/social/reaction"
)

func GetReviewByIdService(id uint) (reviewDTO ReviewResponseDTO, err error) {

	review, err := GetReviewByIdRepository(id)
	if err != nil {
		return reviewDTO, err
	}

	reactionsCount, err := reaction.GetReactionsCountService(id, "review")
	if err != nil {
		return reviewDTO, err
	}

	commentsCount, err := comment.GetCommentsCountService(id)
	if err != nil {
		return reviewDTO, err
	}

	imagesPath, err := image.GetReviewImagesService(id)
	if err != nil {
		return reviewDTO, err
	}

	reviewDTO = ReviewResponseDTO{
		ID:        review.ID,
		UserId:    review.UserId,
		PlaceId:   review.PlaceId,
		Text:      review.Text,
		Rate:      review.Rate,
		Likes:     reactionsCount["likes"],
		Dislikes:  reactionsCount["dislikes"],
		Comments:  commentsCount,
		Images:    imagesPath,
		CreatedAt: review.CreatedAt.String(),
		UpdatedAt: review.UpdatedAt.String(),
	}

	return reviewDTO, nil
}

func CreateReviewService(review ReviewCreateDTO) (id uint, err error) {
	reviewSchema := schema.Review{UserId: review.UserId, PlaceId: review.PlaceId, Text: review.Text, Rate: review.Rate}
	id, err = CreateReviewRepository(reviewSchema)
	if err != nil {
		return id, err
	}
	return id, nil
}

func UpdateReviewService(id uint, review ReviewUpdateDTO) (err error) {
	reviewSchema := schema.Review{Text: review.Text}
	err = UpdateReviewRepository(id, reviewSchema)
	if err != nil {
		return err
	}
	return nil
}

func DeleteReviewService(id uint) (err error) {
	err = DeleteReviewRepository(id)
	if err != nil {
		return err
	}
	return nil
}
