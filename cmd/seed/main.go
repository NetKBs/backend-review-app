package main

import (
	"log"

	"github.com/NetKBs/backend-reviewapp/cmd/seed/fakeSchema"
	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
	"github.com/brianvoe/gofakeit/v7"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	config.SyncDB()
	db := config.DB

	if err := seedDatabase(db); err != nil {
		log.Fatalf("Error al realizar el seeding: %v", err)
	} else {
		log.Println("Database seeded successfully")
	}

}

func seedDatabase(db *gorm.DB) (error error) {
	gofakeit.Seed(777)

	usersId, error := seedUsers(db)
	if error != nil {
		return error
	}

	placesId, error := seedPlaces(db)
	if error != nil {
		return error
	}

	reviewsId, error := seedReviews(db, usersId, placesId)
	if error != nil {
		return error
	}

	commentsId, error := seedComments(db, usersId, reviewsId)
	if error != nil {
		return error
	}

	answersId, error := seedAnswers(db, usersId, commentsId)
	if error != nil {
		return error
	}

	error = seedReviewImages(db, reviewsId)
	if error != nil {
		return error
	}

	error = seedNotifications(db, usersId)
	if error != nil {
		return error
	}

	error = seedReactions(db, usersId, reviewsId, commentsId, answersId)
	if error != nil {
		return error
	}

	error = seedVisitedPlaces(db, usersId, placesId)
	if error != nil {
		return error
	}

	error = seedFollowersAndFollowings(db, usersId)
	if error != nil {
		return error
	}

	error = seedBookmarks(db, usersId, placesId)
	if error != nil {
		return error
	}

	return nil
}

func seedUsers(db *gorm.DB) (usersId []uint, error error) {
	avatarUrl := "https://picsum.photos/200/200"
	userPassword, err := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)
	if err != nil {
		return usersId, err
	}

	for i := 0; i < 10; i++ {
		userFake := fakeSchema.UserFaker{}
		if err := gofakeit.Struct(&userFake); err != nil {
			return usersId, err
		}

		user := schema.User{
			Username:    userFake.Username,
			AvatarUrl:   &avatarUrl,
			DisplayName: userFake.DisplayName,
			Email:       userFake.Email,
			Password:    string(userPassword),
		}

		if err := db.Create(&user).Error; err != nil {
			return usersId, err
		}
		usersId = append(usersId, user.ID)
	}

	return usersId, nil
}

func seedPlaces(db *gorm.DB) (placesId []uint, error error) {
	mapsId := [10]string{
		"512f8a79b18b5e4fc059481220a886962040f00103f9013ebcd3f501000000920318415354524f4d204c6962726572c3ad61202620436166c3a9",
		"5141b62c5f975e4fc059e87472da75962040f00103f9013fbcd3f50100000092030a4d61786920446f6e6173",
		"514ae2f615b55e4fc05992b6c02819962040f00103f901a90452e0010000009203074f76656a697461",
		"51a65e5c4f195f4fc059101c441a92952040f00103f901c5f351e0010000009203074469676974656c",
		"518cb73aa5285f4fc059f7d4b9ac9a952040f00103f9011f6652e0010000009203124d63446f6e616c6427732048656c61646f73",
		"518fbc7328435f4fc059cce7f57b84952040f00103f9015fb11d300100000092030c43696e657320556e69646f73",
		"51b6a7f5c8305f4fc0592eb16d00b3952040f00103f901143c52e0010000009203084d6f766973746172",
		"5136e5af35005f4fc05926df4ca0aa952040f00103f901cf5752e0010000009203054e5926436f",
		"5136a8a21ddc5e4fc059375cc4cbf5952040f00103f9011d5d52e00100000092030fc3937074696361204361726f6ec3ad",
		"51f34d3f4dd25e4fc0594863abc403962040f00103f9012894c390010000009203094d657263616e74696c",
	}

	for _, mapsId := range mapsId {
		place := schema.Place{
			MapsId: mapsId,
		}
		if err := db.Create(&place).Error; err != nil {
			return placesId, err
		}
		placesId = append(placesId, place.ID)
	}

	return placesId, nil
}

func seedReviews(db *gorm.DB, placesId []uint, usersId []uint) (reviewsId []uint, error error) {

	for i := 0; i < 30; i++ {
		reviewFake := fakeSchema.ReviewFaker{}
		if err := gofakeit.Struct(&reviewFake); err != nil {
			return reviewsId, err
		}

		review := schema.Review{
			UserId:  usersId[i%len(usersId)],
			PlaceId: placesId[i%len(placesId)],
			Rate:    reviewFake.Rate,
			Text:    reviewFake.Text,
		}
		if err := db.Create(&review).Error; err != nil {
			return reviewsId, err
		}
		reviewsId = append(reviewsId, review.ID)
	}

	return reviewsId, nil
}

func seedReviewImages(db *gorm.DB, reviewsId []uint) (err error) {
	for _, reviewId := range reviewsId {
		reviewImages := []schema.ReviewImage{
			{ReviewId: reviewId, ImageURL: "https://picsum.photos/600/400"},
			{ReviewId: reviewId, ImageURL: "https://picsum.photos/600/400"},
		}
		if err := db.Create(&reviewImages).Error; err != nil {
			return err
		}
	}
	return nil
}

func seedComments(db *gorm.DB, usersId []uint, reviewsId []uint) (commentsId []uint, err error) {

	for i := 0; i < 30; i++ {
		commentFake := fakeSchema.CommentFaker{}
		if err := gofakeit.Struct(&commentFake); err != nil {
			return commentsId, err
		}

		comment := schema.Comment{
			UserId:   usersId[i%len(usersId)],
			ReviewId: reviewsId[i%len(reviewsId)],
			Text:     commentFake.Text,
		}

		if err := db.Create(&comment).Error; err != nil {
			return commentsId, err
		}
		commentsId = append(commentsId, comment.ID)
	}

	return commentsId, nil
}

func seedAnswers(db *gorm.DB, usersId []uint, commentsId []uint) (answersId []uint, err error) {

	for i := 0; i < 30; i++ {
		answerFake := fakeSchema.AnswerFaker{}
		if err := gofakeit.Struct(&answerFake); err != nil {
			return answersId, err
		}

		answer := schema.Answer{
			UserId:    usersId[i%len(usersId)],
			CommentId: commentsId[i%len(commentsId)],
			Text:      answerFake.Text,
		}
		if err := db.Create(&answer).Error; err != nil {
			return answersId, err
		}
		answersId = append(answersId, answer.ID)
	}

	return answersId, nil
}

func seedNotifications(db *gorm.DB, usersId []uint) (err error) {

	for _, userId := range usersId {
		notificationFake := fakeSchema.NotificationFaker{}
		if err := gofakeit.Struct(&notificationFake); err != nil {
			return err
		}

		notification := schema.Notification{
			UserId: userId,
			Text:   notificationFake.Text,
		}

		if err := db.Create(&notification).Error; err != nil {
			return err
		}
	}

	return nil
}

func seedReactions(db *gorm.DB, usersId []uint, reviewsId []uint, commentsId []uint, answersId []uint) (err error) {
	for i, reviewId := range reviewsId {
		reactionFake := fakeSchema.ReactionFaker{}
		if err := gofakeit.Struct(&reactionFake); err != nil {
			return err
		}

		reaction := schema.Reaction{
			UserId:       usersId[i%len(usersId)],
			ContentId:    reviewId,
			ContentType:  "review",
			ReactionType: reactionFake.ReactionType,
		}

		if err := db.Create(&reaction).Error; err != nil {
			return err
		}
	}

	for i, commentId := range commentsId {
		reactionFake := fakeSchema.ReactionFaker{}
		if err := gofakeit.Struct(&reactionFake); err != nil {
			return err
		}

		reaction := schema.Reaction{
			UserId:       usersId[i%len(usersId)],
			ContentId:    commentId,
			ContentType:  "comment",
			ReactionType: reactionFake.ReactionType,
		}

		if err := db.Create(&reaction).Error; err != nil {
			return err
		}
	}

	for i, answerId := range answersId {
		reactionFake := fakeSchema.ReactionFaker{}
		if err := gofakeit.Struct(&reactionFake); err != nil {
			return err
		}

		reaction := schema.Reaction{
			UserId:       usersId[i%len(usersId)],
			ContentId:    answerId,
			ContentType:  "answer",
			ReactionType: reactionFake.ReactionType,
		}

		if err := db.Create(&reaction).Error; err != nil {
			return err
		}
	}

	return nil
}

func seedVisitedPlaces(db *gorm.DB, usersId []uint, placesId []uint) (err error) {
	for _, userId := range usersId {
		for i := 0; i < gofakeit.Number(1, len(placesId))-1; i++ {

			var place schema.Place
			if err := db.Find(&place, placesId[i]).Error; err != nil {
				return err
			}
			var user schema.User
			if err := db.Find(&user, userId).Error; err != nil {
				return err
			}

			if err := db.Model(&user).Association("VisitedPlaces").Append(&place); err != nil {
				return err
			}
		}
	}

	return nil
}

func seedFollowersAndFollowings(db *gorm.DB, usersId []uint) (err error) {
	for i := 0; i < len(usersId); i++ {
		var user1 schema.User
		if err := db.Find(&user1, usersId[i]).Error; err != nil {
			return err
		}

		for j := 0; j < gofakeit.Number(0, len(usersId)-1); j++ {
			var user2 schema.User
			if err := db.Find(&user2, usersId[j]).Error; err != nil {
				return err
			}

			if err := db.Model(&user1).Association("Followers").Append(&user2); err != nil {
				return err
			}
			if err := db.Model(&user1).Association("Following").Append(&user2); err != nil {
				return err
			}
		}
	}

	return nil
}

func seedBookmarks(db *gorm.DB, usersId []uint, placesId []uint) (err error) {
	for i := 0; i < len(usersId); i++ {
		var user schema.User
		if err := db.Find(&user, usersId[i]).Error; err != nil {
			return err
		}

		for j := 0; j < gofakeit.Number(0, len(placesId)-1); j++ {
			var place schema.Place
			if err := db.Find(&place, placesId[j]).Error; err != nil {
				return err
			}

			if err := db.Model(&user).Association("BookmarkedPlaces").Append(&place); err != nil {
				return err
			}
		}
	}

	return nil
}
