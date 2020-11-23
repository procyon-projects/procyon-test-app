package controller

import (
	"github.com/procyon-projects/procyon-test-app/model/mapper"
	"github.com/procyon-projects/procyon-test-app/model/request"
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
	productService service.ProductService
}

func NewProductController(productService service.ProductService) ImpProductController {
	return ImpProductController{
		productService,
	}
}

func (controller ImpProductController) RegisterHandlers(registry web.HandlerRegistry) {
	registry.RegisterGroup("/api/products",
		web.NewHandler(controller.GetAllProducts),
		web.NewHandler(
			controller.GetProductById, web.WithPath("/:productId"),
			web.WithRequestObject(request.ProductGetRequest{}),
		),
		web.NewHandler(
			controller.CreateProduct,
			web.WithMethod(web.RequestMethodPost),
			web.WithRequestObject(request.ProductCreateRequest{}),
		),
		web.NewHandler(
			controller.UpdateProduct,
			web.WithPath("/:productId"),
			web.WithMethod(web.RequestMethodPut),
			web.WithRequestObject(request.ProductUpdateRequest{}),
		),
		web.NewHandler(
			controller.DeleteProduct,
			web.WithPath("/:productId"),
			web.WithMethod(web.RequestMethodDelete),
			web.WithRequestObject(request.ProductDeleteRequest{}),
		),
	)
}

func (controller ImpProductController) GetAllProducts(context *web.WebRequestContext) {
	products := controller.productService.FindAll(context)
	context.SetBody(mapper.ProductToProductDtoList(products))
}

func (controller ImpProductController) GetProductById(context *web.WebRequestContext) {
	req := &request.ProductGetRequest{}
	context.GetRequest(req)
	product := controller.productService.FindById(context, req.PathVariables.ProductId)
	context.SetBody(mapper.ProductToProductDto(product))
}

func (controller ImpProductController) CreateProduct(context *web.WebRequestContext) {
	req := &request.ProductCreateRequest{}
	context.GetRequest(req)
	product := controller.productService.Save(context, mapper.ProductCreateRequestToProductModel(req))
	context.SetBody(mapper.ProductToProductDto(product))
}

func (controller ImpProductController) UpdateProduct(context *web.WebRequestContext) {
	req := &request.ProductUpdateRequest{}
	context.GetRequest(req)
	product := controller.productService.Update(context, req.PathVariables.ProductId, mapper.ProductUpdateRequestToProductModel(req))
	context.SetBody(mapper.ProductToProductDto(product))
}

func (controller ImpProductController) DeleteProduct(context *web.WebRequestContext) {
	req := &request.ProductDeleteRequest{}
	context.GetRequest(req)
	controller.productService.DeleteById(context, req.PathVariables.ProductId)
}
