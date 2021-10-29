package services

import (
	"crypto/sha1"
	"fmt"

	"github.com/ainurqa95/todo-list-app"
	"github.com/ainurqa95/todo-list-app/pkg/repositories"
)

type AuthService struct {
	repository repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) *AuthService {
	return &AuthService{repository: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repository.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte("salt")))
}
