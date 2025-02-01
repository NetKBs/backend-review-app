package reaction

import (
	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func GetReactionsCountService(id uint, contentType string) (reactionsCount map[string]uint, err error) {
	reactionsCount, err = GetReactionsCountRepository(id, contentType)
	if err != nil {
		return reactionsCount, err
	}

	return reactionsCount, nil
}

// GetAllReactions obtiene todas las reacciones de la base de datos
func GetAllReactionsServices() ([]schema.Reaction, error) {
	reactions, err := GetAllReactionsRepository()
	if err != nil {
		return nil, err
	}
	return reactions, nil
}

// GetReactionByUserIdAndContentId obtiene una reacción por UserId y ContentId
func GetReactionByUserIdAndContentIdService(userId, contentId string) (schema.Reaction, error) {
	reaction, err := GetReactionByUserIdAndContentIdRepository(userId, contentId)
	if err != nil {
		return schema.Reaction{}, err
	}
	return reaction, nil
}

// CreateReaction crea una nueva reacción en la base de datos
func CreateReactionServices(reaction schema.Reaction) (schema.Reaction, error) {
	createdReaction, err := CreateReactionRepository(reaction)
	if err != nil {
		return schema.Reaction{}, err
	}
	return createdReaction, nil
}
