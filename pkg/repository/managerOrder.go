package repository

import (
	"fmt"
	"happyBill/consts"
	"happyBill/dtos"
	"strings"

	"github.com/rs/zerolog/log"
)

func (r *repository) GetAllManagerOrders(id, page int) ([]dtos.ManagerOrder, error) {
	var result []dtos.ManagerOrder
	query := fmt.Sprintf(`Select ord.deadline, ord.status, ord.product_id, usr.name AS client_name,
			usr.username AS client_username from %s ord
			join %s clnt on ord.client_id = clnt.id
			join %s usr on usr.id = clnt.user_id
			where ord.manager_id = $1
			order by ord.created_time desc limit %d offset %d`,
		consts.OrdersTable, consts.ClientsTable, consts.UsersTable,
		consts.PaginationLimit, (page-1)*consts.PaginationLimit)
	log.Info().Msg(query)
	err := r.db.Select(&result, query, id)
	// return result, err
	return nil, err
}

func (r *repository) GetManagerOrderById(id int) (dtos.ManagerOrder, error) {
	var result dtos.ManagerOrder

	query := fmt.Sprintf(`select ord.id, ord.deadline, ord.status, ord.product_id, 
	usr.name as client_name, usr.username as client_username from %s ord join %s clnt on ord.client_id = clnt.id
	join %s usr on usr.id = clnt.user_id
	where ord.id = $1`, consts.OrdersTable, consts.ClientsTable, consts.UsersTable)
	err := r.db.Get(result, query, id)
	return result, err

}

func (r *repository) UpdateManagerOrder(id int, input dtos.UpdateOrder) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Status != "" {
		setValues = append(setValues, fmt.Sprintf("status=$%d", argId))
		args = append(args, input.Status)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = %d",
		consts.OrdersTable, setQuery, id)

	_, err := r.db.Exec(query, args...)
	log.Printf("args: %s query: %s", args, query)
	return err
}
