package controller

import (
	"github.com/procyon-projects/procyon-test-app/model/mapper"
	"github.com/procyon-projects/procyon-test-app/model/request"
	"github.com/procyon-projects/procyon-test-app/service"
	web "github.com/procyon-projects/procyon-web"
)

type ProductController struct {
	productService *service.ProductService
}

func NewProductController(productService *service.ProductService) ProductController {
	return ProductController{
		productService,
	}
}

func (controller ProductController) RegisterHandlers(registry web.HandlerInfoRegistry) {
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

func (controller ProductController) GetAllProducts() (*web.ResponseEntity, error) {
	products, err := controller.productService.FindAll()
	return web.NewResponseEntity(
		web.WithBody(mapper.ProductToProductDtoList(products)),
	), err
}

func (controller ProductController) GetProductById(request *request.ProductGetRequest) (*web.ResponseEntity, error) {
	product, err := controller.productService.FindById(request.PathVariables.ProductId)
	return web.NewResponseEntity(
		web.WithBody(mapper.ProductToProductDto(product)),
	), err
}

func (controller ProductController) CreateProduct(request *request.ProductCreateRequest) (*web.ResponseEntity, error) {
	product, err := controller.productService.Save(mapper.ProductCreateRequestToProductModel(request))
	return web.NewResponseEntity(
		web.WithBody(mapper.ProductToProductDto(product)),
	), err
}

func (controller ProductController) UpdateProduct(request *request.ProductUpdateRequest) (*web.ResponseEntity, error) {
	product, err := controller.productService.Update(request.PathVariables.ProductId, mapper.ProductUpdateRequestToProductModel(request))
	return web.NewResponseEntity(
		web.WithBody(mapper.ProductToProductDto(product)),
	), err
}

func (controller ProductController) DeleteProduct(request *request.ProductDeleteRequest) (*web.ResponseEntity, error) {
	err := controller.productService.DeleteById(request.PathVariables.ProductId)
	return web.NewResponseEntity(), err
}
