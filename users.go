package todoprojectgo

type User struct {
	Id       int    `json: "-" db: "id"`
	Name     string `json: "name" binding: "requires"`
	UserName string `json: "username" binding: "requires"`
	Password string `json: "password" binding: "requires"`
}
