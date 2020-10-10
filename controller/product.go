package controller

import (
	context "github.com/procyon-projects/procyon-context"
	"github.com/procyon-projects/procyon-test-app/model/mapper"
	"github.com/procyon-projects/procyon-test-app/model/request"
	"github.com/procyon-projects/procyon-test-app/service"
	web "github.com/procyon-projects/procyon-web"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return ProductController{
		productService,
	}
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

func (controller ProductController) GetAllProducts(ctx context.Context) (*web.ResponseEntity, error) {
	products, err := controller.productService.FindAll(ctx)
	return web.NewResponseEntity(
		web.WithBody(mapper.ProductToProductDtoList(products)),
	), err
}

func (controller ProductController) GetProductById(ctx context.Context,
	request *request.ProductGetRequest) (*web.ResponseEntity, error) {
	product, err := controller.productService.FindById(ctx, request.PathVariables.ProductId)
	return web.NewResponseEntity(
		web.WithBody(mapper.ProductToProductDto(product)),
	), err
}

func (controller ProductController) CreateProduct(ctx context.Context,
	request *request.ProductCreateRequest) (*web.ResponseEntity, error) {
	product, err := controller.productService.Save(ctx, mapper.ProductCreateRequestToProductModel(request))
	return web.NewResponseEntity(
		web.WithBody(mapper.ProductToProductDto(product)),
	), err
}

func (controller ProductController) UpdateProduct(ctx context.Context,
	request *request.ProductUpdateRequest) (*web.ResponseEntity, error) {
	product, err := controller.productService.Update(ctx, request.PathVariables.ProductId, mapper.ProductUpdateRequestToProductModel(request))
	return web.NewResponseEntity(
		web.WithBody(mapper.ProductToProductDto(product)),
	), err
}

func (controller ProductController) DeleteProduct(ctx context.Context,
	request *request.ProductDeleteRequest) (*web.ResponseEntity, error) {
	err := controller.productService.DeleteById(ctx, request.PathVariables.ProductId)
	return web.NewResponseEntity(), err
}
