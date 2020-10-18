package controller

import (
	context "github.com/procyon-projects/procyon-context"
	"github.com/procyon-projects/procyon-test-app/model/mapper"
	"github.com/procyon-projects/procyon-test-app/model/request"
	"github.com/procyon-projects/procyon-test-app/service"
	web "github.com/procyon-projects/procyon-web"
)

type ProductController interface {
	GetAllProducts(ctx context.Context) (*web.ResponseEntity, error)
	GetProductById(ctx context.Context, request *request.ProductGetRequest) (*web.ResponseEntity, error)
	CreateProduct(ctx context.Context, request *request.ProductCreateRequest) (*web.ResponseEntity, error)
	UpdateProduct(ctx context.Context, request *request.ProductUpdateRequest) (*web.ResponseEntity, error)
	DeleteProduct(ctx context.Context, request *request.ProductDeleteRequest) (*web.ResponseEntity, error)
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
			web.WithPath("/{id}"), web.WithMethod(web.RequestMethodDelete),
		),
	)
}

func (controller ImpProductController) GetAllProducts(ctx context.Context) (*web.ResponseEntity, error) {
	products, err := controller.productService.FindAll(ctx)
	return web.NewResponseEntity(
		web.WithBody(mapper.ProductToProductDtoList(products)),
	), err
}

func (controller ImpProductController) GetProductById(ctx context.Context,
	request *request.ProductGetRequest) (*web.ResponseEntity, error) {
	product, err := controller.productService.FindById(ctx, request.PathVariables.ProductId)
	if err != nil {
		return nil, err
	}
	return web.NewResponseEntity(
		web.WithBody(mapper.ProductToProductDto(product)),
	), nil
}

func (controller ImpProductController) CreateProduct(ctx context.Context,
	request *request.ProductCreateRequest) (*web.ResponseEntity, error) {
	product, err := controller.productService.Save(ctx, mapper.ProductCreateRequestToProductModel(request))
	if err != nil {
		return nil, err
	}
	return web.NewResponseEntity(
		web.WithBody(mapper.ProductToProductDto(product)),
	), nil
}

func (controller ImpProductController) UpdateProduct(ctx context.Context,
	request *request.ProductUpdateRequest) (*web.ResponseEntity, error) {
	product, err := controller.productService.Update(ctx, request.PathVariables.ProductId, mapper.ProductUpdateRequestToProductModel(request))
	if err != nil {
		return nil, err
	}
	return web.NewResponseEntity(
		web.WithBody(mapper.ProductToProductDto(product)),
	), nil
}

func (controller ImpProductController) DeleteProduct(ctx context.Context,
	request *request.ProductDeleteRequest) (*web.ResponseEntity, error) {
	err := controller.productService.DeleteById(ctx, request.PathVariables.ProductId)
	if err != nil {
		return nil, err
	}
	return web.NewResponseEntity(), nil
}
