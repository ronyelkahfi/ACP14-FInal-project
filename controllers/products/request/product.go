package request

import (
	_productsDomain "final-project/business/products"
	"time"
)

type ProductRequest struct {
	Id         uint      `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Name       string    `json:"name"`
	Price      uint      `json:"price"`
	CategoryId uint      `json:"category_id"`
}

func ToDomain(product ProductRequest) _productsDomain.Domain {
	return _productsDomain.Domain{
		Id:         product.Id,
		CreatedAt:  product.CreatedAt,
		UpdatedAt:  product.UpdatedAt,
		Name:       product.Name,
		Price:      uint(product.Price),
		CategoryId: uint(product.CategoryId),
	}
}

// func ToListFromDomain(users []_usersDomain.Domain) []UserRequest {
// 	var result = []UserRequest{}
// 	for _, user := range users {
// 		result = append(result, FromDomain(user))
// 	}
// 	return result
// }
