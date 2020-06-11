package service

import (
	context "github.com/procyon-projects/procyon-context"
	"github.com/procyon-projects/procyon-test-app/err"
	"github.com/procyon-projects/procyon-test-app/model"
	"github.com/procyon-projects/procyon-test-app/repository"
	tx "github.com/procyon-projects/procyon-tx"
)

type ProductService struct {
	productRepository *repository.ProductRepository
	txContext         tx.TransactionalContext
}

func NewProductService(productRepository *repository.ProductRepository, transactionalContext tx.TransactionalContext) ProductService {
	return ProductService{
		productRepository,
		transactionalContext,
	}
}

func (service ProductService) GetServiceMetadata() context.ServiceMetadata {
	return context.ServiceMetadata{}
}

func (service ProductService) FindAll() ([]*model.Product, error) {
	var products []*model.Product
	service.txContext.Block(func() {
		products = service.productRepository.FindAll()
	})
	return products, nil
}

func (service ProductService) FindById(id int) (*model.Product, error) {
	var serviceErr error
	var product *model.Product
	service.txContext.Block(func() {
		product = service.productRepository.FindById(id)
		if product == nil {
			serviceErr = err.NewProductNotFoundError(id)
			return
		}
	})
	return product, serviceErr
}

func (service ProductService) Save(product *model.Product) (*model.Product, error) {
	var savedProduct *model.Product
	service.txContext.Block(func() {
		savedProduct = service.productRepository.Save(product)
	})
	return savedProduct, nil
}

func (service ProductService) Update(id int, updatedProduct *model.Product) (*model.Product, error) {
	var serviceErr error
	var result *model.Product
	service.txContext.Block(func() {
		product := service.productRepository.FindById(id)
		if product == nil {
			serviceErr = err.NewProductNotFoundError(id)
			return
		}
		result = service.productRepository.Update(updatedProduct)
	})
	return result, serviceErr
}

func (service ProductService) DeleteById(id int) error {
	var serviceErr error
	service.txContext.Block(func() {
		product := service.productRepository.FindById(id)
		if product == nil {
			serviceErr = err.NewProductNotFoundError(id)
			return
		}
		service.productRepository.DeleteById(id)
	})
	return serviceErr
}
