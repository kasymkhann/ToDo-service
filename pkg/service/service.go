package service

import (
	todo "to-doProjectGo"
	user "to-doProjectGo"
	"to-doProjectGo/pkg/repository"
)

type Entering interface {
	CreateUser(user user.User) (int, error)
	GenerateTOKEN(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type ToDoList interface {
	Create(userId int, list todo.ToDoList) (int, error)
	GetAll(userId int) ([]todo.ToDoList, error)
	GetById(userId, listId int) (todo.ToDoList, error)
	Update(userId, ListId int, input todo.UpdateListInput) error
	Delete(userId, ListId int) error
}

type ToDoItem interface {
	Create(userId, listId int, item todo.ToDoItem) (int, error)
	GetAll(userId, listId int) ([]todo.ToDoItem, error)
	GetItemId(userId, itemId int) ([]todo.ToDoItem, error)
	Update(userId, itemId int, input todo.UpdateItemListInput) error
	Delete(userId, itemId int) error
}

type Service struct {
	Entering
	ToDoItem
	ToDoList
}

func Servic(r *repository.Repository) *Service {
	return &Service{
		Entering: EnteringService(r.Entering),
		ToDoList: TodoListServise(r.ToDoList),
		ToDoItem: TodoItemService(r.ToDoItem, r.ToDoList),
	}
}
