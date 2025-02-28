package controller

import (
	"bytes"
	"encoding/json"
	"github.com/aronipurwanto/go-restful-api/model/web"
	"github.com/aronipurwanto/go-restful-api/service/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestAppProduct(mockService *mocks.MockProductService) *fiber.App {
	app := fiber.New()
	productController := NewProductController(mockService)

	api := app.Group("/api")
	products := api.Group("/products")
	products.Post("/", productController.Create)
	products.Put("/:productId", productController.Update)
	products.Delete("/:productId", productController.Delete)
	products.Get("/:productId", productController.FindById)
	products.Get("/", productController.FindAll)

	return app
}

func TestProductController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockProductService(ctrl)
	app := setupTestAppProduct(mockService)

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
			name:   "Create product - success",
			method: "POST",
			url:    "/api/products/",
			body:   web.ProductCreateRequest{Name: "Product A", Price: 1000},
			setupMock: func() {
				mockService.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(web.ProductResponse{ProductID: "1", Name: "Product A", Price: 1000}, nil)
			},
			expectedStatus: http.StatusCreated,
			expectedBody: web.WebResponse{
				Code:   http.StatusCreated,
				Status: "Created",
				Data: map[string]interface{}{
					"product_id":  "1",
					"name":        "Product A",
					"description": "",
					"price":       float64(1000),
					"stock_qty":   float64(0),
					"category":    "",
					"sku":         "",
					"tax_rate":    float64(0),
				},
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
			err := json.NewDecoder(resp.Body).Decode(&respBody)
			if err != nil {
				return
			}

			assert.Equal(t, tt.expectedBody, respBody)
		})
	}
}
