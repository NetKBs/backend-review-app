package review

import "github.com/NetKBs/backend-reviewapp/src/schema"

func GetReviewByIdService(id int) (review schema.Review, err error) {

	review, err = GetReviewByIdRepository(id)

	if err != nil {
		return review, err
	}

	return review, nil
}
