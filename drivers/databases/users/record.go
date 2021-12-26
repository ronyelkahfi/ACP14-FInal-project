package users

import (
	_userDomain "final-project/business/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint `gorm:"primarykey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Name      string `gorm:"size:100"`
	Email     string `gorm:"not null;unique"`
	Password  string `gorm:"size:100"`
}

func (user *User) ToDomain() _userDomain.Domain {
	return _userDomain.Domain{
		Id:        user.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
	}
}

func ToListDomain(users []User) []_userDomain.Domain {
	var result = []_userDomain.Domain{}
	for _, user := range users {
		result = append(result, user.ToDomain())
	}
	return result
}
