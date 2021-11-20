package repositories

import (
	"fmt"

	"github.com/ainurqa95/todo-list-app"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UserRepository interface {
	Create(user todo.User) (int, error)
	FindBy(username string, passwordHash string) (todo.User, error)
}

type DbUserRepoistory struct {
	db *sqlx.DB
}

func NewDbUserRepository(db *sqlx.DB) *DbUserRepoistory {
	return &DbUserRepoistory{db: db}
}

func (repos *DbUserRepoistory) Create(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := repos.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (repos *DbUserRepoistory) FindBy(username string, passwordHash string) (todo.User, error) {
	var user todo.User
	logrus.Debug(username, passwordHash)
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)

	err := repos.db.Get(&user, query, username, passwordHash)

	return user, err
}
