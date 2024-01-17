package types

import "github.com/google/uuid"

type Product struct {
	ID    string `gorm:"type:char(36);primary_key;"`
	Name  string `gorm:"type:varchar(256);not null;"`
	Stock int    `gorm:"not null;"`
}

func NewProduct(n string, s int) (product *Product) {
	return &Product{
		ID:    uuid.New().String(),
		Name:  n,
		Stock: s,
	}
}
