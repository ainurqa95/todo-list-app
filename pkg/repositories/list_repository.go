package repositories

import (
	"fmt"

	"github.com/ainurqa95/todo-list-app"
	"github.com/jmoiron/sqlx"
)

type ListRepository interface {
	Create(list todo.TodoList, userId int) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(id int) (todo.TodoList, error)
	Delete(id int) error
	Update(id int, input todo.UpdateListInput) error
}

type DbListRepository struct {
	db *sqlx.DB
}

func NewDbListRepository(db *sqlx.DB) *DbListRepository {
	return &DbListRepository{db: db}
}

func (listRepository *DbListRepository) Create(list todo.TodoList, userId int) (int, error) {
	transaction, err := listRepository.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	queryList := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", todoListsTable)

	row := transaction.QueryRow(queryList, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		transaction.Rollback()
		return 0, err
	}
	queryListUser := fmt.Sprintf("INSERT INTO %s (list_id, user_id) values ($1, $2) RETURNING id", usersListsTable)
	_, err = transaction.Exec(queryListUser, id, userId)
	if err != nil {
		transaction.Rollback()
		return 0, err
	}

	return id, transaction.Commit()
}

func (listRepository *DbListRepository) GetAll(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	query := fmt.Sprintf("SELECT todo_lists.* FROM todo_lists JOIN users_lists ul on todo_lists.id = ul.list_id WHERE ul.user_id=$1")

	err := listRepository.db.Select(&lists, query, userId)

	return lists, err
}

func (listRepository *DbListRepository) GetById(id int) (todo.TodoList, error) {
	var list todo.TodoList

	query := fmt.Sprintf("SELECT id, title, description FROM %s WHERE id=$1", todoListsTable)

	err := listRepository.db.Get(&list, query, id)

	return list, err
}

func (listRepository *DbListRepository) Delete(id int) error {

	query := fmt.Sprintf("Delete FROM %s WHERE id=$1", todoListsTable)

	_, err := listRepository.db.Exec(query, id)

	return err
}

func (listRepository *DbListRepository) Update(id int, input todo.UpdateListInput) error {

	query := fmt.Sprintf("UPDATE %s SET title=$1, description=$2 where id=$3",
		todoListsTable)

	_, err := listRepository.db.Exec(query, input.Title, input.Description, id)

	return err

}
