package service

import (
	"errors"
	"happyBill/models"
	mock_repository "happyBill/pkg/repository/mocks"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestService_CreateOrder(t *testing.T) {
	type mock func(r *mock_repository.MockRepository, id int, order models.Order)

	testTable := []struct {
		name      string
		id        int
		order     models.Order
		mock      mock
		wantId    int
		wantError error
	}{
		{
			name: "OK",
			id:   1,
			mock: func(r *mock_repository.MockRepository, id int, order models.Order) {
				r.EXPECT().GetMostFreeManager().Return(1, nil)
				r.EXPECT().CreateOrder(id, order).Return(1, nil)
			},
			order: models.Order{
				Id:          0,
				OrderedTime: time.Now().Format("2006-01-02"),
				Deadline:    time.Now().AddDate(0, 0, 14).Format("2006-01-02"),
				StartTime:   "2024-05-10",
				EndTime:     "2024-05-25",
				ProductId:   72,
				ManagerId:   1,
				ClientId:    1,
				Price:       150000,
			},
			wantId:    1,
			wantError: nil,
		},
		{
			name: "No managers to take order",
			id:   1,
			mock: func(r *mock_repository.MockRepository, id int, order models.Order) {
				r.EXPECT().GetMostFreeManager().Return(0, nil)
			},
			order: models.Order{
				Id:          0,
				OrderedTime: time.Now().Format("2006-01-02"),
				Deadline:    time.Now().AddDate(0, 0, 14).Format("2006-01-02"),
				StartTime:   "2024-05-10",
				EndTime:     "2024-05-25",
				ProductId:   72,
				ManagerId:   0,
				ClientId:    1,
				Price:       150000,
			},
			wantId:    -1,
			wantError: errors.New("no managers to take order"),
		},
		{
			name: "Error from repository",
			id:   1,
			mock: func(r *mock_repository.MockRepository, id int, order models.Order) {
				r.EXPECT().GetMostFreeManager().Return(0, errors.New("something went wrong in request"))
			},
			order: models.Order{
				Id:          0,
				OrderedTime: time.Now().Format("2006-01-02"),
				Deadline:    time.Now().AddDate(0, 0, 14).Format("2006-01-02"),
				StartTime:   "2024-05-10",
				EndTime:     "2024-05-25",
				ProductId:   72,
				ManagerId:   0,
				ClientId:    1,
				Price:       150000,
			},
			wantId:    -1,
			wantError: errors.New("something wrong with declaring manager for your order"),
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repository := mock_repository.NewMockRepository(c)
			test.mock(repository, test.id, test.order)

			service := &service{repository}

			id, err := service.CreateOrder(test.id, test.order)

			assert.Equal(t, id, test.wantId)
			assert.Equal(t, err, test.wantError)
		})
	}
}
