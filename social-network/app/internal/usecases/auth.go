package usecases

import (
	"fmt"
	"social_network/internal/interfaces/repositories"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthUseCase struct {
	userRepo repositories.UserRepository
}

func NewAuthUseCase(userRepo repositories.UserRepository) *AuthUseCase {
	return &AuthUseCase{userRepo: userRepo}
}

func (uc *AuthUseCase) Login(id, password string) (string, error) {
	user, err := uc.userRepo.GetByID(id)
	if err != nil {
		return "", err
	}

	if err := user.CheckPassword(password); err != nil {
		fmt.Println(err)
		return "", err
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("my_secret_key"))
}
