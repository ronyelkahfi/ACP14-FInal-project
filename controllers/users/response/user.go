package response

import (
	_usersDomain "final-project/business/users"
	"time"
)

type UserResponse struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Email     string    `json:"email,omitempty"`
}

func FromDomain(user _usersDomain.Domain) UserResponse {
	return UserResponse{
		Id:        user.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		Email:     "",
	}
}

func ToListFromDomain(users []_usersDomain.Domain) []UserResponse {
	var result = []UserResponse{}
	for _, user := range users {
		result = append(result, FromDomain(user))
	}
	return result
}
