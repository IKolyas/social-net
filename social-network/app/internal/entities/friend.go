package entities

type Friend struct {
	UserID   string `json:"user_id" gorm:"primaryKey"`
	FriendID string `json:"friend_id"`
}
