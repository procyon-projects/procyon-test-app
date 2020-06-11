package repository

import (
	context "github.com/procyon-projects/procyon-context"
	"github.com/procyon-projects/procyon-test-app/model"
)

type ProductRepository struct {
}

func NewProductRepository() ProductRepository {
	return ProductRepository{}
}

func (repository ProductRepository) GetRepositoryMetadata() context.RepositoryMetadata {
	return context.RepositoryMetadata{}
}

func (repository ProductRepository) FindAll() []*model.Product {
	return nil
}

func (repository ProductRepository) FindById(id int) *model.Product {
	return nil
}

func (repository ProductRepository) Save(product *model.Product) *model.Product {
	return nil
}

func (repository ProductRepository) Update(product *model.Product) *model.Product {
	return nil
}

func (repository ProductRepository) DeleteById(id int) {

}
