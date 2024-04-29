package dtos

type Order struct {
	Id              int    `json:"order_id"`
	Deadline        string `json:"deadline"`
	Status          string `json:"status"`
	ProductId       int    `json:"product_id"`
	ManagerName     string `json:"manager_name"`
	ManagerUsername string `json:"manager_username"`
	ClientName      string `json:"client_name"`
	ClientUsername  string `json:"client_username"`
}

type MyOrder struct {
	Deadline        string `json:"deadline"`
	Status          string `json:"status"`
	ProductId       int    `json:"product_id"`
	ManagerName     string `json:"manager_name" db:"manager_name"`
	ManagerUsername string `json:"manager_username" db:"manager_username"`
}

type ManagerOrder struct {
	Deadline       string `json:"deadline"`
	Status         string `json:"status"`
	ProductId      int    `json:"product_id"`
	ClientName     string `json:"manager_name" db:"manager_name"`
	ClientUsername string `json:"manager_username" db:"manager_username"`
}

type UpdateOrder struct {
	Status string `json:"status"`
}

type GetOrdersResponse struct {
	Data []Order `json:"data"`
}

type GetMyOrdersResponse struct {
	Data []MyOrder `json:"data"`
}
