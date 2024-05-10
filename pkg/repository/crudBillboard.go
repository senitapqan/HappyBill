package repository

import (
	"errors"
	"fmt"
	"happyBill/consts"
	"happyBill/dtos"
	"happyBill/models"
	"strconv"

	"strings"

	"github.com/rs/zerolog/log"
)

func (r *repository) CreateBillboard(product models.Product) (int, error) {
	var productId int

	createBillboardQuery := fmt.Sprintf("INSERT INTO %s (width, height, display_type, location_id, price, main_photo) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", consts.ProductsTable)
	row := r.db.QueryRow(createBillboardQuery, product.Width, product.Height, product.DisplayType, product.LocationId, product.Price, product.MainPhoto)

	err := row.Scan(&productId)
	if err != nil {
		log.Error().Msg(err.Error())
		return -1, errors.New("something wrong with sql request")
	}
	var args string
	for i, val := range product.Photos {
		if i+1 == len(product.Photos) {
			args += "(" + strconv.Itoa(productId) + ", '" + val + "')"
		} else {
			args += "(" + strconv.Itoa(productId) + ", '" + val + "'), "
		}
	}

	query := fmt.Sprintf(`insert into %s (product_id, photo) values %s`, consts.ProductPhotosTable, args)
	_, err = r.db.Exec(query)
	if err != nil {
		return 0, err
	}
	return productId, nil

}

func (r *repository) GetAllSearchedBillboardsFake(filter dtos.Filter) ([]dtos.Product, error) {
	var products []dtos.Product
	query := fmt.Sprintf(`select prod.id, prod.height, prod.width, prod.display_type, prod.price, loc.name as location_name, loc.link as link, prod.main_photo
			from %s prod
			join %s loc on loc.id = prod.location_id
			WHERE prod.archive = false and prod.height >= $1 and prod.height <= $2 and prod.width >= $3 and prod.width <= $4 
			and prod.price >= $5 and prod.price <= $6`, consts.ProductsTable, consts.LocationsTable)
	if err := r.db.Select(&products, query, filter.HeightIn, filter.HeightOut, filter.WidthIn, filter.WidthOut, filter.PriceIn, filter.PriceOut); err != nil {
		return nil, err
	}
	return products, nil
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
	query := fmt.Sprintf(`select prod.id, prod.height, prod.width, prod.display_type, prod.price, loc.name as location_name, loc.link as link, prod.main_photo
			from %s prod
			join %s loc on loc.id = prod.location_id
			WHERE prod.archive = false
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
				join %s loc on loc.id = prod.location_id
				WHERE prod.archive = false`,
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

func (r *repository) GetAllBillboards(page, clinetId int) ([]dtos.Product, dtos.Pagination, error) {
	var products []dtos.Product
	var likedInfo1 string
	var likedInfo2 string
	likedInfo1 = ""
	likedInfo2 = ""
	if clinetId == -1 {
		likedInfo1 = fmt.Sprintf(`left join %s likes on prod.id = likes.product_id and likes.client_id = %d`, consts.ClientProductsTable, clinetId)
		likedInfo2 = `, case when likes.product_id is not null then true else false end as liked`
	}
	query := fmt.Sprintf(`select prod.id, prod.height, prod.width, prod.display_type, prod.price, loc.name as location_name, loc.link as link, prod.main_photo %s
			from %s prod
			join %s loc on loc.id = prod.location_id
			%s
			WHERE prod.archive = false
			order by prod.created_time desc
			limit %d offset %d`,
		likedInfo2, consts.ProductsTable, consts.LocationsTable, likedInfo1, consts.PaginationLimit, (page-1)*consts.PaginationLimit)
	if err := r.db.Select(&products, query); err != nil {
		return nil, dtos.Pagination{}, err
	}

	var pagination dtos.Pagination
	pagination.CurrentPage = page
	var totalRows int

	query = fmt.Sprintf(`select count(*) from %s WHERE archive = false`, consts.ProductsTable)

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
	query := fmt.Sprintf(`select prod.id, prod.height, prod.width, prod.display_type, prod.price, loc.name as location_name, loc.link as link, prod.main_photo
			from %s prod
			join %s clprod on clprod.product_id = prod.id
			join %s loc on loc.id = prod.location_id
			WHERE prod.archive = false and clprod.client_id = $1
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
			WHERE prod.archive = false and clprod.client_id = $1`,
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
	query := fmt.Sprintf(`select prod.id, prod.height, prod.width, prod.display_type, prod.price, loc.name as location_name, loc.link as link, prod.main_photo
						from %s prod
						join %s loc on loc.id = prod.location_id 
						WHERE prod.archive = false and prod.id=$1 LIMIT 1`,
		consts.ProductsTable, consts.LocationsTable)
	if err := r.db.Get(&product, query, id); err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return dtos.Product{}, errors.New("there is no billboards with such id")
		}
		return dtos.Product{}, err
	}
	query = fmt.Sprintf(`select photo from %s where product_id = $1`, consts.ProductPhotosTable)
	if err := r.db.Select(&product.Photos, query, id); err != nil {
		return dtos.Product{}, err
	}

	return product, nil

}

func (r *repository) LikeBillboard(clientId, productId int, action string) error {
	var query string
	if action == "like" {
		query = fmt.Sprintf("insert into %s (client_id, product_id) values ($1, $2)", consts.ClientProductsTable)
	} else {
		query = fmt.Sprintf("delete from %s where client_id = $1 and product_id = $2", consts.ClientProductsTable)
	}

	_, err := r.db.Exec(query, clientId, productId)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return errors.New("you already liked this billboard")
		}
		return err
	}
	return nil
}

func (r *repository) DeleteBillboard(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", consts.ProductsTable)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *repository) UpdateBillboard(id int, input dtos.Product) error {
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
