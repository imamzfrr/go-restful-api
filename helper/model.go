package helper

import (
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

func ToEmployeeResponse(employee domain.Employee) web.EmployeeResponse {
	return web.EmployeeResponse{
		EmployeeID: employee.EmployeeID,
		Name:       employee.Name,
		Role:       employee.Role,
		Email:      employee.Email,
		Phone:      employee.Phone,
		DateHired:  employee.DateHired,
	}
}

func ToEmployeeResponses(employees []domain.Employee) []web.EmployeeResponse {
	var employeeResponses []web.EmployeeResponse
	for _, employee := range employees {
		employeeResponses = append(employeeResponses, ToEmployeeResponse(employee))
	}
	return employeeResponses
}

func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		ProductID:   product.ProductID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		StockQty:    product.StockQty,
		Category:    product.Category,
		SKU:         product.SKU,
		TaxRate:     product.TaxRate,
	}
}

func ToProductResponses(products []domain.Product) []web.ProductResponse {
	var productResponses []web.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}

func ToCustomerResponse(customer domain.Customer) web.CustomerResponse {
	return web.CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Email:      customer.Email,
		Phone:      customer.Phone,
		Address:    customer.Address,
		LoyaltyPts: customer.LoyaltyPts,
	}
}

func ToCustomerResponses(customers []domain.Customer) []web.CustomerResponse {
	var customerResponses []web.CustomerResponse
	for _, customer := range customers {
		customerResponses = append(customerResponses, ToCustomerResponse(customer))
	}
	return customerResponses
}
