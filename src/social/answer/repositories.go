package answer

import (
	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func GetCountAnswersByCommentIdRepository(id uint) (uint, error) {
	db := config.DB
	var count int64
	if err := db.Model(&schema.Answer{}).Where("`comment_id` = ?", id).Count(&count).Error; err != nil {
		return uint(count), err
	}

	return uint(count), nil
}

func GetAnswersByCommentIdRepository(id uint, limit int, cursor uint) (commentAnswers []schema.Answer, err error) {
	db := config.DB

	query := db.Where("comment_id = ?", id).Order("id DESC").Limit(limit)
	if cursor != 0 {
		query = query.Where("id < ?", cursor)
	}

	if err = query.Find(&commentAnswers).Error; err != nil {
		return commentAnswers, err
	}
	return commentAnswers, nil
}

func GetAnswerLikesByIdRepository(id uint64, limit int, lastID uint) ([]schema.Reaction, error) {
	db := config.DB
	likes := []schema.Reaction{}

	query := db.Table("reactions").
		Select("*").
		Where("content_type = 'answer' AND reaction_type = 1 AND content_id = ?", id).
		Order("id DESC").
		Limit(limit)

	if lastID != 0 {
		query = query.Where("id < ?", lastID)
	}

	if err := query.Scan(&likes).Error; err != nil {
		return likes, err
	}

	return likes, nil
}

func GetAnswerDislikesByIdRepository(id uint64, limit int, lastID uint) ([]schema.Reaction, error) {
	db := config.DB
	dislikes := []schema.Reaction{}

	query := db.Table("reactions").
		Select("*").
		Where("content_type = 'answer' AND reaction_type = 0 AND content_id = ?", id).
		Order("id DESC").
		Limit(limit)

	if lastID != 0 {
		query = query.Where("id < ?", lastID)
	}

	if err := query.Scan(&dislikes).Error; err != nil {
		return dislikes, err
	}

	return dislikes, nil
}

func GetAnswerByIdRepository(id uint) (answer schema.Answer, err error) {
	db := config.DB

	if err = db.Where("id = ?", id).First(&answer).Error; err != nil {
		return answer, err
	}
	return answer, nil
}

func CreateAnswerByIdRepository(answer schema.Answer) (id uint, err error) {
	db := config.DB

	if err = db.Create(&answer).Error; err != nil {
		return id, err
	}
	return answer.ID, nil
}

func UpdateAnswerRepository(id uint, answer schema.Answer) (err error) {
	db := config.DB

	if err = db.Where("id = ?", id).First(&schema.Answer{}).Error; err != nil {
		return err
	}

	if err = db.Model(&schema.Answer{}).Where("id = ?", id).Update("Text", answer.Text).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAnswerRepository(id uint) (err error) {
	db := config.DB

	if err = db.Where("id = ?", id).First(&schema.Answer{}).Error; err != nil {
		return err
	}

	if err = db.Where("id = ?", id).Delete(&schema.Answer{}).Error; err != nil {
		return err
	}
	return nil
}
