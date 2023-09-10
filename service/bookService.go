package service

import (
	"github.com/khaniqshahid/book-details-service/domain"
	"github.com/khaniqshahid/book-details-service/dto"
	"github.com/khaniqshahid/book-details-service/errs"
)

// Introduced PRIMARY-PORT  Interface for book service
type BookService interface {
	GetAllBook() ([]dto.BookResponse, *errs.AppError)
	GetBookById(int) (*dto.BookResponse, *errs.AppError)
}

// Introduced the primary port ADAPTER for book service to connect with Domain's secondary port interface
// Using the SECONDARY PORT Interface(have functions to run query in domain repository) in a type Struct here as Default book service to pass or connect with secondary por
type DefaultBookService struct {
	repo domain.BookRepository
}

// Reciever Function takes struct as input and returns from interface instance
// func
func (s DefaultBookService) GetAllBook() ([]dto.BookResponse, *errs.AppError) {
	books, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	bookResponse := make([]dto.BookResponse, 0)
	for _, b := range books {
		bookResponse = append(bookResponse, b.ToDto())
	}
	return bookResponse, nil
}

/*
Reciever Function takes struct object of domain repo with argument id. If it get data from DB, then it will construct
the response object and returns from interface instance
*/
func (s DefaultBookService) GetBookById(id int) (*dto.BookResponse, *errs.AppError) {
	// fetching book from domain's repository db from SECONDARY PORT Interface
	b, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	// construct a dto by a function in domain
	bookResponse := b.ToDto()
	return &bookResponse, nil
}

// Helper function to instantiate DefaultBookService create a new book service instances
func NewCustomBookService(repository domain.BookRepository) DefaultBookService {
	return DefaultBookService{repo: repository}
}
