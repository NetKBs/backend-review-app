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
func GetAllReactionsServices() ([]ReactionResponseDTO, error) {
	reviewReactions := []ReactionResponseDTO{}

	reactions, err := GetAllReactionsRepository()
	if err != nil {
		return nil, err
	}

	for _, reaction := range reactions {

		reviewReactions = append(reviewReactions, ReactionResponseDTO{
			ID:           reaction.ID,
			UserId:       reaction.UserId,
			ContentId:    reaction.ContentId,
			ContentType:  reaction.ContentType,
			ReactionType: reaction.ReactionType,
			CreatedAt:    reaction.CreatedAt.String(),
			UpdatedAt:    reaction.UpdatedAt.String()})

	}

	return reviewReactions, nil
}

// GetReactionByUserIdAndContentId obtiene una reacción por UserId y ContentId
func GetReactionByUserIdAndContentIdService(userId uint, contentId string, contentType string) (ReactionResponseDTO, error) {

	reviewReactions := ReactionResponseDTO{}

	reaction, err := GetReactionByUserIdAndContentIdRepository(userId, contentId, contentType)
	if err != nil {
		return ReactionResponseDTO{}, err
	}

	reviewReactions = ReactionResponseDTO{
		ID:           reaction.ID,
		UserId:       userId,
		ContentId:    reaction.ContentId,
		ContentType:  reaction.ContentType,
		ReactionType: reaction.ReactionType,
		CreatedAt:    reaction.CreatedAt.String(),
		UpdatedAt:    reaction.UpdatedAt.String()}

	return reviewReactions, nil
}

// CreateReaction crea una nueva reacción en la base de datos
func CreateReactionServices(reaction ReactionCreateDTO, userId uint) (schema.Reaction, error) {

	reactionSchema := schema.Reaction{UserId: userId, ContentId: reaction.ContentId, ContentType: reaction.ContentType, ReactionType: reaction.ReactionType}

	createdReaction, err := CreateReactionRepository(reactionSchema)
	if err != nil {
		return schema.Reaction{}, err
	}
	return createdReaction, nil
}

// DeleteReactionServices maneja la lógica de negocio para eliminar una reacción
func DeleteReactionServices(userId uint, contentId uint, contentType string) error {
	return DeleteReactionRepository(userId, contentId, contentType)
}
