package transactionDetails

import (
	_transactionDomain "final-project/business/transactions"
	"time"

	"gorm.io/gorm"
)

type Transaction_detail struct {
	Id            uint `gorm:"primarykey;autoIncrement"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	TransactionId uint `gorm:"not null;unique"`
	ProductId     uint `gorm:"not null;unique"`
	Price         uint `gorm:"not null;"`
	Qty           uint `gorm:"not null;"`
}

func (detail *Transaction_detail) ToDomain() _transactionDomain.Detail {
	return _transactionDomain.Detail{
		Id:            detail.Id,
		CreatedAt:     detail.CreatedAt,
		UpdatedAt:     detail.UpdatedAt,
		TransactionId: detail.TransactionId,
		ProductId:     detail.ProductId,
		Price:         detail.Price,
		Qty:           detail.Qty,
	}
}

func ToListDomain(details []Transaction_detail) []_transactionDomain.Detail {
	var result = []_transactionDomain.Detail{}
	for _, detail := range details {
		result = append(result, detail.ToDomain())
	}
	return result
}
