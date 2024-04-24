package repository

import (
	"fmt"
	"happyBill/consts"
	"happyBill/models"
	"strings"

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

func (r *repository) GetUser(input string) (models.User, error) {
	var user models.User

	queryParam := "username"

	if strings.Contains(input, "@") {
			queryParam = "email"
	}
	query := fmt.Sprintf("select id, username, password from %s where %s = $1", consts.UsersTable, queryParam)
	err := r.db.Get(&user, query, input)
	return user, err
}

func (r *repository) GetUserById(id int) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("select id, username, password, name, surname, phone, email from %s where id = $1", consts.UsersTable)
	err := r.db.Get(&user, query, id)
	return user, err
}

