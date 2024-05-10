package dtos

type GetAllManagersResponse struct {
	Data []User `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type GetManagerResponse struct {
	Data User `json:"data"`
}

type GetAllManagerOrdersResponse struct {
	Data []ManagerOrder `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type GetManagerOrderResponse struct {
	Data ManagerOrder `json:"data"`
}
