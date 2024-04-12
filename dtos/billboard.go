package dtos

import "happyBill/models"

type GetAllBillboardsResponse struct {
	Data []models.Product `json:"data"`
}

type GetBillboardByIdResponse struct {
	Data models.Product `json:"data"`
}
