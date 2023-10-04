package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	userTable     = "users"
	todoListTable = "todo_list"
	userListTable = "users_list"
	todoItemTable = "todo_item"
	listItemTable = "lists_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
	SSLmode  string
}

func PostgresqlDB(cnf Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s  sslmode=%s", cnf.Host, cnf.Port, cnf.Username, cnf.DBname, cnf.Password, cnf.SSLmode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
