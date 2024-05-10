package service

import (
	"errors"
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
	return s.repos.GetAllSearchedBillboards(page, search, filter)
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
