package dtos

type User struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Name     string `json:"name" db:"name"`
	Surname  string `json:"surname" db:"surname"`
	RoleId   int    `json:"roleId" db:"roleid"`
}

