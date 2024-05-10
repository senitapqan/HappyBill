package service

import (
	"errors"
	"fmt"
	"happyBill/consts"
	"happyBill/dtos"
	"happyBill/models"

	"github.com/rs/zerolog/log"
)

func (s *service) CreateBillboard(product models.Product) (int, error) {
	log.Info().Msg("service send request to repository: create billboard request")
	product.MainPhoto = &product.Photos[0]

	id, err := s.repos.CreateBillboard(product)

	if err != nil {
		return -1, err
	}

	return id, nil
}

func (s *service) GetAllBillboards(page, clientId int) ([]dtos.Product, dtos.Pagination, error) {
	log.Info().Msg("service send request to repository: get all billboards request")
	return s.repos.GetAllBillboards(page, clientId)
}

func (s *service) GetAllSearchedBillboards(page int, search dtos.Search, filter dtos.Filter) ([]dtos.Product, dtos.Pagination, error) {
	log.Info().Msg("service send request to repository: get all searched billboards request")
	products, err := s.repos.GetAllSearchedBillboardsFake(filter)
	if err != nil {
		return nil, dtos.Pagination{}, err
	}

	orders, err := s.repos.GetAllOrders()
	if err != nil {
		return nil, dtos.Pagination{}, err
	}

	useless := make(map[int]bool)

	for _, order := range orders {
		if order.StartTime > search.CheckOut || search.CheckIn > order.EndTime {
			continue
		} else {
			useless[order.ProductId] = true
			log.Info().Msg(fmt.Sprintf("this product is shit: %d", order.ProductId))
		}
	}

	skipped := 0
	added := 0
	total := 0
	var result []dtos.Product
	for _, product := range products {
		if useless[product.Id] == false {
			if skipped < consts.PaginationLimit * (page - 1) {
				skipped++
				total++
				continue
			}
			if added < consts.PaginationLimit {
				result = append(result, product)
				added++
			}
			total++
		}
	}

	return result, dtos.Pagination{CurrentPage: page, TotalPage: (total + consts.PaginationLimit - 1) / consts.PaginationLimit}, nil
}

func (s *service) GetMyBillboards(id, page int) ([]dtos.Product, dtos.Pagination, error) {
	log.Info().Msg("service send request to repository: get my billboards (my fav) request")

	myBillboards, pagination, err := s.repos.GetMyBillboards(id, page)
	if err != nil {
		return nil, dtos.Pagination{}, err
	}
	if len(myBillboards) == 0 {
		return nil, dtos.Pagination{}, errors.New("there is no liked billboards")
	}

	return myBillboards, pagination, err
}

func (s *service) DeleteBillboard(id int) error {
	log.Info().Msg("service send request to repository: delete billboards request")
	return s.repos.DeleteBillboard(id)
}

func (s *service) GetBillboardById(id int) (dtos.Product, error) {
	log.Info().Msg("service send request to repository: get billboard by id request")
	return s.repos.GetBillboardById(id)
}

func (s *service) UpdateBillboard(id int, input dtos.Product) error {
	log.Info().Msg("service send request to repository: update billboard request")
	return s.repos.UpdateBillboard(id, input)
}

func (s *service) LikeBillboard(clientId, productId int, action string) error {
	log.Info().Msg("service send request to repository: like/dislike billboard request")
	return s.repos.LikeBillboard(clientId, productId, action)
}
