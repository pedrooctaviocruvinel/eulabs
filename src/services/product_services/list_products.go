package services

import (
	"github.com/pedrooctaviocruvinel/eulabs/src/types"
	"github.com/pedrooctaviocruvinel/eulabs/src/util"
)

func (ps *ProductServices) List() (listProductsResult util.ResultWrapper[[]types.Product]) {
	products := ps.productRepository.List()

	return util.NewResultWrapperSucceded(products)
}
