package service

import (
	"happyBill/dtos"
	"happyBill/models"

	"github.com/rs/zerolog/log"
)

func (s *service) CreateManager(manager models.User) (int, error) {
	manager.Password = s.hashPassword(manager.Password)
	log.Info().Msg("service send request to repository: create manager request")
	manager_id, err := s.repos.CreateManager(manager)
	return manager_id, err
}

func (s *service) GetAllManagers(page int) ([]dtos.User, dtos.Pagination, error) {
	log.Info().Msg("service send request to repository: get all managers request")
	return s.repos.GetAllManagers(page)
}

func (s *service) GetManagerById(id int) (dtos.User, error) {
	log.Info().Msg("service send request to repository: get manager by id request")
	return s.repos.GetManagerById(id)
}

func (s *service) GetMostFreeManager() (int, error)  {
	log.Info().Msg("service send request to repository: get most free manager request")
	return s.repos.GetMostFreeManager()
}
/*
func (s *service) DeleteManager(id int) error {
	return s.repos.DeleteManager(id)
}

func (s *service) UpdateManager(id int, input models.User) error {
	return s.repos.UpdateManager(id, input)
}*/