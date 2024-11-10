package main

import (
	"log"

	"github.com/NetKBs/backend-reviewapp/config"
	"github.com/NetKBs/backend-reviewapp/src/schema"
	"gorm.io/gorm"
)

func main() {
	config.ConnectDB()
	config.SyncDB()
	db := config.DB

	if err := seedDabase(db); err != nil {
		log.Fatalf("Error al realizar el seeding: %v", err)
	} else {
		log.Println("Database seeded successfully")
	}

}

func seedDabase(db *gorm.DB) error {
	var error error

	// Users
	user1 := schema.User{Username: "john_doe", DisplayName: "John Doe", Email: "john@example.com", Password: "securepassword"}
	user2 := schema.User{Username: "jane_smith", DisplayName: "Jane Smith", Email: "jane@example.com", Password: "anothersecurepassword"}
	user3 := schema.User{Username: "alice_jones", DisplayName: "Alice Jones", Email: "alice@example.com", Password: "yetanotherpassword"}
	users := []*schema.User{&user1, &user2, &user3}
	if error = db.Create(&users).Error; error != nil {
		return error
	}

	// Places
	place1 := schema.Place{MapsId: "ChIJRyfoDNb7y40RMHzJsFJjTHQ"}
	place2 := schema.Place{MapsId: "ChIJHwi2ed37y40RN-_sO1osSEg"}
	place3 := schema.Place{MapsId: "ChIJy8sZ-vH6y40RZaRS6OxTPh8"}
	places := []*schema.Place{&place1, &place2, &place3}
	if error = db.Create(&places).Error; error != nil {
		return error
	}

	// Reviews
	review1 := schema.Review{UserId: user1.ID, PlaceId: place1.ID, Text: "Great place to visit!"}
	review2 := schema.Review{UserId: user2.ID, PlaceId: place2.ID, Text: "Not worth the hype."}
	review3 := schema.Review{UserId: user3.ID, PlaceId: place3.ID, Text: "Amazing experience!"}
	reviews := []*schema.Review{&review1, &review2, &review3}
	if error = db.Create(&reviews).Error; error != nil {
		return error
	}

	// Review Images
	reviewImage1 := schema.ReviewImage{ReviewId: review1.ID, ImageURL: "https://picsum.photos/600/400"}
	reviewImage2 := schema.ReviewImage{ReviewId: review2.ID, ImageURL: "https://picsum.photos/600/400"}
	reviewImage3 := schema.ReviewImage{ReviewId: review3.ID, ImageURL: "https://picsum.photos/600/400"}
	reviewImages := []*schema.ReviewImage{&reviewImage1, &reviewImage2, &reviewImage3}
	if error = db.Create(&reviewImages).Error; error != nil {
		return error
	}

	// Comments
	comment1 := schema.Comment{UserId: user2.ID, ReviewId: review1.ID, Text: "I totally agree!"}
	comment2 := schema.Comment{UserId: user3.ID, ReviewId: review2.ID, Text: "I had a different experience."}
	comments := []*schema.Comment{&comment1, &comment2}
	if error = db.Create(&comments).Error; error != nil {
		return error
	}

	// Answers
	answer1 := schema.Answer{UserId: user1.ID, CommentId: comment1.ID, Text: "Thanks for your input!"}
	answer2 := schema.Answer{UserId: user2.ID, CommentId: comment2.ID, Text: "Interesting perspective."}
	answers := []*schema.Answer{&answer1, &answer2}
	if error = db.Create(&answers).Error; error != nil {
		return error
	}

	// Notifications
	notification1 := schema.Notification{UserId: user1.ID, Text: "You have a new follower!"}
	notification2 := schema.Notification{UserId: user2.ID, Text: "Your review received a comment."}
	notifications := []*schema.Notification{&notification1, &notification2}
	if error = db.Create(&notifications).Error; error != nil {
		return error
	}

	// Reactions
	reaction1 := schema.Reaction{UserId: user1.ID, ContentId: review1.ID, ContentType: "review", ReactionType: true}
	reaction2 := schema.Reaction{UserId: user2.ID, ContentId: review2.ID, ContentType: "review", ReactionType: false}
	reaction3 := schema.Reaction{UserId: user3.ID, ContentId: review3.ID, ContentType: "comment", ReactionType: true}
	reaction4 := schema.Reaction{UserId: user3.ID, ContentId: review3.ID, ContentType: "answer", ReactionType: false}
	reactions := []*schema.Reaction{&reaction1, &reaction2, &reaction3, &reaction4}
	if error = db.Create(&reactions).Error; error != nil {
		return error
	}

	// Follows
	if err := db.Model(&user1).Association("Following").Append(&user2, &user3); err != nil {
		return err
	}
	if err := db.Model(&user2).Association("Followers").Append(&user1); err != nil {
		return err
	}
	if err := db.Model(&user3).Association("Followers").Append(&user1); err != nil {
		return err
	}

	// Visited Places
	if err := db.Model(&user1).Association("VisitedPlaces").Append(&place1, &place2); err != nil {
		return err
	}
	if err := db.Model(&user2).Association("VisitedPlaces").Append(&place2, &place3); err != nil {
		return err
	}
	if err := db.Model(&user3).Association("VisitedPlaces").Append(&place1, &place3); err != nil {
		return err
	}

	// Bookmarks
	if err := db.Model(&user1).Association("BookmarkedPlaces").Append(&place1); err != nil {
		return err
	}
	if err := db.Model(&user2).Association("BookmarkedPlaces").Append(&place2); err != nil {
		return err
	}
	if err := db.Model(&user3).Association("BookmarkedPlaces").Append(&place3); err != nil {
		return err
	}

	return nil
}
