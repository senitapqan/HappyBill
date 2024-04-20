package service

import (
	"errors"
	"happyBill/dtos"
)

func (s *service) GetAllOrders(page int) ([]dtos.Order, error) {
	//return s.repos.GetAllOrders()
	return nil, nil
}

func (s *service) GetMyOrders(clientId, page int) ([]dtos.MyOrder, error) {
	myOrders, err := s.repos.GetMyOrders(clientId, page)
	if len(myOrders) == 0 {
		return nil, errors.New("There is no orders")
	}
	return myOrders, err
}

// func (s *service) CreateOrder() (int, error) {
// 	return s.repos.CreateOrder()
// }