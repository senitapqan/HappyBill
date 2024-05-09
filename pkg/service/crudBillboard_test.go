package service

import (
	"errors"
	"happyBill/models"
	mock_repository "happyBill/pkg/repository/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestService_CreateBillboard(t *testing.T) {
	type mock func(r *mock_repository.MockRepository, id int, product models.Product)

	testTable := []struct {
		name  string
		id    int
		product models.Product
		mock mock
		wantId  int
		wantError error
	}{
		{
			name: "OK",
			id: 1,
			mock: func(r *mock_repository.MockRepository, id int, product models.Product) {
				r.EXPECT().CreateBillboard(product).Return(1, nil)
			},
			
			product: models.Product{
				Width:       1,
				Height:      1,
				DisplayType: 1,
				LocationId:  1,
				Price:       100,
			},
			wantId: 1,
			wantError: nil,
		},
		{
			name: "error with db",
			id: 1,
			mock: func(r *mock_repository.MockRepository, id int,  product models.Product) {
				r.EXPECT().CreateBillboard(product).Return(-1, errors.New("something wrong with sql request")) 
			},
			product: models.Product{
				Width:       1,
				Height:      1,
				DisplayType: 1,
				LocationId:  1,
				Price:       100,
			},
			wantId: -1,
			wantError: errors.New("something wrong with sql request"),
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repository := mock_repository.NewMockRepository(c)
			test.mock(repository, test.id, test.product)

			service := &service{repository}

			id, err := service.CreateBillboard(test.product)
			
			assert.Equal(t, id, test.wantId)
			assert.Equal(t, err, test.wantError)
		})
	}
}


func TestService_DeleteBillboard(t *testing.T) {
	type mock func(r *mock_repository.MockRepository, id int, product models.Product)

	testTable := []struct {
		name  string
		id    int
		product models.Product
		mock mock
		wantId  int
		wantError error
	}{
		{
			name: "OK",
			id: 1,
			mock: func(r *mock_repository.MockRepository, id int, product models.Product) {
				r.EXPECT().DeleteBillboard(id).Return(nil)
			},

			
			product: models.Product{
				Width:       1,
				Height:      1,
				DisplayType: 1,
				LocationId:  1,
				Price:       100,
			},
			wantId: 1,
			wantError: nil,
		},
		{
			name: "error with db",
			id: 1,
			mock: func(r *mock_repository.MockRepository, id int,  product models.Product) {
				r.EXPECT().CreateBillboard(id).Return(errors.New("something wrong with sql request")) 
			},
			product: models.Product{
				Width:       1,
				Height:      1,
				DisplayType: 1,
				LocationId:  1,
				Price:       100,
			},
			wantId: -1,
			wantError: errors.New("something wrong with sql request"),
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repository := mock_repository.NewMockRepository(c)
			test.mock(repository, test.id, test.product)

			service := &service{repository}

			id, err := service.CreateBillboard(test.product)
			
			assert.Equal(t, id, test.wantId)
			assert.Equal(t, err, test.wantError)
		})
	}
}


func TestService_UpdateBillboard(t *testing.T) {
	type mock func(r *mock_repository.MockRepository, id int, product models.Product)

	testTable := []struct {
		name  string
		id    int
		product models.Product
		mock mock
		wantId  int
		wantError error
	}{
		{
			name: "OK",
			id: 1,
			mock: func(r *mock_repository.MockRepository, id int, product models.Product) {
				r.EXPECT().CreateBillboard(product).Return(1, nil)
			},
			
			product: models.Product{
				Width:       1,
				Height:      1,
				DisplayType: 1,
				LocationId:  1,
				Price:       100,
			},
			wantId: 1,
			wantError: nil,
		},
		{
			name: "error with db",
			id: 1,
			mock: func(r *mock_repository.MockRepository, id int,  product models.Product) {
				r.EXPECT().CreateBillboard(product).Return(-1, errors.New("something wrong with sql request")) 
			},
			product: models.Product{
				Width:       1,
				Height:      1,
				DisplayType: 1,
				LocationId:  1,
				Price:       100,
			},
			wantId: -1,
			wantError: errors.New("something wrong with sql request"),
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repository := mock_repository.NewMockRepository(c)
			test.mock(repository, test.id, test.product)

			service := &service{repository}

			id, err := service.CreateBillboard(test.product)
			
			assert.Equal(t, id, test.wantId)
			assert.Equal(t, err, test.wantError)
		})
	}
}

