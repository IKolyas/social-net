package entities

import "encoding/json"

type Post struct {
	ID           string `json:"id" gorm:"primaryKey"`
	Text         string `json:"text"`
	AuthorUserID string `json:"author_user_id"`
}

// Добавить методы для маршалинга
func (p *Post) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Post) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, p)
}
