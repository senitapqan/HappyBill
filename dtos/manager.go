package dtos

type GetAllManagersResponse struct {
	Data []User `json:"data"`
}

type GetManagersResponse struct {
	Data User `json:"data"`
}

type GetAllManagerOrdersResponse struct {
	Data []ManagerOrder `json:"data"`
}

type GetManagerOrderResponse struct {
	Data ManagerOrder `json:"data"`
}
