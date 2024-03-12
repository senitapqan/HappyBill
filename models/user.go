package models

type Manager struct {
	Id     int `json:"managerId"`
	UserId int `json:"userId"`
}

type Client struct {
	Id     int `json:"clientId"`
	UserId int `json:"userId"`
}

type Admin struct {
	Id     int `json:"adminId"`
	UserId int `json:"userId"`
}

type User struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" binding:"required" db:"username"`
	Password string `json:"password" binding:"required" db:"password"`
	Email    string `json:"email"    binding:"required"`
	Name     string `json:"name"     binding:"required"`
	Surname  string `json:"surname"  binding:"required"`
}

type Roles struct {
	Id       int    `json:"roleId"`
	RoleName string `json:"role_name"`
}

type RolesHeaders struct {
	Role string
	Id   int
}
