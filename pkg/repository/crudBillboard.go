package repository

import (
	"fmt"
	"happyBill/consts"
	"happyBill/dtos"
	"happyBill/models"
	"log"
	"time"

	"strings"
)

func (r *repository) CreateBillboard(product models.Product) (int, error) {
	var productId int

	createBillboardQuery := fmt.Sprintf("INSERT INTO %s (width, height, display_type, locationId, price, created_time) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", consts.ProductsTable)
	row := r.db.DB.QueryRow(createBillboardQuery, product.Width, product.Height, product.DisplayType, product.LocationId, product.Price, time.Now())

	err := row.Scan(&productId)
	if err != nil {
		return -1, err
	}

	return productId, nil

}

func (r *repository) GetAllBillboards(page int) ([]dtos.Product, error) {
	var products []dtos.Product
	query := fmt.Sprintf(`select prod.id, prod.height, prod.width, prod.display_type, prod.price, loc.name as location_name
			from %s prod
			join %s loc on loc.id = prod.locationid
			order by prod.created_time desc
			limit %d offset %d`, 
			consts.ProductsTable, consts.LocationsTable, consts.PaginationLimit, (page - 1) * consts.PaginationLimit)
	if err := r.db.Select(&products, query); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repository) GetBillboardById(id int) (dtos.Product, error) {
	var product dtos.Product
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1 LIMIT 1", consts.ProductsTable)
	err := r.db.Get(&product, query, id)
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
		setValues = append(setValues, fmt.Sprintf("width=$%d", argId))
		args = append(args, input.Width)
		argId++
	}

	if input.Height != 0 {
		setValues = append(setValues, fmt.Sprintf("height=$%d", argId))
		args = append(args, input.Height)
		argId++
	}

	if input.DisplayType != 0 {
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

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = %d",
		consts.ProductsTable, setQuery, id)

	_, err := r.db.Exec(query, args...)
	log.Printf("args: %s query: %s", args, query)
	return err
}
