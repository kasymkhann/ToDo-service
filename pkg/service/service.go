package service

import (
	user "to-doProjectGo"
	"to-doProjectGo/pkg/repository"
)

type Entering interface {
	CreateUser(user user.User) (int, error)
}

type ToDoList interface {
}

type ToDoItem interface {
}

type Service struct {
	Entering
	ToDoItem
	ToDoList
}

func Servic(r *repository.Repository) *Service {
	return &Service{
		Entering: EnteringService(r.Entering),
	}
}
