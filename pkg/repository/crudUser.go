package repository

import (
	"fmt"
	"happyBill/consts"
	"happyBill/models"
	"github.com/jmoiron/sqlx"
)



func (r *repository) CreateUser(user models.User, tx *sqlx.Tx) (int, error) {
	var userId int
	query := fmt.Sprintf("insert into %s (username, password, name, surname, email) values ($1, $2, $3, $4, $5) returning id", consts.UsersTable)
	row := tx.QueryRow(query, user.Username, user.Password, user.Name, user.Surname, user.Email)
	if err := row.Scan(&userId); err != nil {
		tx.Rollback()
		return 0, err
	}
	return userId, nil
}

func (r *repository) GetUser(username string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("select id, username, password from %s where username = $1", consts.UsersRolesTable)
	err := r.db.Get(&user, query, username)
	return user, err
}