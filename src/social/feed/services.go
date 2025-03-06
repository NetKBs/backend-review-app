package feed

import (
	"strconv"

	"github.com/NetKBs/backend-reviewapp/src/image"
	"github.com/NetKBs/backend-reviewapp/src/social/comment"
	"github.com/NetKBs/backend-reviewapp/src/social/reaction"
	"github.com/NetKBs/backend-reviewapp/src/social/review"
)

func GetFeedService(userId uint, limit int, cursor string) ([]review.ReviewResponseDTO, string, error) {
	var reviewsResponse = []review.ReviewResponseDTO{}
	var lastID uint = 0

	if cursor != "" {
		cursorUint, err := strconv.ParseUint(cursor, 10, 64)
		if err != nil {
			return reviewsResponse, "", err
		}
		lastID = uint(cursorUint)
	}

	reviews, err := GetFeedRepository(userId, limit, lastID)
	if err != nil {
		return reviewsResponse, "", err
	}

	for _, re := range reviews {
		reactionsCount, err := reaction.GetReactionsCountService(re.ID, "review")
		if err != nil {
			return reviewsResponse, "", err
		}

		commentsCount, err := comment.GetCommentsReviewCountService(re.ID)
		if err != nil {
			return reviewsResponse, "", err
		}

		imagesPath, err := image.GetReviewImagesService(re.ID)
		if err != nil {
			return reviewsResponse, "", err
		}

		reviewDTO := review.ReviewResponseDTO{
			ID:        re.ID,
			UserId:    re.UserId,
			PlaceId:   re.PlaceId,
			Text:      re.Text,
			Rate:      re.Rate,
			Likes:     reactionsCount["likes"],
			Dislikes:  reactionsCount["dislikes"],
			Comments:  commentsCount,
			Images:    imagesPath,
			CreatedAt: re.CreatedAt.String(),
			UpdatedAt: re.UpdatedAt.String(),
		}
		reviewsResponse = append(reviewsResponse, reviewDTO)
	}

	nextCursor := ""
	if len(reviews) > 0 {
		nextCursor = strconv.FormatUint(uint64(reviews[len(reviews)-1].ID), 10)
	}

	return reviewsResponse, nextCursor, nil
}
