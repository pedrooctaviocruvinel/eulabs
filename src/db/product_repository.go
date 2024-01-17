package db

import (
	"github.com/pedrooctaviocruvinel/eulabs/src/types"
	"gorm.io/gorm"
)

type ProductRepository struct {
	instance *gorm.DB
}

func newProductRepository(i *gorm.DB) (productRepository *ProductRepository) {
	return &ProductRepository{
		instance: i,
	}
}

func (pr ProductRepository) List() (products []types.Product) {
	pr.instance.Find(&products)

	return products
}
