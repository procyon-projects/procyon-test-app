package request

import (
	"github.com/procyon-projects/procyon-test-app/dto"
)

type CreateEmployeeRequest struct {
	Body dto.CreateEmployeeDto `type:"body"`
}

type GetEmployeeRequest struct {
	EmployeeId int `type:"path" name:"id"`
}

type UpdateEmployeeRequest struct {
	Body       dto.UpdateEmployeeDto `type:"body"`
	EmployeeId int                   `type:"path" name:"id"`
}

type DeleteEmployeeRequest struct {
	EmployeeId int `type:"path" name:"id"`
}
