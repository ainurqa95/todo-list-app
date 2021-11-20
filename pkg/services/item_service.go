package services

import (
	"github.com/ainurqa95/todo-list-app"
	"github.com/ainurqa95/todo-list-app/pkg/repositories"
)

type ItemManager interface {
	CreateItem(item todo.TodoItem, userId int, listId int) (int, error)
}

type ItemService struct {
	itemRepository repositories.ItemRepository
}

func NewItemService(repository repositories.ItemRepository) *ItemService {
	return &ItemService{itemRepository: repository}
}

func (itemService *ItemService) CreateItem(item todo.TodoItem, userId int, listId int) (int, error) {
	return itemService.itemRepository.Create(item, userId, listId)
}
