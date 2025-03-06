package answer

import (
	"strconv"

	"github.com/NetKBs/backend-reviewapp/src/schema"
	"github.com/NetKBs/backend-reviewapp/src/social/reaction"
)

func GetCountAnswersByCommentIdService(id uint) (count uint, err error) {
	return GetCountAnswersByCommentIdRepository(id)
}

func GetAnswersByCommentIdService(id uint, limit int, cursor uint) ([]AnswerResponseDTO, string, error) {
	answerComments := []AnswerResponseDTO{}

	answers, err := GetAnswersByCommentIdRepository(id, limit, cursor)
	if err != nil {
		return answerComments, "", err
	}

	for _, answer := range answers {
		reactions, err := reaction.GetReactionsCountService(answer.ID, "answer")
		if err != nil {
			return answerComments, "", err
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

	nextCursor := ""
	if len(answers) > 0 {
		nextCursor = strconv.FormatUint(uint64(answers[len(answers)-1].ID), 10)
	}

	return answerComments, nextCursor, nil
}

func GetAnswerLikesByIdService(id uint64, limit int, cursor string) ([]reaction.ReactionResponseDTO, string, error) {
	answerLikes := []reaction.ReactionResponseDTO{}
	var lastID uint = 0

	if cursor != "" {
		cursorUint, err := strconv.ParseUint(cursor, 10, 64)
		if err != nil {
			return answerLikes, "", err
		}
		lastID = uint(cursorUint)
	}

	likes, err := GetAnswerLikesByIdRepository(id, limit, lastID)
	if err != nil {
		return answerLikes, "", err
	}

	for _, like := range likes {
		likeDTO := reaction.ReactionResponseDTO{
			ID:           like.ID,
			UserId:       like.UserId,
			ContentId:    like.ContentId,
			ContentType:  like.ContentType,
			ReactionType: like.ReactionType,
			CreatedAt:    like.CreatedAt.String(),
			UpdatedAt:    like.UpdatedAt.String(),
		}

		answerLikes = append(answerLikes, likeDTO)
	}

	nextCursor := ""
	if len(likes) > 0 {
		nextCursor = strconv.FormatUint(uint64(likes[len(likes)-1].ID), 10)
	}

	return answerLikes, nextCursor, nil
}

func GetAnswerDislikesByIdService(id uint64, limit int, cursor string) ([]reaction.ReactionResponseDTO, string, error) {
	answerDislikes := []reaction.ReactionResponseDTO{}
	var lastID uint = 0

	if cursor != "" {
		cursorUint, err := strconv.ParseUint(cursor, 10, 64)
		if err != nil {
			return answerDislikes, "", err
		}
		lastID = uint(cursorUint)
	}

	dislikes, err := GetAnswerDislikesByIdRepository(id, limit, lastID)
	if err != nil {
		return answerDislikes, "", err
	}

	for _, dislike := range dislikes {
		dislikeDTO := reaction.ReactionResponseDTO{
			ID:           dislike.ID,
			UserId:       dislike.UserId,
			ContentId:    dislike.ContentId,
			ContentType:  dislike.ContentType,
			ReactionType: dislike.ReactionType,
			CreatedAt:    dislike.CreatedAt.String(),
			UpdatedAt:    dislike.UpdatedAt.String(),
		}

		answerDislikes = append(answerDislikes, dislikeDTO)
	}

	nextCursor := ""
	if len(dislikes) > 0 {
		nextCursor = strconv.FormatUint(uint64(dislikes[len(dislikes)-1].ID), 10)
	}

	return answerDislikes, nextCursor, nil
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
