package product_services

import (
	"github.com/pedrooctaviocruvinel/eulabs/src/types"
	"github.com/pedrooctaviocruvinel/eulabs/src/util"
)

type CreateProductDTORequest struct {
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

type CreateProductDTOResult struct {
	ID string `json:"id"`
}

func newCreateProductDTOResult(i string) (createProductDTOResult CreateProductDTOResult) {
	return CreateProductDTOResult{
		ID: i,
	}
}

func (ps *ProductServices) Create(p *CreateProductDTORequest) (createProductResult util.ResultWrapper[CreateProductDTOResult]) {
	product := types.NewProduct(p.Name, p.Stock)

	validateResult, errors := util.ValidateType(product)
	if !validateResult {
		return util.NewResultWrapperFailed[CreateProductDTOResult](errors)
	}

	ps.productRepository.Create(*product)

	return util.NewResultWrapperSucceded(newCreateProductDTOResult(product.ID))
}
