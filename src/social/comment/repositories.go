package comment

import (
	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func GetCommentsReviewCountRepository(id uint) (commentsCount uint, err error) {
	db := config.DB
	var _commentsCount int64

	if err = db.Model(&schema.Comment{}).Where("review_id = ?", id).Count(&_commentsCount).Error; err != nil {
		return uint(_commentsCount), err
	}

	return uint(_commentsCount), nil
}

func GetCommentsByIdReviewRepository(id uint) (reviewComments []schema.Comment, err error) {
	db := config.DB

	if err = db.Table("comments").Where("review_id = ?", id).Find(&reviewComments).Error; err != nil {
		return reviewComments, err
	}
	return reviewComments, nil
}

func GetCommentsByIdRepository(id uint) (comment schema.Comment, err error) {
	db := config.DB

	if err = db.Where("id = ?", id).First(&comment).Error; err != nil {
		return comment, err
	}
	return comment, nil
}

func CreateCommentRepository(comment schema.Comment) (id uint, err error) {
	db := config.DB

	if err = db.Create(&comment).Error; err != nil {
		return id, err
	}
	return comment.ID, nil
}

func UpdateCommentRepository(id uint, comment schema.Comment) (err error) {
	db := config.DB

	if err = db.Where("id = ?", id).First(&schema.Comment{}).Error; err != nil {
		return err
	}

	if err = db.Model(&schema.Comment{}).Where("id = ?", id).Update("Text", comment.Text).Error; err != nil {
		return err
	}
	return nil
}

func DeleteCommentRepository(id uint) (err error) {
	db := config.DB

	if err = db.Where("id = ?", id).First(&schema.Comment{}).Error; err != nil {
		return err
	}

	if err = db.Where("id = ?", id).Delete(&schema.Comment{}).Error; err != nil {
		return err
	}
	return nil
}
