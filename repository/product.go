package repository

import (
	context "github.com/procyon-projects/procyon-context"
	"github.com/procyon-projects/procyon-test-app/model"
	web "github.com/procyon-projects/procyon-web"
)

type ProductRepository interface {
	FindAll(ctx *web.WebRequestContext) []*model.Product
	FindById(ctx *web.WebRequestContext, id int) *model.Product
	Save(ctx *web.WebRequestContext, product *model.Product) *model.Product
	Update(ctx *web.WebRequestContext, product *model.Product) *model.Product
	DeleteById(ctx *web.WebRequestContext, id int)
}

type ImpProductRepository struct {
}

func NewProductRepository() ImpProductRepository {
	return ImpProductRepository{}
}

func (repository ImpProductRepository) GetRepositoryMetadata() context.RepositoryMetadata {
	return context.RepositoryMetadata{}
}

func (repository ImpProductRepository) FindAll(ctx *web.WebRequestContext) []*model.Product {
	return nil
}

func (repository ImpProductRepository) FindById(ctx *web.WebRequestContext, id int) *model.Product {
	return nil
}

func (repository ImpProductRepository) Save(ctx *web.WebRequestContext, product *model.Product) *model.Product {
	return nil
}

func (repository ImpProductRepository) Update(ctx *web.WebRequestContext, product *model.Product) *model.Product {
	return nil
}

func (repository ImpProductRepository) DeleteById(ctx *web.WebRequestContext, id int) {

}
