package dtos

type User struct {
	Id       int         `json:"id" db:"id"`
	Username string      `json:"username" db:"username"`
	Email    string      `json:"email" db:"email"`
	Name     string      `json:"name" db:"name"`
	Surname  string      `json:"surname" db:"surname"`
	Phone    interface{} `db:"phone"`
	RoleId   int         `json:"role_id" db:"role_id"`
}

type UpdateUser struct {
	Name        string `json:"name" db:"name"`
	Surname     string `json:"surname" db:"surname"`
	Username    string `json:"username" db:"username"`
	Password    string `json:"password" db:"password"`
	OldPassword string `json:"old_password" binding:"required"`
}

type SignInRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetUserByIdResponse struct {
	Data User `json:"data"`
}
