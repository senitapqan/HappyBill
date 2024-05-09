package repository

import (
	"errors"
	"fmt"
	"happyBill/consts"
	"happyBill/dtos"
	"happyBill/models"
	"strings"

	"github.com/rs/zerolog/log"
)

func (r *repository) CreateManager(manager models.User) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		log.Error().Msg(err.Error())
		return 0, errors.New("something went wrong with repository")
	}

	var managerId int
	userId, err := r.CreateUser(manager, tx)
	if err != nil {
		log.Error().Msg(err.Error())
		return 0, errors.New("something went wrong with repository")
	}

	query := fmt.Sprintf("insert into %s (user_id) values($1) returning id", consts.ManagersTable)
	row := tx.QueryRowx(query, userId)

	if err := row.Scan(&managerId); err != nil {
		tx.Rollback()
		log.Error().Msg(err.Error())
		return 0, errors.New("something went wrong with repository")
	}

	query = fmt.Sprintf("insert into %s (user_id, role_id) values ($1, $2)", consts.UsersRolesTable)
	_, err = tx.Exec(query, userId, consts.ManagerRoleId)

	if err != nil {
		tx.Rollback()
		log.Error().Msg(err.Error())
		return 0, errors.New("something went wrong with repository")
	}

	return managerId, tx.Commit()
}

func (r *repository) GetAllManagers(page int) ([]dtos.User, dtos.Pagination, error) {
	var result []dtos.User
	query := fmt.Sprintf(`select u.id, u.name, u.surname, u.username, u.email, m.id as role_id 
			from %s u join %s m ON m.user_id = u.id 
			order by m.created_time desc
			limit %d offset %d`,
		consts.UsersTable, consts.ManagersTable, consts.PaginationLimit, (page-1)*consts.PaginationLimit)

	if err := r.db.Select(&result, query); err != nil {
		return nil, dtos.Pagination{}, err
	}

	var pagination dtos.Pagination
	pagination.CurrentPage = page
	var totalRows int
	
	query = fmt.Sprintf(`select count(*) from %s u join %s m ON m.user_id = u.id`, 
			consts.UsersTable, consts.ManagersTable)

	if err := r.db.Get(&totalRows, query); err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return nil, dtos.Pagination{}, errors.New("there is no managers")
		}
		return nil, dtos.Pagination{}, err
	}

	pagination.TotalPage = (totalRows + consts.PaginationLimit - 1) / consts.PaginationLimit
	return result, pagination, nil
}

func (r *repository) GetMostFreeManager() (int, error) {
	var id int
	query := fmt.Sprintf("select id from %s order by active_order_count limit 1", consts.ManagersTable)
	err := r.db.Get(&id, query)
	if err != nil {
		log.Error().Msg(err.Error())
		return 0, errors.New("something went wrong in request")
	}
	return id, nil
}

func (r *repository) GetManagerById(id int) (dtos.User, error) {
	var result dtos.User

	query := fmt.Sprintf(`select u.id, u.name, u.surname, u.username, u.email, m.id as role_id from %s u join %s m ON m.user_id = u.id where m.id = $1`,
		consts.UsersRolesTable, consts.ManagersTable)
	err := r.db.Get(result, query, id)
	return result, err
}

func (r *repository) DeleteManager(id int) error {
	return nil
}
