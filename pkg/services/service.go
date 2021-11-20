package services

import (
	"github.com/ainurqa95/todo-list-app/pkg/repositories"
)

type TodoItem interface {
}

type Service struct {
	Authorization
	ListManager
	ItemManager
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.UserRepository),
		ListManager:   NewListService(repos.ListRepository),
		ItemManager:   NewItemService(repos.ItemRepository),
	}
}
