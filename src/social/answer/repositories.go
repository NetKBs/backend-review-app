package answer

import (
	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func GetAnswersByCommentIdRepository(id uint) (commentAnswers []schema.Answer, err error) {
	db := config.DB

	if err = db.Table("answers").Where("`comment_id` = ?", id).Find(&commentAnswers).Error; err != nil {
		return commentAnswers, err
	}
	return commentAnswers, nil
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
