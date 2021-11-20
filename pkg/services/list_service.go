package services

import (
	"errors"

	"github.com/ainurqa95/todo-list-app"
	"github.com/ainurqa95/todo-list-app/pkg/repositories"
)

type ListManager interface {
	CreateList(list todo.TodoList, userId int) (int, error)
	GetAllLists(userId int) ([]todo.TodoList, error)
	GetListById(id int) (todo.TodoList, error)
	UpdateList(id int, input todo.UpdateListInput) error
	DeleteList(id int) error
}

type ListService struct {
	listRepository repositories.ListRepository
}

func NewListService(repository repositories.ListRepository) *ListService {
	return &ListService{listRepository: repository}
}

func (listService *ListService) CreateList(list todo.TodoList, userId int) (int, error) {
	return listService.listRepository.Create(list, userId)
}

func (listService *ListService) GetAllLists(userId int) ([]todo.TodoList, error) {
	return listService.listRepository.GetAll(userId)
}

func (listService *ListService) GetListById(id int) (todo.TodoList, error) {
	return listService.listRepository.GetById(id)
}

func (listService *ListService) DeleteList(id int) error {
	return listService.listRepository.Delete(id)
}

func (listService *ListService) UpdateList(id int, input todo.UpdateListInput) error {
	if err := validate(input); err != nil {
		return err
	}
	return listService.listRepository.Update(id, input)
}

func validate(input todo.UpdateListInput) error {
	if input.Title == nil || input.Description == nil {
		return errors.New("empty params")
	}
	return nil
}
