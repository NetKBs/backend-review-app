package reaction

import (
	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func GetReactionsCountRepository(id uint) (reactionsCount map[string]uint, err error) {
	db := config.DB
	var likesCount, dislikesCount int64

	if err = db.Model(&schema.Reaction{}).Where("content_id = ? AND content_type = 'review' AND reaction_type = true", id).Count(&likesCount).Error; err != nil {
		return map[string]uint{"likes": 0, "dislikes": 0}, err
	}
	if err = db.Model(&schema.Reaction{}).Where("content_id = ? AND content_type = 'review' AND reaction_type = false", id).Count(&dislikesCount).Error; err != nil {
		return map[string]uint{"likes": 0, "dislikes": 0}, err
	}

	return map[string]uint{"likes": uint(likesCount), "dislikes": uint(dislikesCount)}, nil
}
