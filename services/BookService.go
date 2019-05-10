package services

import "go-wire-test/models"

type BookService interface {
	GetAllBooks() []models.Book
}
