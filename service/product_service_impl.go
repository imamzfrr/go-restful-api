package service

import (
	"context"
	"errors"

	"github.com/aronipurwanto/go-restful-api/exception"
	"github.com/aronipurwanto/go-restful-api/helper"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/model/web"
	"github.com/aronipurwanto/go-restful-api/repository"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &ProductServiceImpl{ProductRepository: productRepository}
}

// Create Product
func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) (web.ProductResponse, error) {
	product := domain.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		StockQty:    request.StockQty,
		Category:    request.Category,
		SKU:         request.SKU,
		TaxRate:     request.TaxRate,
	}
	newProduct, err := service.ProductRepository.Save(ctx, product)
	if err != nil {
		return web.ProductResponse{}, err
	}
	return helper.ToProductResponse(newProduct), nil
}

// Update Product
func (service *ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) (web.ProductResponse, error) {
	// Cari product berdasarkan ID
	product, err := service.ProductRepository.FindById(ctx, request.ProductID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return web.ProductResponse{}, exception.NewNotFoundError("Product not found")
	} else if err != nil {
		return web.ProductResponse{}, err
	}

	// Update field-field product
	product.Name = request.Name
	product.Description = request.Description
	product.Price = request.Price
	product.StockQty = request.StockQty
	product.Category = request.Category
	product.SKU = request.SKU
	product.TaxRate = request.TaxRate

	updatedProduct, err := service.ProductRepository.Update(ctx, product)
	if err != nil {
		return web.ProductResponse{}, err
	}
	return helper.ToProductResponse(updatedProduct), nil
}

// Delete Product
func (service *ProductServiceImpl) Delete(ctx context.Context, productId string) error {
	product, err := service.ProductRepository.FindById(ctx, productId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return exception.NewNotFoundError("Product not found")
	} else if err != nil {
		return err
	}
	return service.ProductRepository.Delete(ctx, product)
}

// Find Product By ID
func (service *ProductServiceImpl) FindById(ctx context.Context, productId string) (web.ProductResponse, error) {
	product, err := service.ProductRepository.FindById(ctx, productId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return web.ProductResponse{}, exception.NewNotFoundError("Product not found")
	} else if err != nil {
		return web.ProductResponse{}, err
	}
	return helper.ToProductResponse(product), nil
}

// Find All Products
func (service *ProductServiceImpl) FindAll(ctx context.Context) ([]web.ProductResponse, error) {
	products, err := service.ProductRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return helper.ToProductResponses(products), nil
}
