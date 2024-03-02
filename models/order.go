package models

type Order struct {
	Id        int    `json:"orderId"`
	Deadline  string `json:"deadline"`
	ProductId int    `json:"productId"`
	ManagerId int    `json:"managerId"`
	ClientId  int    `json:"clientId"`
}
