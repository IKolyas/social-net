package usecases

import (
	"social_network/internal/entities"
	"social_network/internal/interfaces/repositories"

	"github.com/google/uuid"
)

type UserUseCase struct {
	userRepo repositories.UserRepository
}

func NewUserUseCase(userRepo repositories.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

func (uc *UserUseCase) Register(user *entities.User) error {
	user.ID = uuid.New().String()
	// Добавить хеширование пароля
	if err := user.HashPassword(user.Password); err != nil {
		return err
	}
	return uc.userRepo.Create(user)
}

func (uc *UserUseCase) GetUser(id string) (*entities.User, error) {
	return uc.userRepo.GetByID(id)
}

func (uc *UserUseCase) SearchUsers(query string) ([]*entities.User, error) {
	return uc.userRepo.Search(query)
}
