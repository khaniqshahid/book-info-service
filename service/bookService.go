package service

import (
	"github.com/khaniqshahid/book-details-service/domain"
	"github.com/khaniqshahid/book-details-service/errs"
)

// Introduced PRIMARY-PORT  Interface for book service
type BookService interface {
	GetAllBook() ([]domain.Book, error)
	GetBookById(int) (*domain.Book, *errs.AppError)
}

// Introduced the primary port ADAPTER for book service to connect with Domain's secondary port interface
// Struct for default book service
type DefaultBookService struct {
	repo domain.BookRepository
}

// Reciever Function takes struct as input and returns from interface instance
func (s DefaultBookService) GetAllBook() ([]domain.Book, error) {
	return s.repo.FindAll()
}

// Reciever Function takes struct as input and returns from interface instance
func (s DefaultBookService) GetBookById(id int) (*domain.Book, *errs.AppError) {
	return s.repo.ById(id)
}

// Helper function to instantiate DefaultBookService create a new book service instances
func NewCustomBookService(repository domain.BookRepository) DefaultBookService {
	return DefaultBookService{repo: repository}
}
