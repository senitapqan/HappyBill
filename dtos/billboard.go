package dtos

type GetAllBillboardsResponse struct {
	Data []Product `json:"data"`
}

type GetBillboardByIdResponse struct {
	Data Product `json:"data"`
}

type Product struct {
	Id           int    `json:"product_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	DisplayType  int    `json:"display_type" db:"display_type"`
	LocationName string `json:"location_id" db:"location_name"`
	Price        int    `json:"price"`
}

type Search struct {
	RegionId int
	CheckIn  string
	CheckOut string
}

type Filter struct {
	PriceIn     int
	PriceOut    int
	DisplayType []int
	WidthIn     int
	WidthOut    int
	HeightIn    int
	HeightOut   int
}
