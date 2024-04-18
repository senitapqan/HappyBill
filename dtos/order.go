package dtos

type Order struct {
	Id              int    `json:"orderId"`
	Deadline        string `json:"deadline"`
	Status          string `json:"status"`
	ProductId       int    `json:"productId"`
	ManagerName     string `json:"managerName"`
	ManagerUsername string `json:"managerUsername"`
	ClientName      string `json:"clientName"`
	ClientUsername  string `json:"clientUsername"`
}

type GetAllOrdersResponse struct {
	Data []Order `json:"data"`
}
