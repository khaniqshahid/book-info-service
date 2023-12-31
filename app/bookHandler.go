package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/khaniqshahid/book-details-service/service"
)

// we are not using the following code in this codebase.
// type BookDetails struct {
// 	Id        uint            `json:"book_id" gorm:"primarykey"`
// 	Title     string          `json:"title" gorm:"title;type:varchar(100);not null"`
// 	Author    string          `json:"author" gorm:"author;type:varchar(100);not null"`
// 	Publisher string          `json:"publisher" gorm:"publisher;type:varchar(100);not null"`
// 	Price     decimal.Decimal `json:"price" gorm:"price;type:decimal(5,2);not null"`
// 	// Price       float64   `json:"price" gorm:"price;type:decimal(5,2);not null"`
// 	// IssuedAt    time.Time `json:"issuedAt" gorm:"issuedAt;type:date;not null"`
// 	IssuedAt    string `json:"issuedAt" gorm:"issuedAt;type:date;not null"`
// 	Description string `json:"description" gorm:"description;type:varchar(300)"`
// }

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

	books, err1 := bh.service.GetAllBook()
	if err1 != nil {
		writeResponse(w, err1.Code, err1.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, books)
	}

	// w.Header().Add("Content-Type", "application/json")
	// // db.Find(&books)
	// json.NewEncoder(w).Encode(books)
}

func (bh *BookInfoHandler) GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["book_id"])
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Book ID", http.StatusBadRequest)
		return
	}
	book, err1 := bh.service.GetBookById(id)
	if err1 != nil {
		// fmt.Println(w, err1.Code, err1.Message)

		writeResponse(w, err1.Code, err1.AsMessage())
		// w.Header().Add("Content-Type", "application/json")
		// w.WriteHeader(err1.Code)
		// json.NewEncoder(w).Encode(err1.AsMessage())

	} else {
		writeResponse(w, http.StatusOK, book)
		// w.Header().Add("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// json.NewEncoder(w).Encode(book)
	}

}

// Helper function to write response so that we can reuse it in different handler.

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
