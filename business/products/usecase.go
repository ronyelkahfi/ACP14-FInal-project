package products

import (
	"context"
	"time"
)

type ProductUsecase struct {
	contextTimeout time.Duration
	repository     Repository
}

func NewProductUsecase(repo Repository, timeout time.Duration) Usecase {
	return &ProductUsecase{
		contextTimeout: timeout,
		repository:     repo,
	}
}

func (uc *ProductUsecase) GetProduct(ctx context.Context) ([]Domain, error) {
	// timeout
	users, err := uc.repository.GetProduct(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return users, nil
}

func (uc *ProductUsecase) CreateProduct(ctx context.Context, domain Domain) (int, error) {
	affectedRow, err := uc.repository.CreateProduct(ctx, domain)
	return affectedRow, err
}
func (uc *ProductUsecase) DeleteProduct(ctx context.Context, id int) (int, error) {
	return uc.repository.DeleteProduct(ctx, id)
}
