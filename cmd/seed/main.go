package main

import (
	"log"

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
		log.Fatalf("Error seeding database: %v", err)
	} else {
		log.Println("Database seeded successfully")
	}

}

func seedDatabase(db *gorm.DB) error {
	log.Println("Seeding database...")
	gofakeit.Seed(777)

	// Places
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
	var places []schema.Place
	for _, mapsId := range mapsId {
		place := schema.Place{
			MapsId: mapsId,
		}
		places = append(places, place)
	}
	if err := db.Create(&places).Error; err != nil {
		return err
	}
	log.Println("Places seeded successfully")

	// User
	avatarUrl := "https://picsum.photos/200/200"
	userPassword, err := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	var users []schema.User
	for i := 0; i < 30; i++ {
		user := schema.User{
			Username:    gofakeit.Username(),
			AvatarUrl:   &avatarUrl,
			DisplayName: gofakeit.Name(),
			Email:       gofakeit.Email(),
			Password:    string(userPassword),
		}
		users = append(users, user)
	}
	if err := db.Create(&users).Error; err != nil {
		return err
	}
	log.Println("Users seeded successfully")

	// Reviews
	var reviews []schema.Review
	for i := 0; i < len(places); i++ {
		for j := 0; j < len(users); j++ {
			review := schema.Review{
				UserId:  users[j].ID,
				PlaceId: places[i].ID,
				Rate:    uint(gofakeit.Number(1, 5)),
				Text:    gofakeit.Sentence(gofakeit.Number(20, 50)),
			}
			reviews = append(reviews, review)
		}
	}
	if err := db.Create(&reviews).Error; err != nil {
		return err
	}
	log.Println("Reviews seeded successfully")

	// Reviews Images
	var reviewImages []schema.ReviewImage
	for i := 0; i < len(reviews); i++ {
		for j := 0; j < 3; j++ {
			reviewImage := schema.ReviewImage{
				ReviewId: reviews[i].ID,
				ImageURL: "https://picsum.photos/600/400",
			}
			reviewImages = append(reviewImages, reviewImage)
		}
	}
	if err := db.Create(&reviewImages).Error; err != nil {
		return err
	}
	log.Println("Reviews Images seeded successfully")

	// Comments
	var comments []schema.Comment
	for i := 0; i < len(reviews); i += 10 {
		for j := 0; j < len(users); j++ {
			comment := schema.Comment{
				UserId:   users[j].ID,
				ReviewId: reviews[i].ID,
				Text:     gofakeit.Sentence(gofakeit.Number(20, 50)),
			}
			comments = append(comments, comment)
		}
	}
	if err := db.Create(&comments).Error; err != nil {
		return err
	}
	log.Println("Comments seeded successfully")

	// Answers
	var answers []schema.Answer
	for i := 0; i < len(comments); i += 30 {
		for j := 0; j < len(users); j++ {
			answer := schema.Answer{
				UserId:    users[j].ID,
				CommentId: comments[i].ID,
				Text:      gofakeit.Sentence(gofakeit.Number(20, 50)),
			}
			answers = append(answers, answer)
		}
	}
	if err := db.Create(&answers).Error; err != nil {
		return err
	}
	log.Println("Answers seeded successfully")

	// Reactions for Reviews, Comments and Answers
	var reactions []schema.Reaction

	for i, review := range reviews {
		reaction := schema.Reaction{
			UserId:       users[i%len(users)].ID,
			ContentId:    review.ID,
			ContentType:  "review",
			ReactionType: gofakeit.Bool(),
		}
		reactions = append(reactions, reaction)
	}

	for i := 0; i < len(comments); i += 10 {
		reaction := schema.Reaction{
			UserId:       users[i%len(users)].ID,
			ContentId:    comments[i].ID,
			ContentType:  "comment",
			ReactionType: gofakeit.Bool(),
		}
		reactions = append(reactions, reaction)
	}

	for i := 0; i < len(answers); i += 10 {
		reaction := schema.Reaction{
			UserId:       users[i%len(users)].ID,
			ContentId:    answers[i].ID,
			ContentType:  "answer",
			ReactionType: gofakeit.Bool(),
		}
		reactions = append(reactions, reaction)
	}

	if err := db.Create(&reactions).Error; err != nil {
		return err
	}
	log.Println("Reactions successfully seeded")

	// Visited places
	for _, user := range users {
		for i := 0; i < len(places); i += 2 {

			var place_model schema.Place
			if err := db.Find(&place_model, places[i].ID).Error; err != nil {
				return err
			}
			var user_model schema.User
			if err := db.Find(&user_model, user.ID).Error; err != nil {
				return err
			}

			if err := db.Model(&user_model).Association("VisitedPlaces").Append(&place_model); err != nil {
				return err
			}
		}
	}
	log.Println("Visited places seeded successfully")

	// Bookmark
	for _, user := range users {
		for i := 0; i < len(places); i += 2 {
			var place_model schema.Place
			if err := db.Find(&place_model, places[i].ID).Error; err != nil {
				return err
			}
			var user_model schema.User
			if err := db.Find(&user_model, user.ID).Error; err != nil {
				return err
			}

			if err := db.Model(&user_model).Association("BookmarkedPlaces").Append(&place_model); err != nil {
				return err
			}
		}
	}
	log.Println("Bookmarks seeded successfully")

	// Followers and followings
	for _, user := range users {
		for i := 0; i < len(users); i += 2 {
			if users[i].ID == user.ID {
				continue
			}
			var user_model schema.User
			if err := db.Find(&user_model, user.ID).Error; err != nil {
				return err
			}
			var user_model2 schema.User
			if err := db.Find(&user_model2, users[i].ID).Error; err != nil {
				return err
			}

			if err := db.Model(&user_model).Association("Following").Append(&user_model2); err != nil {
				return err
			}
			if err := db.Model(&user_model2).Association("Followers").Append(&user_model); err != nil {
				return err
			}
		}
	}

	log.Println("Followers and followings seeded successfully")

	return nil
}
