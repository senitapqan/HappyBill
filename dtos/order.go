package dtos

type Order struct {
	Id              int    `json:"orderId"`
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
	ManagerName     string `json:"manager_name"`
	ManagerUsername string `json:"manager_username"`
}

type GetOrdersResponse struct {
	Data []Order `json:"data"`
}

type GetMyOrdersResponse struct {
	Data []MyOrder `json:"data"`
}
