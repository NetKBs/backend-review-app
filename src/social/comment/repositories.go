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

func GetCommentsByIdReviewRepository(id uint, limit int, cursor uint) (reviewComments []schema.Comment, err error) {
	db := config.DB

	query := db.Where("review_id = ?", id).Order("id DESC").Limit(limit)
	if cursor != 0 {
		query = query.Where("id < ?", cursor)
	}

	if err = query.Find(&reviewComments).Error; err != nil {
		return reviewComments, err
	}
	return reviewComments, nil
}

func GetCommentLikesByIdRepository(id uint64, limit int, lastID uint) ([]schema.Reaction, error) {
	db := config.DB
	likes := []schema.Reaction{}

	query := db.Table("reactions").
		Select("*").
		Where("content_type = 'comment' AND reaction_type = 1 AND content_id = ?", id).
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

func GetCommentDislikesByIdRepository(id uint64, limit int, lastID uint) ([]schema.Reaction, error) {
	db := config.DB
	dislikes := []schema.Reaction{}

	query := db.Table("reactions").
		Select("*").
		Where("content_type = 'comment' AND reaction_type = 0 AND content_id = ?", id).
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
