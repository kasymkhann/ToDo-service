package service

import (
	todo "to-doProjectGo"
	"to-doProjectGo/pkg/repository"
)

type TodoItem struct {
	r     repository.ToDoItem
	listR repository.ToDoList
}

func TodoItemService(r repository.ToDoItem, listR repository.ToDoList) *TodoItem {
	return &TodoItem{r: r, listR: listR}
}

func (t *TodoItem) Create(userId, listId int, item todo.ToDoItem) (int, error) {
	_, err := t.listR.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return t.r.Create(listId, item)
}

func (t *TodoItem) GetAll(userId, listId int) ([]todo.ToDoItem, error) {
	return t.r.GetAll(userId, listId)
}

func (t *TodoItem) GetItemId(userId, itemId int) ([]todo.ToDoItem, error) {
	return t.r.GetItemId(userId, itemId)
}

func (t *TodoItem) Delete(userId, itemId int) error {
	return t.r.Delete(userId, itemId)
}

func (t *TodoItem) Update(userId, itemId int, input todo.UpdateItemListInput) error {
	return t.r.Update(userId, itemId, input)
}
