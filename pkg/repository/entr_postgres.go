package repository

import (
	"fmt"
	user "to-doProjectGo"

	"github.com/jmoiron/sqlx"
)

type EntrPostgres struct {
	db *sqlx.DB
}

func EntrPostgresRepo(db *sqlx.DB) *EntrPostgres {
	return &EntrPostgres{db: db}
}

func (e EntrPostgres) CreateUser(user user.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values($1, $2, $3) RETURNING id", userTable)

	row := e.db.QueryRow(query, user.Name, user.UserName, user.Password)
	if err := row.Scan(id); err != nil {
		return 0, err
	}
	return 0, nil
}
