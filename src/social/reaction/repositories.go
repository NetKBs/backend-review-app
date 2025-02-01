package reaction

import (
	"errors"

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

// GetAllReactions obtiene todas las reacciones de la base de datos
func GetAllReactionsRepository() ([]schema.Reaction, error) {
	db := config.DB
	var reactions []schema.Reaction
	if err := db.Find(&reactions).Error; err != nil {
		return nil, err
	}
	return reactions, nil
}

// GetReactionByUserIdAndContentId obtiene una reacción por UserId y ContentId
func GetReactionByUserIdAndContentIdRepository(userId, contentId string) (schema.Reaction, error) {
	var reaction schema.Reaction
	db := config.DB
	if err := db.Where("user_id = ? AND content_id = ?", userId, contentId).First(&reaction).Error; err != nil {
		return schema.Reaction{}, errors.New("reaction not found")
	}
	return reaction, nil
}

// CreateReaction crea una nueva reacción en la base de datos
func CreateReactionRepository(reaction schema.Reaction) (schema.Reaction, error) {
	db := config.DB
	if err := db.Create(&reaction).Error; err != nil {
		return schema.Reaction{}, err
	}
	return reaction, nil
}
