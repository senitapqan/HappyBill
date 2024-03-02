package service

import (
	"happyBill/dtos"
	"happyBill/models"
)

func (s *service) CreateClient(client models.User) (int, error) {
	client.Password = s.hashPassword(client.Password)
	client_id, err := s.repos.CreateClient(client)
	return client_id, err
}

func (s service) DeleteStudent(lessonId int) error {
	return nil
}

func (s service) GetStudent(lessonId int) (dtos.User, error) {
	var student dtos.User
	return student, nil
}

func (s service) GetStudents() ([]dtos.User,  error) {
	return nil, nil
}