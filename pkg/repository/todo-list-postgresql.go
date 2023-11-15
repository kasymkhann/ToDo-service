package repository

import (
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"

	todo "to-doProjectGo"
)

type TodoListPostgresql struct {
	db *sqlx.DB
}

func TodoListPostgresqlRepo(db *sqlx.DB) *TodoListPostgresql {
	return &TodoListPostgresql{db: db}
}

func (t *TodoListPostgresql) Create(userId int, list todo.ToDoList) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", listItemTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	createUsersListsQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES($1, $2)", userListTable)
	_, err = tx.Exec(createUsersListsQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (t *TodoListPostgresql) GetAll(userId int) ([]todo.ToDoList, error) {
	var lists []todo.ToDoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1", todoListTable, userListTable)
	err := t.db.Select(&lists, query, userId)
	return lists, err
}

func (t *TodoListPostgresql) GetById(userId, listId int) (todo.ToDoList, error) {
	var list todo.ToDoList
	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl INER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`, todoListTable, userListTable)
	err := t.db.Get(&list, query, userId, listId)
	return list, err
}

func (t *TodoListPostgresql) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id =$2", todoListTable, userListTable)
	_, err := t.db.Exec(query, userId, listId)
	return err
}

func (t *TodoListPostgresql) Update(userId, ListId int, input todo.UpdateListInput) error {
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

	setQuery := strings.Join(setValue, ",")

	query := fmt.Sprintf("UPDATE %s tl SET FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d", todoListTable, setQuery, userListTable, argsId, argsId+1)
	args = append(args, ListId, userId)

	log.Fatalf("updateQuery: %s ", query)
	log.Fatalf("args: %s", args)

	_, err := t.db.Exec(query, args...)
	return err
}
