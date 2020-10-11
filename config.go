package main

import (
	core "github.com/procyon-projects/procyon-core"
	"github.com/procyon-projects/procyon-test-app/controller"
	errors "github.com/procyon-projects/procyon-test-app/err"
	"github.com/procyon-projects/procyon-test-app/repository"
	"github.com/procyon-projects/procyon-test-app/service"
)

func init() {
	/* Repositories */
	core.Register(repository.NewProductRepository)
	/* Services */
	core.Register(service.NewProductService)
	/* Controllers */
	core.Register(controller.NewProductController)
	/* Error Adviser */
	core.Register(errors.NewErrorHandlerAdviser)
}
