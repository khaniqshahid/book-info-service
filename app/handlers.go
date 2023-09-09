package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/khaniqshahid/book-details-service/service"
	"github.com/shopspring/decimal"
)

type BookDetails struct {
	Id        uint            `json:"book_id" gorm:"primarykey"`
	Title     string          `json:"title" gorm:"title;type:varchar(100);not null"`
	Author    string          `json:"author" gorm:"author;type:varchar(100);not null"`
	Publisher string          `json:"publisher" gorm:"publisher;type:varchar(100);not null"`
	Price     decimal.Decimal `json:"price" gorm:"price;type:decimal(5,2);not null"`
	// Price       float64   `json:"price" gorm:"price;type:decimal(5,2);not null"`
	// IssuedAt    time.Time `json:"issuedAt" gorm:"issuedAt;type:date;not null"`
	IssuedAt    string `json:"issuedAt" gorm:"issuedAt;type:date;not null"`
	Description string `json:"description" gorm:"description;type:varchar(300)"`
}

// Created a struct which has dependency on Book service to help passing it to Book handler to connect with Domain layer
type BookInfoHandler struct {
	service service.BookService
}

// Connect the right handler to make the request to get all the books detail through the book service ---> domain ---> repository ---> interface --> stub/db
// handler :bh BookDetailHandler
func (bh *BookInfoHandler) getAllBooks(w http.ResponseWriter, r *http.Request) {
	// books := []BookDetails{
	// 	{
	// 		Id:        1,
	// 		Title:     "The Java Programming Language",
	// 		Author:    "Shahid",
	// 		Publisher: "Packt Publishing Ltd",
	// 		// Price:     float64(58.99),
	// 		Price:       decimal.NewFromFloat(58.99),
	// 		IssuedAt:    time.Now(),
	// 		Description: "A practical guide to the Go programming language.",
	// 	},
	// 	{
	// 		Id:        2,
	// 		Title:     "The Go Programming Language",
	// 		Author:    "Kushal",
	// 		Publisher: "Packt Publishing Ltd",
	// 		// Price:     float64(48.99),
	// 		Price:       decimal.NewFromFloat(49.99),
	// 		IssuedAt:    time.Now(),
	// 		Description: "A practical guide to the Go programming language.",
	// 	},
	// }

	books, _ := bh.service.GetAllBook()
	w.Header().Add("Content-Type", "application/json")
	// db.Find(&books)
	json.NewEncoder(w).Encode(books)
}

func (bh *BookInfoHandler) GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["book_id"])
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}
	book, err := bh.service.GetBookById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
		return
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(book)
	}

}
