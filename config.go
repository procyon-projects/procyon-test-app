package main

import (
	core "github.com/procyon-projects/procyon-core"
	"github.com/procyon-projects/procyon-test-app/controller"
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
}
