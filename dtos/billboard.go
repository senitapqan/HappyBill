package dtos

type GetAllBillboardsResponse struct {
	Data []Product `json:"data"`
}

type GetBillboardByIdResponse struct {
	Data Product `json:"data"`
}

type Product struct {
	Id           int    `json:"productId"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	DisplayType  int    `json:"display_type" db:"display_type"`
	LocationName string `json:"locationId" db:"location_name"`
	Price        int    `json:"price"`
}
