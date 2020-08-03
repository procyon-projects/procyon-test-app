package service

import (
	context "github.com/procyon-projects/procyon-context"
	"github.com/procyon-projects/procyon-test-app/err"
	"github.com/procyon-projects/procyon-test-app/model"
	"github.com/procyon-projects/procyon-test-app/repository"
	tx "github.com/procyon-projects/procyon-tx"
)

type ProductService struct {
	ProductRepository *repository.ProductRepository `inject:""`
	TxContext         tx.TransactionalContext       `inject:""`
}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (service ProductService) GetServiceMetadata() context.ServiceMetadata {
	return context.ServiceMetadata{}
}

func (service ProductService) FindAll() ([]*model.Product, error) {
	var products []*model.Product
	service.TxContext.Block(func() {
		products = service.ProductRepository.FindAll()
	})
	return products, nil
}

func (service ProductService) FindById(id int) (*model.Product, error) {
	var serviceErr error
	var product *model.Product
	service.TxContext.Block(func() {
		product = service.ProductRepository.FindById(id)
		if product == nil {
			serviceErr = err.NewProductNotFoundError(id)
			return
		}
	})
	return product, serviceErr
}

func (service ProductService) Save(product *model.Product) (*model.Product, error) {
	var savedProduct *model.Product
	service.TxContext.Block(func() {
		savedProduct = service.ProductRepository.Save(product)
	})
	return savedProduct, nil
}

func (service ProductService) Update(id int, updatedProduct *model.Product) (*model.Product, error) {
	var serviceErr error
	var result *model.Product
	service.TxContext.Block(func() {
		product := service.ProductRepository.FindById(id)
		if product == nil {
			serviceErr = err.NewProductNotFoundError(id)
			return
		}
		result = service.ProductRepository.Update(updatedProduct)
	})
	return result, serviceErr
}

func (service ProductService) DeleteById(id int) error {
	var serviceErr error
	service.TxContext.Block(func() {
		product := service.ProductRepository.FindById(id)
		if product == nil {
			serviceErr = err.NewProductNotFoundError(id)
			return
		}
		service.ProductRepository.DeleteById(id)
	})
	return serviceErr
}
