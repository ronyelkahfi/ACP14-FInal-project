package categories

import (
	"context"
	"time"
)

type CategoryUsecase struct {
	contextTimeout time.Duration
	repository     Repository
}

func NewCategoryUsecase(repo Repository, timeout time.Duration) Usecase {
	return &CategoryUsecase{
		contextTimeout: timeout,
		repository:     repo,
	}
}

func (uc *CategoryUsecase) GetCategory(ctx context.Context) ([]Domain, error) {
	// timeout
	users, err := uc.repository.GetCategory(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return users, nil
}

func (uc *CategoryUsecase) CreateCategory(ctx context.Context, domain Domain) (int, error) {
	var affectedRow int
	var err error
	affectedRow, err = uc.repository.CreateCategory(ctx, domain)
	return affectedRow, err
}

func (uc *CategoryUsecase) DeleteCategory(ctx context.Context, id int) (int, error) {
	var affectedRow int
	var err error
	affectedRow, err = uc.repository.DeleteCategory(ctx, id)
	return affectedRow, err
}
