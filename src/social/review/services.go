package review

import (
	"github.com/NetKBs/backend-reviewapp/src/social/comment"
	"github.com/NetKBs/backend-reviewapp/src/social/reaction"
)

func GetReviewByIdService(id uint) (reviewDTO ReviewResponseDTO, err error) {

	review, err := GetReviewByIdRepository(id)
	if err != nil {
		return reviewDTO, err
	}

	reactionsCount, err := reaction.GetReactionsCountService(id)
	if err != nil {
		return reviewDTO, err
	}

	commentsCount, err := comment.GetCommentsCountService(id)
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
		CreatedAt: review.CreatedAt.String(),
		UpdatedAt: review.UpdatedAt.String(),
	}

	return reviewDTO, nil
}
