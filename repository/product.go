package repository

import (
	context "github.com/procyon-projects/procyon-context"
)

type ProductRepository struct {
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (repository *ProductRepository) GetRepositoryMetadata() context.RepositoryMetadata {
	return context.RepositoryMetadata{}
}

func (repository *ProductRepository) FindAll() {

}

func (repository *ProductRepository) FindById() {

}

func (repository *ProductRepository) Save() {

}

func (repository *ProductRepository) Update() {

}

func (repository *ProductRepository) DeleteById() {

}
