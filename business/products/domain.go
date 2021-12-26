package products

import (
	"context"
	"time"
)

type Domain struct {
	Id         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Name       string
	Price      uint
	CategoryId uint
}

type Usecase interface {
	GetProduct(ctx context.Context) ([]Domain, error)
	CreateProduct(ctx context.Context, data Domain) (int, error)
	DeleteProduct(ctx context.Context, id int) (int, error)
}

type Repository interface {
	GetProduct(ctx context.Context) ([]Domain, error)
	CreateProduct(ctx context.Context, data Domain) (int, error)
	DeleteProduct(ctx context.Context, id int) (int, error)
}
