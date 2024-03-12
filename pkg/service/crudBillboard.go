package service

import "happyBill/models"

func (s *service) CreateBillboard(product models.Product) (int, error) {
	return s.repos.CreateBillboard(product)
}
