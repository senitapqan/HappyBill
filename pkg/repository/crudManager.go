package repository

import (
	"fmt"
	"happyBill/consts"
	"happyBill/dtos"
	"happyBill/models"
)

func (r *repository) CreateManager(manager models.User) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}

	var managerId int
	userId, err := r.CreateUser(manager, tx)
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("insert into %s (user_id) values($1) returning id", consts.ManagersTable)
	row := tx.QueryRowx(query, userId)

	if err := row.Scan(&managerId); err != nil {
		tx.Rollback()
		return 0, err
	}

	query = fmt.Sprintf("insert into %s (user_id, role_id) values ($1, $2)", consts.UsersRolesTable)
	_, err = tx.Exec(query, userId, consts.ManagerRoleId)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return managerId, tx.Commit()
}

func (r *repository) GetAllManagers() ([]dtos.User, error) {
	var result []dtos.User
	query := fmt.Sprintf("select u.id, u.name, u.surname, u.username, u.email, m.id as roleid from %s u join %s m ON m.user_id = u.id", consts.UsersTable, consts.ManagersTable)

	err := r.db.Select(&result, query)
	return result, err
}

func (r *repository) GetManagerById(id int) (dtos.User, error) {
	var result dtos.User

	query := fmt.Sprintf(`select u.id, u.name, u.surname, u.username, u.email, m.id as roleid from %s u join %S m ON m.user_id = u.id where m.id = $1`,
		consts.UsersRolesTable, consts.ManagersTable)
	err := r.db.Get(result, query, id)
	return result, err
}

func (r *repository) DeleteManager(id int)
