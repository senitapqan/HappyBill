package service

import (
	"happyBill/dtos"
	"happyBill/models"
)

func (s *service) CreateManager(manager models.User) (int, error) {
	manager.Password = s.hashPassword(manager.Password)
	manager_id, err := s.repos.CreateManager(manager)
	return manager_id, err
}

func (s *service) GetAllManagers() ([]dtos.User, error) {
	return s.repos.GetAllManagers()
}

func (s *service) GetManagerById(id int) (dtos.User, error) {
	return s.repos.GetManagerById(id)
}

func (s *service) DeleteManager(id int) error {
	return s.repos.DeleteManager(id)
}

func (s *service) UpdateManager(id int, input models.User) error {
	return s.repos.UpdateManager(id, input)
}