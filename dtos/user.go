package dtos

type User struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Name     string `json:"name" db:"name"`
	Surname  string `json:"surname" db:"surname"`
	RoleId   int    `json:"roleId" db:"roleid"`
}

type UpdateUser struct {
	Name     string `json:"name" db:"name"`
	Surname  string `json:"surname" db:"surname"`
	Password string `json:"password" db:"password"`
}

type SignInRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetUserByIdResponse struct {
	Data User `json:"data"`
}
