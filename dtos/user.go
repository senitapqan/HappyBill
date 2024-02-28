package dtos

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"  `
	Name     string `json:"name"   `
	Surname  string `json:"surname"`
	RoleId   int    `json:"roleId"`
}