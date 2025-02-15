package feed

import (
	"strconv"

	"github.com/NetKBs/backend-reviewapp/src/social/follow"
	"github.com/NetKBs/backend-reviewapp/src/social/review"
)

func GetFeedService(userId uint, limit int, cursor string) ([]review.ReviewResponseDTO, string, error) {
	followings, err := follow.GetFollowingsByIdRepository(userId)
	if err != nil {
		return nil, "", err
	}

	var allReviews []review.ReviewResponseDTO
	var lastID uint = 0

	if cursor != "" {
		cursorUint, err := strconv.ParseUint(cursor, 10, 64)
		if err != nil {
			return nil, "", err
		}
		lastID = uint(cursorUint)
	}

	for _, following := range followings {
		reviews, err := review.GetReviewsByUserIdRepositoryCursorService(following.ID, limit, lastID)
		if err != nil {
			return nil, "", err
		}
		allReviews = append(allReviews, reviews...)
	}

	nextCursor := ""
	if len(allReviews) > 0 {
		nextCursor = strconv.FormatUint(uint64(allReviews[len(allReviews)-1].ID), 10)
	}

	return allReviews, nextCursor, nil
}
