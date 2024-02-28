package service

import (
	"happyBill/pkg/repository"
	"happyBill/models"
)

type Service interface {
	GenerateToken(username, password string) (string, error)

	CreateClient(student models.User) (int, error)
}

type service struct {
	repos repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{
		repos: r,
	}
}
