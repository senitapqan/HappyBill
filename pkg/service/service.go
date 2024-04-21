package service

import (
	"happyBill/dtos"
	"happyBill/models"
	"happyBill/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Service interface {
	GenerateToken(username, password string) ([]models.RolesHeaders, string, error)
	ParseToken(accessToken string) (int, []models.RolesHeaders, error)

	CreateClient(client models.User) (int, error)
	GetClientByUserId(id int) (dtos.User, error)
	GetClientById(id int) (dtos.User, error)

	CreateManager(manager models.User) (int, error)
	GetAllManagers(page int) ([]dtos.User, error)
	GetManagerById(id int) (dtos.User, error) /*
		UpdateManager(id int, input models.User) error
		DeleteManager(id int) error
	*/

	//CreateOrder(order models.Order) (int, error)
	GetAllOrders(page int) ([]dtos.Order, error)
	GetMyOrders(clientId, page int) ([]dtos.MyOrder, error)
	UpdateMyProfile(userId int, input dtos.UpdateUser) error

	CreateBillboard(product models.Product) (int, error)
	GetAllBillboards(page int) ([]dtos.Product, error)
	GetMyBillboards(id, page int) ([]dtos.Product, error)
	GetBillboardById(id int) (dtos.Product, error)
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
