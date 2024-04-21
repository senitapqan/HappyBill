package service

import (
	"errors"
	"happyBill/dtos"
	"happyBill/models"
	"strings"
)

func (s *service) CreateClient(client models.User) (int, error) {
	client.Password = s.hashPassword(client.Password)
	client_id, err := s.repos.CreateClient(client)
	if err != nil {
		msg := err.Error()
		if strings.Contains(msg, "duplicate") {
			if strings.Contains(msg, "email") {
				return -1, errors.New("there is already exist account with such email")
			}
			if strings.Contains(msg, "username") {
				return -1, errors.New("there is already exist account with such username")
			}
		}
		return 0, errors.New("error with repository")
	}
	return client_id, err
}

func (s service) DeleteClient(lessonId int) error {
	return nil
}

func (s service) GetClientById(id int) (dtos.User, error) {
	return s.repos.GetClientById(id)
}

func (s service) GetClientByUserId(id int) (dtos.User, error) {
	return s.repos.GetClientByUserId(id)
}

func (s service) GetStudents() ([]dtos.User,  error) {
	return nil, nil
}