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

func (repository ProductRepository) FindAll(ctx context.Context) []*model.Product {
	return nil
}

func (repository ProductRepository) FindById(ctx context.Context, id int) *model.Product {
	return nil
}

func (repository ProductRepository) Save(ctx context.Context, product *model.Product) *model.Product {
	return nil
}

func (repository ProductRepository) Update(ctx context.Context, product *model.Product) *model.Product {
	return nil
}

func (repository ProductRepository) DeleteById(ctx context.Context, id int) {

}
