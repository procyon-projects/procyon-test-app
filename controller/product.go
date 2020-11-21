package controller

import (
	context "github.com/procyon-projects/procyon-context"
	"github.com/procyon-projects/procyon-test-app/model/mapper"
	"github.com/procyon-projects/procyon-test-app/service"
	web "github.com/procyon-projects/procyon-web"
)

type ProductController interface {
	GetAllProducts(context *web.WebRequestContext)
	GetProductById(context *web.WebRequestContext)
	CreateProduct(context *web.WebRequestContext)
	UpdateProduct(context *web.WebRequestContext)
	DeleteProduct(context *web.WebRequestContext)
}

type ImpProductController struct {
	logger         context.Logger
	productService service.ProductService
}

func NewProductController(logger context.Logger, productService service.ProductService) ImpProductController {
	return ImpProductController{
		logger,
		productService,
	}
}

func (controller ImpProductController) RegisterHandlers(registry web.HandlerRegistry) {
	registry.RegisterGroup("/api/products",
		web.NewHandler(controller.GetAllProducts),
		web.NewHandler(
			controller.GetProductById, web.WithPath("/{productId}"),
		),
		web.NewHandler(
			controller.CreateProduct, web.WithMethod(web.RequestMethodPost),
		),
		web.NewHandler(
			controller.UpdateProduct, web.WithPath("/{productId}"), web.WithMethod(web.RequestMethodPut),
		),
		web.NewHandler(
			controller.DeleteProduct, web.WithPath("/{productId}"), web.WithMethod(web.RequestMethodDelete),
		),
	)
}

func (controller ImpProductController) GetAllProducts(context *web.WebRequestContext) {
	products := controller.productService.FindAll(context)
	context.SetBody(mapper.ProductToProductDtoList(products))
}

func (controller ImpProductController) GetProductById(context *web.WebRequestContext) {
	product := controller.productService.FindById(context, 0)
	context.SetBody(mapper.ProductToProductDto(product))
}

func (controller ImpProductController) CreateProduct(context *web.WebRequestContext) {
	product := controller.productService.Save(context, mapper.ProductUpdateRequestToProductModel(nil))
	context.SetBody(mapper.ProductToProductDto(product))
}

func (controller ImpProductController) UpdateProduct(context *web.WebRequestContext) {
	product := controller.productService.Update(context, 0, mapper.ProductUpdateRequestToProductModel(nil))
	context.SetBody(mapper.ProductToProductDto(product))
}

func (controller ImpProductController) DeleteProduct(context *web.WebRequestContext) {
	controller.productService.DeleteById(context, 0)
}
