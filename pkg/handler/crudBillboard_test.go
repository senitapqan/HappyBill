package handler

// import (
// 	"bytes"
// 	_ "errors"
// 	"happyBill/models"
// 	mock_service "happyBill/pkg/service/mocks"
// 	"log"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang/mock/gomock"
// 	"github.com/magiconair/properties/assert"
// )

// func TestHandler_createBillboard(t *testing.T) {
// 	log.Print(time.Date(2024, 4, 19, 14, 49, 0, 0, time.Now().Local().Location()))
// 	// Init Test Table
// 	type mockBehavior func(s *mock_service.MockService, product models.Product)
// 	tests := []struct {
// 		name                 string
// 		inputBody            string
// 		inputProduct         models.Product
// 		mockBehavior         mockBehavior
// 		expectedStatusCode   int
// 		expectedResponseBody string
// 	}{
// 		{
// 			name:      "Ok",
// 			inputBody: `{"width": 1, "height": 1, "display_type": 1, "locationId": 1, "price": 100}`,
// 			inputProduct: models.Product{
// 				Width:       1,
// 				Height:      1,
// 				DisplayType: 1,
// 				LocationId:  1,
// 				Price:       100,
// 			},
// 			mockBehavior: func(r *mock_service.MockService, product models.Product) {
// 				r.EXPECT().CreateBillboard(product).Return(1, nil)
// 			},
// 			expectedStatusCode:   200,
// 			expectedResponseBody: `{"Message":"New billboard was created!","id":1}`,
// 		},
// 		{
// 			name:                 "Wrong input",
// 			inputBody:            `{"width": 1}`,
// 			inputProduct:         models.Product{},
// 			mockBehavior:         func(r *mock_service.MockService, product models.Product) {},
// 			expectedStatusCode:   400,
// 			expectedResponseBody: `{"message":"invalid input body"}`,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			// Init Dependencies
// 			c := gomock.NewController(t)
// 			defer c.Finish()

// 			services := mock_service.NewMockService(c)
// 			test.mockBehavior(services, test.inputProduct)

// 			handler := Handler{services}

// 			// Init Endpoint
// 			r := gin.New()
// 			r.POST("/admin/bill", handler.createBillboard)

// 			// Create Request
// 			w := httptest.NewRecorder()
// 			req := httptest.NewRequest("POST", "/admin/bill",
// 				bytes.NewBufferString(test.inputBody))

// 			// Make Request
// 			r.ServeHTTP(w, req)

// 			// Assert
// 			assert.Equal(t, w.Code, test.expectedStatusCode)
// 			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
// 		})
// 	}
// }
