package user

type UserResponseDTO struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	AvatarUrl   string `json:"avatar_url"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}

/* Preguntar
type User struct {
	ID          uint       `gorm:"primaryKey"`
	Username    string     `gorm:"uniqueIndex"`
	AvatarUrl   string
	DisplayName string
	Email       string     `gorm:"uniqueIndex"`
	Password    string
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime"`
	DeletedAt   *gorm.DeletedAt
}
*/
