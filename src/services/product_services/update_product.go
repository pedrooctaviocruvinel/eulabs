package product_services

import (
	"github.com/pedrooctaviocruvinel/eulabs/src/types"
	"github.com/pedrooctaviocruvinel/eulabs/src/util"
)

type UpdateProductDTORequest struct {
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

func (ps *ProductServices) Update(i string, p *UpdateProductDTORequest) (createProductResult util.ResultWrapper[any]) {
	product := ps.productRepository.GetByID(i)

	if (product == types.Product{}) {
		return util.NewResultWrapperFailed[any]([]string{"product doesn't exists"})
	}

	product.Name = p.Name
	product.Stock = p.Stock

	validateResult, errors := util.ValidateType(product)
	if !validateResult {
		return util.NewResultWrapperFailed[any](errors)
	}

	ps.productRepository.Update(product)

	return util.NewResultWrapperSuccededEmpty[any]()
}
