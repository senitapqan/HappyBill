package repository

import (
	"fmt"
	"happyBill/consts"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func (r *repository) GetRoles(id int) ([]string, error) {
	var roles []string
	query := fmt.Sprintf("select r.role_name from %s r join %s c on c.role_id = r.id where c.user_id = $1", consts.RolesTable, consts.UsersRolesTable)

	log.Error().Msg("EXECUTING SQL QUERY")

	rows, err := r.db.Query(query, id)

	if err != nil {
		return roles, err
	}
	for rows.Next() {
		var role string
		err := rows.Scan(&role)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, err
}

func (r *repository) GetRoleId(role string, userId int) (int, error) {
	var id int
	table := "t_" + strings.ToLower(role)

	query := fmt.Sprintf("select id from %s where user_id = $1", table)
	err := r.db.Get(&id, query, userId)
	return id, err
}

func (r *repository) GetIdByRole(tx *sqlx.Tx, roleId int, roleTable string) (int, error) {
	var userId int
	query := fmt.Sprintf("select user_id from %s where id = $1", roleTable)
	row := r.db.QueryRowx(query, roleId)

	if err := row.Scan(&userId); err != nil {
		tx.Rollback()
		return 0, err
	}

	return userId, nil
}
