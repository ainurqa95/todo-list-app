package repositories

import (
	"fmt"

	"github.com/ainurqa95/todo-list-app"
	"github.com/jmoiron/sqlx"
)

type DbUserRepoistory struct {
	db *sqlx.DB
}

func NewDbUserRepository(db *sqlx.DB) *DbUserRepoistory {
	return &DbUserRepoistory{db: db}
}

func (r *DbUserRepoistory) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
