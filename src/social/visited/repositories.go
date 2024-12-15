package visited

import (
	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
)

func GetVisitedCountRepository(userId uint) (visitedCount uint, err error) {
	db := config.DB
	var user schema.User

	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		return visitedCount, err
	}

	visitedCount = uint(db.Model(&user).Association("VisitedPlaces").Count())
	return visitedCount, nil
}
