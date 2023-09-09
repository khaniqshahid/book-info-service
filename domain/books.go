package domain

import (
	"time"

	"github.com/khaniqshahid/book-details-service/errs"
	"github.com/shopspring/decimal"
)

type Book struct {
	Id          int             `db:"book_id" json:"book_id" `                               // gorm:"primarykey"`
	Title       string          `db:"title";type:varchar(100);not null json:"title"`         // gorm:"title;type:varchar(100);not null"`
	Author      string          `db:"author";type:varchar(100);not null json:"author" `      // gorm:"author;type:varchar(100);not null"`
	Publisher   string          `db:"publisher";type:varchar(100);not null json:"publisher"` // gorm:"publisher;type:varchar(100);not null"`
	Price       decimal.Decimal `db:"price";type:decimal(5,2);not null json:"price"`         // gorm:"price;type:decimal(5,2);not null"`
	IssuedAt    time.Time       `db:"issued_at";type:date;not null json:"issued_at"`         //gorm:"issuedAt;type:date;not null"`
	Description string          `db:"description";type:varchar(300)" json:"description"`     // gorm:"description;type:varchar(300)"`
}

// Introduced a SECONDARY-PORT Interface here.
type BookRepository interface {
	FindAll() ([]Book, *errs.AppError)
	ById(int) (*Book, *errs.AppError)
}
