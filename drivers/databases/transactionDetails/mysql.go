package transactionDetails

import (
	"context"
	_TransactionDomain "final-project/business/transactions"
	"time"

	"gorm.io/gorm"
)

type TransDetailRepository struct {
	db *gorm.DB
}

func NewTransDetailRepository(db *gorm.DB) _TransactionDomain.DetailRepository {
	return TransDetailRepository{
		db: db,
	}
}

func (transdetailRepo TransDetailRepository) CreateDetail(ctx context.Context, data _TransactionDomain.Detail) error {
	dataToInsert := Transaction_detail{
		TransactionId: data.TransactionId,
		ProductId:     data.ProductId,
		Price:         data.Price,
		Qty:           data.Qty,
	}
	result := transdetailRepo.db.Save(&dataToInsert)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (transdetailRepo TransDetailRepository) DeleteDetail(ctx context.Context, id uint) error {
	transaction_detail := Transaction_detail{}
	// fmt.Println(data)
	result := transdetailRepo.db.Model(&transaction_detail).Where("id", id).Update("deleted_at", time.Now())
	return result.Error
}

func (transetailRepo TransDetailRepository) GetByTransactionId(ctx context.Context, transactionId uint) ([]_TransactionDomain.Detail, error) {
	// var transaction_detail []_TransactionDomain.Detail
	// fmt.Println(data)
	var transaction_detail []Transaction_detail
	result := transetailRepo.db.Find(&transaction_detail).Where("transaction_id", transactionId)
	if result.Error != nil {
		return []_TransactionDomain.Detail{}, result.Error
	}
	return ToListDomain(transaction_detail), nil
}
