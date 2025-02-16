package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Структуры из entities
type User struct {
	ID         string
	FirstName  string
	SecondName string
	Birthdate  string
	Biography  string
	City       string
	Password   string
}

type Post struct {
	ID           string
	Text         string
	AuthorUserID string
	CreatedAt    time.Time
}

type Friend struct {
	UserID    string
	FriendID  string
	CreatedAt time.Time
}

func main() {
	db, err := gorm.Open(postgres.Open("postgres://social:social@localhost:5432/social"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Генерация пользователей
	users := generateUsers(1000)
	for _, user := range users {
		db.Create(&user)
	}

	// Генерация постов
	posts := generatePosts(users, 5000)
	for _, post := range posts {
		db.Create(&post)
	}

	// Генерация друзей
	friends := generateFriends(users, 3000)
	for _, friend := range friends {
		db.Create(&friend)
	}
}

func generateUsers(count int) []User {
	firstNames := []string{"Иван", "Петр", "Анна", "Мария", "Александр"}
	lastNames := []string{"Иванов", "Петров", "Сидоров", "Смирнов", "Кузнецов"}
	cities := []string{"Москва", "Санкт-Петербург", "Казань", "Новосибирск"}

	users := make([]User, count)
	for i := 0; i < count; i++ {
		users[i] = User{
			ID:         uuid.New().String(),
			FirstName:  firstNames[rand.Intn(len(firstNames))],
			SecondName: lastNames[rand.Intn(len(lastNames))],
			Birthdate:  randomDate(1970, 2000),
			Biography:  fmt.Sprintf("Bio of user %d", i),
			City:       cities[rand.Intn(len(cities))],
			Password:   "$2a$14$ajq8Q7fbtFRQvXpdCq7Jcuy.Rx1h/L4J60Otx.gyNLbAYctGMJ9tK", // хэш пароля "123456"
		}
	}
	return users
}

func generatePosts(users []User, count int) []Post {
	posts := make([]Post, count)
	for i := 0; i < count; i++ {
		posts[i] = Post{
			ID:           uuid.New().String(),
			Text:         fmt.Sprintf("Post content %d", i),
			AuthorUserID: users[rand.Intn(len(users))].ID,
			CreatedAt:    time.Now().Add(-time.Duration(rand.Intn(100)) * 24 * time.Hour),
		}
	}
	return posts
}

func generateFriends(users []User, count int) []Friend {
	friends := make([]Friend, count)
	for i := 0; i < count; i++ {
		user1 := users[rand.Intn(len(users))]
		user2 := users[rand.Intn(len(users))]
		if user1.ID != user2.ID {
			friends[i] = Friend{
				UserID:    user1.ID,
				FriendID:  user2.ID,
				CreatedAt: time.Now(),
			}
		}
	}
	return friends
}

func randomDate(minYear, maxYear int) string {
	min := time.Date(minYear, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(maxYear, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0).Format("2006-01-02")
}
