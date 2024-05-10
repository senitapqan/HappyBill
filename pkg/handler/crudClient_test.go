package handler

import (
	"bytes"
	"happyBill/dtos"
	mock_service "happyBill/pkg/service/mocks"
	mock_validator "happyBill/pkg/handler/validatorMock"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestHandler_updateMyProfile(t *testing.T) {
	type mockBehavior func(s *mock_service.MockService, id int, input dtos.UpdateUser)
	tests := []struct {
		id                   int
		name                 string
		inputBody            string
		inputUser            dtos.UpdateUser
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			id:        0,
			name:      "Ok",
			inputBody: `{"username": "username", "name": "Test Name", "password": "qwerty", "old_password": "oldqwerty", "surname": "Test Surname"}`,
			inputUser: dtos.UpdateUser{
				Username:    "username",
				Name:        "Test Name",
				Password:    "qwerty",
				OldPassword: "oldqwerty",
				Surname:     "Test Surname",
			},
			mockBehavior: func(r *mock_service.MockService, id int, input dtos.UpdateUser) {
				r.EXPECT().UpdateMyProfile(id, input).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"Message":"Updated succesfully"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			services := mock_service.NewMockService(c)
			validator := mock_validator.NewMockValidator(c)
			test.mockBehavior(services, test.id, test.inputUser)

			
			handler := Handler{services, validator}

			// Init Endpoint
			r := gin.New()
			r.PUT("/my/", handler.updateMyProfile)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/my/",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}

}
