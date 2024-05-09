package handler

import (
	"bytes"
	"errors"
	"happyBill/models"
	mock_service "happyBill/pkg/service/mocks"
	mock_validator "happyBill/pkg/handler/validatorMock"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestHandler_createManager(t *testing.T) {
	type mockBehavior func(s *mock_service.MockService, user models.User)
	tests := []struct {
		name                 string
		inputBody            string
		inputUser            models.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"username": "username", "name": "Test Name", "password": "qwerty", "email": "test@gmail.com", "surname": "Test Surname"}`,
			inputUser: models.User{
				Username: "username",
				Name:     "Test Name",
				Password: "qwerty",
				Email:    "test@gmail.com",
				Surname:  "Test Surname",
			},
			mockBehavior: func(r *mock_service.MockService, user models.User) {
				r.EXPECT().CreateManager(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"new Manager was succesfully added with id":1}`,
		},
		{
			name:                 "Wrong input",
			inputUser:            models.User{},
			mockBehavior:         func(r *mock_service.MockService, user models.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "error",
			inputBody: `{"username": "username", "name": "Test Name", "password": "qwerty", "email": "test@gmail.com", "surname": "Test Surname"}`,
			inputUser: models.User{
				Username: "username",
				Name:     "Test Name",
				Password: "qwerty",
				Email:    "test@gmail.com",
				Surname:  "Test Surname",
			},
			mockBehavior: func(r *mock_service.MockService, user models.User) {
				r.EXPECT().CreateManager(user).Return(0, errors.New("something went wrong with repository"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			services := mock_service.NewMockService(c)
			validator := mock_validator.NewMockValidator(c)
			test.mockBehavior(services, test.inputUser)

			handler := Handler{services, validator}

			// Init Endpoint
			r := gin.New()
			r.POST("/manager/", handler.createManager)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/manager/",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}

}
