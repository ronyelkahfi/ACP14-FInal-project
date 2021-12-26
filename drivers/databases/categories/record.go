package categories

import (
	_categoryDomain "final-project/business/categories"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	Id        uint `gorm:"primarykey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Name      string `gorm:"size:100"`
}

func (category *Category) ToDomain() _categoryDomain.Domain {
	return _categoryDomain.Domain{
		Id:        category.Id,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
		Name:      category.Name,
	}
}

func ToListDomain(categories []Category) []_categoryDomain.Domain {
	var result = []_categoryDomain.Domain{}
	for _, categori := range categories {
		result = append(result, categori.ToDomain())
	}
	return result
}
