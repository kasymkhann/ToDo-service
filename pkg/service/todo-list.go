package service

import (
	todo "to-doProjectGo"
	"to-doProjectGo/pkg/repository"
)

type TodoList struct {
	r repository.ToDoList
}

func TodoListServise(r repository.ToDoList) *TodoList {
	return &TodoList{r: r}
}

func (t *TodoList) Create(userId int, list todo.ToDoList) (int, error) {
	return t.r.Create(userId, list)
}

func (t *TodoList) GetAll(userId int) ([]todo.ToDoList, error) {
	return t.r.GetAll(userId)
}

func (t *TodoList) GetById(userId, listId int) (todo.ToDoList, error) {
	return t.r.GetById(userId, listId)
}
