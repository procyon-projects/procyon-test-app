package repository

import (
	context "github.com/procyon-projects/procyon-context"
	"github.com/procyon-projects/procyon-test-app/model"
)

type ProductRepository interface {
	FindAll(ctx context.Context) []*model.Product
	FindById(ctx context.Context, id int) *model.Product
	Save(ctx context.Context, product *model.Product) *model.Product
	Update(ctx context.Context, product *model.Product) *model.Product
	DeleteById(ctx context.Context, id int)
}

type ImpProductRepository struct {
}

func NewProductRepository() ImpProductRepository {
	return ImpProductRepository{}
}

func (repository ImpProductRepository) GetRepositoryMetadata() context.RepositoryMetadata {
	return context.RepositoryMetadata{}
}

func (repository ImpProductRepository) FindAll(ctx context.Context) []*model.Product {
	return nil
}

func (repository ImpProductRepository) FindById(ctx context.Context, id int) *model.Product {
	return nil
}

func (repository ImpProductRepository) Save(ctx context.Context, product *model.Product) *model.Product {
	return nil
}

func (repository ImpProductRepository) Update(ctx context.Context, product *model.Product) *model.Product {
	return nil
}

func (repository ImpProductRepository) DeleteById(ctx context.Context, id int) {

}
