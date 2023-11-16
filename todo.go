package todoprojectgo

import (
	"errors"
)

type ToDoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type ToDoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       int    `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type ListItem struct {
	Id       int
	ListItem int
	ItemId   int
}
type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (u UpdateListInput) Validate() error {
	if u.Title == nil && u.Description == nil {
		return errors.New("update structure has no value")
	}
	return nil
}

type UpdateItemListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (u UpdateItemListInput) Validate() error {
	if u.Title == nil && u.Description == nil && u.Done == nil {
		return errors.New("update structure has no value")
	}
	return nil
}
