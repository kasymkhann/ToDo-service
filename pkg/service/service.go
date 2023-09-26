package service

import "to-doProjectGo/pkg/repository"

type Entering interface {
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
	return &Service{}
}
