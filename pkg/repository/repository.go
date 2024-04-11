package repository

import (
	"happyBill/models"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	GetUser(username string) (models.User, error)
	GetRoles(id int) ([]string, error)
	GetRoleId(role string, userId int) (int, error)

	CreateClient(student models.User) (int, error)

	CreateBillboard(product models.Product) (int, error)
	GetAllBillboards() ([]models.Product, error)
	GetBillboardById(id int) (models.Product, error)
	DeleteBillboard(id int) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}
