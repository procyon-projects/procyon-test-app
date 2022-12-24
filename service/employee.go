package service

import (
	"context"
	"github.com/procyon-projects/procyon-test-app/dto"
	"github.com/procyon-projects/procyon-test-app/resource"
	"github.com/procyon-projects/procyon/data/page"
)

type EmployeeService interface {
	CreateEmployee(ctx context.Context, createEmployeeDto *dto.CreateEmployeeDto) (*resource.EmployeeResource, error)
	GetAllEmployees(ctx context.Context) (page.Page[*resource.EmployeeResource], error)
	GetEmployee(ctx context.Context, employeeId int) (*resource.EmployeeResource, error)
	UpdateEmployee(ctx context.Context, employeeId int, updateEmployeeDto *dto.UpdateEmployeeDto) (*resource.EmployeeResource, error)
	DeleteEmployee(ctx context.Context, employeeId int) error
}

type employeeService struct {
}

func NewEmployeeService() EmployeeService {
	return &employeeService{}
}

func (c *employeeService) CreateEmployee(ctx context.Context,
	createEmployeeData *dto.CreateEmployeeDto) (*resource.EmployeeResource, error) {
	return &resource.EmployeeResource{
		Id:   1,
		Name: "anyName",
		Role: "anyRole",
	}, nil
}

func (c *employeeService) GetAllEmployees(ctx context.Context) (page.Page[*resource.EmployeeResource], error) {
	return nil, nil
}

func (c *employeeService) GetEmployee(ctx context.Context, employeeId int) (*resource.EmployeeResource, error) {
	return &resource.EmployeeResource{
		Id:   1,
		Name: "anyName",
		Role: "anyRole",
	}, nil
}

func (c *employeeService) UpdateEmployee(ctx context.Context,
	employeeId int,
	updateEmployeeDto *dto.UpdateEmployeeDto) (*resource.EmployeeResource, error) {

	return &resource.EmployeeResource{
		Id:   1,
		Name: "anyName",
		Role: "anyRole",
	}, nil
}

func (c *employeeService) DeleteEmployee(ctx context.Context, employeeId int) error {
	return nil
}
