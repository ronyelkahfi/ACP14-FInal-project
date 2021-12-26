package categories

import (
	"context"
	"time"
)

type Domain struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

type Usecase interface {
	GetCategory(ctx context.Context) ([]Domain, error)
	CreateCategory(ctx context.Context, data Domain) (int, error)
	DeleteCategory(ctx context.Context, id int) (int, error)
}

type Repository interface {
	GetCategory(ctx context.Context) ([]Domain, error)
	CreateCategory(ctx context.Context, data Domain) (int, error)
	DeleteCategory(ctx context.Context, id int) (int, error)
}
