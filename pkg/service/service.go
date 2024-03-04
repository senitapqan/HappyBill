package service

import (
	"happyBill/models"
	"happyBill/pkg/repository"
)

type Service interface {
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, []models.RolesHeaders, error)

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
