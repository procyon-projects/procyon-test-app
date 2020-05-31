package service

import (
	context "github.com/Rollcomp/procyon-context"
	"github.com/Rollcomp/procyon-test-app/repository"
)

type ProductService struct {
	productRepository *repository.ProductRepository
}

func NewProductService(productRepository *repository.ProductRepository) *ProductService {
	return &ProductService{
		productRepository,
	}
}

func (service ProductService) GetServiceMetadata() context.ServiceMetadata {
	return context.ServiceMetadata{}
}
