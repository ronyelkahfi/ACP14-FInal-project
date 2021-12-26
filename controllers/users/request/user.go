package request

import (
	_usersDomain "final-project/business/users"
	"time"
)

type UserRequest struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Token     string    `json:"token"`
}

func ToDomain(user UserRequest) _usersDomain.Domain {
	return _usersDomain.Domain{
		Id:        user.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		Email:     "",
	}
}

// func ToListFromDomain(users []_usersDomain.Domain) []UserRequest {
// 	var result = []UserRequest{}
// 	for _, user := range users {
// 		result = append(result, FromDomain(user))
// 	}
// 	return result
// }
