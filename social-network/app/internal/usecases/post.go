package usecases

import (
	"encoding/json"
	"fmt"
	"social_network/internal/entities"
	"social_network/internal/infrastructure/logger"
	"social_network/internal/interfaces/repositories"
	"time"

	"github.com/google/uuid"
)

type PostUseCase struct {
	postRepo repositories.PostRepository
	cache    repositories.CacheRepository
}

func NewPostUseCase(postRepo repositories.PostRepository, cache repositories.CacheRepository) *PostUseCase {
	return &PostUseCase{postRepo: postRepo, cache: cache}
}

func (uc *PostUseCase) CreatePost(post *entities.Post) error {
	post.ID = uuid.New().String()
	return uc.postRepo.Create(post)
}

func (uc *PostUseCase) UpdatePost(post *entities.Post) error {
	return uc.postRepo.Update(post)
}

func (uc *PostUseCase) DeletePost(id string) error {
	return uc.postRepo.Delete(id)
}

func (uc *PostUseCase) GetPost(id string) (*entities.Post, error) {
	return uc.postRepo.GetByID(id)
}

func (uc *PostUseCase) GetFeed(userID string, offset, limit int) ([]*entities.Post, error) {
	cacheKey := fmt.Sprintf("feed_%s_%d_%d", userID, offset, limit)

	// Пытаемся получить данные из кэша
	if cachedData, err := uc.cache.Get(cacheKey); err == nil {
		var posts []*entities.Post
		if err := json.Unmarshal([]byte(cachedData.(string)), &posts); err == nil {
			return posts, nil
		}
	}

	// Если в кэше нет, получаем из БД
	posts, err := uc.postRepo.GetFeed(userID, offset, limit)
	if err != nil {
		return nil, err
	}

	// Перед сохранением в кэш преобразовать посты в JSON
	postsJSON, err := json.Marshal(posts)
	if err != nil {
		return nil, err
	}

	// Кэшируем результат
	err = uc.cache.Set(cacheKey, postsJSON, time.Hour)
	if err != nil {
		logger.ErrorLogger.Fatal(err.Error())
		return nil, err
	}
	return posts, nil
}
