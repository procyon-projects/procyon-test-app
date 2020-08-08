package controller

import (
	"github.com/procyon-projects/procyon-test-app/model/mapper"
	"github.com/procyon-projects/procyon-test-app/model/request"
	"github.com/procyon-projects/procyon-test-app/service"
	web "github.com/procyon-projects/procyon-web"
)

type ProductController struct {
	ProductService *service.ProductService `inject:""`
}

func NewProductController() *ProductController {
	return &ProductController{}
}

func (controller ProductController) RegisterHandlers(registry web.HandlerRegistry) {
	registry.RegisterGroup("/api/v1/products",
		web.NewHandler(controller.GetAllProducts, web.WithPath("/")),
		web.NewHandler(controller.GetProductById, web.WithPath("/{id}")),
		web.NewHandler(controller.CreateProduct,
			web.WithPath("/"), web.WithMethod(web.RequestMethodPost),
		),
		web.NewHandler(controller.UpdateProduct,
			web.WithPath("/{id}"), web.WithMethod(web.RequestMethodPost),
		),
		web.NewHandler(controller.DeleteProduct,
			web.WithMethod("/{id}"), web.WithMethod(web.RequestMethodDelete),
		),
	)
}

func (controller ProductController) GetAllProducts() (*web.ResponseEntity, error) {
	products, err := controller.ProductService.FindAll()
	return web.NewResponseEntity(
		web.WithBody(mapper.ProductToProductDtoList(products)),
	), err
}

func (controller ProductController) GetProductById(request *request.ProductGetRequest) (*web.ResponseEntity, error) {
	product, err := controller.ProductService.FindById(request.PathVariables.ProductId)
	return web.NewResponseEntity(
		web.WithBody(mapper.ProductToProductDto(product)),
	), err
}

func (controller ProductController) CreateProduct(request *request.ProductCreateRequest) (*web.ResponseEntity, error) {
	product, err := controller.ProductService.Save(mapper.ProductCreateRequestToProductModel(request))
	return web.NewResponseEntity(
		web.WithBody(mapper.ProductToProductDto(product)),
	), err
}

func (controller ProductController) UpdateProduct(request *request.ProductUpdateRequest) (*web.ResponseEntity, error) {
	product, err := controller.ProductService.Update(request.PathVariables.ProductId, mapper.ProductUpdateRequestToProductModel(request))
	return web.NewResponseEntity(
		web.WithBody(mapper.ProductToProductDto(product)),
	), err
}

func (controller ProductController) DeleteProduct(request *request.ProductDeleteRequest) (*web.ResponseEntity, error) {
	err := controller.ProductService.DeleteById(request.PathVariables.ProductId)
	return web.NewResponseEntity(), err
}
