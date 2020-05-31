package config

import (
	core "github.com/Rollcomp/procyon-core"
	"github.com/Rollcomp/procyon-test-app/controller"
	"github.com/Rollcomp/procyon-test-app/repository"
	"github.com/Rollcomp/procyon-test-app/service"
)

func init() {
	/* Repositories */
	core.Register(repository.NewProductRepository)
	/* Services */
	core.Register(service.NewProductService)
	/* Controllers */
	core.Register(controller.NewProductController)
}
