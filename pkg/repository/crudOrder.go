package repository

import (
	"fmt"
	"happyBill/consts"
	"happyBill/dtos"
)

func (r *repository) GetMyOrders(id, page int) ([]dtos.MyOrder, error) {
	var result []dtos.MyOrder
	query := fmt.Sprintf(`Select ord.deadline, ord.status, ord.product_id,
			usr.name AS manager_name, usr.username AS manager_username
			from %s ord join %s mng on ord.managerid = mng.id
			join %s usr on usr.id = mng.user_id
			where ord.clientid = $1
			order by ord.created_time desc
			limit %d offset %d`, consts.OrdersTable, consts.ManagersTable, consts.UsersTable, 
			consts.PaginationLimit, (page - 1) * consts.PaginationLimit)
	err := r.db.Select(&result, query, id)
	return result, err
}