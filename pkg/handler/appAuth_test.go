package handler

// import (
// 	"bytes"
// 	"errors"
// 	"happyBill/models"
// 	mock_service "happyBill/pkg/service/mocks"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang/mock/gomock"
// 	"github.com/magiconair/properties/assert"
// )

// func TestHandler_signUp(t *testing.T) {
// 	// Init Test Table
// 	type mockBehavior func(s *mock_service.MockService, user models.User)
// 	tests := []struct {
// 		name                 string
// 		inputBody            string
// 		inputUser            models.User
// 		mockBehavior         mockBehavior
// 		expectedStatusCode   int
// 		expectedResponseBody string
// 	}{
// 		{
// 			name:      "Ok",
// 			inputBody: `{"username": "username", "name": "Test Name", "password": "qwerty", "email": "test@gmail.com", "surname": "Test Surname"}`,
// 			inputUser: models.User{
// 				Username: "username",
// 				Name:     "Test Name",
// 				Password: "qwerty",
// 				Email:    "test@gmail.com",
// 				Surname:  "Test Surname",
// 			},
// 			mockBehavior: func(r *mock_service.MockService, user models.User) {
// 				r.EXPECT().CreateClient(user).Return(1, nil)
// 			},
// 			expectedStatusCode:   200,
// 			expectedResponseBody: `{"new user was succesfully added with id":1}`,
// 		},
// 		{
// 			name:                 "Wrong Input",
// 			inputBody:            `{"username": "username"}`,
// 			inputUser:            models.User{},
// 			mockBehavior:         func(r *mock_service.MockService, user models.User) {},
// 			expectedStatusCode:   400,
// 			expectedResponseBody: `{"message":"invalid input body"}`,
// 		},
// 		{
// 			name:      "Service Error",
// 			inputBody: `{"username": "username", "name": "Test Name", "password": "qwerty", "email": "test@gmail.com", "surname": "Test Surname"}`,
// 			inputUser: models.User{
// 				Username: "username",
// 				Name:     "Test Name",
// 				Password: "qwerty",
// 				Email:    "test@gmail.com",
// 				Surname:  "Test Surname",
// 			},
// 			mockBehavior: func(r *mock_service.MockService, user models.User) {
// 				r.EXPECT().CreateClient(user).Return(0, errors.New("error with repository"))
// 			},
// 			expectedStatusCode:   500,
// 			expectedResponseBody: `{"message":"something went wrong: error with repository"}`,
// 		},
// 		{
// 			name:      "Existing username",
// 			inputBody: `{"username": "username", "name": "Test Name", "password": "qwerty", "email": "test@gmail.com", "surname": "Test Surname"}`,
// 			inputUser: models.User{
// 				Username: "username",
// 				Name:     "Test Name",
// 				Password: "qwerty",
// 				Email:    "test@gmail.com",
// 				Surname:  "Test Surname",
// 			},
// 			mockBehavior: func(r *mock_service.MockService, user models.User) {
// 				r.EXPECT().CreateClient(user).Return(0, errors.New("there is already exist account with such username"))
// 			},
// 			expectedStatusCode:   500,
// 			expectedResponseBody: `{"message":"something went wrong: there is already exist account with such username"}`,
// 		},
// 		{
// 			name:      "Existing email",
// 			inputBody: `{"username": "username", "name": "Test Name", "password": "qwerty", "email": "test@gmail.com", "surname": "Test Surname"}`,
// 			inputUser: models.User{
// 				Username: "username",
// 				Name:     "Test Name",
// 				Password: "qwerty",
// 				Email:    "test@gmail.com",
// 				Surname:  "Test Surname",
// 			},
// 			mockBehavior: func(r *mock_service.MockService, user models.User) {
// 				r.EXPECT().CreateClient(user).Return(0, errors.New("there is already exist account with such email"))
// 			},
// 			expectedStatusCode:   500,
// 			expectedResponseBody: `{"message":"something went wrong: there is already exist account with such email"}`,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			// Init Dependencies
// 			c := gomock.NewController(t)
// 			defer c.Finish()

// 			services := mock_service.NewMockService(c)
// 			test.mockBehavior(services, test.inputUser)

// 			handler := Handler{services}

// 			// Init Endpoint
// 			r := gin.New()
// 			r.POST("/sign-up", handler.signUp)

// 			// Create Request
// 			w := httptest.NewRecorder()
// 			req := httptest.NewRequest("POST", "/sign-up",
// 				bytes.NewBufferString(test.inputBody))

// 			// Make Request
// 			r.ServeHTTP(w, req)

// 			// Assert
// 			assert.Equal(t, w.Code, test.expectedStatusCode)
// 			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
// 		})
// 	}
// }
