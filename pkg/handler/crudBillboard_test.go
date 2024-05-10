package handler

import (
	"bytes"
	_ "errors"
	"happyBill/dtos"
	"happyBill/models"
	mock_service "happyBill/pkg/service/mocks"
	mock_validator "happyBill/pkg/handler/validatorMock"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestHandler_createBillboard(t *testing.T) {
	// Init Test Table
	type mockBehavior func(s *mock_service.MockService, product models.Product)
	tests := []struct {
		name                 string
		inputBody            string
		inputProduct         models.Product
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"width": 1, "height": 1, "display_type": 1, "location_id": 1, "price": 100}`,
			inputProduct: models.Product{
				Width:       1,
				Height:      1,
				DisplayType: 1,
				LocationId:  1,
				Price:       100,
			},
			mockBehavior: func(r *mock_service.MockService, product models.Product) {
				r.EXPECT().CreateBillboard(product).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"Message":"New billboard was created!","id":1}`,
		},
		{
			name:                 "Wrong input",
			inputBody:            `{"width": 1}`,
			inputProduct:         models.Product{},
			mockBehavior:         func(r *mock_service.MockService, product models.Product) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			services := mock_service.NewMockService(c)
			validator := mock_validator.NewMockValidator(c)
			test.mockBehavior(services, test.inputProduct)

			handler := Handler{services, validator}

			// Init Endpoint
			r := gin.New()
			r.POST("/admin/bill", handler.createBillboard)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/admin/bill",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_getAllSearchedBillboards(t *testing.T) {
	type mockBehavior func(s *mock_service.MockService, v *mock_validator.MockValidator, page int, search dtos.Search, filter dtos.Filter)
	tests := []struct {
		name                 string
		inputBody            string
		filter               dtos.Filter
		page                 int
		search               dtos.Search
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: "",
			filter:    dtos.Filter{},
			search:    dtos.Search{},
			mockBehavior: func(s *mock_service.MockService, v *mock_validator.MockValidator, page int, search dtos.Search, filter dtos.Filter) {
				v.EXPECT().ValidatePage(&gin.Context{}).Return(1)
				s.EXPECT().GetAllSearchedBillboards(page, search, filter).Return(1, nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			services := mock_service.NewMockService(c)
			validator := mock_validator.NewMockValidator(c)
			test.mockBehavior(services, validator, test.page, test.search, test.filter)

			handler := Handler{services, validator}

			// Init Endpoint
			r := gin.New()
			r.POST("/home/search", handler.getAllSearchedBillboards)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/home/search",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
