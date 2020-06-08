package service

import (
	context "github.com/procyon-projects/procyon-context"
	"github.com/procyon-projects/procyon-test-app/repository"
	tx "github.com/procyon-projects/procyon-tx"
)

type ProductService struct {
	productRepository *repository.ProductRepository
	txContext         tx.TransactionalContext
}

func NewProductService(productRepository *repository.ProductRepository, transactionalContext tx.TransactionalContext) *ProductService {
	return &ProductService{
		productRepository,
		transactionalContext,
	}
}

func (service *ProductService) GetServiceMetadata() context.ServiceMetadata {
	return context.ServiceMetadata{}
}

func (service *ProductService) FindAll() {
	service.txContext.Block(func() {
		service.productRepository.FindAll()
	})
}

func (service *ProductService) FindById() {
	service.txContext.Block(func() {
		service.productRepository.FindById()
		/* etc...*/
	})
}

func (service *ProductService) Save() {
	service.txContext.Block(func() {
		service.productRepository.Save()
		/* etc...*/
	})
}

func (service *ProductService) Update() {
	service.txContext.Block(func() {
		service.productRepository.Update()
		/* etc...*/
	})
}

func (service *ProductService) DeleteById() {
	service.txContext.Block(func() {
		service.productRepository.Update()
		/* etc...*/
	})
}
