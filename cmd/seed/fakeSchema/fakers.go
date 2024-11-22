package fakeSchema

type UserFaker struct {
	Username    string `fake:"{username}"`
	AvatarUrl   string
	DisplayName string `fake:"{name}"`
	Email       string `fake:"{email}"`
	Password    string
}

type ReviewFaker struct {
	UserId  uint
	PlaceId uint
	Rate    uint   `fake:"{number:1,5}"`
	Text    string `fake:"{sentence:50}"`
}

type CommentFaker struct {
	UserId   uint
	ReviewId uint
	Text     string `fake:"{sentence:10}"`
}

type AnswerFaker struct {
	UserId    uint
	CommentId uint
	Text      string `fake:"{sentence:7}"`
}

type NotificationFaker struct {
	UserId uint
	Text   string `fake:"{sentence:7}"`
}

type ReactionFaker struct {
	UserId       uint
	ContentId    uint
	ContentType  string
	ReactionType bool `fake:"{bool}"`
}
