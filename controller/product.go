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
		web.Get(controller.GetAllProducts),
		web.Get(
			controller.GetProductById,
			web.Path("/:productId"),
			web.RequestObject(request.ProductGetRequest{}),
		),
		web.Post(
			controller.CreateProduct,
			web.RequestObject(request.ProductCreateRequest{}),
		),
		web.Put(
			controller.UpdateProduct,
			web.Path("/:productId"),
			web.RequestObject(request.ProductUpdateRequest{}),
		),
		web.Delete(
			controller.DeleteProduct,
			web.Path("/:productId"),
			web.RequestObject(request.ProductDeleteRequest{}),
		),
	)
}

func (controller ImpProductController) GetAllProducts(context *web.WebRequestContext) {
	products := controller.productService.FindAll(context)
	context.Ok().SetModel(mapper.ProductsToProductDtoList(products))
}

func (controller ImpProductController) GetProductById(context *web.WebRequestContext) {
	req := &request.ProductGetRequest{}
	context.BindRequest(req)
	product := controller.productService.FindById(context, req.PathVariables.ProductId)
	context.Ok().SetModel(mapper.ProductToProductDto(product)).SetResponseContentType(web.MediaTypeApplicationJson)
}

func (controller ImpProductController) CreateProduct(context *web.WebRequestContext) {
	req := &request.ProductCreateRequest{}
	context.BindRequest(req)
	product := controller.productService.Save(context, mapper.ProductCreateRequestToProductModel(req))
	context.Ok().SetModel(mapper.ProductToProductDto(product))
}

func (controller ImpProductController) UpdateProduct(context *web.WebRequestContext) {
	req := &request.ProductUpdateRequest{}
	context.BindRequest(req)
	product := controller.productService.Update(context, req.PathVariables.ProductId, mapper.ProductUpdateRequestToProductModel(req))
	context.Ok().SetModel(mapper.ProductToProductDto(product))
}

func (controller ImpProductController) DeleteProduct(context *web.WebRequestContext) {
	req := &request.ProductDeleteRequest{}
	context.BindRequest(req)
	controller.productService.DeleteById(context, req.PathVariables.ProductId)
}
