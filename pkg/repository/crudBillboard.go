package repository

import (
	"fmt"
	"happyBill/consts"
	"happyBill/models"
)

func (r *repository) CreateBillboard(product models.Product) (int, error) {
	var productId int

	createBillboardQuery := fmt.Sprintf("INSERT INTO %s (size, display_type, locationId) VALUES ($1, $2, $3) RETURNING id", consts.ProductsTable)
	row := r.db.DB.QueryRow(createBillboardQuery, product.Size, product.DisplayType, product.LocationId)

	err := row.Scan(&productId)
	if err != nil {
		return -1, err
	}

	return productId, nil

}

func (r *repository) GetAllBillboards() ([]models.Product, error) {
	var products []models.Product
	query := fmt.Sprintf("SELECT * FROM %s", consts.ProductsTable)
	err := r.db.Get(&products, query)
	return products, err

}

func (r *repository) DeleteBillboard(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", consts.ProductsTable)
	_, err := r.db.Exec(query, id)
	return err
}
