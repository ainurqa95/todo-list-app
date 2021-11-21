package services

import (
	"github.com/ainurqa95/todo-list-app"
	"github.com/ainurqa95/todo-list-app/pkg/repositories"
)

type ItemManager interface {
	CreateItem(item todo.TodoItem, listId int) (int, error)
	GetAllItems(listId int, userId int) ([]todo.TodoItem, error)
}

type ItemService struct {
	itemRepository repositories.ItemRepository
}

func NewItemService(repository repositories.ItemRepository) *ItemService {
	return &ItemService{itemRepository: repository}
}

func (itemService *ItemService) CreateItem(item todo.TodoItem, listId int) (int, error) {
	return itemService.itemRepository.Create(item, listId)
}

func (itemService *ItemService) GetAllItems(listId int, userId int) ([]todo.TodoItem, error) {
	return itemService.itemRepository.GetAll(userId, listId)
}
