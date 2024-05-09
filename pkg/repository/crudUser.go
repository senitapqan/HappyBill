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
	query := fmt.Sprintf("insert into %s (username, password, name, surname, email, phone) values ($1, $2, $3, $4, $5, $6) returning id", consts.UsersTable)
	row := tx.QueryRow(query, user.Username, user.Password, user.Name, user.Surname, user.Email, user.Phone)
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
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return models.User{}, fmt.Errorf("there is no such user with username/email: %s", input)
		}
		return models.User{}, err
	}
	return user, nil
}

func (r *repository) GetUserById(id int) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("select id, username, password, name, surname, email from %s where id = $1", consts.UsersTable)

	err := r.db.Get(&user, query, id)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return models.User{}, fmt.Errorf("there is no such user with id: %d", id)
		}
		return models.User{}, err
	}
	return user, nil
}
