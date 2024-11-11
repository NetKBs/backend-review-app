package schema

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username         string
	AvatarUrl        *string
	DisplayName      string
	Email            string
	Password         string
	Reviews          []Review       `gorm:"foreignKey:UserId"`
	Comments         []Comment      `gorm:"foreignKey:UserId"`
	Answers          []Answer       `gorm:"foreignKey:UserId"`
	Reactions        []Reaction     `gorm:"foreignKey:UserId"`
	Notifications    []Notification `gorm:"foreignKey:UserId"`
	Followers        []User         `gorm:"many2many:follow;joinForeignKey:FollowerID;joinReferences:FollowedID"`
	Following        []User         `gorm:"many2many:follow;joinForeignKey:FollowedID;joinReferences:FollowerID"`
	VisitedPlaces    []Place        `gorm:"many2many:place_visitors"`
	BookmarkedPlaces []Place        `gorm:"many2many:bookmark"`
}

type Reaction struct {
	gorm.Model
	UserId       uint
	ContentId    uint
	ContentType  string `gorm:"check:content_type IN ('review', 'comment', 'answer')"`
	ReactionType bool
}

type Review struct {
	gorm.Model
	UserId   uint
	PlaceId  uint
	Rate     uint `gorm:"check:rate BETWEEN 1 AND 5"`
	Text     string
	Images   []ReviewImage `gorm:"foreignKey:ReviewId"`
	Comments []Comment     `gorm:"foreignKey:ReviewId"`
}

type ReviewImage struct {
	gorm.Model
	ReviewId uint
	ImageURL string
}

type Place struct {
	gorm.Model
	MapsId  string
	Reviews []Review `gorm:"foreignKey:PlaceId"`
}

type Notification struct {
	gorm.Model
	UserId uint
	Text   string
}

type Comment struct {
	gorm.Model
	UserId   uint
	ReviewId uint
	Text     string
	Answers  []Answer `gorm:"foreignKey:CommentId"`
}

type Answer struct {
	gorm.Model
	UserId    uint
	CommentId uint
	Text      string
}
