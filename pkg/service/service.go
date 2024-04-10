package service

import (
	"happyBill/models"
	"happyBill/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Service interface {
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, []models.RolesHeaders, error)

	CreateClient(student models.User) (int, error)

	CreateBillboard(product models.Product) (int, error)
	GetAllBillboards() ([]models.Product, error)
	DeleteBillboard(id int) error
}

type service struct {
	repos repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{
		repos: r,
	}
}
