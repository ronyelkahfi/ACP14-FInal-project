package request

import (
	_transactionsDomain "final-project/business/transactions"
	"time"
)

type TransactionRequest struct {
	Id           uint `json:"id"`
	ProductId    uint `json:"product_id"`
	ProductPrice uint `json:"product_price`
	OrderQty     uint `json:"order_qty"`
}

func DetailToDomain(detail TransactionRequest) _transactionsDomain.Detail {
	return _transactionsDomain.Detail{
		Id:            detail.Id,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     time.Time{},
		TransactionId: 0,
		ProductId:     detail.ProductId,
		Qty:           detail.OrderQty,
		Price:         detail.ProductPrice,
	}
}

func DetailToListToDomain(details []TransactionRequest) []_transactionsDomain.Detail {
	var result = []_transactionsDomain.Detail{}
	for _, detail := range details {
		result = append(result, DetailToDomain(detail))
	}
	return result
}
