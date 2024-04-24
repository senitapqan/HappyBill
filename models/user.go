package models

import "time"

type Manager struct {
	Id     int `json:"manager_id"`
	UserId int `json:"user_id"`
	CreatedAt time.Time `json:"created"`
}

type Client struct {
	Id        int       `json:"client_id"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created"`
}

type Admin struct {
	Id        int       `json:"admin_id"`
	UserId    int       `json:"user_id"`
	OrderCount int	
	ActiveOrderCount int
	CreatedAt time.Time `json:"created"`
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
	Id       int    `json:"role_id"`
	RoleName string `json:"role_name"`
}

type RolesHeaders struct {
	Role string
	Id   int
}
