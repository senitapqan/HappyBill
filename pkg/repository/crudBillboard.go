package repository

import (
	"fmt"
	"happyBill/consts"
	"happyBill/models"
	"log"
	"strings"

	"github.com/sirupsen/logrus"
)

func (r *repository) CreateBillboard(product models.Product) (int, error) {
	var productId int

	createBillboardQuery := fmt.Sprintf("INSERT INTO %s (width, height, display_type, locationId, price) VALUES ($1, $2, $3, $4, $5) RETURNING id", consts.ProductsTable)
	row := r.db.DB.QueryRow(createBillboardQuery, product.Width, product.Height, product.DisplayType, product.LocationId, product.Price)

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

func (r *repository) GetBillboardById(id int) (models.Product, error) {
	var product models.Product
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1 LIMIT 1", consts.ProductsTable)
	err := r.db.Get(&product, query, id)
	log.Print(err)
	return product, err

}

func (r *repository) DeleteBillboard(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", consts.ProductsTable)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *repository) UpdateBillboard(id int, input models.Product) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Width != 0 {
		setValues = append(setValues, fmt.Sprintf("size=$%d", argId))
		args = append(args, input.Width)
		argId++
	}

	if input.Height != 0 {
		setValues = append(setValues, fmt.Sprintf("size=$%d", argId))
		args = append(args, input.Height)
		argId++
	}

	if input.DisplayType != "" {
		setValues = append(setValues, fmt.Sprintf("display_type=$%d", argId))
		args = append(args, input.DisplayType)
		argId++
	}

	if input.Price != 0 {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, input.Price)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tp SET %s WHERE tp.id = $%d",
		consts.ProductsTable, setQuery, id)

	args = append(args, id)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
