package reaction

import (
	"errors"
	"fmt"

	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
	"gorm.io/gorm"
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
func GetReactionByUserIdAndContentIdRepository(userId uint, contentId string, contentType string) (schema.Reaction, error) {
	var reaction schema.Reaction
	db := config.DB
	if err := db.Where("user_id = ? AND content_id = ? AND content_type = ?", userId, contentId, contentType).First(&reaction).Error; err != nil {
		return schema.Reaction{}, errors.New("reaction not found")
	}
	return reaction, nil
}

func CreateReactionRepository(reaction schema.Reaction) (schema.Reaction, error) {
	db := config.DB

	// Validar si ya existe una reacción con el mismo UserId, ContentId y ContentType
	var existingReaction schema.Reaction
	err := db.Where("user_id = ? AND content_id = ? AND content_type = ?",
		reaction.UserId, reaction.ContentId, reaction.ContentType).First(&existingReaction).Error

	if err == nil {
		// Si no es nil, significa que ya existe una reacción
		return schema.Reaction{}, fmt.Errorf("ya existe una reacción para este usuario y contenido con el tipo %s", reaction.ContentType)
	}

	if err != gorm.ErrRecordNotFound {
		// Si hay un error diferente a "no encontrado", lo devolvemos
		return schema.Reaction{}, err
	}

	// Si no existe, proceder a crear la nueva reacción
	if err := db.Create(&reaction).Error; err != nil {
		return schema.Reaction{}, err
	}
	return reaction, nil
}

// DeleteReactionRepository elimina una reacción de la base de datos
func DeleteReactionRepository(userId uint, contentId uint, contentType string) error {
	db := config.DB
	if err := db.Where("user_id = ? AND content_id = ? AND content_type = ?", userId, contentId, contentType).Delete(&schema.Reaction{}).Error; err != nil {
		return err
	}
	return nil
}
