package comment

import (
	"github.com/NetKBs/backend-reviewapp/src/schema"
	"github.com/NetKBs/backend-reviewapp/src/social/answer"
	"github.com/NetKBs/backend-reviewapp/src/social/reaction"
)

func GetCommentsReviewCountService(id uint) (commentsCount uint, err error) {

	commentsCount, err = GetCommentsReviewCountRepository(id)
	if err != nil {
		return commentsCount, err
	}

	return commentsCount, nil
}

func GetCommentsByIdReviewService(id uint, limit int, page int) ([]CommentResponseDTO, schema.Pagination, error) {
	reviewComments := []CommentResponseDTO{}

	comments, total, err := GetCommentsByIdReviewRepository(id, limit, page)
	if err != nil {
		return reviewComments, schema.Pagination{}, err
	}

	for _, comment := range comments {
		reactions, err := reaction.GetReactionsCountService(comment.ID, "comment")
		if err != nil {
			return reviewComments, schema.Pagination{}, err
		}

		replies, err := answer.GetCountAnswersByCommentIdService(comment.ID)
		if err != nil {
			return reviewComments, schema.Pagination{}, err
		}

		reviewComments = append(reviewComments, CommentResponseDTO{
			ID:        comment.ID,
			UserId:    comment.UserId,
			ReviewId:  comment.ReviewId,
			Text:      comment.Text,
			Likes:     reactions["likes"],
			Dislikes:  reactions["dislikes"],
			Answers:   replies,
			CreatedAt: comment.CreatedAt.String(),
			UpdatedAt: comment.UpdatedAt.String(),
		})
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))
	pagination := schema.Pagination{
		TotalItems:  total,
		TotalPages:  totalPages,
		Limit:       limit,
		Page:        page,
		HasNextPage: page < totalPages,
		HasPrevPage: page > 1,
	}

	return reviewComments, pagination, nil
}

func GetCommentByIdService(id uint) (commentDTO CommentResponseDTO, err error) {
	comment, err := GetCommentsByIdRepository(id)
	if err != nil {
		return commentDTO, err
	}

	reactions, err := reaction.GetReactionsCountService(id, "comment")
	if err != nil {
		return commentDTO, err
	}

	replies, err := answer.GetCountAnswersByCommentIdService(id)
	if err != nil {
		return commentDTO, err
	}

	commentDTO = CommentResponseDTO{
		ID:        comment.ID,
		UserId:    comment.UserId,
		ReviewId:  comment.ReviewId,
		Text:      comment.Text,
		Likes:     reactions["likes"],
		Dislikes:  reactions["dislikes"],
		Answers:   replies,
		CreatedAt: comment.CreatedAt.String(),
		UpdatedAt: comment.UpdatedAt.String(),
	}

	return commentDTO, nil
}

func CreateCommentService(comment CommentCreateDTO, userId uint) (id uint, err error) {
	commentSchema := schema.Comment{UserId: userId, ReviewId: comment.ReviewId, Text: comment.Text}

	id, err = CreateCommentRepository(commentSchema)
	if err != nil {
		return id, err
	}

	return id, nil
}

func UpdateCommentService(id uint, comment CommentUpdateDTO) (err error) {
	commentSchema := schema.Comment{Text: comment.Text}
	err = UpdateCommentRepository(id, commentSchema)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCommentService(id uint) (err error) {
	err = DeleteCommentRepository(id)
	if err != nil {
		return err
	}
	return nil
}
