package models

type Order struct {
	Id        int    `json:"orderId"`
	Deadline  string `json:"deadline"`
	Status    string `json:"status"`
	ProductId int    `json:"productId"`
	ManagerId int    `json:"managerId"`
	ClientId  int    `json:"clientId"`
}
