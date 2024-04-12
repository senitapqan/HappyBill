package service

import (
	"happyBill/dtos"
	"happyBill/models"
	"happyBill/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Service interface {
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, []models.RolesHeaders, error)

	CreateClient(client models.User) (int, error)

	CreateManager(manager models.User) (int, error)
	GetAllManagers() ([]dtos.User, error)
	GetManagerById(id int) (dtos.User, error)
	UpdateManager(id int, input models.User) error
	DeleteManager(id int) error
 
	CreateBillboard(product models.Product) (int, error)
	GetAllBillboards() ([]models.Product, error)
	GetBillboardById(id int) (models.Product, error)
	DeleteBillboard(id int) error
	UpdateBillboard(id int, input models.Product) error
}

type service struct {
	repos repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{
		repos: r,
	}
}
