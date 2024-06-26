package service

import (
	"errors"
	"happyBill/dtos"
	"happyBill/models"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *service) GetAllOrders(page int) ([]dtos.Order, dtos.Pagination, error) {
	//return s.repos.GetAllOrders()
	return nil, dtos.Pagination{},  nil
}

func (s *service) GetMyOrders(clientId, page int, status string) ([]dtos.MyOrder, dtos.Pagination, error) {
	log.Info().Msg("service send request to repository: get my orders request")
	myOrders, pagination, err := s.repos.GetMyOrders(clientId, page, status)
	if err != nil {
		return nil, dtos.Pagination{}, err
	}
	if len(myOrders) == 0 {
		return nil, dtos.Pagination{}, errors.New("there is no orders")
	}
	return myOrders, pagination, err
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

func (s *service) GetAllManagerOrders(id, page int) ([]dtos.ManagerOrder, dtos.Pagination, error) {
	log.Info().Msg("service send request to repository: get all manager orders request")
	return s.repos.GetAllManagerOrders(id, page)
}

func (s *service) GetManagerOrderById(id int) (dtos.ManagerOrder, error) {
	log.Info().Msg("service send request to repository: get manager order by id request")
	return s.repos.GetManagerOrderById(id)
}

func (s *service) UpdateManagerOrder(id int, input dtos.UpdateOrder) error {
	log.Info().Msg("service send request to repository: update manager order request")
	return s.repos.UpdateManagerOrder(id, input)
}

