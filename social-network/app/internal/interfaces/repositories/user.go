package repositories

import (
	"social_network/internal/entities"
	"social_network/internal/infrastructure/database"
)

type UserRepository interface {
	Create(user *entities.User) error
	GetByID(id string) (*entities.User, error)
	Search(query string) ([]*entities.User, error)
}

type UserRepositoryImpl struct {
	db *database.PostgresRepository
}

func NewUserRepository(db *database.PostgresRepository) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(user *entities.User) error {
	return r.db.CreateUser(user)
}

func (r *UserRepositoryImpl) GetByID(id string) (*entities.User, error) {
	return r.db.GetUserByID(id)
}

func (r *UserRepositoryImpl) Search(query string) ([]*entities.User, error) {
	return r.db.SearchUsers(query)
}
