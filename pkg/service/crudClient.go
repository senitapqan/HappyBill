package service

import (
	"happyBill/dtos"
	"happyBill/models"
)

func (s *service) CreateClient(student models.User) (int, error) {
	student.Password = s.hashPassword(student.Password)
	student_id, err := s.repos.CreateClient(student)
	return student_id, err
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