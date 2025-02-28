package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestEmployeeRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockEmployeeRepository(ctrl)
	ctx := context.Background()

	tests := []struct {
		name      string
		mock      func()
		method    func() (interface{}, error)
		expect    interface{}
		expectErr bool
	}{
		{
			name: "Save Success",
			mock: func() {
				employee := domain.Employee{EmployeeID: "E001", Name: "John Doe", Role: "Developer"}
				repo.EXPECT().Save(ctx, employee).Return(employee, nil)
			},
			method: func() (interface{}, error) {
				return repo.Save(ctx, domain.Employee{EmployeeID: "E001", Name: "John Doe", Role: "Developer"})
			},
			expect:    domain.Employee{EmployeeID: "E001", Name: "John Doe", Role: "Developer"},
			expectErr: false,
		},
		{
			name: "Save Failure",
			mock: func() {
				repo.EXPECT().Save(ctx, gomock.Any()).Return(domain.Employee{}, errors.New("error saving"))
			},
			method: func() (interface{}, error) {
				return repo.Save(ctx, domain.Employee{EmployeeID: "E002", Name: "Invalid"})
			},
			expect:    domain.Employee{},
			expectErr: true,
		},
		{
			name: "Update Success",
			mock: func() {
				employee := domain.Employee{EmployeeID: "E001", Name: "John Doe Updated", Role: "Senior Developer"}
				repo.EXPECT().Update(ctx, employee).Return(employee, nil)
			},
			method: func() (interface{}, error) {
				return repo.Update(ctx, domain.Employee{EmployeeID: "E001", Name: "John Doe Updated", Role: "Senior Developer"})
			},
			expect:    domain.Employee{EmployeeID: "E001", Name: "John Doe Updated", Role: "Senior Developer"},
			expectErr: false,
		},
		{
			name: "FindById Success",
			mock: func() {
				repo.EXPECT().FindById(ctx, "E001").Return(domain.Employee{EmployeeID: "E001", Name: "John Doe"}, nil)
			},
			method: func() (interface{}, error) {
				return repo.FindById(ctx, "E001")
			},
			expect:    domain.Employee{EmployeeID: "E001", Name: "John Doe"},
			expectErr: false,
		},
		{
			name: "FindById Not Found",
			mock: func() {
				repo.EXPECT().FindById(ctx, "E999").Return(domain.Employee{}, errors.New("employee not found"))
			},
			method: func() (interface{}, error) {
				return repo.FindById(ctx, "E999")
			},
			expect:    domain.Employee{},
			expectErr: true,
		},
		{
			name: "FindAll Success",
			mock: func() {
				repo.EXPECT().FindAll(ctx).Return([]domain.Employee{{EmployeeID: "E001", Name: "John Doe"}}, nil)
			},
			method: func() (interface{}, error) {
				return repo.FindAll(ctx)
			},
			expect:    []domain.Employee{{EmployeeID: "E001", Name: "John Doe"}},
			expectErr: false,
		},
		{
			name: "Delete Success",
			mock: func() {
				repo.EXPECT().Delete(ctx, domain.Employee{EmployeeID: "E001"}).Return(nil)
			},
			method: func() (interface{}, error) {
				return nil, repo.Delete(ctx, domain.Employee{EmployeeID: "E001"})
			},
			expect:    nil,
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			result, err := tt.method()

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, result)
			}
		})
	}
}
