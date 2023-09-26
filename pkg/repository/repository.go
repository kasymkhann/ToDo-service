package repository

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

func Repositories() *Repository {
	return &Repository{}
}
