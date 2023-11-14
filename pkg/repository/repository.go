package repository

import (
	user "to-doProjectGo"

	"github.com/jmoiron/sqlx"

	todo "to-doProjectGo"
)

type Entering interface {
	CreateUser(user user.User) (int, error)
	GetUser(username, password string) (user.User, error)
}

type ToDoList interface {
	Create(userId int, list todo.ToDoList) (int, error)
	GetAll(userId int) ([]todo.ToDoList, error)
	GetById(userId, listId int) (todo.ToDoList, error)
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
		ToDoList: TodoListPostgresqlRepo(db),
	}
}
