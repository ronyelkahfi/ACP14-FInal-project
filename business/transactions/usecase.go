package transactions

import (
	"context"
	"math/rand"
	"time"
)

type TransactionUsecase struct {
	contextTimeout   time.Duration
	repository       Repository
	detailrepository DetailRepository
}

func NewTransactionUsecase(repo Repository, detailrepo DetailRepository, timeout time.Duration) Usecase {
	return &TransactionUsecase{
		contextTimeout:   timeout,
		repository:       repo,
		detailrepository: detailrepo,
	}
}

func (trxUsecase TransactionUsecase) NewTransaction(ctx context.Context, userId uint, detail Detail) ([]Detail, error) {

	transactionId, err := trxUsecase.repository.IsOpenTransaction(ctx, userId)

	if err != nil {
		return []Detail{}, err
	} else {
		var total uint
		// productdomain.Repository.GetProductById(detail.Id)
		total += detail.Price

		if transactionId == 0 {
			// create new transaction
			dataTransaction := Domain{
				UserId:        userId,
				InvoiceNumber: RandStringBytes(5),
				InvoiceAmount: total,
			}
			var err error
			transactionId, err = trxUsecase.repository.CreateTransaction(ctx, dataTransaction)
			if err.Error != nil {
				return []Detail{}, err
			}
		}
		detail.TransactionId = transactionId
		err := trxUsecase.detailrepository.CreateDetail(ctx, detail)
		if err != nil {
			return []Detail{}, err
		}
		// var err error
		details, err := trxUsecase.detailrepository.GetByTransactionId(ctx, transactionId)
		return details, err
	}
}

func (trxUsecase TransactionUsecase) DeleteTransaction(ctx context.Context, id int) (transactionId int, error error) {

	return
}

func (trxUsecase TransactionUsecase) DeleteDetailTransaction(ctx context.Context, id int) (detailId int, error error) {

	return
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// func (uc *UserUsecase) GetUser(ctx context.Context) ([]Domain, error) {
// 	// timeout
// 	users, err := uc.repository.GetUser(ctx)
// 	if err != nil {
// 		return []Domain{}, err
// 	}
// 	return users, nil
// }

// func (uc *UserUsecase) Register(ctx context.Context, domain Domain) (int, error) {
// 	var affectedRow int
// 	var err error
// 	affectedRow, err = uc.repository.Register(ctx, domain)
// 	return affectedRow, err
// }

// func (uc *UserUsecase) Login(ctx context.Context, email string, password string) (string, error) {
// 	var err error
// 	var UserData Domain
// 	UserData, err = uc.repository.GetByEmail(ctx, email)
// 	if err != nil {
// 		return "", err
// 	}

// 	if UserData.Password != password {
// 		return "", errors.New("Password mismatch!")
// 	} else {

// 		return CreateToken(UserData)
// 	}
// 	//affectedRow, err := uc.repository.GetByEmail(ctx, email)
// 	// return token, err
// }

// func CreateToken(data Domain) (string, error) {
// 	claims := _jwt.MapClaims{}
// 	claims["authorized"] = true
// 	claims["userId"] = data.Id
// 	claims["exp"] = time.Now().Add(time.Hour * 6)
// 	token := _jwt.NewWithClaims(_jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte("ThisisSecretGais"))
// }
