package transactions

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	Id            uint `gorm:"primarykey;autoIncrement"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	UserId        uint   `gorm:"not null;"`
	InvoiceNumber string `gorm:"size:50;unique"`
	InvoiceAmount uint   `gorm:"not null;"`
	CartStatus    string `gorm:"size:20;"`
	PaymentDate   time.Time
}
