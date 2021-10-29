package repositories

import (
	"github.com/ainurqa95/todo-list-app"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(user todo.User) (int, error)
}
type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	UserRepository
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{UserRepository: NewDbUserRepository(db)}
}
