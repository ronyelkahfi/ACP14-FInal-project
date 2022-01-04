package users

import (
	"context"
	"errors"
	_middleware "final-project/apps/middlewares"
	"time"
)

type UserUsecase struct {
	contextTimeout time.Duration
	repository     Repository
}

func NewUserUsecase(repo Repository, timeout time.Duration) Usecase {
	return &UserUsecase{
		contextTimeout: timeout,
		repository:     repo,
	}
}

func (uc *UserUsecase) GetUser(ctx context.Context) ([]Domain, error) {
	// timeout
	users, err := uc.repository.GetUser(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return users, nil
}

func (uc *UserUsecase) Register(ctx context.Context, domain Domain) (int, error) {
	var affectedRow int
	var err error
	affectedRow, err = uc.repository.Register(ctx, domain)
	return affectedRow, err
}

func (uc *UserUsecase) Login(ctx context.Context, email string, password string) (string, error) {
	var err error
	var UserData Domain
	UserData, err = uc.repository.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if UserData.Password != password {
		return "", errors.New("Password mismatch!")
	} else {
		return _middleware.GenerateToken(UserData.Id)
	}
	//affectedRow, err := uc.repository.GetByEmail(ctx, email)
	// return token, err
}

// func CreateToken(data Domain) (string, error) {
// 	claims := _jwt.MapClaims{}
// 	claims["authorized"] = true
// 	claims["userId"] = data.Id
// 	claims["exp"] = time.Now().Add(time.Hour * 6)
// 	token := _jwt.NewWithClaims(_jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte("ThisisSecretGais"))
// }
