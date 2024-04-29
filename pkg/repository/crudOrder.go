package repository

import (
	"errors"
	"fmt"
	"happyBill/consts"
	"happyBill/dtos"
	"happyBill/models"
	"time"

	"github.com/rs/zerolog/log"
)

func (r *repository) GetMyOrders(id, page int, status string) ([]dtos.MyOrder, error) {
	var result []dtos.MyOrder
	query := fmt.Sprintf(`Select ord.deadline, ord.status, ord.product_id, usr.name AS manager_name, 
			usr.username AS manager_username from %s ord 
			join %s mng on ord.manager_id = mng.id 
			join %s usr on usr.id = mng.user_id 
			where ord.client_id = $1 and ord.status = '%s' 
			order by ord.created_time desc limit %d offset %d`, 
			consts.OrdersTable, consts.ManagersTable, consts.UsersTable, status,
		consts.PaginationLimit, (page-1)*consts.PaginationLimit)
	log.Info().Msg(query)
	err := r.db.Select(&result, query, id)
	return result, err
}

func (r *repository) CreateOrder(id int, order models.Order) (int, error) {
	OrderedTime, err := time.Parse("2006-01-02", order.OrderedTime)

	if err != nil {
		return -1, errors.New("wrong format of ordered time")
	}

	Deadline, err := time.Parse("2006-01-02", order.Deadline)

	if err != nil {
		return -1, errors.New("wrong format of deadline time")
	}

	if Deadline.Sub(OrderedTime).Hours() / 24 != 14 {
		return -1, errors.New("wrong deadline is setted")
	}

	var orderId int
	query := fmt.Sprintf(`insert into %s (ordertime, deadline, startdate, enddate, product_id, client_id, manager_id, price)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8) returning id`, consts.OrdersTable)

	row := r.db.QueryRow(query, order.OrderedTime, order.Deadline, order.StartTime, order.EndTime, order.ProductId,
		order.ClientId, order.ManagerId, order.Price)

	if err := row.Scan(&orderId); err != nil {
		return -1, err
	}
	return orderId, nil
}