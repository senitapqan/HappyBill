package service

import "happyBill/models"

func (s *service) CreateBillboard(product models.Product) (int, error) {
	return s.repos.CreateBillboard(product)
}

func (s *service) GetAllBillboards() ([]models.Product, error) {
	return s.repos.GetAllBillboards()
}

func (s *service) DeleteBillboard(id int) error {
	return s.repos.DeleteBillboard(id)
}
