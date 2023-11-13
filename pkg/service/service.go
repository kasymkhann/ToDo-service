package service

import (
	user "to-doProjectGo"
	"to-doProjectGo/pkg/repository"
)

type Entering interface {
	CreateUser(user user.User) (int, error)
	GenerateTOKEN(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
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
