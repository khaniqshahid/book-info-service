package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

// type BookResponse struct {
// 	Id          int             `db:"book_id" json:"book_id" `                               // gorm:"primarykey"`
// 	Title       string          `db:"title";type:varchar(100);not null json:"title"`         // gorm:"title;type:varchar(100);not null"`
// 	Author      string          `db:"author";type:varchar(100);not null json:"author" `      // gorm:"author;type:varchar(100);not null"`
// 	Publisher   string          `db:"publisher";type:varchar(100);not null json:"publisher"` // gorm:"publisher;type:varchar(100);not null"`
// 	Price       decimal.Decimal `db:"price";type:decimal(5,2);not null json:"price"`         // gorm:"price;type:decimal(5,2);not null"`
// 	IssuedAt    time.Time       `db:"issued_at";type:date;not null json:"issued_at"`         //gorm:"issuedAt;type:date;not null"`
// 	Description string          `db:"description";type:varchar(300)" json:"description"`     // gorm:"description;type:varchar(300)"`
// }

type BookResponse struct {
	Id          int             `json:"book_id" `    // gorm:"primarykey"`
	Title       string          `json:"title"`       // gorm:"title;type:varchar(100);not null"`
	Author      string          `json:"author" `     // gorm:"author;type:varchar(100);not null"`
	Publisher   string          `json:"publisher"`   // gorm:"publisher;type:varchar(100);not null"`
	Price       decimal.Decimal `json:"price"`       // gorm:"price;type:decimal(5,2);not null"`
	IssuedAt    time.Time       `json:"issued_at"`   //gorm:"issuedAt;type:date;not null"`
	Description string          `json:"description"` // gorm:"description;type:varchar(300)"`
}
