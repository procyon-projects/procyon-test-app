package mapper

import (
	"github.com/procyon-projects/procyon-test-app/model"
	"github.com/procyon-projects/procyon-test-app/model/request"
	"github.com/procyon-projects/procyon-test-app/model/response"
)

func ProductCreateRequestToProductModel(product *request.ProductCreateRequest) *model.Product {
	return &model.Product{
		Name:     product.Body.Name,
		Category: product.Body.Category,
	}
}

func ProductUpdateRequestToProductModel(product *request.ProductUpdateRequest) *model.Product {
	return &model.Product{
		Name:     product.Body.Name,
		Category: product.Body.Category,
	}
}

func ProductToProductDto(product *model.Product) *response.ProductDto {
	return &response.ProductDto{
		Id:       0,
		Name:     product.Name,
		Category: product.Category,
	}
}

func ProductToProductDtoList(products []*model.Product) []*response.ProductDto {
	list := make([]*response.ProductDto, len(products))
	for index, product := range products {
		list[index] = ProductToProductDto(product)
	}
	return list
}
