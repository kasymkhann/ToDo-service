package todoprojectgo

type User struct {
	Id       int    `json: "-"`
	Name     string `json: "name" binding: "requires"`
	UserName string `json: "username" binding: "requires"`
	Password string `json: "password" binding: "requires"`
}
