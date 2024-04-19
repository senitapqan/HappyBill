package models

import "time"

type Product struct {
	Id          int       `json:"productId"`
	Width       int       `json:"width" binding:"required"`
	Height      int       `json:"height" binding:"required"`
	DisplayType int       `json:"display_type" binding:"required" db:"display_type"`
	LocationId  int       `json:"locationId" binding:"required"`
	Price       int       `json:"price" binding:"required"`
	CreatedAt   time.Time `json:"created"`
}

type Location struct {
	Id       int    `json:"locationId"`
	Name     string `json:"name"`
	RegionId int    `json:"regionId"`
	Link     string `json:"link"`
}

type Region struct {
	Id   int    `json:"regionId"`
	Name string `json:"name"`
}
