package repo

import (
	"fmt"

	todoapp "gin-todo-prc/src/models"

	"github.com/jmoiron/sqlx"
)

type AuthDb struct {
	db *sqlx.DB
}

func NewAuthDb(db *sqlx.DB) *AuthDb {
	return &AuthDb{db: db}
}

func (r *AuthDb) CreateUser(user todoapp.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthDb) GetUser(username, password string) (todoapp.User, error) {
	var user todoapp.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
