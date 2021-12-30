package transactions

import (
	"context"
	_TransactionDomain "final-project/business/transactions"
	"time"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) _TransactionDomain.Repository {
	return TransactionRepository{
		db: db,
	}
}

func (transactionRepo TransactionRepository) CreateTransaction(ctx context.Context, data _TransactionDomain.Domain) (uint, error) {
	dataToInsert := Transaction{
		UserId:        data.UserId,
		InvoiceNumber: data.InvoiceNumber,
		InvoiceAmount: data.InvoiceAmount,
		CartStatus:    data.CartStatus,
		PaymentDate:   data.PaymentDate,
	}
	result := transactionRepo.db.Save(&dataToInsert)
	if result.Error != nil {
		return 0, result.Error
	} else {
		return dataToInsert.Id, nil
	}
}

func (transactionRepo TransactionRepository) DeleteTransaction(ctx context.Context, id uint) error {
	transaction := Transaction{}
	// fmt.Println(data)
	result := transactionRepo.db.Model(&transaction).Where("id", id).Update("deleted_at", time.Now())
	return result.Error
}

func (transactionRepo TransactionRepository) IsOpenTransaction(ctx context.Context, userId uint) (uint, error) {
	transaction := Transaction{}
	result := transactionRepo.db.First(&transaction).Where("UserId", userId)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return 0, nil
		} else {
			return 0, result.Error
		}
	} else {
		return transaction.Id, nil
	}
}
