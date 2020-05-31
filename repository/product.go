package repository

import (
	context "github.com/Rollcomp/procyon-context"
)

type ProductRepository struct {
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (repository ProductRepository) GetRepositoryMetadata() context.RepositoryMetadata {
	return context.RepositoryMetadata{}
}
