package dtos

type GetAllBillboardsResponse struct {
	Data       []Product  `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type GetBillboardByIdResponse struct {
	Data Product `json:"data"`
}

type Product struct {
	Id           int      `json:"product_id"`
	Width        int      `json:"width"`
	Height       int      `json:"height"`
	DisplayType  int      `json:"display_type" db:"display_type"`
	LocationName string   `json:"location_name" db:"location_name"`
	LocationLink string   `json:"location_link" db:"link"`
	Price        int      `json:"price"`
	MainPhoto    *string  `json:"main_photo" db:"main_photo"`
	Photos       []string `json:"photos_link" db:"photos_link"`
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
