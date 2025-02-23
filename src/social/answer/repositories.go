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

func GetAnswersByCommentIdRepository(id uint, limit int, page int) (commentAnswers []schema.Answer, total int64, err error) {
	db := config.DB
	offset := (page - 1) * limit

	if err = db.Model(&schema.Answer{}).Where("`comment_id` = ?", id).Count(&total).Error; err != nil {
		return commentAnswers, 0, err
	}

	if err = db.Where("comment_id = ?", id).Limit(limit).Offset(offset).Order("created_at DESC").Find(&commentAnswers).Error; err != nil {
		return commentAnswers, 0, err
	}
	return commentAnswers, total, nil
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
