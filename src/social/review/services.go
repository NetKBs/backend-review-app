package review

import (
	"math"

	"github.com/NetKBs/backend-reviewapp/src/image"
	"github.com/NetKBs/backend-reviewapp/src/schema"
	"github.com/NetKBs/backend-reviewapp/src/social/comment"
	"github.com/NetKBs/backend-reviewapp/src/social/reaction"
)

func GetCountReviewsByUserIdService(id uint) (uint, error) {
	return GetCountReviewsByUserIdRepository(id)
}

func GetReviewByIdService(id uint) (reviewDTO ReviewResponseDTO, err error) {

	review, err := GetReviewByIdRepository(id)
	if err != nil {
		return reviewDTO, err
	}

	reactionsCount, err := reaction.GetReactionsCountService(id, "review")
	if err != nil {
		return reviewDTO, err
	}

	commentsCount, err := comment.GetCommentsReviewCountService(id)
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

func GetReviewsByUserIdService(userId uint, limit int, page int) ([]ReviewResponseDTO, schema.Pagination, error) {
	reviews, total, err := GetReviewsByUserIdRepository(userId, limit, page)
	if err != nil {
		return []ReviewResponseDTO{}, schema.Pagination{}, err
	}

	var reviewDTOs []ReviewResponseDTO
	for _, review := range reviews {
		reactionsCount, err := reaction.GetReactionsCountService(userId, "review")
		if err != nil {
			return []ReviewResponseDTO{}, schema.Pagination{}, err
		}

		commentsCount, err := comment.GetCommentsReviewCountService(userId)
		if err != nil {
			return []ReviewResponseDTO{}, schema.Pagination{}, err
		}

		imagesPath, err := image.GetReviewImagesService(userId)
		if err != nil {
			return []ReviewResponseDTO{}, schema.Pagination{}, err
		}

		reviewDTO := ReviewResponseDTO{
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
		reviewDTOs = append(reviewDTOs, reviewDTO)
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	hasNextPage := page < totalPages
	hasPreviousPage := page > 1

	pagination := schema.Pagination{
		TotalItems:  total,
		TotalPages:  totalPages,
		Limit:       limit,
		Page:        page,
		HasNextPage: hasNextPage,
		HasPrevPage: hasPreviousPage,
	}

	return reviewDTOs, pagination, nil
}

func GetReviewsByUserIdRepositoryCursorService(userId uint, limit int, lastID uint) ([]ReviewResponseDTO, error) {
	reviews, err := GetReviewsByUserIdRepositoryCursor(userId, limit, lastID)
	if err != nil {
		return []ReviewResponseDTO{}, err
	}

	var reviewDTOs []ReviewResponseDTO
	for _, review := range reviews {
		reactionsCount, err := reaction.GetReactionsCountService(userId, "review")
		if err != nil {
			return []ReviewResponseDTO{}, err
		}

		commentsCount, err := comment.GetCommentsReviewCountService(userId)
		if err != nil {
			return []ReviewResponseDTO{}, err
		}

		imagesPath, err := image.GetReviewImagesService(userId)
		if err != nil {
			return []ReviewResponseDTO{}, err
		}

		reviewDTO := ReviewResponseDTO{
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
		reviewDTOs = append(reviewDTOs, reviewDTO)
	}

	return reviewDTOs, nil
}

func CreateReviewService(review ReviewCreateDTO, userId uint) (id uint, err error) {
	reviewSchema := schema.Review{UserId: userId, PlaceId: review.PlaceId, Text: review.Text, Rate: review.Rate}
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
