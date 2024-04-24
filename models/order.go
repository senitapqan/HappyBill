package models

type Order struct {
	Id          int    `json:"order_id"`
	OrderedTime string `json:"ordertime"`
	Deadline    string `json:"deadline"`
	StartTime   string `json:"startdate"`
	EndTime     string `json:"enddate"`
	Status      string `json:"status"`
	ProductId   int    `json:"product_id"`
	ManagerId   int    `json:"manager_id"`
	ClientId    int    `json:"client_id"`
	Price       int    `json:"price"`
}
