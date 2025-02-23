package answer

import (
	"github.com/NetKBs/backend-reviewapp/src/schema"
	"github.com/NetKBs/backend-reviewapp/src/social/reaction"
)

func GetCountAnswersByCommentIdService(id uint) (count uint, err error) {
	return GetCountAnswersByCommentIdRepository(id)
}

func GetAnswersByCommentIdService(id uint, limit int, page int) ([]AnswerResponseDTO, schema.Pagination, error) {
	answerComments := []AnswerResponseDTO{}

	answers, total, err := GetAnswersByCommentIdRepository(id, limit, page)
	if err != nil {
		return answerComments, schema.Pagination{}, err
	}

	for _, answer := range answers {
		reactions, err := reaction.GetReactionsCountService(answer.ID, "answer")
		if err != nil {
			return answerComments, schema.Pagination{}, err
		}

		answerComments = append(answerComments, AnswerResponseDTO{
			ID:        answer.ID,
			UserID:    answer.UserId,
			CommentID: answer.CommentId,
			Text:      answer.Text,
			Likes:     reactions["likes"],
			Dislikes:  reactions["dislikes"],
			CreatedAt: answer.CreatedAt.String(),
			UpdatedAt: answer.UpdatedAt.String(),
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

	return answerComments, pagination, nil
}

func GetAnswerByIdService(id uint) (answerDTO AnswerResponseDTO, err error) {

	answer, err := GetAnswerByIdRepository(id)
	if err != nil {
		return answerDTO, err
	}

	answerDTO = AnswerResponseDTO{
		ID:        answer.ID,
		UserID:    answer.UserId,
		CommentID: answer.CommentId,
		Text:      answer.Text,
		CreatedAt: answer.CreatedAt.String(),
		UpdatedAt: answer.UpdatedAt.String(),
	}

	return answerDTO, nil
}

func CreateAnswerService(answer AnswerCreateDTO, userId uint) (id uint, err error) {
	answerSchema := schema.Answer{UserId: userId, CommentId: answer.CommentID, Text: answer.Text}

	id, err = CreateAnswerByIdRepository(answerSchema)
	if err != nil {
		return id, err
	}

	return id, nil
}

func UpdateAnswerService(id uint, answer AnswerUpdateDTO) (err error) {
	answerSchema := schema.Answer{Text: answer.Text}

	err = UpdateAnswerRepository(id, answerSchema)
	if err != nil {
		return err
	}

	return nil
}

func DeleteAnswerService(id uint) (err error) {
	err = DeleteAnswerRepository(id)
	if err != nil {
		return err
	}

	return nil
}
