package response

import (
	_transactionsDomain "final-project/business/transactions"
	"time"
)

type TransactionDetailResponse struct {
	Id            uint      `json:"id"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	TransactionId uint      `json:transactionId`
	ProductId     uint      `json:productId`
	Price         uint      `json:price`
	Qty           uint      `json:qty`
}

func FromDomain(detail _transactionsDomain.Detail) TransactionDetailResponse {
	return TransactionDetailResponse{
		Id:            detail.Id,
		CreatedAt:     detail.CreatedAt,
		UpdatedAt:     detail.UpdatedAt,
		TransactionId: detail.TransactionId,
		ProductId:     detail.ProductId,
		Price:         detail.Price,
		Qty:           detail.Qty,
	}
}

func ToListFromDomain(details []_transactionsDomain.Detail) []TransactionDetailResponse {
	var result = []TransactionDetailResponse{}
	for _, detail := range details {
		result = append(result, FromDomain(detail))
	}
	return result
}
