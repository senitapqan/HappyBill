package repository

import (
	"fmt"
	"happyBill/consts"
	"happyBill/dtos"
	"happyBill/models"
	"strings"

	"github.com/rs/zerolog/log"
)

func (r *repository) CreateClient(client models.User) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return 0, err
	}
	log.Info().Msg("mglksjdf;vjksfdks")
	var clientId int
	userId, err := r.CreateUser(client, tx)
	if err != nil {
		return 0, err
	}
	log.Info().Msg("mglksjdf;vjksfdks")

	query := fmt.Sprintf("insert into %s (user_id) values($1) returning id", consts.ClientsTable)
	row := tx.QueryRowx(query, userId)
	log.Info().Msg(fmt.Sprintf("userid: %d ", userId))
	if err := row.Scan(&clientId); err != nil {
		tx.Rollback()
		return 0, err
	}
	log.Info().Msg("mglksjdf;vjksfdks")
	query = fmt.Sprintf("insert into %s (user_id, role_id) values ($1, $2)", consts.UsersRolesTable)
	log.Info().Msg(query)
	_, err = tx.Exec(query, userId, consts.ClientRoleId)
	log.Error().Msg(err.Error() + "kjbjno")
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	log.Info().Msg("mglksjdf;vjksfdks")
	return clientId, tx.Commit()
}

func (r *repository) GetClientByUserId(id int) (dtos.User, error) {
	var result dtos.User
	query := fmt.Sprintf(`select usr.name, usr.surname, usr.phone, usr.email, usr.username from %s usr
						where usr.id = $1`, consts.UsersRolesTable)
	err := r.db.Get(&result, query, id)
	return result, err
}

func (r *repository) GetClientById(id int) (dtos.User, error) {
	var result dtos.User
	query := fmt.Sprintf(`select usr.name, usr.surname, usr.phone, usr.email, usr.username from %s clnt
						join %s usr ON usr.id = clnt.user_id
						where clnt.id = $1`, consts.ClientsTable, consts.UsersRolesTable)
	err := r.db.Get(&result, query, id)
	return result, err
}

func (r *repository) UpdateMyProfile(userId int, input models.User) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, input.Name)
		argId++
	}

	if input.Surname != "" {
		setValues = append(setValues, fmt.Sprintf("surname=$%d", argId))
		args = append(args, input.Surname)
		argId++
	}

	if input.Password != "" {
		setValues = append(setValues, fmt.Sprintf("password=$%d", argId))
		args = append(args, input.Password)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = %d",
		consts.UsersTable, setQuery, userId)

	_, err := r.db.Exec(query, args...)

	return err
}
