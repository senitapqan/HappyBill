package models

type Product struct {
	Id          int    `json:"productId"`
	Size        int    `json:"size"`
	DisplayType string `json:"display_type"`
	LocationId  int    `json:"locationId"`
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
