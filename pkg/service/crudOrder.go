package service

import (
	"errors"
	"happyBill/dtos"
	"happyBill/models"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *service) GetAllOrders(page int) ([]dtos.Order, error) {
	//return s.repos.GetAllOrders()
	return nil, nil
}

func (s *service) GetMyOrders(clientId, page int, status string) ([]dtos.MyOrder, error) {
	log.Info().Msg("service send request to repository: get my orders request")
	myOrders, err := s.repos.GetMyOrders(clientId, page, status)
	if err != nil {
		return nil, err
	}
	if len(myOrders) == 0 {
		return nil, errors.New("there is no orders")
	}
	return myOrders, err
}

func (s *service) CreateOrder(id int, order models.Order) (int, error) {
	OrderedTime := time.Now()
	order.OrderedTime = OrderedTime.Format("2006-01-02")

	Deadline := OrderedTime.AddDate(0, 0, 14)
	order.Deadline = Deadline.Format("2006-01-02")

	log.Info().Msg("service send request to repository: get most free manager request")
	managerId, err := s.GetMostFreeManager()

	if err != nil {
		return -1, errors.New("something wrong with declaring manager for your order")
	}

	if managerId == 0 {
		return -1, errors.New("no managers to take order")
	}

	order.ManagerId = managerId
	order.ClientId = id

	log.Info().Msg("service send request to repository: create order request")
	return s.repos.CreateOrder(id, order)
}
