package domain

import (
	"time"

	"github.com/khaniqshahid/book-details-service/errs"
	"github.com/shopspring/decimal"
)

type Book struct {
	Id        int             `json:"book_id"`   // gorm:"primarykey"`
	Title     string          `json:"title"`     // gorm:"title;type:varchar(100);not null"`
	Author    string          `json:"author"`    // gorm:"author;type:varchar(100);not null"`
	Publisher string          `json:"publisher"` // gorm:"publisher;type:varchar(100);not null"`
	Price     decimal.Decimal `json:"price"`     // gorm:"price;type:decimal(5,2);not null"`
	// Price       float64   `json:"price" gorm:"price;type:decimal(5,2);not null"`
	// IssuedAt string `json:"issuedAt"` // gorm:"issuedAt;type:date;not null"`
	IssuedAt    time.Time `json:"issuedAt" gorm:"issuedAt;type:date;not null"`
	Description string    `json:"description"` // gorm:"description;type:varchar(300)"`
}

// Introduced a SECONDARY-PORT Interface here.
type BookRepository interface {
	FindAll() ([]Book, *errs.AppError)
	ById(int) (*Book, *errs.AppError)
}
