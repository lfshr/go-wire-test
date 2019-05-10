package services

import (
	"go-wire-test/models"
	"go-wire-test/repositories"
)

type TestBookService struct {
	BookRepository repositories.BookRepository
}

func NewTestBookService(br repositories.BookRepository) BookService {
	return TestBookService{
		BookRepository: br,
	}
}

func (bs TestBookService) GetAllBooks() []models.Book {
	return bs.BookRepository.GetAllBooks()
}
