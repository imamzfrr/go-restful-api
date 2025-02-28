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

func TestCreateEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockEmployeeRepository(ctrl)
	mockValidator := validator.New()
	employeeService := service.NewEmployeeService(mockRepo, mockValidator)

	tests := []struct {
		name      string
		input     web.EmployeeCreateRequest
		mock      func()
		expect    web.EmployeeResponse
		expectErr bool
	}{
		{
			name:  "success",
			input: web.EmployeeCreateRequest{Name: "Alice", Role: "Developer", Email: "alice@example.com", Phone: "2534356", DateHired: "12/10/2025"},
			mock: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(domain.Employee{EmployeeID: "1", Name: "Alice", Role: "Developer", Email: "alice@example.com"}, nil)
			},
			expect:    web.EmployeeResponse{EmployeeID: "1", Name: "Alice", Role: "Developer", Email: "alice@example.com"},
			expectErr: false,
		},
		{
			name:  "repository error",
			input: web.EmployeeCreateRequest{Name: "Alice", Role: "Developer", Email: "alice@example.com", Phone: "2534356", DateHired: "12/10/2025"},
			mock: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(domain.Employee{}, errors.New("repository error"))
			},
			expect:    web.EmployeeResponse{},
			expectErr: true,
		},
		{
			name:      "validation error",
			input:     web.EmployeeCreateRequest{Name: "", Role: "Developer", Email: "alice@example.com", Phone: "2534356", DateHired: "12/10/2025"},
			mock:      func() {},
			expect:    web.EmployeeResponse{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			resp, err := employeeService.Create(context.Background(), tt.input)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, resp)
			}
		})
	}
}

func TestDeleteEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockEmployeeRepository(ctrl)
	employeeService := service.NewEmployeeService(mockRepo, validator.New())

	tests := []struct {
		name      string
		id        string
		mock      func()
		expectErr bool
	}{
		{
			name: "success",
			id:   "1",
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), "1").Return(domain.Employee{EmployeeID: "1", Name: "Alice"}, nil)
				mockRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectErr: false,
		},
		{
			name: "not found",
			id:   "1",
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), "1").Return(domain.Employee{}, errors.New("not found"))

			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := employeeService.Delete(context.Background(), tt.id)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestFindByIdEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockEmployeeRepository(ctrl)
	employeeService := service.NewEmployeeService(mockRepo, validator.New())

	tests := []struct {
		name      string
		id        string
		mock      func()
		expect    web.EmployeeResponse
		expectErr bool
	}{
		{
			name: "success",
			id:   "1",
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), "1").Return(domain.Employee{EmployeeID: "1", Name: "Alice"}, nil)
			},
			expect:    web.EmployeeResponse{EmployeeID: "1", Name: "Alice"},
			expectErr: false,
		},
		{
			name: "not found",
			id:   "2",
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), "2").Return(domain.Employee{}, errors.New("not found"))
			},
			expect:    web.EmployeeResponse{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			resp, err := employeeService.FindById(context.Background(), tt.id)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, resp)
			}
		})
	}
}

func TestUpdateEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockEmployeeRepository(ctrl)
	mockValidator := validator.New()
	employeeService := service.NewEmployeeService(mockRepo, mockValidator)

	tests := []struct {
		name      string
		input     web.EmployeeUpdateRequest
		mock      func()
		expect    web.EmployeeResponse
		expectErr bool
	}{
		{
			name:  "success",
			input: web.EmployeeUpdateRequest{EmployeeID: "1", Name: "Alice", Role: "Developer", Email: "alice@example.com", Phone: "2534356", DateHired: "12/10/2025"},
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(domain.Employee{EmployeeID: "1", Name: "Alice", Role: "Developer", Email: "alice@example.com", Phone: "2534356", DateHired: "12/10/2025"}, nil)
				mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(domain.Employee{EmployeeID: "1", Name: "Alice", Role: "Developer", Email: "alice@example.com", Phone: "2534356", DateHired: "12/10/2025"}, nil)

			},
			expect:    web.EmployeeResponse{EmployeeID: "1", Name: "Alice", Role: "Developer", Email: "alice@example.com", Phone: "2534356", DateHired: "12/10/2025"},
			expectErr: false,
		},
		{
			name:  "repository error",
			input: web.EmployeeUpdateRequest{EmployeeID: "1", Name: "Alice", Role: "Developer", Email: "alice@example.com", Phone: "2534356", DateHired: "12/10/2025"},
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(domain.Employee{EmployeeID: "1", Name: "Alice", Role: "Developer", Email: "alice@example.com", Phone: "2534356", DateHired: "12/10/2025"}, nil)
				mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(domain.Employee{}, errors.New("repository error"))
			},
			expect:    web.EmployeeResponse{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			resp, err := employeeService.Update(context.Background(), tt.input)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, resp)
			}
		})
	}
}

func TestFindAllEmployees(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockEmployeeRepository(ctrl)
	employeeService := service.NewEmployeeService(mockRepo, validator.New())

	mockRepo.EXPECT().FindAll(gomock.Any()).Return([]domain.Employee{{EmployeeID: "1", Name: "Alice"}}, nil)

	resp, err := employeeService.FindAll(context.Background())

	assert.NoError(t, err)
	assert.Len(t, resp, 1)
	assert.Equal(t, "Alice", resp[0].Name)
}
