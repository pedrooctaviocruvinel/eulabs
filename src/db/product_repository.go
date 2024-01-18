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

func (pr ProductRepository) GetByID(i string) (product types.Product) {
	pr.instance.Where(types.Product{ID: i}).First(&product)

	return product
}

func (pr ProductRepository) Create(p types.Product) {
	pr.instance.Create(p)
}

func (pr ProductRepository) Update(p types.Product) {
	pr.instance.Save(p)
}

func (pr ProductRepository) Delete(p types.Product) {
	pr.instance.Delete(p)
}
