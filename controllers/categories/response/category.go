package response

import (
	_categoriesDomain "final-project/business/categories"
	"time"
)

type CategoryResponse struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
}

func FromDomain(category _categoriesDomain.Domain) CategoryResponse {
	return CategoryResponse{
		Id:        category.Id,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
		Name:      category.Name,
	}
}

func ToListFromDomain(categories []_categoriesDomain.Domain) []CategoryResponse {
	var result = []CategoryResponse{}
	for _, category := range categories {
		result = append(result, FromDomain(category))
	}
	return result
}
