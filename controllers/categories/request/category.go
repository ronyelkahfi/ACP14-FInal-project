package request

import (
	_categoriesDomain "final-project/business/categories"
	"time"
)

type CategoryRequest struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
}

func ToDomain(category CategoryRequest) _categoriesDomain.Domain {
	return _categoriesDomain.Domain{
		Id:        category.Id,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
		Name:      category.Name,
	}
}

// func ToListFromDomain(users []_usersDomain.Domain) []UserRequest {
// 	var result = []UserRequest{}
// 	for _, user := range users {
// 		result = append(result, FromDomain(user))
// 	}
// 	return result
// }
