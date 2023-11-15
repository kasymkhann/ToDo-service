package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	todo "to-doProjectGo"
)

type TodoItemPostgresql struct {
	db *sqlx.DB
}

func TodoItemPostgresqlRepo(db *sqlx.DB) *TodoItemPostgresql {
	return &TodoItemPostgresql{db: db}
}

func (t *TodoItemPostgresql) Create(listId int, item todo.ToDoItem) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createQueryItem := fmt.Sprintf("INSERT INTO %s (title , description) VALUES ($1, $2)", todoItemTable)

	row := tx.QueryRow(createQueryItem, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	createQueryListItems := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listItemTable)
	_, err = tx.Exec(createQueryListItems, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (t *TodoItemPostgresql) GetAll(userId, listId int) ([]todo.ToDoItem, error) {
	var items []todo.ToDoItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2 ", todoItemTable, listItemTable, userListTable)
	if err := t.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}
	return items, nil

}

func (t *TodoItemPostgresql) GetItemId(userId, itemId int) ([]todo.ToDoItem, error) {
	var item []todo.ToDoItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id INNER JOIN %s ul on ul.list_id = li.list_id WHERE ti.id = $1 AND ul.user_id = $2 ", todoItemTable, listItemTable, userListTable)
	if err := t.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}
	return item, nil
}
func (t *TodoItemPostgresql) Delete(userId, itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s ti USING %s li, %s lu WHERE ti.id = li.item_id = AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2 ", todoItemTable, listItemTable, userListTable)
	_, err := t.db.Exec(query, userId, itemId)
	return err
}

func (t *TodoItemPostgresql) Update(userId, itemId int, input todo.UpdateItemListInput) error {
	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argsId := 1
	if input.Title != nil {
		setValue = append(setValue, fmt.Sprintf("title=$%d", argsId))
		args = append(args, *input.Title)
		argsId++
	}
	if input.Description != nil {
		setValue = append(setValue, fmt.Sprintf("description=$%d", argsId))
		args = append(args, *input.Description)
		argsId++
	}
	if input.Done != nil {
		setValue = append(setValue, fmt.Sprintf("done=$%d", argsId))
		args = append(args, *input.Done)
		argsId++
	}

	setQuery := strings.Join(setValue, ",")

	query := fmt.Sprintf("UPDATE %s ti SET %s FROM %s li, %s ul WHERE ti.id = li.item_id AND li.list_id = ul.list AND ul.user_id = $%d AND ti.id = $%d",
		todoItemTable, setQuery, listItemTable, userListTable, argsId, argsId+1)
	args = append(args, userId, itemId)

	_, err := t.db.Exec(query, args...)
	return err
}
