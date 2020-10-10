package service

import (
	context "github.com/procyon-projects/procyon-context"
	"github.com/procyon-projects/procyon-test-app/err"
	"github.com/procyon-projects/procyon-test-app/model"
	"github.com/procyon-projects/procyon-test-app/repository"
	tx "github.com/procyon-projects/procyon-tx"
)

type ProductService struct {
	productRepository    repository.ProductRepository
	transactionalContext tx.TransactionalContext
}

func NewProductService(productRepository repository.ProductRepository,
	transactionalContext tx.TransactionalContext) *ProductService {
	return &ProductService{
		productRepository,
		transactionalContext,
	}
}

func (service ProductService) GetServiceMetadata() context.ServiceMetadata {
	return context.ServiceMetadata{}
}

func (service ProductService) FindAll(ctx context.Context) ([]*model.Product, error) {
	var products []*model.Product
	service.transactionalContext.Block(ctx, func() {
		products = service.productRepository.FindAll(ctx)
	})
	return products, nil
}

func (service ProductService) FindById(ctx context.Context, id int) (*model.Product, error) {
	var serviceErr error
	var product *model.Product
	service.transactionalContext.Block(ctx, func() {
		product = service.productRepository.FindById(ctx, id)
		if product == nil {
			serviceErr = err.NewProductNotFoundError(id)
			return
		}
	})
	return product, serviceErr
}

func (service ProductService) Save(ctx context.Context, product *model.Product) (*model.Product, error) {
	var savedProduct *model.Product
	service.transactionalContext.Block(ctx, func() {
		savedProduct = service.productRepository.Save(ctx, product)
	})
	return savedProduct, nil
}

func (service ProductService) Update(ctx context.Context, id int, updatedProduct *model.Product) (*model.Product, error) {
	var serviceErr error
	var result *model.Product
	service.transactionalContext.Block(ctx, func() {
		product := service.productRepository.FindById(ctx, id)
		if product == nil {
			serviceErr = err.NewProductNotFoundError(id)
			return
		}
		result = service.productRepository.Update(ctx, updatedProduct)
	})
	return result, serviceErr
}

func (service ProductService) DeleteById(ctx context.Context, id int) error {
	var serviceErr error
	service.transactionalContext.Block(ctx, func() {
		product := service.productRepository.FindById(ctx, id)
		if product == nil {
			serviceErr = err.NewProductNotFoundError(id)
			return
		}
		service.productRepository.DeleteById(ctx, id)
	})
	return serviceErr
}
