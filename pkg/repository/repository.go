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
	Update(userId, ListId int, input todo.UpdateListInput) error
	Delete(userId, listId int) error
}

type ToDoItem interface {
	Create(listId int, item todo.ToDoItem) (int, error)
	GetAll(userId, listId int) ([]todo.ToDoItem, error)
	GetItemId(userId, itemId int) ([]todo.ToDoItem, error)
	Update(userId, itemId int, input todo.UpdateItemListInput) error
	Delete(userId, itemId int) error
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
		ToDoItem: TodoItemPostgresqlRepo(db),
	}
}
