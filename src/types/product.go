package types

import "github.com/google/uuid"

type Product struct {
	ID    string `gorm:"type:char(36);primary_key;" valid:"required, uuid"`
	Name  string `gorm:"type:varchar(256);not null;" valid:"required, maxstringlength(256)"`
	Stock int    `gorm:"not null;" valid:"required"`
}

func NewProduct(n string, s int) (product *Product) {
	return &Product{
		ID:    uuid.New().String(),
		Name:  n,
		Stock: s,
	}
}

type IProductRepository interface {
	List() (products []Product)
	GetByID(i string) (product Product)
	Create(p Product)
	Update(p Product)
	Delete(p Product)
}
