package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/model/web"
	"github.com/aronipurwanto/go-restful-api/repository/mocks"
	"github.com/aronipurwanto/go-restful-api/service"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	mockValidator := validator.New()
	productService := service.NewProductService(mockRepo, mockValidator)

	tests := []struct {
		name      string
		input     web.ProductCreateRequest
		mock      func()
		expect    web.ProductResponse
		expectErr bool
	}{
		{
			name:  "success",
			input: web.ProductCreateRequest{Name: "Laptop", Description: "Gaming Laptop", Price: 15000000, StockQty: 100, Category: "Gaming Laptop", SKU: "4"},
			mock: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(domain.Product{ProductID: "1", Name: "Laptop", Description: "Gaming Laptop", Price: 15000000, StockQty: 100, Category: "Gaming Laptop", SKU: "4"}, nil)
			},
			expect:    web.ProductResponse{ProductID: "1", Name: "Laptop", Description: "Gaming Laptop", Price: 15000000, StockQty: 100, Category: "Gaming Laptop", SKU: "4"},
			expectErr: false,
		},
		{
			name:  "repository error",
			input: web.ProductCreateRequest{Name: "Laptop", Description: "Gaming Laptop", Price: 15000000},
			mock: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(domain.Product{}, errors.New("repository error"))
			},
			expect:    web.ProductResponse{},
			expectErr: true,
		},
		{
			name:      "validation error",
			input:     web.ProductCreateRequest{Name: "Laptop", Description: "Gaming Laptop", Price: 15000000, StockQty: 100, Category: "Gaming Laptop", SKU: "4"},
			mock:      func() {},
			expect:    web.ProductResponse{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			resp, err := productService.Create(context.Background(), tt.input)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, resp)
			}
		})
	}
}

func TestDeleteProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	mockValidator := validator.New()
	productService := service.NewProductService(mockRepo, mockValidator)

	tests := []struct {
		name      string
		inputID   string
		mock      func()
		expectErr bool
	}{
		{
			name:    "success",
			inputID: "1",
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), "1").Return(domain.Product{ProductID: "1", Name: "Laptop"}, nil)
				mockRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectErr: false,
		},
		{
			name:    "not found",
			inputID: "999",
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), "999").Return(domain.Product{}, errors.New("product not found"))
			},
			expectErr: true,
		},
		{
			name:    "repository error on delete",
			inputID: "2",
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), "2").Return(domain.Product{ProductID: "2", Name: "Mouse"}, nil)
				mockRepo.EXPECT().Delete(gomock.Any(), domain.Product{ProductID: "2", Name: "Mouse"}).Return(errors.New("failed to delete product"))

			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := productService.Delete(context.Background(), tt.inputID)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestFindAllProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	productService := service.NewProductService(mockRepo, validator.New())

	mockRepo.EXPECT().FindAll(gomock.Any()).Return([]domain.Product{{ProductID: "1", Name: "Alice"}}, nil)

	resp, err := productService.FindAll(context.Background())

	assert.NoError(t, err)
	assert.Len(t, resp, 1)
	assert.Equal(t, "Alice", resp[0].Name)
}

func TestUpdateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	mockValidator := validator.New()
	productService := service.NewProductService(mockRepo, mockValidator)

	tests := []struct {
		name      string
		input     web.ProductUpdateRequest
		mock      func()
		expect    web.ProductResponse
		expectErr bool
	}{
		{
			name:  "success",
			input: web.ProductUpdateRequest{ProductID: "1", Name: "Laptop", Description: "Gaming Laptop", Price: 15000000, StockQty: 100, Category: "Gaming Laptop", SKU: "4"},
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(domain.Product{ProductID: "1", Name: "Laptop", Description: "Gaming Laptop", Price: 15000000, StockQty: 100, Category: "Gaming Laptop", SKU: "4"}, nil)
				mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(domain.Product{ProductID: "1", Name: "Laptop", Description: "Gaming Laptop", Price: 15000000, StockQty: 100, Category: "Gaming Laptop", SKU: "4"}, nil)

			},
			expect:    web.ProductResponse{ProductID: "1", Name: "Laptop", Description: "Gaming Laptop", Price: 15000000, StockQty: 100, Category: "Gaming Laptop", SKU: "4"},
			expectErr: false,
		},
		{
			name:  "repository error",
			input: web.ProductUpdateRequest{ProductID: "1", Name: "Laptop", Description: "Gaming Laptop", Price: 15000000, StockQty: 100, Category: "Gaming Laptop", SKU: "4"},
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(domain.Product{ProductID: "1", Name: "Laptop", Description: "Gaming Laptop", Price: 15000000, StockQty: 100, Category: "Gaming Laptop", SKU: "4"}, nil)
				mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(domain.Product{}, errors.New("repository error"))
			},
			expect:    web.ProductResponse{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			resp, err := productService.Update(context.Background(), tt.input)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, resp)
			}
		})
	}
}

func TestFindByIdProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	productService := service.NewProductService(mockRepo, validator.New())

	tests := []struct {
		name      string
		id        string
		mock      func()
		expect    web.ProductResponse
		expectErr bool
	}{
		{
			name: "success",
			id:   "1",
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), "1").Return(domain.Product{ProductID: "1", Name: "Alice"}, nil)
			},
			expect:    web.ProductResponse{ProductID: "1", Name: "Alice"},
			expectErr: false,
		},
		{
			name: "not found",
			id:   "2",
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), "2").Return(domain.Product{}, errors.New("not found"))
			},
			expect:    web.ProductResponse{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			resp, err := productService.FindById(context.Background(), tt.id)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, resp)
			}
		})
	}
}
