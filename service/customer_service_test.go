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

func TestCreateCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCustomerRepository(ctrl)
	mockValidator := validator.New()
	customerService := service.NewCustomerService(mockRepo, mockValidator)

	tests := []struct {
		name      string
		input     web.CustomerCreateRequest
		mock      func()
		expect    web.CustomerResponse
		expectErr bool
	}{
		{
			name:  "success",
			input: web.CustomerCreateRequest{Name: "John Doe", Email: "john@example.com", Phone: "123456789", Address: "Street 123", LoyaltyPts: 10},
			mock: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(domain.Customer{CustomerID: "1", Name: "John Doe", Email: "john@example.com", Phone: "123456789", Address: "Street 123", LoyaltyPts: 10}, nil)
			},
			expect:    web.CustomerResponse{CustomerID: "1", Name: "John Doe", Email: "john@example.com", Phone: "123456789", Address: "Street 123", LoyaltyPts: 10},
			expectErr: false,
		},
		{
			name:      "validation error",
			input:     web.CustomerCreateRequest{Name: ""},
			mock:      func() {},
			expect:    web.CustomerResponse{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			resp, err := customerService.Create(context.Background(), tt.input)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, resp)
			}
		})
	}
}

func TestUpdateCustomer(t *testing.T) {
	tests := []struct {
		name      string
		input     web.CustomerUpdateRequest
		mock      func(mockCustomerRepo *mocks.MockCustomerRepository)
		expect    web.CustomerResponse
		expectErr bool
	}{
		{
			name: "success",
			input: web.CustomerUpdateRequest{
				CustomerID: "1", Name: "John Doe Updated", Email: "johnupdated@example.com",
				Phone: "987654321", Address: "New Street 123", LoyaltyPts: 20,
			},
			mock: func(mockCustomerRepo *mocks.MockCustomerRepository) {
				gomock.InOrder(
					mockCustomerRepo.EXPECT().FindById(gomock.Any(), "1").Return(domain.Customer{CustomerID: "1"}, nil),
					mockCustomerRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(
						domain.Customer{
							CustomerID: "1", Name: "John Doe Updated", Email: "johnupdated@example.com",
							Phone: "987654321", Address: "New Street 123", LoyaltyPts: 20,
						}, nil,
					),
				)
			},
			expect: web.CustomerResponse{
				CustomerID: "1", Name: "John Doe Updated", Email: "johnupdated@example.com",
				Phone: "987654321", Address: "New Street 123", LoyaltyPts: 20,
			},
			expectErr: false,
		},
		{
			name:  "not found",
			input: web.CustomerUpdateRequest{CustomerID: "99", Name: "Unknown", Email: "lalala@gmail.com", Phone: "7662346"},
			mock: func(mockCustomerRepo *mocks.MockCustomerRepository) {
				mockCustomerRepo.EXPECT().FindById(gomock.Any(), "99").Return(domain.Customer{}, errors.New("not found"))
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockCustomerRepo := mocks.NewMockCustomerRepository(ctrl)
			tt.mock(mockCustomerRepo)

			service := service.NewCustomerService(mockCustomerRepo, validator.New())
			resp, err := service.Update(context.Background(), tt.input)

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, resp)
			}
		})
	}
}

func TestDeleteCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCustomerRepository(ctrl)
	customerService := service.NewCustomerService(mockRepo, validator.New())

	tests := []struct {
		name       string
		customerId string
		mock       func()
		expectErr  bool
	}{
		{
			name:       "success",
			customerId: "1",
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), "1").Return(domain.Customer{CustomerID: "1", Name: "John Doe"}, nil)
				mockRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectErr: false,
		},
		{
			name:       "not found",
			customerId: "99",
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), "99").Return(domain.Customer{}, errors.New("not found"))
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := customerService.Delete(context.Background(), tt.customerId)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
func TestFindCustomerById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCustomerRepository(ctrl)
	customerService := service.NewCustomerService(mockRepo, validator.New())

	tests := []struct {
		name       string
		customerId string
		mock       func()
		expect     web.CustomerResponse
		expectErr  bool
	}{
		{
			name:       "success",
			customerId: "1",
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), "1").Return(domain.Customer{CustomerID: "1", Name: "John Doe", Email: "john@example.com", Phone: "123456789", Address: "Street 123", LoyaltyPts: 10}, nil)
			},
			expect:    web.CustomerResponse{CustomerID: "1", Name: "John Doe", Email: "john@example.com", Phone: "123456789", Address: "Street 123", LoyaltyPts: 10},
			expectErr: false,
		},
		{
			name:       "not found",
			customerId: "99",
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), "99").Return(domain.Customer{}, errors.New("not found"))
			},
			expect:    web.CustomerResponse{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			resp, err := customerService.FindById(context.Background(), tt.customerId)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, resp)
			}
		})
	}
}

func TestFindAllCustomers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCustomerRepository(ctrl)
	customerService := service.NewCustomerService(mockRepo, validator.New())

	mockRepo.EXPECT().FindAll(gomock.Any()).Return([]domain.Customer{
		{CustomerID: "1", Name: "John Doe"},
		{CustomerID: "2", Name: "Jane Doe"},
	}, nil)

	resp, err := customerService.FindAll(context.Background())
	assert.NoError(t, err)
	assert.Len(t, resp, 2)
}
