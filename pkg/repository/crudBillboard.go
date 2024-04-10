package repository

import (
	"fmt"
	"happyBill/consts"
	"happyBill/models"
	"strings"

	"github.com/sirupsen/logrus"
)

func (r *repository) CreateBillboard(product models.Product) (int, error) {
	var productId int

	createBillboardQuery := fmt.Sprintf("INSERT INTO %s (size, display_type, locationId, price) VALUES ($1, $2, $3, $4) RETURNING id", consts.ProductsTable)
	row := r.db.DB.QueryRow(createBillboardQuery, product.Size, product.DisplayType, product.LocationId, product.Price)

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

func (r *repository) UpdateBillboard(id int, input models.Product) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Size != 0 {
		setValues = append(setValues, fmt.Sprintf("size=$%d", argId))
		args = append(args, input.Size)
		argId++
	}

	if input.DisplayType != "" {
		setValues = append(setValues, fmt.Sprintf("display_type=$%d", argId))
		args = append(args, input.DisplayType)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id = $%d AND ul.user_id = $%d",
		todoListsTable, setQuery, usersListsTable, argId, argId+1)

	args = append(args, listId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
