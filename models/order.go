package models

type Order struct {
	Id          int    `json:"order_id" db:"id"`
	OrderedTime string `json:"ordertime" db:"ordertime"`
	Deadline    string `json:"deadline" db:"deadline"`
	StartTime   string `json:"startdate" db:"startdate"`
	EndTime     string `json:"enddate" db:"enddate"`
	Status      string `json:"status" db:"status"`
	ProductId   int    `json:"product_id" db:"product_id"`
	ManagerId   int    `json:"manager_id" db:"manager_id"`
	ClientId    int    `json:"client_id" db:"client_id"`
	Price       int    `json:"price" db:"price"`
	Archive     bool   `json:"archive" db:"archive"`
}
