package product_services

import "github.com/pedrooctaviocruvinel/eulabs/src/types"

type ProductServices struct {
	productRepository types.IProductRepository
}

func NewProductServices(pr types.IProductRepository) (productServices *ProductServices) {
	return &ProductServices{
		productRepository: pr,
	}
}
