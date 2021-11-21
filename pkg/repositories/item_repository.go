package repositories

import (
	"fmt"

	"github.com/ainurqa95/todo-list-app"
	"github.com/jmoiron/sqlx"
)

type ItemRepository interface {
	Create(item todo.TodoItem, listId int) (int, error)
	GetAll(userId int, listId int) ([]todo.TodoItem, error)
}

type DBListItemRepository struct {
	db *sqlx.DB
}

func NewDbListItemRepository(db *sqlx.DB) *DBListItemRepository {
	return &DBListItemRepository{db: db}
}

func (itemRepository *DBListItemRepository) Create(item todo.TodoItem, listId int) (int, error) {
	transaction, err := itemRepository.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", todoItemsTable)

	row := transaction.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		transaction.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)", listsItemsTable)
	_, err = transaction.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		transaction.Rollback()
		return 0, err
	}

	return itemId, transaction.Commit()
}

func (itemRepository *DBListItemRepository) GetAll(userId int, listId int) ([]todo.TodoItem, error) {

	var items []todo.TodoItem

	query := fmt.Sprintf("select ti.* from lists_items li join todo_items ti on ti.id = li.item_id join users_lists ul on li.list_id = ul.list_id where li.list_id = $1 and ul.user_id = $2")

	err := itemRepository.db.Select(&items, query, listId, userId)

	return items, err
}
