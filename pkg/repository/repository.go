package repository

import (
	"happyBill/dtos"
	"happyBill/models"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetUser(username string) (models.User, error)
	GetRoles(id int) ([]string, error)
	GetRoleId(role string, userId int) (int, error)

	CreateClient(client models.User) (int, error)

	CreateManager(manager models.User) (int, error)
	GetAllManagers() ([]dtos.User, error)
	GetManagerById(id int) (dtos.User, error)
	/*UpdateManager(id int, input models.User) error
	DeleteManager(id int) error*/

	CreateBillboard(product models.Product) (int, error)
	GetAllBillboards() ([]models.Product, error)
	GetBillboardById(id int) (models.Product, error)
	DeleteBillboard(id int) error
	UpdateBillboard(id int, input models.Product) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}
