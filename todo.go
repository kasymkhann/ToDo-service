package todoprojectgo

type ToDoList struct {
	Id          int    `json: "id"`
	Title       int    `json: "title"`
	Description string `json: "description"`
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
