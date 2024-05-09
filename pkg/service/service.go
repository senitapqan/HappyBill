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
	GetAllManagers(page int) ([]dtos.User, dtos.Pagination, error)
	GetManagerById(id int) (dtos.User, error) 

	CreateOrder(id int, order models.Order) (int, error)
	GetAllOrders(page int) ([]dtos.Order, dtos.Pagination, error)
	GetMyOrders(clientId, page int, status string) ([]dtos.MyOrder, dtos.Pagination, error)
	UpdateMyProfile(userId int, input dtos.UpdateUser) error

	CreateBillboard(product models.Product) (int, error)
	GetAllBillboards(page int) ([]dtos.Product, dtos.Pagination, error)
	GetAllSearchedBillboards(page int, search dtos.Search, filter dtos.Filter) ([]dtos.Product, dtos.Pagination, error)
	GetBillboardById(id int) (dtos.Product, error)
	DeleteBillboard(id int) error
	UpdateBillboard(id int, input models.Product) error
	GetMyBillboards(id, page int) ([]dtos.Product, dtos.Pagination, error)
	LikeBillboard(clientId, productId int, action string) error

	GetAllManagerOrders(id, page int) ([]dtos.ManagerOrder, dtos.Pagination, error)
	GetManagerOrderById(id int) (dtos.ManagerOrder, error)
	UpdateManagerOrder(id int, input dtos.UpdateOrder) error
}

type service struct {
	repos repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{
		repos: r,
	}
}
