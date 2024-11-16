package comment

import (
	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func GetCommentsCountRepository(id uint) (commentsCount uint, err error) {
	db := config.DB
	var _commentsCount int64

	if err = db.Model(&schema.Comment{}).Where("review_id = ?", id).Count(&_commentsCount).Error; err != nil {
		return uint(_commentsCount), err
	}

	return uint(_commentsCount), nil
}
