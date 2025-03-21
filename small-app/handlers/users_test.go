package handlers

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"rest-api/middleware"
	"rest-api/models"
	"rest-api/models/mockmodels"
	"strings"
	"testing"
)

func TestSignup(t *testing.T) {

	// sample newUser that we would get after reading JSON body
	newUser := models.NewUser{
		Name:     "John Doe",
		Age:      25,
		Email:    "johndoe@example.com",
		Password: "your_secure_password",
	}
	// User that we would get from CreateUser method of models package
	mockUser := models.User{
		Id:           "8a0e6759-edbc-4e83-902f-499f7afa9ad3",
		Email:        "johndoe@example.com",
		Name:         "John Doe",
		Age:          25,
		PasswordHash: "$2a$10$dJ7UpJeX3AN0YVfvXu3WOu07IdIID1tcQCp4xo6WR.wkjaclM5wfi",
	}

	tt := []struct {
		name             string
		body             []byte
		expectedStatus   int
		expectedResponse string
		mockStore        func(m *mockmodels.MockService)
	}{
		{
			name: "OK",
			body: []byte(`{
  						"name": "John Doe",
 						 "email": "johndoe@example.com",
						"age": 25,
						"password": "your_secure_password"}`),
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"id":"8a0e6759-edbc-4e83-902f-499f7afa9ad3","email":"johndoe@example.com","name":"John Doe","age":25,"password_hash":"$2a$10$dJ7UpJeX3AN0YVfvXu3WOu07IdIID1tcQCp4xo6WR.wkjaclM5wfi"}`,
			mockStore: func(m *mockmodels.MockService) {
				m.EXPECT().CreateUser(gomock.Eq(newUser)).Return(mockUser, nil).Times(1)
			},
		},
		{
			name: "Fail_NoEmail",
			body: []byte(`{
  						"Name": "John Doe",
  						"Age": 30,
  						"Password": "abc"
  				}`),
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: `{"error":"Key: 'NewUser.Email' Error:Field validation for 'Email' failed on the 'required' tag"}`,
			mockStore: func(m *mockmodels.MockService) {
				m.EXPECT().CreateUser(gomock.Any()).Times(0)
			},
		},
	}
	ctrl := gomock.NewController(t)

	//NewMockService would return the implementation of the interface
	mockService := mockmodels.NewMockService(ctrl)
	h := handler{conn: mockService}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockStore(mockService)
			traceId := `fake-test-trace-id`
			ctx := context.Background()
			ctx = context.WithValue(ctx, middleware.TraceIdKey, traceId)

			r := httptest.NewRequestWithContext(ctx, http.MethodPost, "/user/signup", bytes.NewReader(tc.body))
			w := httptest.NewRecorder()

			h.Signup(w, r)

			require.Equal(t, tc.expectedStatus, w.Code)
			require.Equal(t, tc.expectedResponse, strings.TrimSpace(w.Body.String()))
		})
	}

}
