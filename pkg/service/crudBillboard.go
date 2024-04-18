package service

import (
	"happyBill/dtos"
	"happyBill/models"
)

func (s *service) CreateBillboard(product models.Product) (int, error) {
	return s.repos.CreateBillboard(product)
}

func (s *service) GetAllBillboards(page int) ([]dtos.Product, error) {
	return s.repos.GetAllBillboards(page)
}

func (s *service) DeleteBillboard(id int) error {
	return s.repos.DeleteBillboard(id)
}

func (s *service) GetBillboardById(id int) (dtos.Product, error) {
	return s.repos.GetBillboardById(id)
}

func (s *service) UpdateBillboard(id int, input models.Product) error {
	return s.repos.UpdateBillboard(id, input)
}
