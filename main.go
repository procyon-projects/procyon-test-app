package main

import (
	"github.com/procyon-projects/procyon"
	"github.com/procyon-projects/procyon-test-app/controller"
	"github.com/procyon-projects/procyon-test-app/service"
	"github.com/procyon-projects/procyon/app/component"
)

func init() {
	component.Register(controller.NewEmployeeController)
	component.Register(service.NewEmployeeService)
}

func main() {
	procyon.New().Run()
}
