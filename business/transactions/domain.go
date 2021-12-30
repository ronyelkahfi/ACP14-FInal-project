package transactions

import (
	"context"
	"time"
)

type Domain struct {
	Id            uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	UserId        uint
	InvoiceNumber string
	InvoiceAmount uint
	CartStatus    string
	PaymentStatus string
	PaymentDate   time.Time
}

type Detail struct {
	Id            uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
	TransactionId uint
	ProductId     uint
	Qty           uint
	Price         uint
}

type Usecase interface {
	NewTransaction(ctx context.Context, userId uint, detail Detail) ([]Detail, error)
	DeleteTransaction(ctx context.Context, id int) (int, error)
	DeleteDetailTransaction(ctx context.Context, id int) (int, error)
}

type Repository interface {
	CreateTransaction(ctx context.Context, data Domain) (uint, error)
	DeleteTransaction(ctx context.Context, id uint) error
	IsOpenTransaction(ctx context.Context, userId uint) (transactionid uint, error error)
}

type DetailRepository interface {
	CreateDetail(ctx context.Context, data Detail) error
	DeleteDetail(ctx context.Context, id uint) error
	GetByTransactionId(ctx context.Context, transactionId uint) ([]Detail, error)
}
