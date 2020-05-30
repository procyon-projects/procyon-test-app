package main

import (
	core "github.com/Rollcomp/procyon-core"
	"github.com/Rollcomp/procyon-test-app/controller"
)

func init() {
	/* Controllers */
	core.Register(controller.NewProductController)
}
