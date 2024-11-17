package reaction

import (
	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func GetReactionsCountRepository(id uint, contentType string) (reactionsCount map[string]uint, err error) {
	db := config.DB
	var likesCount, dislikesCount int64

	if err = db.Model(&schema.Reaction{}).Where("content_id = ? AND content_type = ? AND reaction_type = true", id, contentType).Count(&likesCount).Error; err != nil {
		return map[string]uint{"likes": 0, "dislikes": 0}, err
	}
	if err = db.Model(&schema.Reaction{}).Where("content_id = ? AND content_type = ? AND reaction_type = false", id, contentType).Count(&dislikesCount).Error; err != nil {
		return map[string]uint{"likes": 0, "dislikes": 0}, err
	}

	return map[string]uint{"likes": uint(likesCount), "dislikes": uint(dislikesCount)}, nil
}
