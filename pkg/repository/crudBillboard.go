package repository

import (
	"errors"
	"fmt"
	"happyBill/consts"
	"happyBill/dtos"
	"happyBill/models"

	"strings"

	"github.com/rs/zerolog/log"
)

func (r *repository) CreateBillboard(product models.Product) (int, error) {
	var productId int

	createBillboardQuery := fmt.Sprintf("INSERT INTO %s (width, height, display_type, location_id, price) VALUES ($1, $2, $3, $4, $5) RETURNING id", consts.ProductsTable)
	row := r.db.QueryRow(createBillboardQuery, product.Width, product.Height, product.DisplayType, product.LocationId, product.Price)

	err := row.Scan(&productId)
	if err != nil {
		log.Error().Msg(err.Error())
		return -1, errors.New("something wrong with sql request")
	}

	return productId, nil

}

func (r *repository) GetAllSearchedBillboards(page int, search dtos.Search, filter dtos.Filter) ([]dtos.Product, dtos.Pagination, error) {
	var products []dtos.Product
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	setValues = append(setValues, "1 = 1")
	if search.RegionId != -1 {
		setValues = append(setValues, fmt.Sprintf("region_id = $%d", argId))
		args = append(args, search.RegionId)
		argId++
	}
	if search.CheckIn != "" {
		setValues = append(setValues, fmt.Sprintf("check_in > $%d", argId))
		args = append(args, search.CheckIn)
		argId++
	}
	if search.CheckOut != "" {
		setValues = append(setValues, fmt.Sprintf("check_out = $%d", argId))
		args = append(args, search.CheckOut)
		argId++
	}
	query := fmt.Sprintf(`select prod.id, prod.height, prod.width, prod.display_type, prod.price, loc.name as location_name, loc.link as link
			from %s prod
			join %s loc on loc.id = prod.location_id
			order by prod.created_time desc
			limit %d offset %d`,
		consts.ProductsTable, consts.LocationsTable, consts.PaginationLimit, (page-1)*consts.PaginationLimit)
		
	if err := r.db.Select(&products, query); err != nil {
		return nil, dtos.Pagination{}, err
	}

	var pagination dtos.Pagination
	pagination.CurrentPage = page
	var totalRows int
	
	query = fmt.Sprintf(`select count(*)
				from %s prod
				join %s loc on loc.id = prod.location_id`, 
				consts.ProductsTable, consts.LocationsTable)
	
	if err := r.db.Get(&totalRows, query); err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return nil, dtos.Pagination{}, errors.New("there is no billboards with such filters")
		}
		return nil, dtos.Pagination{}, err
	}
	pagination.TotalPage = (totalRows + consts.PaginationLimit - 1) / consts.PaginationLimit
	return products, pagination, nil
}

func (r *repository) GetAllBillboards(page int) ([]dtos.Product, dtos.Pagination, error) {
	var products []dtos.Product
	query := fmt.Sprintf(`select prod.id, prod.height, prod.width, prod.display_type, prod.price, loc.name as location_name, loc.link as link
			from %s prod
			join %s loc on loc.id = prod.location_id
			order by prod.created_time desc
			limit %d offset %d`,
		consts.ProductsTable, consts.LocationsTable, consts.PaginationLimit, (page-1)*consts.PaginationLimit)
	if err := r.db.Select(&products, query); err != nil {
		return nil, dtos.Pagination{}, err
	}

	var pagination dtos.Pagination
	pagination.CurrentPage = page
	var totalRows int
	
	query = fmt.Sprintf(`select count(*)
				from %s prod
				join %s loc on loc.id = prod.location_id`, 
				consts.ProductsTable, consts.LocationsTable)
	
	if err := r.db.Get(&totalRows, query); err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return nil, dtos.Pagination{}, errors.New("there is no billboards with such filters")
		}
		return nil, dtos.Pagination{}, err
	}
	pagination.TotalPage = (totalRows + consts.PaginationLimit - 1) / consts.PaginationLimit
	return products, pagination, nil
}

func (r *repository) GetMyBillboards(clientId, page int) ([]dtos.Product, dtos.Pagination, error) {
	var products []dtos.Product
	query := fmt.Sprintf(`select prod.id, prod.height, prod.width, prod.display_type, prod.price, loc.name as location_name, loc.link as link
			from %s prod
			join %s clprod on clprod.product_id = prod.id
			join %s loc on loc.id = prod.location_id
			where clprod.client_id = $1
			order by clprod.created_time desc
			limit %d offset %d`,
		consts.ProductsTable, consts.ClientProductsTable, consts.LocationsTable, consts.PaginationLimit, (page-1)*consts.PaginationLimit)
	if err := r.db.Select(&products, query, clientId); err != nil {
		return nil, dtos.Pagination{}, err
	}
	
	var pagination dtos.Pagination
	pagination.CurrentPage = page
	var totalRows int
	
	query = fmt.Sprintf(`select count(*) from %s prod
			join %s clprod on clprod.product_id = prod.id
			join %s loc on loc.id = prod.location_id
			where clprod.client_id = $1`, 
		consts.ProductsTable, consts.ClientProductsTable, consts.LocationsTable)

	if err := r.db.Get(&totalRows, query, clientId); err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return nil, dtos.Pagination{}, errors.New("there is no billboards with such filters")
		}
		return nil, dtos.Pagination{}, err
	}

	pagination.TotalPage = (totalRows + consts.PaginationLimit - 1) / consts.PaginationLimit
	return products, pagination, nil
}

func (r *repository) GetBillboardById(id int) (dtos.Product, error) {
	var product dtos.Product
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1 LIMIT 1", consts.ProductsTable)
	err := r.db.Get(&product, query, id)
	return product, err

}

func (r *repository) LikeBillboard(clientId, productId int, action string) error {
	var query string
	if action == "like" {
		query = fmt.Sprintf("insert into %s (client_id, product_id) values ($1, $2)", consts.ClientProductsTable)
	} else {
		query = fmt.Sprintf("delete from %s where client_id = $1 and product_id = $2", consts.ClientProductsTable)
	}

	_, err := r.db.Exec(query, clientId, productId)
	return err
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
