package main

import (
	core "github.com/procyon-projects/procyon-core"
	"github.com/procyon-projects/procyon-test-app/controller"
	"github.com/procyon-projects/procyon-test-app/repository"
	"github.com/procyon-projects/procyon-test-app/service"
)

func init() {
	// controller
	core.Register(controller.NewProductController)
	// service
	core.Register(service.NewProductService)
	// repository
	core.Register(repository.NewProductRepository)
	// error handler
	core.Register(NewErrorHandler)
	// interceptor
	core.Register(NewCustomInterceptor)
}
