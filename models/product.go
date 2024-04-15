package models

type Product struct {
	Id          int `json:"productId"`
	Width       int `json:"width"`
	Height      int `json:"height"`
	DisplayType int `json:"display_type" db:"display_type"`
	LocationId  int `json:"locationId"`
	Price       int `json:"price"`
}

type Location struct {
	Id       int    `json:"locationId"`
	Name     string `json:"name"`
	RegionId int    `json:"regionId"`
}

type Region struct {
	Id   int    `json:"regionId"`
	Name string `json:"name"`
}
