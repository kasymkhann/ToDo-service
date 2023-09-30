package repository

import "github.com/jmoiron/sqlx"

type Entering interface {
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
	return &Repository{}
}
