package service

import (
	"happyBill/models"
	mock_repository "happyBill/pkg/repository/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestService_CreateOrder(t *testing.T) {
	type mock1 func(r *mock_repository.MockRepository, id int, order models.Order)

	testTable := []struct {
		name  string
		id    int
		order models.Order
		mock1 mock1
		wantId  int
		wantError error
	}{
		{
			name: "OK",
			id: 1,
			mock1: func(r *mock_repository.MockRepository, id int, order models.Order) {
				r.EXPECT().GetMostFreeManager().Return(1, nil)
				r.EXPECT().CreateOrder(id, order).Return(1, nil)
			},
			order: models.Order{
				Id: 0,
				OrderedTime: "2024-04-23",
				Deadline: "2024-05-07",
				StartTime: "2024-05-10",
				EndTime: "2024-05-25",
				ProductId: 72,
				ManagerId: 1,
				ClientId: 1,
				Price: 150000,
			},
			wantId: 1,
			wantError: nil,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repository := mock_repository.NewMockRepository(c)
			test.mock1(repository, test.id, test.order)

			service := &service{repository}

			id, err := service.CreateOrder(test.id, test.order)
			
			assert.Equal(t, id, test.wantId)
			assert.Equal(t, err, test.wantError)
		})
	}
}
