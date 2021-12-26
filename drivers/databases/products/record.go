package products

import (
	_productDomain "final-project/business/products"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id         uint `gorm:"primarykey;autoIncrement"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
	Name       string `gorm:"size:100"`
	Price      uint   `gorm:"size:7;not null"`
	CategoryId uint   `gorm:"size:3;not null"`
}

func (product *Product) ToDomain() _productDomain.Domain {
	return _productDomain.Domain{
		Id:         product.Id,
		CreatedAt:  product.CreatedAt,
		UpdatedAt:  product.UpdatedAt,
		Name:       product.Name,
		Price:      product.Price,
		CategoryId: product.CategoryId,
	}
}

func ToListDomain(products []Product) []_productDomain.Domain {
	var result = []_productDomain.Domain{}
	for _, user := range products {
		result = append(result, user.ToDomain())
	}
	return result
}
