package handler

import (
	"bytes"
	"happyBill/models"
	mock_service "happyBill/pkg/service/mocks"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestHandler_CreateMyOrder(t *testing.T) {
	type mock func(r *mock_service.MockService, id int, order models.Order)

	testTable := []struct {
		name                 string
		id                   int
		inputBody            string
		order                models.Order
		mock                 mock
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK",
			id:   1,
			mock: func(r *mock_service.MockService, clientId int, order models.Order) {

				r.EXPECT().CreateOrder(clientId, order).Return(1, nil)
			},
			inputBody: `{"startdate": "2024-05-10", "enddate": "2024-05-25",  "product_id": 1}`,
			order: models.Order{
				StartTime: "2024-05-10",
				EndTime:   "2024-05-25",
				ProductId: 1,
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"message":"Order with id 1 was created and sent to manager for pending"}`,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			services := mock_service.NewMockService(c)
			test.mock(services, 1, test.order)

			handler := Handler{services}

			r := gin.New()
			r.POST("/my/order/:id", handler.createMyOrder)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/my/order/1",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}

}
