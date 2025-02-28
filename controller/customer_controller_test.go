package controller_test

import (
	"bytes"
	"encoding/json"
	"github.com/aronipurwanto/go-restful-api/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aronipurwanto/go-restful-api/model/web"
	"github.com/aronipurwanto/go-restful-api/service/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func setupTestAppCustomer(mockService *mocks.MockCustomerService) *fiber.App {
	app := fiber.New()
	customerController := controller.NewCustomerController(mockService)

	api := app.Group("/api")
	customers := api.Group("/customers")
	customers.Post("/", customerController.Create)
	customers.Put("/:customerId", customerController.Update)
	customers.Delete("/:customerId", customerController.Delete)
	customers.Get("/:customerId", customerController.FindById)
	customers.Get("/", customerController.FindAll)

	return app
}

func TestCustomerController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockCustomerService(ctrl)
	app := setupTestAppCustomer(mockService)

	tests := []struct {
		name           string
		method         string
		url            string
		body           interface{}
		setupMock      func()
		expectedStatus int
		expectedBody   web.WebResponse
	}{
		{
			name:   "Create customer - success",
			method: "POST",
			url:    "/api/customers/",
			body:   web.CustomerCreateRequest{Name: "John Doe", Email: "john@example.com"},
			setupMock: func() {
				mockService.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(web.CustomerResponse{CustomerID: "1", Name: "John Doe", Email: "john@example.com"}, nil)
			},
			expectedStatus: http.StatusCreated,
			expectedBody: web.WebResponse{
				Code:   http.StatusCreated,
				Status: "Created",
				Data:   web.CustomerResponse{CustomerID: "1", Name: "John Doe", Email: "john@example.com"},
			},
		},
		{
			name:   "Find customer by ID - success",
			method: "GET",
			url:    "/api/customers/1",
			body:   nil,
			setupMock: func() {
				mockService.EXPECT().
					FindById(gomock.Any(), "1").
					Return(web.CustomerResponse{CustomerID: "1", Name: "John Doe", Email: "john@example.com"}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: web.WebResponse{
				Code:   http.StatusOK,
				Status: "OK",
				Data:   web.CustomerResponse{CustomerID: "1", Name: "John Doe", Email: "john@example.com"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			var reqBody []byte
			if tt.body != nil {
				reqBody, _ = json.Marshal(tt.body)
			}

			req := httptest.NewRequest(tt.method, tt.url, bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")

			resp, _ := app.Test(req)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			var respBody web.WebResponse
			json.NewDecoder(resp.Body).Decode(&respBody)

			if dataMap, ok := respBody.Data.(map[string]interface{}); ok {
				customerID, _ := dataMap["customer_id"].(string)
				respBody.Data = web.CustomerResponse{
					CustomerID: customerID,
					Name:       dataMap["name"].(string),
					Email:      dataMap["email"].(string),
				}
			}

			assert.Equal(t, tt.expectedBody, respBody)
		})
	}
}
