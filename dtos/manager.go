package dtos

type GetAllManagersResponse struct {
	Data []User `json:"data"`
}

type GetManagersResponse struct {
	Data User `json:"data"`
}
