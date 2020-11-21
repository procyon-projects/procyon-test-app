package service

import (
	configure "github.com/procyon-projects/procyon-configure"
	context "github.com/procyon-projects/procyon-context"
	errors "github.com/procyon-projects/procyon-test-app/err"
	"github.com/procyon-projects/procyon-test-app/model"
	"github.com/procyon-projects/procyon-test-app/repository"
	tx "github.com/procyon-projects/procyon-tx"
	web "github.com/procyon-projects/procyon-web"
)

type ProductService interface {
	FindAll(ctx *web.WebRequestContext) []*model.Product
	FindById(ctx *web.WebRequestContext, id int) *model.Product
	Save(ctx *web.WebRequestContext, product *model.Product) *model.Product
	Update(ctx *web.WebRequestContext, id int, updatedProduct *model.Product) *model.Product
	DeleteById(ctx *web.WebRequestContext, id int)
}

type ImpProductService struct {
	logger               context.Logger
	productRepository    repository.ProductRepository
	transactionalContext tx.TransactionalContext
}

func NewProductService(logger context.Logger,
	productRepository repository.ProductRepository,
	transactionalContext tx.TransactionalContext, properties configure.WebServerProperties) ImpProductService {
	return ImpProductService{
		logger,
		productRepository,
		transactionalContext,
	}
}

func (service ImpProductService) GetServiceMetadata() context.ServiceMetadata {
	return context.ServiceMetadata{}
}

func (service ImpProductService) FindAll(ctx *web.WebRequestContext) []*model.Product {
	products, _ := service.transactionalContext.Block(ctx, func() (interface{}, error) {
		return service.productRepository.FindAll(ctx), nil
	})
	return products.([]*model.Product)
}

func (service ImpProductService) FindById(ctx *web.WebRequestContext, id int) *model.Product {
	result, _ := service.transactionalContext.Block(ctx, func() (interface{}, error) {
		product := service.productRepository.FindById(ctx, id)
		if product == nil {
			ctx.ThrowError(errors.NewProductNotFoundError(id))
		}
		return product, nil
	})
	return result.(*model.Product)
}

func (service ImpProductService) Save(ctx *web.WebRequestContext, product *model.Product) *model.Product {
	result, _ := service.transactionalContext.Block(ctx, func() (interface{}, error) {
		return service.productRepository.Save(ctx, product), nil
	})
	return result.(*model.Product)
}

func (service ImpProductService) Update(ctx *web.WebRequestContext, id int, updatedProduct *model.Product) *model.Product {
	result, _ := service.transactionalContext.Block(ctx, func() (interface{}, error) {
		product := service.productRepository.FindById(ctx, id)
		if product == nil {
			ctx.ThrowError(errors.NewProductNotFoundError(id))
		}
		return service.productRepository.Update(ctx, updatedProduct), nil
	})
	return result.(*model.Product)
}

func (service ImpProductService) DeleteById(ctx *web.WebRequestContext, id int) {
	service.transactionalContext.Block(ctx, func() (interface{}, error) {
		product := service.productRepository.FindById(ctx, id)
		if product == nil {
			ctx.ThrowError(errors.NewProductNotFoundError(id))
		}
		service.productRepository.DeleteById(ctx, id)
		return nil, nil
	})
}
