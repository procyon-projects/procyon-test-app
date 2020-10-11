package service

import (
	context "github.com/procyon-projects/procyon-context"
	errors "github.com/procyon-projects/procyon-test-app/err"
	"github.com/procyon-projects/procyon-test-app/model"
	"github.com/procyon-projects/procyon-test-app/repository"
	tx "github.com/procyon-projects/procyon-tx"
)

type ProductService interface {
	FindAll(ctx context.Context) ([]*model.Product, error)
	FindById(ctx context.Context, id int) (*model.Product, error)
	Save(ctx context.Context, product *model.Product) (*model.Product, error)
	Update(ctx context.Context, id int, updatedProduct *model.Product) (*model.Product, error)
	DeleteById(ctx context.Context, id int) error
}

type ImpProductService struct {
	productRepository    repository.ProductRepository
	transactionalContext tx.TransactionalContext
}

func NewProductService(productRepository repository.ProductRepository,
	transactionalContext tx.TransactionalContext) *ImpProductService {
	return &ImpProductService{
		productRepository,
		transactionalContext,
	}
}

func (service ImpProductService) GetServiceMetadata() context.ServiceMetadata {
	return context.ServiceMetadata{}
}

func (service ImpProductService) FindAll(ctx context.Context) ([]*model.Product, error) {
	products, err := service.transactionalContext.Block(ctx, func() (interface{}, error) {
		return service.productRepository.FindAll(ctx), nil
	})
	return products.([]*model.Product), err
}

func (service ImpProductService) FindById(ctx context.Context, id int) (*model.Product, error) {
	result, err := service.transactionalContext.Block(ctx, func() (interface{}, error) {
		product := service.productRepository.FindById(ctx, id)
		if product == nil {
			return nil, errors.NewProductNotFoundError(id)
		}
		return product, nil
	})
	return result.(*model.Product), err
}

func (service ImpProductService) Save(ctx context.Context, product *model.Product) (*model.Product, error) {
	result, err := service.transactionalContext.Block(ctx, func() (interface{}, error) {
		return service.productRepository.Save(ctx, product), nil
	})
	return result.(*model.Product), err
}

func (service ImpProductService) Update(ctx context.Context, id int, updatedProduct *model.Product) (*model.Product, error) {
	result, err := service.transactionalContext.Block(ctx, func() (interface{}, error) {
		product := service.productRepository.FindById(ctx, id)
		if product == nil {
			return nil, errors.NewProductNotFoundError(id)
		}
		return service.productRepository.Update(ctx, updatedProduct), nil
	})
	return result.(*model.Product), err
}

func (service ImpProductService) DeleteById(ctx context.Context, id int) error {
	_, err := service.transactionalContext.Block(ctx, func() (interface{}, error) {
		product := service.productRepository.FindById(ctx, id)
		if product == nil {
			return nil, errors.NewProductNotFoundError(id)
		}
		service.productRepository.DeleteById(ctx, id)
		return nil, nil
	})
	return err
}
