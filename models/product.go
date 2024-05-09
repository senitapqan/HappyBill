package models

import "time"

type Product struct {
	Id          int       `json:"product_id"`
	Width       int       `json:"width" binding:"required"`
	Height      int       `json:"height" binding:"required"`
	DisplayType int       `json:"display_type" binding:"required" db:"display_type"`
	LocationId  int       `json:"location_id" binding:"required"`
	Price       int       `json:"price" binding:"required"`
	CreatedAt   time.Time `json:"created"`
	Archive     bool      `json:"archive" db:"archive"`
}

type Location struct {
	Id       int    `json:"location_id"`
	Name     string `json:"name"`
	RegionId int    `json:"region_id"`
	Link     string `json:"link"`
}

type Region struct {
	Id   int    `json:"region_id"`
	Name string `json:"name"`
}
