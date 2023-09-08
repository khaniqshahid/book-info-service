package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

// Created an ADAPTER for the BookRepository interface port.
type BookRepositoryStub struct {
	books []Book
}

func (s BookRepositoryStub) FindAll() ([]Book, error) {
	return s.books, nil
}

func NewBookRepositoryStub() BookRepositoryStub {
	books := []Book{
		{
			Id:        1,
			Title:     "The Go Programming Language",
			Author:    "Aidan",
			Publisher: "Packt Publishing Ltd",
			Price:     decimal.NewFromFloat(100),
			IssuedAt:  time.Now(),
			// IssuedAt: "1980-11-02",

			Description: "The Go Programming Language is a high-level, imperative, object-oriented programming language.",
		},
		{
			Id:        2,
			Title:     "The Python Programming Language",
			Author:    "David",
			Publisher: "Packt Publishing Ltd",
			Price:     decimal.NewFromFloat(100),
			// IssuedAt:  "1985-11-02",
			IssuedAt:    time.Now(),
			Description: "The Go Programming Language is a high-level, imperative, object-oriented programming language.",
		},
	}
	return BookRepositoryStub{books}
}
