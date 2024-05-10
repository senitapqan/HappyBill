package repository

import (
	"happyBill/dtos"
	"happyBill/models"

	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go
type Repository interface {
	GetUser(username string) (models.User, error)
	GetUserById(id int) (models.User, error)
	GetRoles(id int) ([]string, error)
	GetRoleId(role string, userId int) (int, error)

	CreateClient(client models.User) (int, error)
	GetClientById(id int) (dtos.User, error)
	GetClientByUserId(id int) (dtos.User, error)

	CreateManager(manager models.User) (int, error)
	GetAllManagers(page int) ([]dtos.User, dtos.Pagination, error)
	GetManagerById(id int) (dtos.User, error)
	GetMostFreeManager() (int, error)
	/*UpdateManager(id int, input models.User) error
	DeleteManager(id int) error*/

	CreateBillboard(product models.Product) (int, error)
	GetAllBillboards(page int) ([]dtos.Product, dtos.Pagination, error)
	GetAllSearchedBillboards(page int, search dtos.Search, filter dtos.Filter) ([]dtos.Product, dtos.Pagination, error)
	GetMyBillboards(id, page int) ([]dtos.Product, dtos.Pagination, error)
	GetBillboardById(id int) (dtos.Product, error)
	DeleteBillboard(id int) error
	UpdateBillboard(id int, input dtos.Product) error
	LikeBillboard(clientId, productId int, action string) error

	//GetAllOrders() ([]dtos.Order, error)
	CreateOrder(clientId int, order models.Order) (int, error)
	GetMyOrders(clientId, page int, status string) ([]dtos.MyOrder, dtos.Pagination, error)
	UpdateMyProfile(userId int, input dtos.UpdateUser) error

	GetAllManagerOrders(id, page int) ([]dtos.ManagerOrder, dtos.Pagination,error)
	GetManagerOrderById(id int) (dtos.ManagerOrder, error)
	UpdateManagerOrder(id int, input dtos.UpdateOrder) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}
