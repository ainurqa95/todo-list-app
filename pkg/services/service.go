package services

import (
	"github.com/ainurqa95/todo-list-app"
	"github.com/ainurqa95/todo-list-app/pkg/repositories"
)

//go:generate $HOME/go/bin/mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type ItemManager interface {
	CreateItem(item todo.TodoItem, listId int) (int, error)
	GetAllItems(listId int, userId int) ([]todo.TodoItem, error)
}

type ListManager interface {
	CreateList(list todo.TodoList, userId int) (int, error)
	GetAllLists(userId int) ([]todo.TodoList, error)
	GetListById(id int) (todo.TodoList, error)
	UpdateList(id int, input todo.UpdateListInput) error
	DeleteList(id int) error
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
