package repository

import (
	user "to-doProjectGo"

	"github.com/jmoiron/sqlx"
)

type Entering interface {
	CreateUser(user user.User) (int, error)
}

type ToDoList interface {
}

type ToDoItem interface {
}

type Repository struct {
	Entering
	ToDoItem
	ToDoList
}

func Repositories(db *sqlx.DB) *Repository {
	return &Repository{
		Entering: EntrPostgresRepo(db),
	}
}
