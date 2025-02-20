package answer

import (
	"github.com/NetKBs/backend-reviewapp/src/schema"
	"github.com/NetKBs/backend-reviewapp/src/social/reaction"
)

func GetCountAnswersByCommentIdService(id uint) (count uint, err error) {
	return GetCountAnswersByCommentIdRepository(id)
}

func GetAnswersByCommentIdService(id uint) (answerComments []AnswerResponseDTO, err error) {
	anscomments, err := GetAnswersByCommentIdRepository(id)
	if err != nil {
		return answerComments, err
	}

	if len(anscomments) == 0 {
		return []AnswerResponseDTO{}, nil
	}

	reactions, err := reaction.GetReactionsCountService(id, "answer")
	if err != nil {
		return answerComments, err
	}

	for _, anscomment := range anscomments {
		answerComments = append(answerComments, AnswerResponseDTO{
			ID:        anscomment.ID,
			UserID:    anscomment.UserId,
			CommentID: anscomment.CommentId,
			Text:      anscomment.Text,
			Likes:     reactions["likes"],
			Dislikes:  reactions["dislikes"],
			CreatedAt: anscomment.CreatedAt.String(),
			UpdatedAt: anscomment.UpdatedAt.String(),
		})
	}
	return answerComments, nil
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
