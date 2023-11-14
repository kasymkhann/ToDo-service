package todoprojectgo

type ToDoList struct {
	Id          int    `json: "id" db: "id"`
	Title       int    `json: "title" db: "title" binding: "required"`
	Description string `json: "description" db: "description"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type ToDoItem struct {
	Id          int    `json: "id"`
	Title       int    `json: "title"`
	Description string `json: "description"`
	Done        bool   `json: "done"`
}

type ListItem struct {
	Id       int
	ListItem int
	ItemId   int
}
