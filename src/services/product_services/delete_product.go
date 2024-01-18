package product_services

import (
	"github.com/pedrooctaviocruvinel/eulabs/src/types"
	"github.com/pedrooctaviocruvinel/eulabs/src/util"
)

func (ps *ProductServices) Delete(i string) (createProductResult util.ResultWrapper[any]) {
	product := ps.productRepository.GetByID(i)

	if (product == types.Product{}) {
		return util.NewResultWrapperFailed[any]([]string{"product doesn't exists"})
	}

	ps.productRepository.Delete(product)

	return util.NewResultWrapperSuccededEmpty[any]()
}
