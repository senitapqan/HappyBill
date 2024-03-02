package repository

import (
	"fmt"
	"happyBill/models"
)

func (r *repository) CreateClient(client models.User) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	var clientId int
	userId, err := r.CreateUser(client, tx)
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("insert into %s (user_id) values($1) returning id", clientsTable)
	row := tx.QueryRowx(query, userId)

	if err := row.Scan(&clientId); err != nil {
		tx.Rollback()
		return 0, err
	}

	return clientId, tx.Commit()
}
