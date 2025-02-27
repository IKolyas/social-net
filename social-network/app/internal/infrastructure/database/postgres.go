package database

import (
	"fmt"
	"social_network/internal/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresMaster struct {
	db *gorm.DB
}

type PostgresSlave struct {
	db *gorm.DB
}

func NewPostgresMaster(dsn string) (*PostgresMaster, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &PostgresMaster{db: db}, nil
}

func NewPostgresSlave(dsn string) (*PostgresSlave, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &PostgresSlave{db: db}, nil
}

// Методы для записи (используют Master)
func (m *PostgresMaster) CreatePost(post *entities.Post) error {
	return m.db.Create(post).Error
}

func (m *PostgresMaster) UpdatePost(post *entities.Post) error {
	return m.db.Save(post).Error
}

func (m *PostgresMaster) DeletePost(postID string) error {
	return m.db.Delete(&entities.Post{}, postID).Error
}

func (m *PostgresMaster) CreateUser(user *entities.User) error {
	return m.db.Create(user).Error
}

func (m *PostgresMaster) AddFriend(userID, friendID string) error {
	friend := entities.Friend{UserID: userID, FriendID: friendID}
	return m.db.Create(&friend).Error
}

func (m *PostgresMaster) DeleteFriend(userID, friendID string) error {
	return m.db.Where("user_id = $1 AND friend_id = $2", userID, friendID).Delete(&entities.Friend{}).Error
}

// Методы для чтения (используют Slave)
func (s *PostgresSlave) GetFeed(userID string, offset, limit int) ([]*entities.Post, error) {
	var posts []*entities.Post
	err := s.db.Where("author_user_id IN (SELECT friend_id FROM friends WHERE user_id = $1)", userID).
		Offset(offset).Limit(limit).Find(&posts).Error
	return posts, err
}

func (s *PostgresSlave) GetPostByID(postID string) (*entities.Post, error) {
	var post entities.Post
	err := s.db.Where("id = $1", postID).First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil

}

func (s *PostgresSlave) GetUserByID(userID string) (*entities.User, error) {
	var user entities.User
	err := s.db.Where("id = $1", userID).First(&user).Error
	if err != nil {
		fmt.Println("Error getting user by ID:", err)
		return nil, err
	}
	return &user, nil
}

func (s *PostgresSlave) SearchUsers(query string) ([]*entities.User, error) {
	var users []*entities.User
	err := s.db.Where("first_name LIKE $1 OR second_name LIKE $2", query+"%", query+"%").
		Find(&users).Error
	return users, err
}
