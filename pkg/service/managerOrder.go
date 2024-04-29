package service

import (
	"happyBill/dtos"

	"github.com/rs/zerolog/log"
)

func (s *service) GetAllManagerOrders(id, page int) ([]dtos.ManagerOrder, error) {
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
