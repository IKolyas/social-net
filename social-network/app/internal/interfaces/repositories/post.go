package repositories

import (
	"social_network/internal/entities"
	"social_network/internal/infrastructure/database"
)

type PostRepository interface {
	Create(post *entities.Post) error
	Update(post *entities.Post) error
	Delete(id string) error
	GetByID(id string) (*entities.Post, error)
	GetFeed(userID string, offset, limit int) ([]*entities.Post, error)
}

type PostRepositoryImpl struct {
	db *database.PostgresRepository
}

func NewPostRepository(db *database.PostgresRepository) *PostRepositoryImpl {
	return &PostRepositoryImpl{db: db}
}

func (r *PostRepositoryImpl) Create(post *entities.Post) error {
	return r.db.CreatePost(post)
}

func (r *PostRepositoryImpl) Update(post *entities.Post) error {
	return r.db.UpdatePost(post)
}

func (r *PostRepositoryImpl) Delete(id string) error {
	return r.db.DeletePost(id)
}

func (r *PostRepositoryImpl) GetByID(id string) (*entities.Post, error) {
	return r.db.GetPostByID(id)
}

func (r *PostRepositoryImpl) GetFeed(userID string, offset, limit int) ([]*entities.Post, error) {
	return r.db.GetFeed(userID, offset, limit)
}
