package controller

import (
	"fmt"
	"github.com/procyon-projects/procyon-test-app/controller/request"
	"github.com/procyon-projects/procyon-test-app/resource"
	"github.com/procyon-projects/procyon-test-app/service"
	"github.com/procyon-projects/procyon/data/page"
	"github.com/procyon-projects/procyon/void"
	"github.com/procyon-projects/procyon/web"
)

type EmployeeController interface {
	CreateEmployee(ctx *web.RestContext[request.CreateEmployeeRequest, *resource.EmployeeResource]) error
	GetAllEmployees(ctx *web.RestContext[void.Void, page.Page[*resource.EmployeeResource]]) error
	GetEmployee(ctx *web.RestContext[request.GetEmployeeRequest, *resource.EmployeeResource]) error
	UpdateEmployee(ctx *web.RestContext[request.UpdateEmployeeRequest, *resource.EmployeeResource]) error
	DeleteEmployee(ctx *web.RestContext[request.DeleteEmployeeRequest, void.Void]) error
}

type employeeController struct {
	employeeService service.EmployeeService
}

func NewEmployeeController(employeeService service.EmployeeService) EmployeeController {
	return &employeeController{
		employeeService: employeeService,
	}
}

func (c *employeeController) Routes() *web.RouterGroup {
	routerGroup := web.Routes(web.Path("/api/v1/employees"))

	routerGroup.Handler(web.MethodPost, web.RestHandler(c.CreateEmployee))
	routerGroup.Handler(web.MethodGet, web.RestHandler(c.GetAllEmployees))
	routerGroup.GET("/{id}", web.RestHandler(c.GetEmployee))
	routerGroup.PATCH("/{id}", web.RestHandler(c.UpdateEmployee))
	routerGroup.DELETE("/{id}", web.RestHandler(c.DeleteEmployee))
	return routerGroup
}

func (c *employeeController) CreateEmployee(ctx *web.RestContext[request.CreateEmployeeRequest, *resource.EmployeeResource]) error {
	req := ctx.Get()

	employee, err := c.employeeService.CreateEmployee(ctx, &req.Body)

	if err != nil {
		return err
	}

	ctx.Created(fmt.Sprintf("%s/%d", ctx.Path(), employee.Id)).Body(employee)
	return nil
}

func (c *employeeController) GetAllEmployees(ctx *web.RestContext[void.Void, page.Page[*resource.EmployeeResource]]) error {
	employees, err := c.employeeService.GetAllEmployees(ctx)
	if err != nil {
		return err
	}

	ctx.Ok().Body(employees)
	return nil
}

func (c *employeeController) GetEmployee(ctx *web.RestContext[request.GetEmployeeRequest, *resource.EmployeeResource]) error {
	req := ctx.Get()
	employeeId := req.EmployeeId

	employee, err := c.employeeService.GetEmployee(ctx, employeeId)
	if err != nil {
		return err
	}

	ctx.Ok().Body(employee)
	return nil
}

func (c *employeeController) UpdateEmployee(ctx *web.RestContext[request.UpdateEmployeeRequest, *resource.EmployeeResource]) error {
	req := ctx.Get()
	employeeId := req.EmployeeId

	employee, err := c.employeeService.UpdateEmployee(ctx, employeeId, &req.Body)
	if err != nil {
		return err
	}

	ctx.Ok().Body(employee)
	return nil
}

func (c *employeeController) DeleteEmployee(ctx *web.RestContext[request.DeleteEmployeeRequest, void.Void]) error {
	req := ctx.Get()
	employeeId := req.EmployeeId

	err := c.employeeService.DeleteEmployee(ctx, employeeId)
	if err != nil {
		return err
	}

	ctx.NoContent()
	return nil
}
