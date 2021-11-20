package repositories

import (
	"github.com/jmoiron/sqlx"
)

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	UserRepository
	ListRepository
	ItemRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepository: NewDbUserRepository(db),
		ListRepository: NewDbListRepository(db),
		ItemRepository: NewDbListItemRepository(db),
	}
}
