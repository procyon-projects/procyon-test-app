package controller

import web "github.com/Rollcomp/procyon-web"

type ProductController struct {
}

func NewProductController() *ProductController {
	return &ProductController{}
}

func (controller *ProductController) RegisterHandlers(registry web.HandlerInfoRegistry) {
	registry.RegisterGroup("/api/v1/products",
		web.NewHandlerInfo(controller.GetAllProducts, web.WithPath("/")),
		web.NewHandlerInfo(controller.GetProductById, web.WithPath("/{id}")),
		web.NewHandlerInfo(controller.CreateProduct,
			web.WithPath("/"), web.WithMethod(web.HttpMethodPost),
		),
		web.NewHandlerInfo(controller.UpdateProduct,
			web.WithPath("/{id}"), web.WithMethod(web.HttpMethodPost),
		),
		web.NewHandlerInfo(controller.DeleteProduct,
			web.WithMethod("/{id}"), web.WithMethod(web.HttpMethodDelete),
		),
	)
}

func (controller *ProductController) GetAllProducts() interface{} {
	return nil
}

func (controller *ProductController) GetProductById() interface{} {
	return nil
}

func (controller *ProductController) CreateProduct() interface{} {
	return nil
}

func (controller *ProductController) UpdateProduct() interface{} {
	return nil
}

func (controller *ProductController) DeleteProduct() interface{} {
	return nil
}
