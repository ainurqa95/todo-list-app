package services

import (
	"github.com/ainurqa95/todo-list-app"
	"github.com/ainurqa95/todo-list-app/pkg/repositories"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.UserRepository),
	}
}
